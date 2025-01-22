package main

type Post struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func All() []Post {
	return getPosts()
}

func findByID(postID int) Post {
	for _, post := range getPosts() {
		if post.ID == postID {
			return post
		}
	}
	return Post{}
}

func getPosts() []Post {
	return []Post{
		{ID: 1, Title: "post1", Content: "this is the content of post one", CreatedAt: currentTimestamp(), UpdatedAt: currentTimestamp()},
		{ID: 2, Title: "post2", Content: "this is the content of post one", CreatedAt: currentTimestamp(), UpdatedAt: currentTimestamp()},
		{ID: 3, Title: "post3", Content: "this is the content of post one", CreatedAt: currentTimestamp(), UpdatedAt: currentTimestamp()},
		{ID: 4, Title: "post4", Content: "this is the content of post one", CreatedAt: currentTimestamp(), UpdatedAt: currentTimestamp()},
	}
}
