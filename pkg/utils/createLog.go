package utils

import (
	"io"
	"log"
	"os"
)

func DefineLogger() {
	var f *os.File
	if _, err := os.Stat("log"); os.IsNotExist(err) {
		_ = os.Mkdir("log", os.ModePerm)
	}

	f, _ = os.Create("log/music-learner.log")

	var offset int64 = 0
	if _, seekErr := f.Seek(offset, io.SeekEnd); seekErr != nil {
		log.Panicln("unable to go to end of file music-learner.log", seekErr)
	}

	log.SetOutput(io.MultiWriter(os.Stdout, f))
}
