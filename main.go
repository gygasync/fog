package main

import (
	"fmt"
	"fog/learning"
)

func main() {
	logger := learning.Getlogger(false)
	defer logger.Sync()

	learning.Log("This is only visible in the file log")
	learning.Log("Application started")
	learning.Hello()

	var quote = learning.Getquote()
	fmt.Println(quote)
	learning.Log(quote)

	learning.TestPageOps()
	learning.TestHttpHandler()

	learning.Log("Exiting from application")
}
