package learning

import (
	"fmt"
	"os"
	"time"
)

var logger = Getlogger(false)
var filename = "./logs/log_" + fmt.Sprintf("%d", time.Now().UnixMilli()) + ".txt"
var err = os.Mkdir("logs", 0755)

func Log(text string) {
	logger.Info(text)

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE, 0666)
	//defer file.Close()
	if err == nil {
		file.Write([]byte(time.Now().String() + " " + text + "\n"))
	} else {
		logger.Fatal("Cannot LOG to file" + err.Error())
	}
}
