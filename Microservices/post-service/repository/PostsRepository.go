package repository

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"post_service/model"
)

type PostsRepository struct {
	Database *redis.Client
}

func(repo *PostsRepository) Create(post *model.Post)  error {
	result, err :=  repo.Database.Get(post.USERNAME).Result()
	var posts[] model.Post
	if err != nil {
		posts = append(posts, *post)
		jsonPosts, _ := json.Marshal(posts)
		newErr := repo.Database.Set(post.USERNAME, jsonPosts, 0).Err()
		if newErr != nil {
			fmt.Println(result)
		}
	} else{
		bytes := []byte(result)
		json.Unmarshal(bytes, &posts)
		posts = append(posts, *post)
		jsonPosts, _ := json.Marshal(posts)
		newErr := repo.Database.Set(post.USERNAME, jsonPosts, 0).Err()
		if newErr != nil {
			fmt.Println(result)
		}
	}
	return nil
}

func(repo *PostsRepository) GetByKey(key string) *model.Post {
	fmt.Println("Key je " + key)
	var post *model.Post
	result, _ :=  repo.Database.Get(key).Result()
	bytes := []byte(result)
	json.Unmarshal(bytes, &post)
	return post

}