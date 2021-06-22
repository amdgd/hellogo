package hellogo

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
	"rsc.io/quote"
)

func hellogo(name string) (string string){
	message := fmt.Println("Hello ", name "!")
	quote := fmt.Println(quote.Go())
	return message, quote
}
