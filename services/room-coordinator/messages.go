package messages

import "encoding/json"

type Message struct {
	Type    string          `json:"type"`
	RoomID  string          `json:"roomId,omitempty"`
	UserID  string          `json:"userId,omitempty"`
	Payload json.RawMessage `json:"payload,omitempty"`
}

type CursorUpdatePayload struct {
	Line   int `json:"line"`
	Column int `json:"column"`
}
