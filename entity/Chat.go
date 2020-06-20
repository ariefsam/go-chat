package entity

type Chat struct {
	SenderID   string
	ReceiverID string
	ChatType   string
	Message    string
	Timestamp  int64
}
