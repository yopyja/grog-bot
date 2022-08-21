package listBot

import (
	"encoding/json"
	"github.com/bwmarrin/discordgo"
	"grog-bot/config"
	"grog-bot/list"
	"grog-bot/list/grab"
	"grog-bot/list/test"
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

	BotID = user.ID
	goBot.AddHandler(messageHandler)
	err = goBot.Open()
	if err != nil {
		return
	}
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == BotID {
		return
	}
	// Only commands can be issued in #general
	if m.ChannelID == config.General {
		// Checks if message has correct cmd.prefix
		if strings.HasPrefix(m.Content, config.Prefix) {
			// Help Command
			if strings.HasPrefix(m.Content, config.Prefix+"help") {
				_, _ = s.ChannelMessageSend(config.General, "```md\n"+config.Prefix+"list            - List current watchlist\n"+config.Prefix+"add <sku>       - Adds <sku> to watchlist\n"+config.Prefix+"remove <sku>    - Removes <sku> from watchlist\n"+config.Prefix+"sub <sku>       - Subscribes to <sku> pings\n"+config.Prefix+"unsub <sku>     - Unsubscribes from <sku> pings```")

				// List Command
			} else if strings.HasPrefix(m.Content, config.Prefix+"list") {
				_, _ = s.ChannelMessageSend(config.General, grab.GrabList())

				// Add Command
			} else if strings.HasPrefix(m.Content, config.Prefix+"add") {
				// If message has 6 numbers exactly
				if len(listTest.IsValidSKU(m.Content)) == 6 {
					_, _ = s.ChannelMessageSend(config.General, list.ListManager(0, m.Content))

					// SKU is Invalid
				} else {
					_, _ = s.ChannelMessageSend(config.General, "This SKU is invalid")
				}

				// Remove Command
			} else if strings.HasPrefix(m.Content, config.Prefix+"remove") {
				// If message has 6 numbers exactly
				if len(listTest.IsValidSKU(m.Content)) == 6 {
					_, _ = s.ChannelMessageSend(config.General, list.ListManager(1, m.Content))

					// SKU is Invalid
				} else {
					_, _ = s.ChannelMessageSend(config.General, "This SKU is invalid")
				}

				// Sub Command
			} else if strings.HasPrefix(m.Content, config.Prefix+"sub") {
				// if isdupe
				//if (len(listTest.IsValidSKU(m.Content)) == 6) && stockTest.IsRealItem(listTest.IsValidSKU(m.Content)) {
				//
				//}

				_, _ = s.ChannelMessageSend(config.General, "This cmd is under construction")

				// Unsub Command
			} else if strings.HasPrefix(m.Content, config.Prefix+"unsub") {
				_, _ = s.ChannelMessageSend(config.General, "This cmd is under construction")

				//	Force stock update
			} else if strings.HasPrefix(m.Content, config.Prefix+"force") && m.Author.ID == config.OwnerID {
				stock.StockRunner()

				messageFile, err := os.Open("./json/message.json")
				if err != nil {
					log.Fatal(err)
					return
				}
				byteValueDrink, _ := io.ReadAll(messageFile)
				var messages stockMessage
				json.Unmarshal(byteValueDrink, &messages)
				if len(messages.Message) > 0 {
					_, _ = s.ChannelMessageSend(config.General, "<@&" + config.RoleID + ">")
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
						_, _ = s.ChannelMessageSendEmbed(config.General, &messageEmbed)
					}
				}

				// Invalid Command Dingus
			} else {
				_, _ = s.ChannelMessageSend(config.General, "Invalid command. Type "+config.Prefix+"help for list of commands.")
			}
		}
	}
}
