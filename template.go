package gotemplate

type Templater interface {
	AddFunc(f Function)
}

type Config struct {
	Address  string // mysql address url
	Database string // database name
	Table    string // table name
}

type template struct {
	Functions []Function
}

func NewTemplate(config Config) Templater {
	return &template{}
}

func (t *template) AddFunc(f Function) {
	t.Functions = append(t.Functions, f)
}

type Function interface {
}

type function struct {
	Name string
	Arg  any
	Ret  any
}

func NewFunc(name string) Function {
	return &function{
		Name: name,
	}
}
