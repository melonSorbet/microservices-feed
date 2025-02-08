package handlers

import (
	"encoding/json"
	"feed/src/models"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func GetPostsByTags(tags []string) (models.Post, error) {
	baseURL := "http://host.docker.internal:8000/posts/feed/"
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		fmt.Println("Could not parse URL")
		return models.Post{}, err
	}
	q := parsedURL.Query()

	for _, tag := range tags {

		q.Add("tags", tag)
	}
	parsedURL.RawQuery = q.Encode()
	request, err := http.Get(parsedURL.String())
	if err != nil {
		fmt.Println("Could not reach Post MS")
		return models.Post{}, err
	}
	post := models.Post{}
	body, err := io.ReadAll(request.Body)
	if err != nil {
		fmt.Println("Could not read request")
		return models.Post{}, err
	}

	err = json.Unmarshal(body, &post)
	if err != nil {
		fmt.Println("could not unmarshal request")

		return models.Post{}, err
	}
	return post, nil
}

func GetPostsByFriends(users models.Users) ([]models.Post, error) {

	var posts []models.Post
	for _, users := range users.Following {

		request, err := http.Get("http://host.docker.internal:8000/posts/users/" + users.UserID)

		if err != nil {
			fmt.Println("Could not do request")

			return []models.Post{}, err
		}
		if request.StatusCode == http.StatusNotFound {
			continue
		}
		post := models.Post{}
		body, err := io.ReadAll(request.Body)
		if err != nil {
			fmt.Println("Could not read request")
			return []models.Post{}, err
		}
		err = json.Unmarshal(body, &post)

		if err != nil {
			fmt.Println("User not found")
			// Return empty posts or handle this case differently
			return []models.Post{}, nil
		}

		posts = append(posts, post)

	}
	return posts, nil
}
