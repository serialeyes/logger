package logger

import (
    "fmt"
    "os"
    "log"
)

var Version string = "1.0"

func Log(mess string) {
    fmt.Println("[LOG] " + mess)
}

func LogToFile(mess string, dir string) {
    err := os.WriteFile("log.txt",[]byte(mess), 0666)
    if err != nil {
        log.Fatal(err)
    }
}
