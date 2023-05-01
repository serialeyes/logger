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
    f, err := os.OpenFile("log.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil{
        log.Fatal(err)
    }
    if _, err := f.Write([]byte(mess)); err != nil {
        log.Fatal(err)
    }
    if err := f.Close(); err != nil {
        log.Fatal(err)
    }
}
