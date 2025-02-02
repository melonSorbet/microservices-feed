package service

import (
	"encoding/json"
	"feed/src/handlers"
	"fmt"
	"net/http"
)

func GetExplorePage(w http.ResponseWriter, r *http.Request) {
	tags := r.URL.Query()["tags"]
	posts, err := handlers.GetPostsByTags(tags)
	if err != nil {
		return
	}
	SortPostsAfterCreationDate(posts)
	postEncoding, err := json.Marshal(posts)
	if err != nil {
		return
	}

	fmt.Fprintf(w, string(postEncoding))
}
