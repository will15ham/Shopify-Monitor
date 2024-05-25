package util

import (
	"log"
	"os"
)

var (
    logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
)

func Log(message string) {
    logger.Println(message)
}
