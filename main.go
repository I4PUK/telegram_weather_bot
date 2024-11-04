package main

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"telegram_weather_bot/models"
)

type Config struct {
	TelegramBotToken string
	ApiKey           string
}

func main() {
	file, err := os.Open("configs/config.json")

	defer file.Close()

	decoder := json.NewDecoder(file)
	configuration := Config{}
	err = decoder.Decode(&configuration)

	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Successfully got a configuration!")

	bot, err := tgbotapi.NewBotAPI(configuration.TelegramBotToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			log.Printf("Got message from [%s] with text: %s", update.Message.From.UserName, update.Message.Text)

			//var msg tgbotapi.MessageConfig = procedureCommand(update.Message.Command(), update.Message.Chat.ID)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}

	}
}

func procedureCommand(command string, chatID int64) tgbotapi.MessageConfig {
	switch command {
	case "start":
		return tgbotapi.NewMessage(chatID, "Welcome! I am your bot.")
	case "help":
		return tgbotapi.NewMessage(chatID, "I can help you with the following commands:\n/start - Start the bot\n/help - Display this help message")
	default:
		return tgbotapi.NewMessage(chatID, "I don't know that command")
	}
}

func callForecastApi(lan, lon float32) string {
	strLan := strconv.FormatFloat(float64(lan), 'f', 2, 32)
	strLon := strconv.FormatFloat(float64(lon), 'f', 2, 32)

	url := "https://api.openweathermap.org/data/2.5/forecast/hourly?" + "lan=" + strLan + "lon=" + strLon + "&APPID="
	//+ ApiKey + "&lang=en"
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic("Weather request was not OK: " + resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	response := parseResponseForecast(body)

	return response
}

func parseResponseForecast(body []byte) string {
	var s = new(models.ForecastAPIResponse)
	err := json.Unmarshal(body, &s)
	if err != nil {
		panic(err)
	}
	var message = ""
	for _, value := range s.List[0:3] {
		var partialMessage = " ⌚️ " + value.DtTxt + " " + value.Weather[0].Description + ". 🌡  Temperature (ºC) " + fmt.Sprintf("%.0f", toCelsius(value.Main.Temp)) + ". Feels like " + fmt.Sprintf("%.0f", toCelsius(value.Main.FeelsLike)) + " (L " + fmt.Sprintf("%.0f", toCelsius(value.Main.TempMin)) + " - H " + fmt.Sprintf("%.0f", toCelsius(value.Main.TempMax)) + "), " + fmt.Sprintf("%2d", value.Main.Humidity) + " humidity."
		message += "\n" + partialMessage
	}
	return message
}

func toCelsius(kelvin float64) float64 {
	return math.Round(kelvin - 273.15)
}
