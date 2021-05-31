package repository

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"post_service/model"
)

type CommentsRepository struct {
	Database *redis.Client
}

func(repo *CommentsRepository) Create(comment *model.Comment)  error {
	fmt.Println(comment)
	s, _ := json.Marshal(comment)
	fmt.Println(s)
	result := repo.Database.Set( comment.ID.String(), s, 0).Err()
	if result != nil {
		fmt.Println(result)
	}
	return nil
}

func(repo *CommentsRepository) GetByKey(key string) *model.Comment {
	fmt.Println("Key je " + key)
	var comment *model.Comment
	result, _ :=  repo.Database.Get(key).Result()
	bytes := []byte(result)
	json.Unmarshal(bytes, &comment)
	return comment

}