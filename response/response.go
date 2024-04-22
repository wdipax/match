package response

type Response struct {
	Messages []*Message
}

func New() *Response {
	return &Response{}
}

type Message struct {
	Text string
}
