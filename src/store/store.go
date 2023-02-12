package store

import (
	"errors"
	"os"
	"path/filepath"
    "github.com/ProtonMail/gopenpgp/v2/helper"
)

type Store struct {
	Dir string
}

func GetStore() (*Store, error) {
	// Try to get pwm store directory
	dir, err := os.UserHomeDir()

	if err != nil {
		return nil, err
	}

	dir = filepath.Join(dir, "pwm")

	err = os.Mkdir(dir, 0750)

	if err != nil && !os.IsExist(err) {
		return nil, err
	}

	return &Store{
		Dir: dir,
	}, nil
}

func (s Store) IsFileExists(fname string) bool {
	finfo, err := os.Stat(filepath.Join(s.Dir, fname))

	if err != nil || finfo.IsDir() {
		return false
	}

	return true
}

func (s Store) WriteFile(fname string, data []byte) error {
	return os.WriteFile(filepath.Join(s.Dir, fname), data, 0777)
}

func (s Store) ReadFile(fname string) ([]byte, error) {
	// If exists, show alert
	if !s.IsFileExists(fname) {
		return nil, errors.New("Specified file doesn't exist in the store")
	}

	return os.ReadFile(filepath.Join(s.Dir, fname))
}

func (s Store) ListFiles() ([]os.DirEntry, error) {
    return os.ReadDir(s.Dir)
}

func (s Store) RemoveFile(fname string) error {
    return os.Remove(filepath.Join(s.Dir, fname))
}

func (s Store) Decrypt(fname string, pass []byte) (string, error) {
	content, err := s.ReadFile(fname)

	if err != nil {
		return "", err
	}

	return helper.DecryptMessageWithPassword(pass, string(content))
}
