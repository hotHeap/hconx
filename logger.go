package hconx

import (
	"io/ioutil"
	"log"
	"os"
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatalln("Get path failed.")
	}

	path = path + "/log"
	os.Chdir("~")

	file, err := os.OpenFile("error.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatalln("open file failed : ", err)
	}

	Trace = log.New(ioutil.Discard, "Trace : ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(os.Stdout, "Info : ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, "Warning : ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(file, "Error : ", log.Ldate|log.Ltime|log.Lshortfile) // Save to log file
}
