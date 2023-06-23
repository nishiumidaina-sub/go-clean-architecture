package lib

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func GenerateKey() (string, error) {
	id, err := gonanoid.New(12)
	if err != nil {
		return "", err
	}
	
	return id, nil
}
