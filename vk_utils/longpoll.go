package vk_utils

import (
	"encoding/json"
	"fmt"
)

type LPResponse struct {
	Ts      string   `json:"ts,omitempty"`
	Updates []Update `json:"updates"`
}
type Update struct {
	Type    string `json:"type,omitempty"`
	EventId string `json:"event_id,omitempty"`
	V       string `json:"v,omitempty"`
	Object  `json:"object"`
}

type Object struct {
	Id      int `json:"id,omitempty"`
	FromId  int `json:"from_id,omitempty"`
	Message `json:"message,omitempty"`
}

type Message struct {
	Id          int          `json:"id,omitempty"`
	Date        int          `json:"date,omitempty"`
	PeerId      int          `json:"peer_id,omitempty"`
	FromId      int          `json:"from_id,omitempty"`
	Text        string       `json:"text,omitempty"`
	Attachments []JSONObject `json:"attachments,omitempty"`
}

type JSONObject map[string]any

type AttachObject struct {
	Id int `json:"id,omitempty"`
}

func (vk *VKBot) GetUpdates() (resp *LPResponse, err error) {
	url := fmt.Sprintf("%s?act=a_check&key=%s&ts=%s&wait=%d", vk.lpServer, vk.key, vk.ts, 25)

	body, err := HTTPGetBody(url)
	if err != nil {
		return
	}

	resp = &LPResponse{}
	err = json.Unmarshal(body, resp)
	vk.ts = resp.Ts

	return
}

func (vk *VKBot) LongPoll(upd chan *LPResponse) {
	for {
		resp, _ := vk.GetUpdates()
		upd <- resp
	}
}
