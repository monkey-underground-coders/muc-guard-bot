package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Check dotenv is present in the project
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Fatal: Dotenv file in directory root not found.\n" +
			"Make sure that .env file is presented in the root and have TG_KEY property filled")
		return
	}

	telegramKey, telegramKeyExist := os.LookupEnv("TG_KEY")

	if !telegramKeyExist || len(telegramKey) <= 0 {
		fmt.Printf("Fatal: Dotenv file is presented but TG_KEY is not.\n" +
			"Example: TG_KEY=bot123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11")
		return
	}

	subscribeListeners(telegramKey)
}

func subscribeListeners(apiKey string) {
	fmt.Printf(apiKey)
}
