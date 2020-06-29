package sms_sender

import "log"

type SMSSender struct{}

func (s *SMSSender) Send(phoneNumber string, message string) (err error) {
	/*
		Implementation Here
	*/
	log.Println(phoneNumber, message)
	return
}
