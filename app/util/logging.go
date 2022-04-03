package util

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	// 読み書き、ファイル作成、ファイル追記
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatal(err)
	}

	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	// 日付、時間、ファイル名で出力する
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}
