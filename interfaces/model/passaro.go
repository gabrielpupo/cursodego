package model

//Galinha representa uma ave do tipo Galinha
type Galinha interface {
	Cacareja() string
}

//Pata representa uma ave do tipo Pato
type Pata interface {
	Quack() string
}

//Ave representa um animal
type Ave struct {
	Nome string
}

//Cacareja retorna o som emitido por uma galinha
func (a Ave) Cacareja() string {
	return "Cocoricó..."
}

//Quack retorna o som emitido por uma pata
func (a Ave) Quack() string {
	return "Quack, quack..."
}
