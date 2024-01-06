package vk_utils

// if returns anything - bot sends this as answer
func (vk *VKBot) commandHandler(event Update) (answer string) {
	switch text := event.Text; text {
	case "ping", "пинг":
		return "понг"
	case "event", "ивент":
		return event.rawObject
	}
	return

}
