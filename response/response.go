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

type MessageType int

const (
	Text MessageType = iota
	BoysLink
	GirlsLink
	TeamRegistration
	KnowEachOther
	Voting
)

type Message struct {
	ChatID int64
	Text   string
	Type   MessageType
}
