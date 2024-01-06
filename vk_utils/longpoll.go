package vk_utils

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/tidwall/pretty"
	"net/url"
)

type LPResponse struct {
	Ts      string   `json:"ts,omitempty"`
	Updates []Update `json:"updates"`
}

type Update struct {
	Type      string `json:"type,omitempty"`
	EventId   string `json:"event_id,omitempty"`
	V         string `json:"v,omitempty"`
	Object    `json:"object"`
	rawObject string
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

func setRawObject(json string, resp *LPResponse) {
	res := gjson.Get(json, "updates").Array()
	for i, re := range res {
		resp.Updates[i].rawObject = url.QueryEscape(re.Raw)
	}

}

func (vk *VKBot) GetUpdates() (resp *LPResponse, err error) {
	addr := fmt.Sprintf("%s?act=a_check&key=%s&ts=%s&wait=%d", vk.lpServer, vk.key, vk.ts, 25)

	body, err := HTTPGetBody(addr)
	if err != nil {
		return
	}

	resp = &LPResponse{}
	err = json.Unmarshal(body, resp)
	vk.ts = resp.Ts

	setRawObject(string(
		pretty.Pretty(body)), resp)

	return
}

func (vk *VKBot) LongPoll(upd chan *LPResponse) error {
	for {
		resp, err := vk.GetUpdates()
		if err != nil {
			return err
		}
		upd <- resp
	}
}
