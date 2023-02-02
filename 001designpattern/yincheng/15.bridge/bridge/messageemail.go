package bridge

import "fmt"

type MessageEmail struct{}

func (msms *MessageEmail) Send(text, to string) { fmt.Println("发送->邮件->接收", text, to) }

func ViaEmail() MessageImlementer { return &MessageEmail{} }
