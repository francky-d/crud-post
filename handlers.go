package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

func GetAllHandler(response http.ResponseWriter, request *http.Request) {
	posts := All()
	postsToJson, err := json.Marshal(posts)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(ErrSomethingWentWrong.Error()))
	}
	response.Write(postsToJson)
}

func FindPostHandler(response http.ResponseWriter, request *http.Request) {
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

	post := findByID(postID)

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
}

func AddPost(response http.ResponseWriter, request *http.Request) {
	data, err := io.ReadAll(request.Body)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("Something went wrong while parsing body"))
		return
	}

	postToCreate := Post{}
	err = json.Unmarshal(data, &postToCreate)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("Somthing went wrong while unmarshaling"))
		return
	}

	if postToCreate.Title == "" {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("title is required"))
	}

	if postToCreate.Content == "" {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("content is required"))
		return
	}

	postToCreate.CreatedAt = currentTimestamp()
	postToCreate.UpdatedAt = currentTimestamp()
	postToCreate.ID = getPosts()[len(getPosts())-1].ID + 1

	response.WriteHeader(http.StatusCreated)
	//TODO : save the post
	json.NewEncoder(response).Encode(postToCreate)
}
