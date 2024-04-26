package response

const (
	Text = iota
	BoysToken
	GirlsToken
)

type Message struct {
	To   int64
	Data string
	Type int
}

type Response struct {
	Messages []Message
}

func New() *Response {
	return &Response{}
}
