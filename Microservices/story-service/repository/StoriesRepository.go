package repository

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"storyservice/model"
	"time"
)

type StoriesRepository struct {
	Database *redis.Client
}

func(repo *StoriesRepository) Create(story *model.Story)  error {
	fmt.Println(story)
	s, _ := json.Marshal(story)
	fmt.Println(s)
	result := repo.Database.Set( story.ID.String(), s, time.Second*20).Err()
	if result != nil {
		fmt.Println(result)
	}
	return nil
}

func(repo *StoriesRepository) GetByKey(key string) *model.Story {
	fmt.Println("Key je " + key)
	var story *model.Story
	result, _ :=  repo.Database.Get(key).Result()
	bytes := []byte(result)
	json.Unmarshal(bytes, &story)
	fmt.Println(story.HASHTAG)
	return story

}
