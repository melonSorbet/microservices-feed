package handlers

import (
	"encoding/json"
	"feed/src/models"
	"fmt"
	"io"
	"net/http"
)

func GetFollowing(userId string) (models.Users, error) {
	url := fmt.Sprintf("http://host.docker.internal:8001/users/%s/following", userId)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {

		return models.Users{}, err
	}
	cool, err := CreateToken(userId)
	if err != nil {
		return models.Users{}, err
	}
	// Set custom headers
	req.Header.Set("Authorization", cool)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return models.Users{}, err
	}

	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {

	}

	// Print the response body as a string (for debugging purposes)
	fmt.Println("Response Body:", string(body))

	// Unmarshal the JSON response into a struct
	users := models.Users{}
	err = json.Unmarshal(body, &users)
	if err != nil {

	}

	// Return the users or handle them as needed
	fmt.Println("Parsed Users:", users)
	// You could send a response or process further depending on your use case
	return users, nil
}
