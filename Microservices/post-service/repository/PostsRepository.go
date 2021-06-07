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
	result, err :=  repo.Database.Get(post.USERID.String()).Result()
	var posts[] model.Post
	if err != nil {
		posts = append(posts, *post)
		jsonPosts, _ := json.Marshal(posts)
		newErr := repo.Database.Set(post.USERID.String(), jsonPosts, 0).Err()
		if newErr != nil {
			fmt.Println(result)
		}
	} else{
		bytes := []byte(result)
		json.Unmarshal(bytes, &posts)
		posts = append(posts, *post)
		jsonPosts, _ := json.Marshal(posts)
		newErr := repo.Database.Set(post.USERID.String(), jsonPosts, 0).Err()
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

func(repo *PostsRepository) GetFeed(id string) []model.Post {
	fmt.Println("Id je " + id)
	var posts []model.Post
	var feedInputs []model.Feed
	result, _ :=  repo.Database.Get(id + "_feed").Result()
	bytes := []byte(result)
	json.Unmarshal(bytes, &feedInputs)
	for i := range feedInputs {
		var userPosts []model.Post
		result, _ :=  repo.Database.Get(feedInputs[i].UserId.String()).Result()
		bytes := []byte(result)
		json.Unmarshal(bytes, &userPosts)
		for j := range  userPosts {
			if userPosts[j].ID == feedInputs[i].PostId {
				posts = append(posts, userPosts[j])
				break
			}
		}
	}
	return posts

}

func(repo *PostsRepository) AddPostToFeed(keys []string, post *model.Post) error {
	for i := range keys {
		result, err := repo.Database.Get(keys[i] + "_feed").Result()
		var feedInputs []model.Feed
		var feed model.Feed
		feed.PostId = post.ID
		feed.UserId = post.USERID
		feedInputs = append(feedInputs, feed)

		if err == nil {
			bytes := []byte(result)
			var tmp []model.Feed
			json.Unmarshal(bytes, &tmp)
			feedInputs = append(feedInputs, tmp...)
		}

		jsonFeed, _ := json.Marshal(feedInputs)
		newErr := repo.Database.Set(keys[i] + "_feed", jsonFeed, 0).Err()
		if newErr != nil {
			fmt.Println(result)
		}
	}
	return nil
}