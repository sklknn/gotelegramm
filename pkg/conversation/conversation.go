package conversation

import (
	"fmt"
	"html"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/conversation"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/message"
)

//Просто конастанты для таблицы действий, значение не оказывает влияния
const (
	NAME   = "Имя"
	PIZZA  = "Пицца"
	ADDRESS = "Адрес"
    CONFIRMATION = "Подтверждение"
)

var (
    orderName string
    orderPizza string
    orderAdress string
)

func OrderPizza(dispatcher *ext.Dispatcher) {
	dispatcher.AddHandler(handlers.NewConversation(
		[]ext.Handler{handlers.NewCommand("pizza", pizzaConversation)},
		map[string][]ext.Handler{
			NAME:   {handlers.NewMessage(noCommands, name)},
			PIZZA:  {handlers.NewMessage(noCommands, pizza)},
			ADDRESS: {handlers.NewMessage(noCommands, address)},
            CONFIRMATION: {handlers.NewMessage(noCommands, confirmOrder)},
		},
		&handlers.ConversationOpts{
			Exits:        []ext.Handler{handlers.NewCommand("cancel_order", endConversation)},
			StateStorage: conversation.NewInMemoryStorage(conversation.KeyStrategySenderAndChat),
			AllowReEntry: true,
		},
	))
}

func noCommands(msg *gotgbot.Message) bool {
	return message.Text(msg) && !message.Command(msg)
}

func pizzaConversation(bot *gotgbot.Bot, context *ext.Context) error {
	_, err := context.EffectiveMessage.Reply(bot, fmt.Sprintf("Хочешь заказать пиццы? Хорошо! \n Тогда давай уточним твои данные \n Как тебя зовут?"), &gotgbot.SendMessageOpts{
		ParseMode: "HTML",
	})
	if err != nil {
		return fmt.Errorf("Не удалось отрпавить сообщение 'Начало заказа': %w", err)
	}
	return handlers.NextConversationState(NAME)
}

func endConversation(bot *gotgbot.Bot, context *ext.Context) error {
	_, err := context.EffectiveMessage.Reply(bot, "Хорошо, заказ отменен!", &gotgbot.SendMessageOpts{
		ParseMode: "HTML",
	})
	if err != nil {
		return fmt.Errorf("Не удалось отправить сообщение 'Отмена заказа': %w", err)
	}
	return handlers.EndConversation()
}

func name(bot *gotgbot.Bot, context *ext.Context) error {
	inputName := context.EffectiveMessage.Text
	_, err := context.EffectiveMessage.Reply(bot, fmt.Sprintf("Хорошо %s \n Какую пиццу хочешь заказать?", html.EscapeString(inputName)), &gotgbot.SendMessageOpts{
		ParseMode: "HTML",
	})
	if err != nil {
		return fmt.Errorf("Не удалось отправить сообщение 'Имя' %w", err)
	}
    orderName = inputName
	return handlers.NextConversationState(PIZZA)
}

func pizza(bot *gotgbot.Bot, context *ext.Context) error {
	inputPizza := context.EffectiveMessage.Text
	_, err := context.EffectiveMessage.Reply(bot, fmt.Sprintf("Хорошо, выбрали %s \n Куда её привезти?\n Подскажи адрес доставки", html.EscapeString(inputPizza)), &gotgbot.SendMessageOpts{
		ParseMode: "HTML",
	})
	if err != nil {
		return fmt.Errorf("Не удалось отправить сообщение 'Пицца' %w", err)
	}
    orderPizza = inputPizza
	return handlers.NextConversationState(ADDRESS)
}

func address(bot *gotgbot.Bot, context *ext.Context) error {
    inputAddress := context.EffectiveMessage.Text
    _, err := context.EffectiveMessage.Reply(bot, fmt.Sprintf("Отлично! \n Заказываю на %s", html.EscapeString(inputAddress)), &gotgbot.SendMessageOpts{
    ParseMode : "HTML",
    })
    if err != nil {
        return fmt.Errorf("Не удалось отправить сообщение 'Адресс' %w", err)
    }
    orderAdress = inputAddress
    return handlers.NextConversationState(CONFIRMATION)
}

func confirmOrder(bot *gotgbot.Bot, context *ext.Context) error {
    _, err := context.EffectiveMessage.Reply(bot, fmt.Sprintf("Давай проверим данные\n 1. Заказ на имя %s \n Пицца: %s \n По адресу: %s \n Верно? Да/Нет",html.EscapeString(orderName),html.EscapeString(orderPizza),html.EscapeString(orderAdress)), &gotgbot.SendMessageOpts{
ParseMode : "HTML",})
    if err != nil {
        return fmt.Errorf("Не удалось сверить заказ %w", err)
    }
    return handlers.EndConversation()
} 
