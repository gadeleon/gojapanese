package gojapanese

type Word interface {
	Negative(w string) string
}

type Noun struct {
	Name string
}

func (N *Noun) Negative() string{
	return N.Name + "じゃない"
}

func (N *Noun) IsNegative(s string) bool {
	return s == N.Negative()
}

