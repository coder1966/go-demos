package bridge

// SMS
// Email

// 做了2个抽象

type AbstractMessage interface{ SendMessage(text, to string) } // 抽象：发送快，发送慢普通
type MessageImlementer interface{ Send(text, to string) }      // 抽象： 短信|邮件
