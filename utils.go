package main

import (
	"fmt"
	"strconv"
)

const basicURL = "https://api.telegram.org/bot"

func constructBotURL(apiKey string) string {
	return fmt.Sprintf("%s%s", basicURL, apiKey)
}

func getUpdatesURL(constructedBotURL string, updateOffset int) string {
	return fmt.Sprintf("%s%s%s%s", constructedBotURL, "/getUpdates", "?offset=", strconv.Itoa(updateOffset))
}

func sendMessageURL(constructedBotURL string) string {
	return fmt.Sprintf("%s%s", constructedBotURL, "/sendMessage")
}
