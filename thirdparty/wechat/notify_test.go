package wechat

import (
	"context"
	"testing"
)

func TestNewWeChatNotify(t *testing.T) {
	client := NewWeChatNotify("none")
	_, err := client.Send(context.Background(), WeChatNotifyTextMessage{
		ToUser:  "czasg",
		Content: "hello",
	})
	if err == nil {
		t.Error("invalid wechat")
	}
}
