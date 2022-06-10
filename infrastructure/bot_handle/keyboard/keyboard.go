package keyboard

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type KBoard struct {
	CommandKeyboard tgbotapi.ReplyKeyboardMarkup
	RatingKeyboard  tgbotapi.InlineKeyboardMarkup
	TimeKeyboard    tgbotapi.InlineKeyboardMarkup
}

func NewKeyboard()KBoard{

	CommandKeyboard :=tgbotapi.NewReplyKeyboard(
					[]tgbotapi.KeyboardButton{{Text: "/set_rating"},{Text: "/get_rating"}},
					[]tgbotapi.KeyboardButton{{Text: "/help"},{Text: "/change_rating"}},
					[]tgbotapi.KeyboardButton{{Text: ""}},
				)
	CommandKeyboard.OneTimeKeyboard=true


	RatingKeyboard:=tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("☹", "1"),
			tgbotapi.NewInlineKeyboardButtonData("😕", "2"),
			tgbotapi.NewInlineKeyboardButtonData("😐", "3"),
			tgbotapi.NewInlineKeyboardButtonData("🙂", "4"),
			tgbotapi.NewInlineKeyboardButtonData("😊", "5"),
		))
	TimeKeyboard:=tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("1:00", "1"),
			tgbotapi.NewInlineKeyboardButtonData("2:00", "2"),
			tgbotapi.NewInlineKeyboardButtonData("3:00", "3"),
			tgbotapi.NewInlineKeyboardButtonData("4:00", "4"),
			tgbotapi.NewInlineKeyboardButtonData("5:00", "5"),
			tgbotapi.NewInlineKeyboardButtonData("6:00", "6"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("7:00", "7"),
			tgbotapi.NewInlineKeyboardButtonData("8:00", "8"),
			tgbotapi.NewInlineKeyboardButtonData("9:00", "9"),
			tgbotapi.NewInlineKeyboardButtonData("10:00", "10"),
			tgbotapi.NewInlineKeyboardButtonData("11:00", "11"),
			tgbotapi.NewInlineKeyboardButtonData("12:00", "12"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("13:00", "13"),
			tgbotapi.NewInlineKeyboardButtonData("14:00", "14"),
			tgbotapi.NewInlineKeyboardButtonData("15:00", "15"),
			tgbotapi.NewInlineKeyboardButtonData("16:00", "16"),
			tgbotapi.NewInlineKeyboardButtonData("17:00", "17"),
			tgbotapi.NewInlineKeyboardButtonData("18:00", "18"),
		),
			tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("19:00", "19"),
			tgbotapi.NewInlineKeyboardButtonData("20:00", "20"),
			tgbotapi.NewInlineKeyboardButtonData("21:00", "21"),
			tgbotapi.NewInlineKeyboardButtonData("22:00", "22"),
			tgbotapi.NewInlineKeyboardButtonData("23:00", "23"),
			tgbotapi.NewInlineKeyboardButtonData("00:00", "0"),
		))
	return KBoard{
		CommandKeyboard: CommandKeyboard,
		RatingKeyboard: RatingKeyboard,
		TimeKeyboard: TimeKeyboard,
	}
}