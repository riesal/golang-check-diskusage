package main

import (
"os"
"net/http"
"fmt"
"io/ioutil"
)
func main() {

apipath := "https://api.ipify.org"
req, err := http.NewRequest("GET", apipath, nil)
    req.Header.Set("User-Agent", "")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Printf(string(body))
}
