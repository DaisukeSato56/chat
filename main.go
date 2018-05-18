package main

import (
  "log"
  "net/http"
)

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.write([]byte(`
      <html>
        <head>
          <title>チャット</title>
        </head>
      </html>
      `))
  })
}
