package entity

type Chat struct {
	ID         string
	SenderID   string
	ReceiverID string
	ChatType   string
	Message    string
	Timestamp  int64
}
