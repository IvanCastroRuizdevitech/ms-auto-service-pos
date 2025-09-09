package infraestructura_db_cliente

import (
	"context"
	"fmt"
	"genexis/pos/autoservicios/domain/constants"
	"net/url"
	"regexp"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ClientDB struct {
	Conn *pgxpool.Pool

	UrlConnection string
	context       context.Context
}

func (CDB *ClientDB) Select(query string, arguments []any) ([][]interface{}, error) {
	connection, err := CDB.Conn.Acquire(CDB.context)

	if err != nil {
		return nil, err
	}
	defer connection.Release()
	err = connection.Ping(CDB.context)
	if err != nil {
		return nil, err
	}
	rows, err := connection.Query(CDB.context, query, arguments...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result [][]interface{}
	for rows.Next() {
		valores, err := rows.Values()

		if err != nil {
			return nil, err
		}
		result = append(result, valores)

	}

	return result, nil

}

func (CDB *ClientDB) Exec(query string, arguments []any) ([][]interface{}, error) {
	connection, err := CDB.Conn.Acquire(CDB.context)
	if err != nil {
		return nil, fmt.Errorf("error al adquirir conexión: %w", err)
	}
	defer connection.Release()

	if err := connection.Ping(CDB.context); err != nil {
		return nil, fmt.Errorf("error al hacer ping a la base de datos: %w", err)
	}

	rows, err := connection.Query(CDB.context, query, arguments...)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta: %w", err)
	}
	defer rows.Close()

	var result [][]interface{}
	for rows.Next() {
		valores, err := rows.Values()
		if err != nil {
			return nil, fmt.Errorf("error al leer valores de la fila: %w", err)
		}
		result = append(result, valores)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("error iterando sobre filas: %w", rows.Err())
	}

	return result, nil
}

const defaultMaxConns = int32(4)
const defaultMinConns = int32(0)
const defaultMaxConnLifetime = time.Hour
const defaultMaxConnIdleTime = time.Minute * 30
const defaultHealthCheckPeriod = time.Minute
const defaultConnectTimeout = time.Second * 5

func CleanConectionString(rawConn string) (string, error) {
	re := regexp.MustCompile(`^postgres(?:ql)?://([^:@/]+):([^@]+)@([^/]+)(/[^?]+)?(\?.*)?$`)

	matches := re.FindStringSubmatch(rawConn)
	if len(matches) == 0 {
		return "", fmt.Errorf("formato inválido de conexión")
	}

	user := matches[1]
	pass := matches[2]
	hostPort := matches[3]
	dbPath := matches[4]
	query := matches[5]

	escapedUser := url.QueryEscape(user)
	escapedPass := url.QueryEscape(pass)

	finalURL := fmt.Sprintf("postgres://%s:%s@%s%s%s",
		escapedUser,
		escapedPass,
		hostPort,
		dbPath,
		query,
	)

	return finalURL, nil
}

func InitializeClient(UrlConn string) (*ClientDB, error) {
	urlConnection, _err := CleanConectionString(constants.DB_CON)
	if _err != nil {
		return nil, _err
	}

	client := &ClientDB{
		context:       context.Background(),
		UrlConnection: urlConnection,
	}

	if dbConfig, err := pgxpool.ParseConfig(client.UrlConnection); err != nil {
		return nil, err
	} else {
		dbConfig.MaxConns = defaultMaxConns
		dbConfig.MinConns = defaultMinConns
		dbConfig.MaxConnLifetime = defaultMaxConnLifetime
		dbConfig.MaxConnIdleTime = defaultMaxConnIdleTime
		dbConfig.HealthCheckPeriod = defaultHealthCheckPeriod
		dbConfig.ConnConfig.ConnectTimeout = defaultConnectTimeout
		client.Conn, err = pgxpool.NewWithConfig(client.context, dbConfig)

		if err != nil {
			return nil, err
		}
	}

	return client, nil
}
func (CDB *ClientDB) BeginTx(ctx context.Context) (pgx.Tx, error) {
	return CDB.Conn.Begin(ctx)
}
