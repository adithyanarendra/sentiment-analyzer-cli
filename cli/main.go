package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type SentimentRequest struct {
	Text string `json:"text"`
}

type SentimentResponse struct {
	Sentiment string  `json:"sentiment"`
	Polarity  float64 `json:"polarity"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: sentiment-analyzer \"your sentence here\"")
		return
	}

	text := os.Args[1]
	reqBody, _ := json.Marshal(SentimentRequest{Text: text})

	resp, err := http.Post("http://localhost:8000/analyze", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer resp.Body.Close()

	var result SentimentResponse
	json.NewDecoder(resp.Body).Decode(&result)

	fmt.Printf("Sentiment : %s (%.2f)\n", result.Sentiment, result.Polarity)
}
