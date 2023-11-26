package usecase

import (
	"context"
	"strconv"
	"time"

	"github.com/Markuysa/courceWorkBackendDev/config"
	"github.com/Markuysa/courceWorkBackendDev/internal/telegram/repository"
	"github.com/Markuysa/courceWorkBackendDev/pkg/telegram"
	"github.com/Markuysa/courceWorkBackendDev/utils/oteltrace"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/gofiber/fiber/v2/log"
)

type Telegram struct {
	cfg  config.Config
	bot  *tgbotapi.BotAPI
	repo repository.Repository
}

func New(
	cfg config.Config,
	repo repository.Repository,
) Usecase {
	// tak luche ne delat
	bot, err := tgbotapi.NewBotAPI(cfg.Telegram.Token)
	if err != nil {
		log.Fatal(err)
	}

	return &Telegram{
		cfg:  cfg,
		bot:  bot,
		repo: repo,
	}
}

func (t *Telegram) ListenMessages(ctx context.Context) (err error) {
	ctx, span := oteltrace.NewSpan(ctx, "ListenMessages")
	defer span.End()

	updatesChan, err := t.bot.GetUpdatesChan(tgbotapi.UpdateConfig{})
	if err != nil {
		return err
	}

	for update := range updatesChan {
		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case telegram.StartCommand:
				{
					clientID := update.Message.CommandArguments()

					id, err := strconv.Atoi(clientID)
					if err != nil {
						log.Error(err)
					}

					err = t.handleStart(ctx, id, update.Message.From.ID)
					if err != nil {
						log.Error(err)
					}
				}
			}
		}
	}

	return nil
}

func (t *Telegram) handleStart(ctx context.Context, clientID int, chatID int) (err error) {
	ctx, span := oteltrace.NewSpan(ctx, "handleStart")
	defer span.End()

	message := telegram.StartMessage

	err = t.repo.LinkTelegram(ctx, clientID, chatID)
	if err != nil {
		message = telegram.FailStartMessage
	}

	msg := tgbotapi.NewMessage(int64(chatID), message)
	_, err = t.bot.Send(msg)

	return err
}

func (t *Telegram) SendTaskNotifications(ctx context.Context) (err error) {
	ctx, span := oteltrace.NewSpan(ctx, "SendTaskNotifications")
	defer span.End()

	messages, err := t.repo.GetNotificationMessages(ctx)
	if err != nil {
		return err
	}

	for chat, message := range messages {
		message := message
		chat := chat

		go func() {
			text := telegram.FormatTaskTgItems(message)
			msg := tgbotapi.NewMessage(int64(chat), text)

			_, err = t.bot.Send(msg)
			if err != nil {
				log.Error(err)
			}
		}()
	}

	return nil
}

func (t *Telegram) StartWorker(ctx context.Context) {
	ctx, span := oteltrace.NewSpan(ctx, "StartWorker")
	defer span.End()

	location, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		log.Fatal(err)
	}
	schedule := time.Date(
		time.Now().Year(),
		time.Now().Month(),
		time.Now().Day(),
		9,
		0,
		0,
		0,
		location)

	for {
		now := time.Now()

		duration := schedule.Sub(now)
		if duration < 0 {
			schedule = schedule.Add(24 * time.Hour)
			duration = schedule.Sub(now)
		}

		timer := time.NewTimer(duration)
		<-timer.C

		err := t.SendTaskNotifications(ctx)
		if err != nil {
			log.Error(err)
		}
	}

}
