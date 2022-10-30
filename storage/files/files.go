package files

import (
	"encoding/gob"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"read-advisor-bot/pkg/e"
	"read-advisor-bot/storage"
	"time"
)

const defaultPerm = 0774

type Storage struct {
	basePath string
}

func New(basePath string) *Storage {
	return &Storage{basePath: basePath}
}

func (s *Storage) Save(page *storage.Page) error {
	fPath := filepath.Join(s.basePath, page.UserName)

	if err := os.MkdirAll(fPath, defaultPerm); err != nil {
		return e.Wrap("can't save page", err)
	}

	fName, err := fileName(page)
	if err != nil {
		return e.Wrap("can't save page", err)
	}

	fPath = filepath.Join(fPath, fName)

	file, err := os.Create(fPath)

	defer file.Close()

	err = gob.NewEncoder(file).Encode(page)
	if err != nil {
		return e.Wrap("can't save page", err)
	}

	return nil
}

func (s *Storage) PickRandom(UserName string) (*storage.Page, error) {
	fPath := filepath.Join(s.basePath, UserName)

	if _, err := os.Stat(fPath); os.IsNotExist(err) {
		return nil, e.Wrap("can't pick random page", storage.ErrStorageNotExists)
	}

	files, err := os.ReadDir(fPath)
	if err != nil {
		return nil, e.Wrap("can't pick random page", err)
	}

	if len(files) == 0 {
		return nil, e.Wrap("can't pick random page", storage.ErrNoSavedPages)
	}

	rand.Seed(time.Now().Unix())
	n := rand.Intn(len(files))

	file := files[n]

	return s.decodePage(filepath.Join(fPath, file.Name()))
}

func (s *Storage) Remove(p *storage.Page) error {
	fName, err := fileName(p)
	if err != nil {
		return e.Wrap("can't remove page", err)
	}

	fPath := filepath.Join(s.basePath, p.UserName, fName)

	if err := os.Remove(fPath); err != nil {
		return e.Wrap(fmt.Sprintf("can't remove page %s", fPath), err)
	}

	return nil
}

func (s *Storage) IsExists(p *storage.Page) (bool, error) {
	fName, err := fileName(p)
	if err != nil {
		return false, e.Wrap("can't check if page exists", err)
	}

	fPath := filepath.Join(s.basePath, p.UserName, fName)

	switch _, err := os.Stat(fPath); {
	case errors.Is(err, os.ErrNotExist):
		return false, nil
	case err != nil:
		return false, e.Wrap(fmt.Sprintf("can't check if page %s exists", fPath), err)
	}

	return true, nil
}

func (s *Storage) decodePage(filePath string) (*storage.Page, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, e.Wrap("can't decode page", err)
	}

	defer f.Close()

	var p storage.Page

	if err := gob.NewDecoder(f).Decode(&p); err != nil {
		return nil, e.Wrap("can't decode page", err)
	}

	return &p, nil
}

func fileName(p *storage.Page) (string, error) {
	return p.Hash()
}
