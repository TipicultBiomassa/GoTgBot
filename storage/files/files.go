package files

import (
	"GoTgBot/lib/e"
	"GoTgBot/storage"
)

type Storage struct {
	basePath string
}

func New(basePath string) Storage {
	return Storage{basePath: basePath}
}

func (s Storage) Save(page *storage.Page) (err error) {
	defer func() { err = e.WrapIfNil("can't save", err) }()
}
