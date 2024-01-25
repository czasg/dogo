package wechat

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type WeChatNotifyMessage interface {
	Body() ([]byte, error)
}

type WeChatNotifyTextMessage struct {
	ToUser  string `json:"touser"`
	ToParty string `json:"toparty"`
	ToTag   string `json:"totag"`
	MsgType string `json:"msgtype"`
	AgentId int    `json:"agentid"`
	Content string `json:"-"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
	Safe                   int `json:"safe"`
	EnableIDTrans          int `json:"enable_id_trans"`
	EnableDuplicateCheck   int `json:"enable_duplicate_check"`
	DuplicateCheckInterval int `json:"duplicate_check_interval"`
}

func (w WeChatNotifyTextMessage) Body() ([]byte, error) {
	w.MsgType = "text"
	if w.Content != "" {
		w.Text.Content = w.Content
	}
	return json.Marshal(w)
}

type WeChatNotifyResponse struct {
	ErrCode        int    `json:"errcode"`
	ErrMsg         string `json:"errmsg"`
	InvalidUser    string `json:"invaliduser"`
	InvalidParty   string `json:"invalidparty"`
	InvalidTag     string `json:"invalidtag"`
	UnlicensedUser string `json:"unlicenseduser"`
	MsgId          string `json:"msgid"`
	ResponseCode   string `json:"response_code"`
}

func NewWeChatNotify(url string) WeChatNotify {
	return WeChatNotify{url: url}
}

func NewWeChatNotifyWithToken(token string) WeChatNotify {
	return NewWeChatNotify(fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/message/send?key=%s", token))
}

type WeChatNotify struct {
	url string
}

func (w WeChatNotify) Send(ctx context.Context, message WeChatNotifyMessage) (*WeChatNotifyResponse, error) {
	body, err := message.Body()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, w.url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	notifyResponse := WeChatNotifyResponse{}
	err = json.Unmarshal(body, &notifyResponse)
	if err != nil {
		return nil, err
	}
	if notifyResponse.ErrCode != 0 {
		return nil, errors.New(notifyResponse.ErrMsg)
	}
	return &notifyResponse, nil
}
