package utils

import (
	"log"
	"os"
	"strconv"
	"time"
)

const (
	layoutISO = "2006-01-02"
)

func openFile() (*os.File, error) {
	t := time.Now()
	temp := t.Format(layoutISO)
	temp = "logs/" + temp + ".logs"
	log.Println("temp:", temp)
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		os.Mkdir("logs", os.ModePerm)
	}
	res := temp
	i := 0
	for {

		f, err := os.OpenFile(res, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

		fi, err := f.Stat()
		if err != nil {
			return nil, err
		}

		if fi.Size() > 1024*1024*8 {
			res = temp
			i++
			res += "_" + strconv.Itoa(i)
		} else {
			return f, nil
		}

		f.Close()

	}

}

//LogErrToFile ...
func LogErrToFile(text string) {

	f, err := openFile()

	if err != nil {
		log.Println(err)
	}

	defer f.Close()

	log.SetOutput(f)
	log.Println(text)
}
