package sms_sender_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ariefsam/go-chat/sms_sender"
)

func TestSend(t *testing.T) {
	var sms sms_sender.SMSSender
	err := sms.Send("6285212323", "Testing sms sender")
	assert.NoError(t, err)
}
