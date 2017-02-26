package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)


type webhook struct {
	Action     string
	Repository struct {
		ID       string
		FullName string
	}
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	webhookData := make(map[string]interface{})
	err := json.NewDecoder(r.Body).Decode(&webhookData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Println("Got webhook payload: ")
	for k, v := range webhookData {
		fmt.Printf("%s : %v\n", k, v)
	}
	fmt.Printf("headers: %v\n", r.Header)
}

func main() {
	log.Println("server started")
	http.HandleFunc("/webhook", handleWebhook)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
