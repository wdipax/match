package response

const (
	Text = iota
	Control
	BoysToken
	GirlsToken
	RestrictedForAdmin
	Restricted
	Joined
	Failed
	Success
	ViewBoys
	ViewGirls
	KnowEachother
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
