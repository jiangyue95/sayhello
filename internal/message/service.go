package message

func CreateMessage(msg *Message) error {
	return Create(msg)
}

func ListMessages() ([]Message, error) {
	return GetAll()
}
