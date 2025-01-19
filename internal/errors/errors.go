package errors

import (
	"fmt"
	"log"
)

const help = "Please make a new issue if you think this is a bug: https://lichess-tui/issues."

func RequestError(err error) {
	fmt.Println(help)
	log.Fatalf("Error making request: %v", err)
}
