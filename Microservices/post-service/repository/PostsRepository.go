package repository

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"post_service/model"
	"post_service/model/dto"
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

func(repo *PostsRepository) GetByKey(key string) []model.Post {
	fmt.Println("Key je " + key)
	var posts []model.Post
	result, _ :=  repo.Database.Get(key).Result()
	bytes := []byte(result)
	json.Unmarshal(bytes, &posts)
	return posts

}

func(repo *PostsRepository) GetPublic(keys []string) []model.Post {
	var posts []model.Post
	for i := range keys {
		posts = append(posts, repo.GetByKey(keys[i])...)
	}
	return posts
}

func(repo *PostsRepository) LikePost(likeReq dto.LikeDto) error {
	var posts []model.Post
	result, _ :=  repo.Database.Get(likeReq.OWNERID.String()).Result()
	bytes := []byte(result)
	err := json.Unmarshal(bytes, &posts)
	if err != nil {
		return err
	}

	for i := range posts {
		flag := true
		if posts[i].ID == likeReq.POSTID {
			for j := range posts[i].DISLIKES {
				if posts[i].DISLIKES[j].UserID == likeReq.USERID {
					posts[i].DISLIKES = removeDislike(posts[i].DISLIKES, j)
				}
			}
			for k := range posts[i].LIKES {
				if posts[i].LIKES[k].UserID == likeReq.USERID {
					posts[i].LIKES = removeLike(posts[i].LIKES, k)
					flag = false
				}
			}
			if flag {
				posts[i].LIKES = append(posts[i].LIKES, model.Like{likeReq.USERID})
			}
			break
		}
	}
	jsonPosts, _ := json.Marshal(posts)
	newErr := repo.Database.Set(likeReq.OWNERID.String(), jsonPosts, 0).Err()
	if newErr != nil {
		return newErr
	}
	return err
}

func(repo *PostsRepository) DislikePost(dislikeReq dto.LikeDto) error {
	var posts []model.Post
	result, _ :=  repo.Database.Get(dislikeReq.OWNERID.String()).Result()
	bytes := []byte(result)
	err := json.Unmarshal(bytes, &posts)
	if err != nil {
		return err
	}

	for i := range posts {
		flag := true
		if posts[i].ID == dislikeReq.POSTID {
			for j := range posts[i].DISLIKES {
				if posts[i].DISLIKES[j].UserID == dislikeReq.USERID {
					posts[i].DISLIKES = removeDislike(posts[i].DISLIKES, j)
					flag = false
				}
			}
			for k := range posts[i].LIKES {
				if posts[i].LIKES[k].UserID == dislikeReq.USERID {
					posts[i].LIKES = removeLike(posts[i].LIKES, k)
				}
			}
			if flag {
				posts[i].DISLIKES = append(posts[i].DISLIKES, model.Dislike{dislikeReq.USERID})
			}
			break
		}
	}
	jsonPosts, _ := json.Marshal(posts)
	newErr := repo.Database.Set(dislikeReq.OWNERID.String(), jsonPosts, 0).Err()
	if newErr != nil {
		return newErr
	}
	return err
}

func removeDislike(slice []model.Dislike, s int) []model.Dislike {
	return append(slice[:s], slice[s+1:]...)
}

func removeLike(slice []model.Like, s int) []model.Like {
	return append(slice[:s], slice[s+1:]...)
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