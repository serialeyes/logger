package logger

import (
    "fmt"
)

var Version string = "1.0"

func Log(mess string) {
    fmt.Println("[LOG] " + mess)
}

func LogToFile(mess string, dir string) {
    f, err := os.Create(dir + "/log.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    f.WriteString(mess + "\n")
}
