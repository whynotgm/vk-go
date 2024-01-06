package vk_utils

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strings"
)

//func (vk *VKBot) PostRequestMethod(method string, params map[string]any) (body []byte, err error) {
//	params["access_token"] = vk.token
//	params["v"] = vk.Version
//	data, err := json.Marshal(params)
//	if err != nil {
//		return
//	}
//
//	response, err := http.Post(vk.ApiUrl+method, "application/x-www-form-encoded", bytes.NewBuffer(data))
//	if err != nil {
//		return
//	}
//
//	body, err = io.ReadAll(response.Body)
//	defer response.Body.Close()
//	return
//}

func (vk *VKBot) RequestMethod(method string, params ...string) (body []byte, err error) {
	address := fmt.Sprintf("%s%s?access_token=%s&v=%s", vk.ApiUrl, method, vk.token, vk.Version)

	// params like: "key=value" TODO: Maybe POST request?
	if len(params) != 0 {
		address += "&" + strings.Join(params, "&")
	}

	return HTTPGetBody(address)

}

func HTTPGetBody(link string) (body []byte, err error) {
	response, err := http.Get(link)
	if err != nil {
		return
	}
	defer response.Body.Close()
	return io.ReadAll(response.Body)
}

type VKBot struct {
	Config
	token string
}

type Config struct {
	ApiUrl            string
	lpServer, key, ts string
	Version           string
	GroupId           string
}

func NewVKBot(config *Config, token string) *VKBot {
	VK := &VKBot{
		Config: *config,
		token:  token,
	}

	resp, err := VK.GetLPServer()
	if err != nil {
		log.Fatal(err)
	}

	// configure LP
	VK.SetupLP(resp.Response.Server, resp.Response.Key, resp.Response.Ts)
	return VK
}

func (vk *VKBot) sendMessage(event Update, text string) (resp []byte, err error) {
	peerId := fmt.Sprintf("peer_id=%v", event.PeerId)
	text = fmt.Sprintf("message=%s", text)
	randID := fmt.Sprintf("random_id=%v", rand.Uint32())

	resp, err = vk.RequestMethod("messages.send", peerId, text, randID)
	return
}
