package main

import (
	"fmt"
	"fog/learning"
)

func main() {
	learning.Hello()

	var quote = learning.Getquote()
	fmt.Println(quote)
}
