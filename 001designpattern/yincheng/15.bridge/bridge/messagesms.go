package bridge

import "fmt"

type MessageSMS struct{}

func (msms *MessageSMS) Send(text, to string) { fmt.Println("发送->短信->接收", text, to) }

func ViaSMS() MessageImlementer { return &MessageSMS{} }
