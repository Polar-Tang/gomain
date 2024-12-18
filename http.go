package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// type Response struct {
// 	Status   int        `json:"Status"`
// 	TC       bool       `json:"TC"`
// 	RD       bool       `json:"RD"`
// 	RA       bool       `json:"RA"`
// 	AD       bool       `json:"AD"`
// 	CD       bool       `json:"CD"`
// 	Question []Question `json:"Question"`
// 	Answer   []Answer   `json:"Answer"`
// }

func getIPAddress(domain string) (string, error) {
	url := fmt.Sprintf("https://cloudflare-dns.com/dns-query?name=%s&type=A", domain)

	var DNSResponse DNSResponse

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("accept", "application/dns-json")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	if err := json.Unmarshal(body, &DNSResponse); err != nil {
		return "", fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	var answerString string

	for _, ans := range DNSResponse.Answer {
		answerString += ans.Data + " "
	}

	IPdomain := fmt.Sprintf("The domain %s resolves to %s", domain, answerString)

	return IPdomain, nil
}
