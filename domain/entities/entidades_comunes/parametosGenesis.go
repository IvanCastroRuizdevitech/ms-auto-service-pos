package comunes_entidades

type Parametros struct {
	Id           int64
	Codigo       string
	Valor        string
	Tipo         string
	Opciones     string
	ValorDefault string
	Descripcion  string
}

type Parametrizacion struct {
	Valor string
}
