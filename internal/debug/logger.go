package debug

import (
	"log"
	"os"
)

var Log *log.Logger

func Init() {
	file, err := os.OpenFile("termwords.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666,
	)
	if err != nil {
		return
	}
	Log = log.New(file, "", log.LstdFlags)
}
