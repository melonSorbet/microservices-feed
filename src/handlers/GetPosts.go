package handlers

import (
	"encoding/json"
	"feed/src/models"
	"fmt"
	"io"
	"net/http"
)

func GetPostsByTags() {
	// TODO: GetPostsByUserTags
}

func GetPostsByFriends(users models.Users) ([]models.Post, error) {

	var posts []models.Post
	for _, users := range users.Following {

		request, err := http.Get("http://host.docker.internal:8000/posts/users/" + users.UserID)
		if err != nil {
			fmt.Println("Could not do request")

			return []models.Post{}, err
		}

		post := models.Post{}
		body, err := io.ReadAll(request.Body)
		if err != nil {
			fmt.Println("Could not read request")
			return []models.Post{}, err
		}

		err = json.Unmarshal(body, &post)
		if err != nil {
			fmt.Println("could not unmarshal request")

			return nil, err
		}

		posts = append(posts, post)

	}
	return posts, nil
}
