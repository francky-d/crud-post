package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"
)

func getPosts() []Post {
	return []Post{
		{ID: 1, Title: "post1", Content: "this is the content of post one", CreatedAt: currentTimestamp(), UpdatedAt: currentTimestamp()},
		{ID: 2, Title: "post2", Content: "this is the content of post one", CreatedAt: currentTimestamp(), UpdatedAt: currentTimestamp()},
		{ID: 3, Title: "post3", Content: "this is the content of post one", CreatedAt: currentTimestamp(), UpdatedAt: currentTimestamp()},
		{ID: 4, Title: "post4", Content: "this is the content of post one", CreatedAt: currentTimestamp(), UpdatedAt: currentTimestamp()},
	}
}

func Router() *http.ServeMux {
	router := http.NewServeMux()
	posts := getPosts()
	router.HandleFunc("GET /posts", func(response http.ResponseWriter, request *http.Request) {
		postsToJson, err := json.Marshal(posts)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte(ErrSomethingWentWrong.Error()))
		}
		response.Write(postsToJson)
	})

	router.HandleFunc("GET /posts/{id}", func(response http.ResponseWriter, request *http.Request) {
		ID := request.PathValue("id")
		if ID == "" {
			response.WriteHeader(http.StatusBadRequest)
			response.Write([]byte("id is required"))
			return
		}

		postID, err := strconv.Atoi(ID)
		if err != nil {
			response.WriteHeader(http.StatusBadRequest)
			response.Write([]byte("id must be an integer"))
			return
		}

		post := findPost(postID, posts)

		if post.ID == 0 {
			response.WriteHeader(http.StatusNotFound)
			response.Write([]byte(ErrNotFound.Error()))
		}

		postJson, err := json.Marshal(post)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte(ErrSomethingWentWrong.Error()))
		}
		response.Write(postJson)
	})

	return router
}

var ErrNotFound = errors.New("resource not found")
var ErrSomethingWentWrong = errors.New("something went wrong")

type Post struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func currentTimestamp() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

func findPost(postID int, allPosts []Post) Post {
	for _, post := range allPosts {
		if post.ID == postID {
			return post
		}
	}
	return Post{}
}

func main() {

	router := Router()
	log.Fatal(http.ListenAndServe(":8000", router))

}
