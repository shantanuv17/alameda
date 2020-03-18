package utils

import (
	"os"
	"path/filepath"
	"time"
)

func TouchFile(fileLoc string) error {
	fileDir := filepath.Dir(fileLoc)
	if _, err := os.Stat(fileDir); os.IsNotExist(err) {
		err = os.MkdirAll(fileDir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	_, err := os.Stat(fileLoc)
	if os.IsNotExist(err) {
		file, err := os.Create(fileLoc)
		if err != nil {
			return err
		}
		defer file.Close()
	} else {
		now := time.Now().Local()
		err = os.Chtimes(fileLoc, now, now)
		return err
	}
	return nil
}

func FileNotUpdateTimeSec(fileLoc string) (int64, error) {
	file, err := os.Stat(fileLoc)
	if err != nil {
		return 0, err
	}
	return time.Now().Unix() - file.ModTime().Unix(), nil
}
