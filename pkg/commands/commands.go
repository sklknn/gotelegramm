package commands

import (
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func Start(bot *gotgbot.Bot, context *ext.Context) error {
	_, err := context.EffectiveMessage.Reply(bot, fmt.Sprintf("Привет, Я %s, \nЯ простой бот чтобы продемонстрировать как работает gotgbotapi, а еще попрактиковаться в программировании на Go", bot.User.Username), &gotgbot.SendMessageOpts{
		ParseMode: "HTML",
	})
	if err != nil {
		return fmt.Errorf("Не удалось отправить сообщение start: %w", err)
	}
	return nil
}

func Help(bot *gotgbot.Bot, context *ext.Context) error {
	_, err := context.EffectiveMessage.Reply(bot, fmt.Sprintf("Я умею заказывать пиццу /order, веселить людей, а еще хвалить моего создателя! Он превосходный начинающий программист и системый администратор"), &gotgbot.SendMessageOpts{
		ParseMode: "HTML",
	})
	if err != nil {
		return fmt.Errorf("Не удалось отправить сообщение help: %w", err)
	}
	return nil
}
