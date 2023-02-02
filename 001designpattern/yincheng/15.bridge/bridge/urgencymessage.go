package bridge

// urgency 紧迫性

type UrgencyMessage struct{ method MessageImlementer }

func NewUrgencyMessage(method MessageImlementer) *UrgencyMessage {
	return &UrgencyMessage{method: method}
}

func (m *UrgencyMessage) SendMessage(text, to string) {
	m.method.Send("发送到["+text+"]", to) // 很快速发
}
