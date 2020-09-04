package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/go-github/github"
)

var (
	secretKey = "asecretonlyiknow"
)

func printEvents(w http.ResponseWriter, r *http.Request) {
	body, err := github.ValidatePayload(r, []byte(secretKey))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	defer r.Body.Close()

	// Let's make the JSON we receive pretty...
	data := make(map[string]interface{})
	if err := json.Unmarshal(body, &data); err == nil {
		b := bytes.Buffer{}
		encoder := json.NewEncoder(&b)
		encoder.SetIndent("", "\t")
		if err = encoder.Encode(data); err == nil {
			// Override the body variable when no errors are present
			body = b.Bytes()
		} else {
			fmt.Println("Error: " + err.Error())
		}
	} else {
		fmt.Println("Error: " + err.Error())
	}

	// Print the event to console
	fmt.Printf("Event on %s:\n%s\n", time.Now().Format("Jan 2 15:04:05"),
		string(body))
}

func main() {
	http.HandleFunc("/", printEvents)
	http.ListenAndServe(":8000", nil)
}
