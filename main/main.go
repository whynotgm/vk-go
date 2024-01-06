package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"vk-go/vk_utils"
)

var config = &vk_utils.Config{
	ApiUrl:  "https://api.vk.com/method/",
	Version: "5.223",
	GroupId: "217369918",
}

func main() {
	// load .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Can't load .env file\t:(")
	}

	// getting dotenv variables
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("can't get token\t:(")
	}

	// Initialize bot
	VK := vk_utils.NewVKBot(config, token)
	updates := make(chan *vk_utils.LPResponse)
	go func() {
		err := VK.LongPoll(updates)
		if err != nil {
			log.Fatal(err)
		}
	}()

	for v := range updates {
		go VK.HandleUpdates(v.Updates)
	}

}
