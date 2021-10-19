package main

import (
	"fmt"
	"fog/learning"
)

func main() {
	logger := learning.Getlogger(false)
	defer logger.Sync()

	logger.Info("Application started")
	learning.Hello()

	var quote = learning.Getquote()
	fmt.Println(quote)
	logger.Info(learning.Getinfo())
	logger.Warn(learning.Getwarning())
	logger.Error(learning.Geterror())
	logger.Info("Exiting from application")
}
