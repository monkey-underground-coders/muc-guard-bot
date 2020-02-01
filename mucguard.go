package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	// Check dotenv is presented in the project
	if err := godotenv.Load(); err != nil {
		panic("Dotenv file in directory root not found.\n" +
			"Make sure that .env file is presented in the root and have TG_KEY property filled")
	}

	telegramKey, telegramKeyExist := os.LookupEnv("TG_KEY")

	// Check dotenv TG_KEY property is presented
	if !telegramKeyExist || len(telegramKey) <= 0 {
		panic("Dotenv file is presented but TG_KEY is not.\n" +
			"Example: TG_KEY=bot123456:ABC-DEF1234ghIkl-zyx57W2v1u1	23ew11")
	}

	subscribeListeners(telegramKey)
}

func subscribeListeners(apiKey string) {
	var botURL = constructBotURL(apiKey)

	createUpdateListener(botURL)
}

func createUpdateListener(botURL string) {
	var updateOffset = 0
	for {
		updates, _ := updateListener(botURL, updateOffset)
		for _, update := range updates {
			sendMessage(botURL, update)
			updateOffset = update.ID + 1
		}
		fmt.Println(updates)
	}
}

func handleKeywords(message string) string {
	var keyphrases = []string{
		"пойду гулять",
		"гулять пойду",
		"иду пить",
		"пить иду",
		"пью",
		"пиво",
		"пивас",
		"иду спать",
		"спать иду",
		"пойду посплю",
		"спать хочется",
		"хочется спать",
		"иду бухать",
		"бухать иду",
		"пойду почилю",
		"пойду почиллю",
		"чилю",
		"гулять",
		"спать ща пойду",
		"пойду ща спать",
		"пойду спать",
		"спать пойду"}
	for _, phrase := range keyphrases {
		if strings.Contains(message, phrase) {
			return "Это все конечно хорошо, а кодить ты когда будешь, пес?"
		}
	}
	return ""
}

func updateListener(constructedBotURL string, updateOffset int) ([]TGUpdate, error) {
	// Make a request to getUpdates
	response, responseError := http.Get(getUpdatesURL(constructedBotURL, updateOffset))
	if responseError != nil {
		return nil, responseError
	}

	// Make sure we close the connection after function end
	defer response.Body.Close()

	// Get response body in bytes
	responseBody, responseBodyError := ioutil.ReadAll(response.Body)
	if responseBodyError != nil {
		return nil, responseBodyError
	}

	// Parse response from bytes to JSON format
	var jsonResponse JSONResponse
	jsonResponseError := json.Unmarshal(responseBody, &jsonResponse)
	if jsonResponseError != nil {
		return nil, jsonResponseError
	}

	return jsonResponse.Result, nil
}

func sendMessage(constructedBotURL string, update TGUpdate) error {
	var botMessage TGBotMessage
	botMessage.ChatID = update.Msg.Chat.ChatID
	botMessage.Text = handleKeywords(strings.ToLower(update.Msg.Text))
	botMessage.MessageID = update.Msg.ID

	if len(botMessage.Text) <= 0 {
		return nil
	}

	// Serialize buffer to json format
	buf, bufErr := json.Marshal(botMessage)
	if bufErr != nil {
		return bufErr
	}

	// Send a message
	_, sendMessageError := http.Post(sendMessageURL(constructedBotURL), "application/json", bytes.NewBuffer(buf))
	if sendMessageError != nil {
		return sendMessageError
	}

	return nil
}
