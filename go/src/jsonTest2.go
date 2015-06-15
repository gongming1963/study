package main

import (
    "encoding/json"
    "fmt"
    "io"
    "log"
    "strings"
)

func main() {
    const jsonStream = `
        {"ip": "Ed", "Text": 23}
        {"ip": "Ed", "Text": 12}
    `
    type Message struct {
        Ip string
        Text int
    }
    dec := json.NewDecoder(strings.NewReader(jsonStream))
    for {
        var m Message
        if err := dec.Decode(&m); err == io.EOF {
            break
        } else if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("%s: %d\n", m.Ip, m.Text)
    }
}

