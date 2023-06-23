package log

import (
	"os"
)

func GeneratApiLog() (fp *os.File) {
	fp, err := os.OpenFile("log/api_debug.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	return fp
}