package message_sender

import (
	"context"
	"fmt"
	"github.com/satont/twitch-notifier/internal/services/db/db_models"
	"github.com/satont/twitch-notifier/internal/test_utils"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestMessageSender_SendMessage(t *testing.T) {
	t.Parallel()

	chat := &db_models.Chat{
		ChatID:  "123",
		Service: db_models.ChatServiceTelegram,
	}

	table := []struct {
		name         string
		chat         *db_models.Chat
		opts         *MessageOpts
		createServer func(*testing.T) *httptest.Server
	}{
		{
			name: "should call send message method",
			chat: chat,
			opts: &MessageOpts{
				Text: "test",
			},
			createServer: func(t *testing.T) *httptest.Server {
				return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					body, err := io.ReadAll(r.Body)
					assert.NoError(t, err)
					query, err := url.ParseQuery(string(body))
					assert.NoError(t, err)

					assert.Equal(t, "test", query.Get("text"))
					assert.Equal(t, "123", query.Get("chat_id"))

					assert.Equal(t, http.MethodPost, r.Method)
					assert.Equal(
						t,
						fmt.Sprintf("/bot%s/sendMessage", test_utils.TelegramClientToken),
						r.URL.Path,
					)

					w.WriteHeader(http.StatusOK)
					_, _ = w.Write([]byte(test_utils.TelegramOkResponse))
				}))
			},
		},
		{
			name: "should call send photo method",
			chat: chat,
			opts: &MessageOpts{
				Text:     "test",
				ImageURL: "https://example.com/image.jpg",
			},
			createServer: func(t *testing.T) *httptest.Server {
				return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					body, err := io.ReadAll(r.Body)
					assert.NoError(t, err)
					query, err := url.ParseQuery(string(body))
					assert.NoError(t, err)

					assert.Equal(t, "test", query.Get("caption"))
					assert.Equal(t, "https://example.com/image.jpg", query.Get("photo"))
					assert.Equal(t, "123", query.Get("chat_id"))

					assert.Equal(t, http.MethodPost, r.Method)
					assert.Equal(
						t,
						fmt.Sprintf("/bot%s/sendPhoto", test_utils.TelegramClientToken),
						r.URL.Path,
					)

					w.WriteHeader(http.StatusOK)
					_, _ = w.Write([]byte(test_utils.TelegramOkResponse))
				}))
			},
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			server := tt.createServer(t)
			tgClient := test_utils.NewTelegramClient(server)
			sender := NewMessageSender(tgClient)

			err := sender.SendMessage(context.Background(), tt.chat, tt.opts)
			assert.NoError(t, err)
		})
	}
}
