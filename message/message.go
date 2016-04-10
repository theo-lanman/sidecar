package message

import (
	"encoding/binary"
	"time"
)

type Message struct {
	Attempts    int
	CreatedAt   time.Time
	AttemptedAt time.Time
	Id          uint64
	Body        []byte
}

type Batch []Message

func NewMessage(id uint64, body []byte) Message {
	return Message{Attempts: 0, CreatedAt: time.Now(), Id: id, Body: body}
}

func NewBatch() Batch {
	return make([]Message, 0, 20)
}

// IdBytes returns the message's id as a big-endian slice of bytes, such that
// the resulting byte slice sort order matches the integer sort order.
func (msg Message) IdBytes() []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, msg.Id)
	return b
}
