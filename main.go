package main

import (
        "fmt"
        "net/http"
        "time"
)

func main() {
        mux := http.NewServeMux()
        mux.HandleFunc("/", events)

        handler := enableCORS(mux)

        err := http.ListenAndServe(":3001", handler)
        if err != nil {
                fmt.Println("Error:", err)
        }
}

func enableCORS(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
                w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
                w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

                if r.Method == "OPTIONS" {
                        w.WriteHeader(http.StatusOK)
                        return
                }

                next.ServeHTTP(w, r)
        })
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
