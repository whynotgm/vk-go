package vk_utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func (vk *VKBot) RequestMethod(method string, params ...string) (body []byte, err error) {
	url := fmt.Sprintf("%s%s?access_token=%s&v=%s", vk.ApiUrl, method, vk.token, vk.version)

	// append params to url
	if len(params) != 0 {
		url += "&" + strings.Join(params, "&")
	}

	return HTTPGetBody(url)

}

func HTTPGetBody(url string) (body []byte, err error) {
	response, err := http.Get(url)
	if err != nil {
		return
	}
	return io.ReadAll(response.Body)
}

type VKBot struct {
	Config
	token   string
	groupId string
	version string
}

type Config struct {
	ApiUrl            string
	lpServer, key, ts string
}

func NewVKBot(config *Config, token, groupId string) *VKBot {
	VK := &VKBot{
		Config:  *config,
		token:   token,
		groupId: groupId,
		version: "5.223",
	}

	resp, err := VK.GetLPServer()
	if err != nil {
		log.Fatal(err)
	}

	// configure LP
	VK.SetupLP(resp.Response.Server, resp.Response.Key, resp.Response.Ts)
	return VK
}
