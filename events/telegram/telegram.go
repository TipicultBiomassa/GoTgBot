package telegram

import (
	"GoTgBot/clients/telegram"
	"GoTgBot/events"
	"GoTgBot/lib/e"
	"GoTgBot/storage"
	"errors"
)

type Processor struct {
	tg      *telegram.Client
	offset  int
	storage storage.Storage
}
type Meta struct {
	ChatID   int
	Username string
}

var (
	ErrUnknownMetaType  = errors.New("unknown meta type")
	ErrUnknownEventType = errors.New("unknown event type")
)

func New(client *telegram.Client, storage storage.Storage) *Processor {
	return &Processor{
		tg:      client,
		storage: storage,
	}
}
func (p *Processor) Fetch(limit int) ([]events.Event, error) {
	update, err := p.tg.Updates(p.offset, limit)
	if err != nil {
		return nil, e.WrapIfNil("can't get events", err)
	}
	if len(update) == 0 {
		return nil, nil
	}
	res := make([]events.Event, 0, len(update))

	for _, u := range update {
		res = append(res, event(u))
	}

	p.offset = update[len(update)-1].ID + 1

	return res, nil
}

func (p Processor) Process(event events.Event) error {
	switch events.Type {
	case events.Message:
		p.processMessage(event)
	default:
		return e.Wrap("can't process message", ErrUnknownEventType)

	}
}

func (p *Processor) processMessage(event events.Event) error {
	meta, err := meta(event)
	if err != nil {
		return e.Wrap("can't process msg", err)
	}

	return
}

func meta(event2 events.Event) (Meeta, error) {
	res, ok := event.Meta.(Meta)
	if !ok {
		return Meta{}, e.Wrap("Can't get meta", ErrUnknownMetaType)
	}
	return res, nil
}
func event(upd telegram.Update) events.Event {
	updType := fetchType(upd)
	res := events.Event{
		Type: updType,
		Text: fetchText(upd),
	}
	if updType == events.Message {
		res.Meta = Meta{
			ChatID:   upd.Message.Chat.ID,
			Username: upd.Message.From.Username,
		}
	}

	return res
	// chatID username
}
func fetchType(upd telegram.Update) events.Type {
	if upd.Message == nil {
		return events.Unknown
	}
	return events.Message
}

func fetchText(upd telegram.Update) string {
	if upd.Message == nil {
		return ""
	}
	return upd.Message.Text
}
