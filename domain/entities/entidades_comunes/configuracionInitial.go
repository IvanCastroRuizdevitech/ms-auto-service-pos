package comunes_entidades

import "encoding/json"

type ConfiguracionInicial struct {
	Status  int               `json:"status"`
	Mensaje string            `json:"mensaje"`
	Data    ConfiguracionData `json:"data"`
}

type ConfiguracionData struct {
	Surtidores         []Surtidor        `json:"surtidores"`
	SurtidoresDetalles []SurtidorDetalle `json:"surtidores_detalles"`
	Productos          []Producto        `json:"productos"`
	FamiliasPrecios    []ProductoFamilia `json:"familias_precios"`
	Empresas           []Empresa         `json:"empresas"`
	Equipos            []Equipo          `json:"equipos"`
}

type Surtidor struct {
	Id                            int64    `json:"id"`
	Surtidor                      int32    `json:"surtidor"`
	IslasId                       int64    `json:"islas_id"`
	Estado                        string   `json:"estado"`
	SurtidoresTiposId             int64    `json:"surtidores_tipos_id"`
	SurtidoresProtocolosId        int64    `json:"surtidores_protocolos_id"`
	Mac                           string   `json:"mac"`
	Ip                            string   `json:"ip"`
	Port                          string   `json:"port"`
	CreateUser                    int32    `json:"create_user"`
	CreateDate                    string   `json:"create_date"`
	UpdateUser                    *int32   `json:"update_user"`
	UpdateDate                    *string  `json:"update_date"`
	Token                         string   `json:"token"`
	FactorVolumenParcial          float64  `json:"factor_volumen_parcial"`
	FactorImporteParcial          float64  `json:"factor_importe_parcial"`
	FactorPrecio                  float64  `json:"factor_precio"`
	LectorIp                      *string  `json:"lector_ip"`
	LectorPort                    *int32   `json:"lector_port"`
	ImpresoraIp                   *string  `json:"impresora_ip"`
	ImpresoraPort                 *int32   `json:"impresora_port"`
	FactorInventario              *float64 `json:"factor_inventario"`
	EmpresasId                    *int64   `json:"empresas_id"`
	Controlador                   *int32   `json:"controlador"`
	DebugEstado                   *string  `json:"debug_estado"`
	DebugTramas                   *string  `json:"debug_tramas"`
	TieneEcho                     *string  `json:"tiene_echo"`
	FactorPredeterminacionVolumen int32    `json:"factor_predeterminacion_volumen"`
	BytesTotalizador              *int32   `json:"bytes_totalizador"`
}

type SurtidorDetalle struct {
	Id                            int64    `json:"id"`
	SurtidoresId                  int64    `json:"surtidores_id"`
	Surtidor                      int32    `json:"surtidor"`
	Cara                          *int32   `json:"cara"`
	Manguera                      int32    `json:"manguera"`
	Grado                         *int32   `json:"grado"`
	ProductosId                   int64    `json:"productos_id"`
	AcumuladoVenta                *float64 `json:"acumulado_venta"`
	AcumuladoCantidad             *float64 `json:"acumulado_cantidad"`
	UltimaConexion                *string  `json:"ultima_conexion"`
	LectorPuerto                  *int32   `json:"lector_puerto"`
	SaltoLectura                  *string  `json:"salto_lectura"`
	AcumuladoVentaSurt            *float64 `json:"acumulado_venta_surt"`
	AcumuladoCantidadSurt         *float64 `json:"acumulado_cantidad_surt"`
	Estado                        *int32   `json:"estado"`
	EstadoPublico                 *int32   `json:"estado_publico"`
	BodegasId                     *int64   `json:"bodegas_id"`
	LectorRfid                    *string  `json:"lector_rfid"`
	Bloqueo                       *string  `json:"bloqueo"`
	MotivoBloqueo                 *string  `json:"motivo_bloqueo"`
	Conexion                      *int32   `json:"conexion"`
	Puerto                        *string  `json:"puerto"`
	BytesTotalizador              *int32   `json:"bytes_totalizador"`
	FactorPredeterminacionImporte int32    `json:"factor_predeterminacion_importe"`
}

type Producto struct {
	Id               int64    `json:"id"`
	Descripcion      string   `json:"descripcion"`
	Estado           string   `json:"estado"`
	ProductosTiposId int64    `json:"productos_tipos_id"`
	EmpresasId       int64    `json:"empresas_id"`
	Familias         *int64   `json:"familias"`
	Publico          *int64   `json:"publico"`
	Precio           *float64 `json:"precio"`
	Precio2          *float64 `json:"precio2"`
	UnidadMedidaId   *int64   `json:"unidad_medida_id"`
	Sku              *string  `json:"sku"`
	TipoNegocioId    *int32   `json:"tipo_negocio_id"`
	EstadoId         int16    `json:"estado_id"`
}

type ProductoFamilia struct {
	Id        int64            `json:"id"`
	Codigo    string           `json:"codigo"`
	Atributos *json.RawMessage `json:"atributos"`
}

type Empresa struct {
	Id                     int64            `json:"id"`
	RazonSocial            string           `json:"razon_social"`
	Nit                    string           `json:"nit"`
	Direccion              *string          `json:"direccion"`
	Telefono               *string          `json:"telefono"`
	Estado                 string           `json:"estado"`
	Correo                 *string          `json:"correo"`
	ContactoNombre         *string          `json:"contacto_nombre"`
	ContactoTelefono       *string          `json:"contacto_telefono"`
	ContactoCorreo         *string          `json:"contacto_correo"`
	FechaCreacion          *string          `json:"fecha_creacion"`
	Localizacion           *string          `json:"localizacion"`
	EmpresasId             *int64           `json:"empresas_id"`
	CantidadSucursales     *int32           `json:"cantidad_sucursales"`
	CiudadesId             *int64           `json:"ciudades_id"`
	CreateUser             *int64           `json:"create_user"`
	CreateDate             *string          `json:"create_date"`
	UpdateUser             *int64           `json:"update_user"`
	UpdateDate             *string          `json:"update_date"`
	EmpresasTiposId        *int64           `json:"empresas_tipos_id"`
	UrlFoto                *string          `json:"url_foto"`
	Dominio                *string          `json:"dominio"`
	Alias                  *string          `json:"alias"`
	CodigoEmpresa          *string          `json:"codigo_empresa"`
	NegocioId              *int64           `json:"negocio_id"`
	DominioId              *int64           `json:"dominio_id"`
	CiudadesDescripcion    *string          `json:"ciudades_descripcion"`
	Atributos              *json.RawMessage `json:"atributos"`
	IdTipoEmpresa          *int64           `json:"id_tipo_empresa"`
	ProveedorTecnologicoId *int64           `json:"proveedor_tecnologico_id"`
}

type Equipo struct {
	Id                  int64   `json:"id"`
	EmpresasId          *int64  `json:"empresas_id"`
	SerialEquipo        string  `json:"serial_equipo"`
	AlmacenamientosId   *int64  `json:"almacenamientos_id"`
	Estado              string  `json:"estado"`
	EquiposTiposId      *int64  `json:"equipos_tipos_id"`
	EquiposProtocolosId *int64  `json:"equipos_protocolos_id"`
	Mac                 string  `json:"mac"`
	Ip                  *string `json:"ip"`
	Port                *string `json:"port"`
	CreateUser          int32   `json:"create_user"`
	CreateDate          string  `json:"create_date"`
	UpdateUser          *int32  `json:"update_user"`
	UpdateDate          *string `json:"update_date"`
	Token               string  `json:"token"`
	Password            string  `json:"password"`
	FactorPrecio        *int32  `json:"factor_precio"`
	FactorImporte       *int32  `json:"factor_importe"`
	FactorInventario    *int32  `json:"factor_inventario"`
	LectorIp            *string `json:"lector_ip"`
	LectorPort          *int32  `json:"lector_port"`
	ImpresoraIp         *string `json:"impresora_ip"`
	ImpresoraPort       *int32  `json:"impresora_port"`
	UrlFoto             *string `json:"url_foto"`
	Autorizado          string  `json:"autorizado"`
}
