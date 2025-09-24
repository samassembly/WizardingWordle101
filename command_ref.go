package main

import (
    "fmt"
  "strings"
  "net/http"
  "io"
)

func commandRef() error {
	url := "https://wiki.wizard101central.com/wiki/api.php?action=parse&format=json&page=Spell%3ADeathblade"
  method := "GET"

  payload := strings.NewReader("action=parse&format=json&uselang=user&page=Spell%3ADeathblade")

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)
  if err != nil {
    fmt.Println(err)
    return err
  }
  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
  req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/140.0.0.0 Safari/537.36 Edg/140.0.0.0")

  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return err
  }
  defer res.Body.Close()

  body, err := io.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return err 
  }
  fmt.Println(string(body))
  return nil
}