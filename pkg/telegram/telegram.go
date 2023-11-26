package telegram

import (
	"fmt"

	"github.com/Markuysa/courceWorkBackendDev/internal/models"
)

const (
	StartCommand = "start"
)

const (
	TgLink           = "https://t.me/Tasks_MarkuysaBot?start=%v"
	StartMessage     = "Ваш телеграм успешно привязан!"
	FailStartMessage = "Ваш телеграм был ранее привязан!"
)

func FormatTaskTgItems(items []models.TaskTgItem) string {
	var formattedText string

	if len(items) == 0 {
		formattedText = "Нет задач."
	} else {
		formattedText = "Список задач:\n"
		for _, item := range items {
			if item.Deadline.Valid {
				formattedText += fmt.Sprintf("Описание: %s\nДедлайн: %s\n\n", item.Description, item.Deadline.Time.Format("2006-01-02 15:04:05"))
			}

			formattedText += fmt.Sprintf("Описание: %s\nДедлайн: %s\n\n", item.Description, "не установлен")
		}
	}

	formattedText += "Продуктивного вам рабочего дня!"

	return formattedText
}
