package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", events)
	err := http.ListenAndServe(":3001", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func events(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	tokens := []string{"this", "is", "me", "lauda", "lasan", "bhonsada", "mara", "tumhara"}

	for _, token := range tokens {
		content := fmt.Sprintf("data: %s\n\n", token)
		_, err := w.Write([]byte(content))
		if err != nil {
			fmt.Println("Error writing to response:", err)
			return
		}

		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}

		time.Sleep(time.Millisecond * 400)
	}
}
