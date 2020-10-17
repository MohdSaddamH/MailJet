package main

import (
    "log"
    "os"
)

func main() {
    file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

    log.SetOutput(file)
    log.Println("Logging to a file in Go!")


    //Set Environment Variables
    os.Setenv("APIKEY_PUBLIC", "<Your API KEY>")
	os.Setenv("APIKEY_PRIVATE", "<Your Secret KEY>")

    handler()
}