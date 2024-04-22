package response

type Response struct {
	Messages []*Message
}

func New() *Response {
	return &Response{}
}

func (r *Response) GetMessages() []*Message {
	if r == nil {
		return nil
	}

	return r.Messages
}

type Message struct {
	ChatID int64
	Text   string
}
