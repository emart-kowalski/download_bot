package main

import (
	"log"

	tele "gopkg.in/telebot.v4"
)

func register_handlers(bot *tele.Bot) {
	bot.Handle("/start", func(c tele.Context) error {
		return c.Send("Привет! Отправь ссылку на видео tiktok, youtube, instagram и я скачаю его")
	})

	bot.Handle(tele.OnText, func(c tele.Context) error {
		user := c.Chat().Username
		msg_text := c.Message().Text
		log.Print("User " + user + " send text: " + msg_text)

		domain, err := check_domain(msg_text)
		if err != nil {
			log.Print(err)
			return c.Send("Введите корректную ссылку")
		}

		message, service, err := check_correct_service(domain)
		if err != nil {
			log.Print(err)
			return c.Send("Неизвестный сервис. Отправьте ссылку на Tiktok, Youtube, Instagram")
		}
		log.Print(service)
		return c.Send(message)

	})
}
