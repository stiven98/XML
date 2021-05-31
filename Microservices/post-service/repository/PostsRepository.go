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
	fmt.Println(post)
	jsonPost, _ := json.Marshal(post)
	fmt.Println(jsonPost)
	result := repo.Database.Set( post.ID.String(), jsonPost, 0).Err()
	if result != nil {
		fmt.Println(result)
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