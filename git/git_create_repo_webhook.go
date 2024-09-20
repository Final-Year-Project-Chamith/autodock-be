package git

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateWebHookRepository(owner, repo, token, webhookURL, secret string) error {
	payload := map[string]interface{}{
		"name":   "web",
		"active": true,
		"events": []string{"push", "pull_request"},
		"config": map[string]string{
			"url":          webhookURL,
			"content_type": "json",
			"secret":       secret,
		},
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("Error marshalling JSON payload: %v", err)
	}

	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/hooks", owner, repo)
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}
	req.Header.Set("Authorization", "token "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}
	if resp.StatusCode == http.StatusCreated {
		fmt.Println("Webhook created successfully!")
		fmt.Println("Response:", string(body))
	} else {
		fmt.Printf("Failed to create webhook: %s\n", resp.Status)
		fmt.Println("Response:", string(body))
	}
	return nil
}
