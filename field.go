package form

type ValidationFunc func(string) bool

type Field struct {
	Name     string
	Text     string
	Value    string
	Validate ValidationFunc
}
