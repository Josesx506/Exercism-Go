package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

func streamEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Content-Type", "text/event-stream") // primary header for event stream

	tokens := []string{
		"Server sent events(SSE)", "is a pushing", "technology that", "enables", "pushing",
		"notification/message/events", "from the server to", "the client(s) via HTTP",
		"connection.", "While you are developing", "real-time projects,", "there is",
		"always", "a", "one-question mark", "on “how to send", "messages/updates",
		"from server to", "client”. We", "can talk about", "three different ways to",
		"perform server-to-client", "updates: Client polling,", "Web Socket,", "Server-Sent",
		"Events (SSE).",
		"If you have", "another instance", "of your app", "running (or another", "service",
		"like Docker,", "Node, or a", "zombie Go process)", "on the", "same", "port,", "the",
		"server will fail", "immediately."}

	for _, token := range tokens {
		// Prefix with "data: " and with TWO newlines "\n\n" unless the browser
		// EventSource won't trigger
		content := fmt.Sprintf("data: %s\n\n", string(token))
		w.Write([]byte(content))
		w.(http.Flusher).Flush() // flush the buffer

		time.Sleep((time.Millisecond * 200)) // simulate delay
	}

}

func main() {
	addr := flag.String("addr", ":8080", "http service address")
	flag.Parse()

	http.HandleFunc("/events", streamEvents)
	fmt.Printf("Starting server on %s\n", *addr)
	http.ListenAndServe(*addr, nil)

	// Run it with
	// go run api.go -addr=:9020
	// go tool air sse.go // Live reload with air
}
