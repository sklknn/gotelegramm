package main

import (
	"log"
	"os"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
    "github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/conversation"
	"github.com/sklknn/gotelegramm/pkg/commands"
)

func main() {

	token, err := os.ReadFile("./TOKEN.txt")
	if err != nil {
		log.Println("Не удалось получить токен: " + err.Error())
	}

	log.Println("Токен успешно получен")

	bot, err := gotgbot.NewBot(string(token)[:len(token)-1], nil)
	if err != nil {
		panic("Не удалось создать нового бота: " + err.Error())
	}

	dispatcher := ext.NewDispatcher(&ext.DispatcherOpts{
		// if error is returned by a handler - continue
		Error: func(bot *gotgbot.Bot, context *ext.Context, err error) ext.DispatcherAction {
			log.Println("Ошибка во время обновления хэндлера: ", err.Error())
			return ext.DispatcherActionNoop
		},
		MaxRoutines: ext.DefaultMaxRoutines,
	})
	updater := ext.NewUpdater(dispatcher, nil)
	//introduce
	dispatcher.AddHandler(handlers.NewCommand("start", commands.Start))
    //user help
    dispatcher.AddHandler(handlers.NewCommand("help", commands.Help))
    //conversation 
    dispatcher.AddHandler(handlers.NewConversation(
        []ext.Handler{handlers.NewCommand("pizza", commands.Conversation)},
        map[string][]ext.Handler{
            NAME : {handlers.NewMessage(noCommands, name)},
            PIZZA : {handlers.NewMessage(noCommands, pizza)},
            ADDRESS : {handlers.NewMessage(noCommands, address)},
        },
        &handlers.ConversationOpts{
            Exits : []ext.Handler{handlers.NewCommand("cancel_order", commands.EndConversation)},
            StateStorage : conversation.NewInMemoryStorage(conversation.KeyStrategySenderAndChat),
    AllowReEntry : true,
        },
        ))

	err = updater.StartPolling(bot, &ext.PollingOpts{
		DropPendingUpdates: true,
		GetUpdatesOpts: &gotgbot.GetUpdatesOpts{
			Timeout: 10,
			RequestOpts: &gotgbot.RequestOpts{
				Timeout: time.Second * 15,
			},
		},
	})
	if err != nil {
		panic("Не удалось начать пуллинг: " + err.Error())
	}
	log.Printf("%s был запущен...\n", bot.User.Username)

	updater.Idle()
}
