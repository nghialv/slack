package slack

import "encoding/json"

// AckMessage is used for messages received in reply to other messages
type AckMessage struct {
	ReplyTo   int    `json:"reply_to"`
	Timestamp string `json:"ts"`
	Text      string `json:"text"`
	SlackWSResponse
}

type SlackWSResponse struct {
	Ok    bool          `json:"ok"`
	Error *SlackWSError `json:"error"`
}

type SlackWSError string

func (s SlackWSError) Error() string {
	return string(s)
}

type MessageEvent Message

// SlackEvent is the main wrapper. You will find all the other messages attached
type SlackEvent struct {
	Type string
	Data interface{}
}

// HelloEvent represents the hello event
type HelloEvent struct{}

// PresenceChangeEvent represents the presence change event
type PresenceChangeEvent struct {
	Type     string `json:"type"`
	Presence string `json:"presence"`
	User     string `json:"user"`
}

// UserTypingEvent represents the user typing event
type UserTypingEvent struct {
	Type    string `json:"type"`
	User    string `json:"user"`
	Channel string `json:"channel"`
}

type PrefChangeEvent struct {
	Type  string          `json:"type"`
	Name  string          `json:"name"`
	Value json.RawMessage `json:"value"`
}

// ManualPresenceChangeEvent represents the manual presence change event
type ManualPresenceChangeEvent struct {
	Type     string `json:"type"`
	Presence string `json:"presence"`
}

// UserChangeEvent represents the user change event
type UserChangeEvent struct {
	Type string `json:"type"`
	User User   `json:"user"`
}

// EmojiChangedEvent represents the emoji changed event
type EmojiChangedEvent struct {
	Type           string         `json:"type"`
	EventTimestamp JSONTimeString `json:"event_ts"`
}

// CommandsChangedEvent represents the commands changed event
type CommandsChangedEvent struct {
	Type           string         `json:"type"`
	EventTimestamp JSONTimeString `json:"event_ts"`
}

// EmailDomainChangedEvent represents the email domain changed event
type EmailDomainChangedEvent struct {
	Type           string         `json:"type"`
	EventTimestamp JSONTimeString `json:"event_ts"`
	EmailDomain    string         `json:"email_domain"`
}

// BotAddedEvent represents the bot added event
type BotAddedEvent struct {
	Type string `json:"type"`
	Bot  Bot    `json:"bot"`
}

// BotChangedEvent represents the bot changed event
type BotChangedEvent struct {
	Type string `json:"type"`
	Bot  Bot    `json:"bot"`
}

// AccountsChangedEvent represents the accounts changed event
type AccountsChangedEvent struct {
	Type string `json:"type"`
}
