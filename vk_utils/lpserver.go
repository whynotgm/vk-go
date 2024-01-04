package vk_utils

import "encoding/json"

type GetLpResponse struct {
	Response struct {
		Key    string `json:"key,omitempty"`
		Server string `json:"server,omitempty"`
		Ts     string `json:"ts,omitempty"`
	} `json:"response,omitempty"`
}

func (vk *VKBot) GetLPServer() (resp *GetLpResponse, err error) {
	response, err := vk.RequestMethod("groups.getLongPollServer", "group_id="+vk.groupId)
	if err != nil {
		return
	}

	// response schema
	resp = &GetLpResponse{}

	err = json.Unmarshal(response, resp)

	return
}

func (vk *VKBot) SetupLP(serv, key, ts string) {
	vk.lpServer = serv
	vk.key = key
	vk.ts = ts
}
