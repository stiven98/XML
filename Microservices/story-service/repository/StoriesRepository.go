package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"storyservice/model"
	"time"
)

type StoriesRepository struct {
	Database *redis.Client
}

func(repo *StoriesRepository) Create(story *model.Story)  error {
	result, err :=  repo.Database.Get(story.USERID.String()).Result()
	var stories [] model.Story
	if err != nil {
		stories = append(stories, *story)
		jsonPosts, _ := json.Marshal(stories)
		newErr := repo.Database.Set(story.USERID.String(), jsonPosts, 0).Err()
		if newErr != nil {
			fmt.Println(result)
		}
	} else{
		bytes := []byte(result)
		json.Unmarshal(bytes, &stories)
		stories = append(stories, *story)
		jsonPosts, _ := json.Marshal(stories)
		newErr := repo.Database.Set(story.USERID.String(), jsonPosts, 0).Err()
		if newErr != nil {
			fmt.Println(result)
		}
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

func (repo *StoriesRepository) AddStoryToFeed(keys []string, story *model.Story) error {
	for i := range keys {
		result, err := repo.Database.Get(keys[i] + "_feed").Result()
		var feedInputs []model.Feed
		var feed model.Feed
		feed.StoryId = story.ID
		feed.UserId = story.USERID
		feedInputs = append(feedInputs, feed)

		if err == nil {
			bytes := []byte(result)
			var tmp []model.Feed
			json.Unmarshal(bytes, &tmp)
			feedInputs = append(feedInputs, tmp...)
		}

		jsonFeed, _ := json.Marshal(feedInputs)
		newErr := repo.Database.Set(keys[i] + "_feed", jsonFeed, 24 * time.Hour).Err()
		if newErr != nil {
			fmt.Println(result)
		}
	}
	return nil
}

func (repo *StoriesRepository) GetFeed(id string) []model.Story {
	fmt.Println("Id je " + id)
	var stories []model.Story
	var feedInputs []model.Feed
	result, _ :=  repo.Database.Get(id + "_feed").Result()
	bytes := []byte(result)
	json.Unmarshal(bytes, &feedInputs)
	for i := range feedInputs {
		var userPosts []model.Story
		result, _ :=  repo.Database.Get(feedInputs[i].UserId.String()).Result()
		bytes := []byte(result)
		json.Unmarshal(bytes, &userPosts)
		for j := range  userPosts {
			if userPosts[j].ID == feedInputs[i].StoryId {
				stories = append(stories, userPosts[j])
				break
			}
		}
	}
	return stories
}

func (repo *StoriesRepository) GetMyStories(id string) []model.Story {
	fmt.Println("Id je " + id)
	var stories []model.Story
	result, _ :=  repo.Database.Get(id).Result()
	bytes := []byte(result)
	json.Unmarshal(bytes, &stories)
	return stories
}

func (repo *StoriesRepository) AddToHighlights(highlight model.Highlight) error {
	result, err := repo.Database.Get(highlight.UserId.String() + "_highlight").Result()
	var highlights []model.Highlight
	if err == nil {
		bytes := []byte(result)
		err = json.Unmarshal(bytes, &highlights)
		if err != nil {
		return err
		}
		flag := true
		for i := range highlights {
			if highlights[i].StoryId == highlight.StoryId {
				flag = false
				break
			}
		}
		if !flag {
			return errors.New("already highlighted")
		}
	}
	highlights = append(highlights, highlight)
	jsonHighlights, _ := json.Marshal(highlights)
	newErr := repo.Database.Set(highlight.UserId.String() + "_highlight", jsonHighlights, 0).Err()
	if newErr != nil {
		return newErr
	}
	return nil
}


func (repo *StoriesRepository) RemoveFromHighlights(highlight model.Highlight) error {
	var userHighlights []model.Highlight
	var newHighlights []model.Highlight
	result, err := repo.Database.Get(highlight.UserId.String() + "_highlight").Result()
	if err != nil {
		fmt.Println("error")
		fmt.Println(err)
		return err
	}
	bytes := []byte(result)
	json.Unmarshal(bytes, &userHighlights)
	for i := range userHighlights {
		if userHighlights[i].StoryId != highlight.StoryId {
			newHighlights = append(newHighlights, userHighlights[i])
		}
	}
	err = repo.Database.Del(highlight.UserId.String() + "_highlight").Err()
	json, _ := json.Marshal(newHighlights)
	err = repo.Database.Set(highlight.UserId.String() + "_highlight", json, 0).Err()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (repo *StoriesRepository) GetHighlights(id string) []model.Story {
	fmt.Println("Id je " + id)
	var stories []model.Story
	var highlightInputs []model.Highlight
	result, _ :=  repo.Database.Get(id + "_highlight").Result()
	bytes := []byte(result)
	json.Unmarshal(bytes, &highlightInputs)
	for i := range highlightInputs {
		var userPosts []model.Story
		result, _ :=  repo.Database.Get(highlightInputs[i].UserId.String()).Result()
		bytes := []byte(result)
		json.Unmarshal(bytes, &userPosts)
		for j := range  userPosts {
			if userPosts[j].ID == highlightInputs[i].StoryId {
				stories = append(stories, userPosts[j])
				break
			}
		}
	}
	return stories
}

