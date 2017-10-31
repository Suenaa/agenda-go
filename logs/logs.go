package logs

import (
	"log"
	"os"
	"io"
)

var (
	out, e = os.OpenFile("logs/logs.log", os.O_APPEND | os.O_WRONLY | os.O_CREATE, 0666)
	outWriter = io.Writer(out)
	logger = log.New(outWriter, "[error] ", log.LstdFlags)
)

func Log(err error) {
	if e != nil {
		log.Fatalln("log file error")
	}
	defer out.Close()
	logger.Println(err)
}