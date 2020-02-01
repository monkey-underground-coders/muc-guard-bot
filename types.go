package main

// TGUpdate represents an update from tg chat
type TGUpdate struct {
	ID  int       `json:"update_id"`
	Msg TGMessage `json:"message"`
}

// JSONResponse represents result of update request in JSON format
type JSONResponse struct {
	Result []TGUpdate
}

// TGBotMessage represents result of bot message request
type TGBotMessage struct {
	ChatID    int    `json:"chat_id"`
	Text      string `json:"text"`
	MessageID int    `json:"reply_to_message_id"`
}

// TGMessageUser represents user model from message
type TGMessageUser struct {
	ID        int    `json:"id"`
	IsBot     bool   `json:"is_bot"`
	FirstName string `json:"first_name"`
}

// TGMessage represents a message from tg chat
type TGMessage struct {
	ID       int           `json:"message_id"`
	Chat     TGChat        `json:"chat"`
	Text     string        `json:"text"`
	UserFrom TGMessageUser `json:"from"`
}

// TGChat represents id of tg chat
type TGChat struct {
	ChatID int `json:"id"`
}
