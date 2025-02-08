package service

import (
	"encoding/json"
	"feed/src/handlers"
	"feed/src/models"
	"fmt"
	"net/http"
	"sort"
)

func GetFriendsFeed(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("userId")
	if userId == "" {
		fmt.Fprintf(w, "no userId given")
		return
	}
	users, err := handlers.GetFollowing(userId)
	if err != nil {
		fmt.Println("damn", err)
		return
	}
	posts, err := handlers.GetPostsByFriends(users)
	if err != nil {
		fmt.Println(err)
		return
	}

	var all_posts models.Post
	for userIndex := 0; userIndex < len(posts); userIndex++ {
		for postIndex := 0; postIndex < len(posts[userIndex]); postIndex++ {
			for _, user := range users.Following {
				if posts[userIndex][postIndex].UserID == user.UserID {
					// Add user profile information to the post
					posts[userIndex][postIndex].UserProfile = struct {
						Username         string `json:"username"`
						DisplayName      string `json:"display_name"`
						ProfilePicURL    string `json:"profile_pic_url"`
						ProfilePicWidth  int    `json:"profile_pic_width"`
						ProfilePicHeight int    `json:"profile_pic_height"`
					}{
						Username:         user.Username,
						DisplayName:      user.DisplayName,
						ProfilePicURL:    user.ProfilePicURL,
						ProfilePicWidth:  user.ProfilePicWidth,
						ProfilePicHeight: user.ProfilePicHeight,
					}
					break
				}
			}
			all_posts = append(all_posts, posts[userIndex][postIndex])
		}
	}
	SortPostsAfterCreationDate(all_posts)
	response, err := json.Marshal(all_posts)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(w, string(response))
}

func SortPostsAfterCreationDate(posts models.Post) {
	// Sort by CreatedAt using sort.Slice() with custom comparison logic
	sort.Slice(posts, func(i, j int) bool {
		// Compare CreatedAt times
		return posts[i].CreatedAt.After(posts[j].CreatedAt)
	})
}
