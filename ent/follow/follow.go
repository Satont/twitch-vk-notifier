// Code generated by ent, DO NOT EDIT.

package follow

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the follow type in the database.
	Label = "follow"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldChannelID holds the string denoting the channel_id field in the database.
	FieldChannelID = "channel_id"
	// FieldChatID holds the string denoting the chat_id field in the database.
	FieldChatID = "chat_id"
	// EdgeChannel holds the string denoting the channel edge name in mutations.
	EdgeChannel = "channel"
	// EdgeChat holds the string denoting the chat edge name in mutations.
	EdgeChat = "chat"
	// Table holds the table name of the follow in the database.
	Table = "follows"
	// ChannelTable is the table that holds the channel relation/edge.
	ChannelTable = "follows"
	// ChannelInverseTable is the table name for the Channel entity.
	// It exists in this package in order to avoid circular dependency with the "channel" package.
	ChannelInverseTable = "channels"
	// ChannelColumn is the table column denoting the channel relation/edge.
	ChannelColumn = "channel_id"
	// ChatTable is the table that holds the chat relation/edge.
	ChatTable = "follows"
	// ChatInverseTable is the table name for the Chat entity.
	// It exists in this package in order to avoid circular dependency with the "chat" package.
	ChatInverseTable = "chats"
	// ChatColumn is the table column denoting the chat relation/edge.
	ChatColumn = "chat_id"
)

// Columns holds all SQL columns for follow fields.
var Columns = []string{
	FieldID,
	FieldChannelID,
	FieldChatID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
