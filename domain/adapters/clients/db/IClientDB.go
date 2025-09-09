package domain_adapters_clients_db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type IClientDB interface {
	Select(query string, arguments []any) ([][]interface{}, error)
	Exec(query string, arguments []any) ([][]interface{}, error)
	BeginTx(ctx context.Context) (pgx.Tx, error)
}
