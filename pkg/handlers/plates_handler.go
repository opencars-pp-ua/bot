package handlers

import (
	"log"
	"strings"

	"github.com/opencars/bot/internal/bot"
)

func (h OpenCarsHandler) PlatesHandler(msg *bot.Event) {
	if err := msg.SetStatus(bot.ChatTyping); err != nil {
		log.Printf("action error: %s", err.Error())
	}

	plates := strings.TrimPrefix(msg.Message.Text, "/plates ")
	text, err := h.getInfoByPlates(plates)
	if err != nil {
		log.Println(err.Error())
	}

	if err := msg.SendHTML(text); err != nil {
		log.Printf("send error: %s\n", err.Error())
	}
}
