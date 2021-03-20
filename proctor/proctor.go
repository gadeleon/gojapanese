package gojapanese


type question struct {
	// Where do we store questions?
	Prompt string

}
// Give back a question struct but also give
// a question upon creation
func NewProctor() *question {
	b := &question{}
	b.GetQuestion()
	return b
}

// Plop in the prompt to the created question struct
func (q *question) GetQuestion() {
	q.Prompt = "Hello World - this is the question"
}

