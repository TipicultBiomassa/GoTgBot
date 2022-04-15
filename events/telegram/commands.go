package telegram

import (
	"log"
	"strings"
)

func (p *Processor) doCmd(text string, chatID int, username string) error {
	text = strings.TrimSpace(text)

	log.Printf("New command '%s' from '%s'", text, username)

}
