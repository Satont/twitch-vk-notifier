// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/satont/twitch-notifier/ent/chat"
	"github.com/satont/twitch-notifier/ent/chatsettings"
)

// Chat is the model entity for the Chat schema.
type Chat struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// ChatID holds the value of the "chat_id" field.
	ChatID string `json:"chat_id,omitempty"`
	// Service holds the value of the "service" field.
	Service chat.Service `json:"service,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ChatQuery when eager-loading is set.
	Edges ChatEdges `json:"edges"`
}

// ChatEdges holds the relations/edges for other nodes in the graph.
type ChatEdges struct {
	// Settings holds the value of the settings edge.
	Settings *ChatSettings `json:"settings,omitempty"`
	// Follows holds the value of the follows edge.
	Follows []*Follow `json:"follows,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// SettingsOrErr returns the Settings value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ChatEdges) SettingsOrErr() (*ChatSettings, error) {
	if e.loadedTypes[0] {
		if e.Settings == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: chatsettings.Label}
		}
		return e.Settings, nil
	}
	return nil, &NotLoadedError{edge: "settings"}
}

// FollowsOrErr returns the Follows value or an error if the edge
// was not loaded in eager-loading.
func (e ChatEdges) FollowsOrErr() ([]*Follow, error) {
	if e.loadedTypes[1] {
		return e.Follows, nil
	}
	return nil, &NotLoadedError{edge: "follows"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Chat) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case chat.FieldChatID, chat.FieldService:
			values[i] = new(sql.NullString)
		case chat.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Chat", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Chat fields.
func (c *Chat) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case chat.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				c.ID = *value
			}
		case chat.FieldChatID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field chat_id", values[i])
			} else if value.Valid {
				c.ChatID = value.String
			}
		case chat.FieldService:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field service", values[i])
			} else if value.Valid {
				c.Service = chat.Service(value.String)
			}
		}
	}
	return nil
}

// QuerySettings queries the "settings" edge of the Chat entity.
func (c *Chat) QuerySettings() *ChatSettingsQuery {
	return NewChatClient(c.config).QuerySettings(c)
}

// QueryFollows queries the "follows" edge of the Chat entity.
func (c *Chat) QueryFollows() *FollowQuery {
	return NewChatClient(c.config).QueryFollows(c)
}

// Update returns a builder for updating this Chat.
// Note that you need to call Chat.Unwrap() before calling this method if this Chat
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Chat) Update() *ChatUpdateOne {
	return NewChatClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Chat entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Chat) Unwrap() *Chat {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Chat is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Chat) String() string {
	var builder strings.Builder
	builder.WriteString("Chat(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("chat_id=")
	builder.WriteString(c.ChatID)
	builder.WriteString(", ")
	builder.WriteString("service=")
	builder.WriteString(fmt.Sprintf("%v", c.Service))
	builder.WriteByte(')')
	return builder.String()
}

// Chats is a parsable slice of Chat.
type Chats []*Chat
