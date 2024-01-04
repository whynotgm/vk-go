package main

import (
	"github.com/joho/godotenv"
	"gotests/vk_utils"
	"log"
	"os"
)

var config = &vk_utils.Config{
	ApiUrl: "https://api.vk.com/method/",
}

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Can't load .env file\t:(")
	}

	// getting dotenv variables
	token, groupId := os.Getenv("BOT_TOKEN"), os.Getenv("GROUP_ID")
	if token == "" || groupId == "" {
		log.Fatal("can't get variables\t:(")
	}

	// Initialize bot
	VK := vk_utils.NewVKBot(config, token, groupId)
	updates := make(chan *vk_utils.LPResponse)
	go VK.LongPoll(updates)
	for v := range updates {
		go VK.HandleUpdates(v.Updates)
	}

}
