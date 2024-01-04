package vk_utils

import "fmt"

func (vk *VKBot) HandleUpdates(updates []Update) {
	for _, update := range updates {
		if &update == nil {
			continue
		}
		switch update.Type {
		case "message_new":
			messageNewHandler(update)

		}

	}
}

func messageNewHandler(event Update) {
	fmt.Println(event.Text)
}
