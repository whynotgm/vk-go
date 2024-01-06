package vk_utils

import (
	"fmt"
)

func (vk *VKBot) HandleUpdates(updates []Update) {
	for _, update := range updates {
		if &update == nil {
			continue
		}
		switch update.Type {
		case "message_new":
			err := vk.messageNewHandler(update)
			if err != nil {
				fmt.Println("Error:", err)
			}
		default:
			fmt.Println("Unknown event type!")
		}

	}
}

func (vk *VKBot) messageNewHandler(event Update) (err error) {
	answer := vk.commandHandler(event)
	if answer == "" {
		return
	}

	_, err = vk.sendMessage(event, answer)
	return
}
