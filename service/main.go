package main

import (
	"errors"
	"log"
	"os"
)

func main() {

	f, err := os.OpenFile("./logs/testlogfile.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	e := errors.New("test err")

	log.Printf("This is a test log entry: %+v", e)
}
