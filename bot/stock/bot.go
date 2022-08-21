package stockBot

import (
	"encoding/json"
	"github.com/bwmarrin/discordgo"
	"grog-bot/config"
	"grog-bot/stock"
	"io"
	"log"
	"os"
	"strings"
)

var BotID string
var goBot *discordgo.Session

type stockMessage struct {
	Message []string `json:"message"`
}

func Run() {
	// create bot session
	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		log.Fatal(err)
		return
	}
	// make the bot a user
	user, err := goBot.User("@me")
	if err != nil {
		log.Fatal(err)
		return
	}

	stock.StockRunner()

	messageFile, err := os.Open("./json/message.json")

	// Load old stock
	byteValueDrink, _ := io.ReadAll(messageFile)
	var messages stockMessage
	json.Unmarshal(byteValueDrink, &messages)

	for i := 0; i < len(messages.Message); i++ {
		c := strings.Split(messages.Message[i], "||")
		var messageEmbed = discordgo.MessageEmbed{
			Title:       c[0],
			Description: c[1],
			Color:       0x6aa84f,
			URL:         c[2],
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: "https://webapps2.abc.utah.gov/ProdApps/ProductLocatorCore/images/DABS_Logo_StateOnly_WH_TB.png",
			},
		}
		goBot.ChannelMessageSendEmbed(config.General, &messageEmbed)
	}

	BotID = user.ID

	err = goBot.Open()
	if err != nil {
		return
	}
}
