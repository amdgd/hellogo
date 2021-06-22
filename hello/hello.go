package main

import (
	"fmt"

	"github.com/ankitind/hellogo/greetings"
)

//"github.com/ankitind/hellogo/greetings"

func main() {
	message := greetings.Hello("Ankit M")
	fmt.Println(message)
}
