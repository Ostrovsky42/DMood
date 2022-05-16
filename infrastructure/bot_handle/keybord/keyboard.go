package keybord

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type KBoard struct {
	CommandKeyboard tgbotapi.ReplyKeyboardMarkup
	RatingKeyboard  tgbotapi.InlineKeyboardMarkup
}

func NewKeyboard()KBoard{
	// CommandKeyboard := tgbotapi.NewInlineKeyboardMarkup(
	// 	tgbotapi.NewInlineKeyboardRow(
	// 		tgbotapi.NewInlineKeyboardButtonData("Set mood", "set_rating"),
	// 		tgbotapi.NewInlineKeyboardButtonData("Get mood", "get_rating"),
	// 	),
	// 	tgbotapi.NewInlineKeyboardRow(
	// 		tgbotapi.NewInlineKeyboardButtonData("Change mood", "change_rating"),
	// 	),tgbotapi.NewInlineKeyboardRow(
	// 		tgbotapi.NewInlineKeyboardButtonData("Help", "help"),
	// 		tgbotapi.NewInlineKeyboardButtonData("⚡", "fuck"),
	// 		))

	k:=tgbotapi.NewReplyKeyboard(
					[]tgbotapi.KeyboardButton{tgbotapi.KeyboardButton{Text: "/help"},tgbotapi.KeyboardButton{Text: "/fuck"}},
					[]tgbotapi.KeyboardButton{tgbotapi.KeyboardButton{Text: ""}},
					[]tgbotapi.KeyboardButton{tgbotapi.KeyboardButton{Text: "/set_rating"},tgbotapi.KeyboardButton{Text: "/get_rating"}},
					[]tgbotapi.KeyboardButton{tgbotapi.KeyboardButton{Text: "/change_rating"},},
					[]tgbotapi.KeyboardButton{tgbotapi.KeyboardButton{Text: "/fuck"},tgbotapi.KeyboardButton{Text: "/fuck"}},
				)
	k.OneTimeKeyboard=true
	CommandKeyboard:=tgbotapi.ReplyKeyboardMarkup(k)

	RatingKeyboard:=tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("☹", "1"),
			tgbotapi.NewInlineKeyboardButtonData("😕 ", "2"),
			tgbotapi.NewInlineKeyboardButtonData("😐", "3"),
			tgbotapi.NewInlineKeyboardButtonData("🙂", "4"),
			tgbotapi.NewInlineKeyboardButtonData("😊", "5"),
		))
	return KBoard{
		CommandKeyboard: CommandKeyboard,
		RatingKeyboard: RatingKeyboard,
	}
}