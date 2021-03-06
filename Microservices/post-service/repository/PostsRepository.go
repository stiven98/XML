package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"post_service/model"
	"post_service/model/dto"
	"time"
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
func(repo *PostsRepository) CreateCampaign(campaign *model.Campaign)  error {
	result, err :=  repo.Database.Get(campaign.USERID.String()+ "_campaign").Result()
	var campaigns[] model.Campaign
	if err != nil {
		campaigns = append(campaigns, *campaign)
		jsonPosts, _ := json.Marshal(campaigns)
		newErr := repo.Database.Set(campaign.USERID.String() + "_campaign", jsonPosts, 0).Err()
		if newErr != nil {
			fmt.Println(result)
		}
	} else{
		bytes := []byte(result)
		json.Unmarshal(bytes, &campaigns)
		campaigns = append(campaigns, *campaign)
		jsonPosts, _ := json.Marshal(campaigns)
		newErr := repo.Database.Set(campaign.USERID.String()+ "_campaign", jsonPosts, 0).Err()
		if newErr != nil {
			fmt.Println(result)
		}
	}
	return nil
}
func(repo *PostsRepository) CreateCampaignForInfluencer(campaign *model.Campaign)  error {
	result, err :=  repo.Database.Get(campaign.USERID.String()+ "_campaignInf").Result()
	var campaigns[] model.Campaign
	if err != nil {
		campaigns = append(campaigns, *campaign)
		jsonPosts, _ := json.Marshal(campaigns)
		newErr := repo.Database.Set(campaign.USERID.String() + "_campaignInf", jsonPosts, 0).Err()
		if newErr != nil {
			fmt.Println(result)
		}
	} else{
		bytes := []byte(result)
		json.Unmarshal(bytes, &campaigns)
		campaigns = append(campaigns, *campaign)
		jsonPosts, _ := json.Marshal(campaigns)
		newErr := repo.Database.Set(campaign.USERID.String()+ "_campaignInf", jsonPosts, 0).Err()
		if newErr != nil {
			fmt.Println(result)
		}
	}
	return nil
}

func(repo *PostsRepository) GetCampaignsInf(id string) []model.Campaign {

	var campaigns []model.Campaign
	result1, _ :=  repo.Database.Get(id+ "_campaignInf").Result()
	bytes1 := []byte(result1)
	json.Unmarshal(bytes1, &campaigns)


	return campaigns

}


func(repo *PostsRepository) CreateCampaignRequest(campaignReq *dto.CampaignRequestDto)  error {
	result, err :=  repo.Database.Get(campaignReq.INFLUENCERID.String()+ "_campaignReq").Result()
	var campaignReqs[] dto.CampaignRequestDto
	if err != nil {
		campaignReqs = append(campaignReqs, *campaignReq)
		jsonPosts, _ := json.Marshal(campaignReqs)
		newErr := repo.Database.Set(campaignReq.INFLUENCERID.String() + "_campaignReq", jsonPosts, 0).Err()
		if newErr != nil {
			fmt.Println(result)
		}
	} else{
		bytes := []byte(result)
		json.Unmarshal(bytes, &campaignReqs)
		campaignReqs = append(campaignReqs, *campaignReq)
		jsonPosts, _ := json.Marshal(campaignReqs)
		newErr := repo.Database.Set(campaignReq.INFLUENCERID.String()+ "_campaignReq", jsonPosts, 0).Err()
		if newErr != nil {
			fmt.Println(result)
		}
	}
	return nil
}

func(repo *PostsRepository) CreateTemporaryCampaign(campaign *model.Campaign)  error {
	result, err :=  repo.Database.Get(campaign.USERID.String()+ "_campaignTemp").Result()
	var campaigns[] model.Campaign
	if err != nil {
		campaigns = append(campaigns, *campaign)
		jsonPosts, _ := json.Marshal(campaigns)
		newErr := repo.Database.Set(campaign.USERID.String() + "_campaignTemp", jsonPosts, 50 * time.Second).Err()
		if newErr != nil {
			fmt.Println(result)
		}
	} else{
		bytes := []byte(result)
		json.Unmarshal(bytes, &campaigns)
		campaigns = append(campaigns, *campaign)
		jsonPosts, _ := json.Marshal(campaigns)
		newErr := repo.Database.Set(campaign.USERID.String()+ "_campaignTemp", jsonPosts, 50 * time.Second).Err()
		if newErr != nil {
			fmt.Println(result)
		}
	}
	return nil
}
func(repo *PostsRepository) GetCampaigns(id string) []model.Campaign {
	fmt.Println("Id je " + id)
	var tempCampaigns []model.Campaign
	result, _ :=  repo.Database.Get(id+ "_campaignTemp").Result()
	bytes := []byte(result)
	json.Unmarshal(bytes, &tempCampaigns)

	var campaigns []model.Campaign
	result1, _ :=  repo.Database.Get(id+ "_campaign").Result()
	bytes1 := []byte(result1)
	json.Unmarshal(bytes1, &campaigns)



	for i:= range campaigns {
		for j:= range tempCampaigns {
			if campaigns[i].ID == tempCampaigns[j].ID {
				fmt.Println("usao")
				fmt.Println(campaigns[i].DESCRIPTION)
				campaigns[i] = tempCampaigns[j]
				fmt.Println(campaigns[i].DESCRIPTION)
			}
		}
	}

	return campaigns

}

func(repo *PostsRepository) GetCampaignReqs(id string) []dto.CampaignRequestDto {
	var campaignReqs []dto.CampaignRequestDto
	result, _ :=  repo.Database.Get(id+ "_campaignReq").Result()
	bytes := []byte(result)
	json.Unmarshal(bytes, &campaignReqs)

	return campaignReqs

}

func(repo *PostsRepository) GetTemporaryCampaigns(id string) []model.Campaign {
	fmt.Println("Id je " + id)
	var campaigns []model.Campaign
	result, _ :=  repo.Database.Get(id+ "_campaignTemp").Result()
	bytes := []byte(result)
	json.Unmarshal(bytes, &campaigns)
	for i := range campaigns {
		var userPosts []model.Campaign
		result, _ := repo.Database.Get(campaigns[i].USERID.String() + "_campaignTemp").Result()
		bytes := []byte(result)
		json.Unmarshal(bytes, &userPosts)
	}
	return campaigns

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
func(repo *PostsRepository) ReportPost(reportReq dto.ReportDto) error {
	var posts []model.Post
	result, _ :=  repo.Database.Get(reportReq.OWNERID.String()).Result()
	bytes := []byte(result)
	err := json.Unmarshal(bytes, &posts)
	if err != nil {
		return err
	}

	for i := range posts {
		if posts[i].ID == reportReq.POSTID {
			posts[i].REPORTS = append(posts[i].REPORTS, model.ReportedBy{reportReq.USERID})
			break
		}
	}
	jsonPosts, _ := json.Marshal(posts)
	newErr := repo.Database.Set(reportReq.OWNERID.String(), jsonPosts, 0).Err()
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
func(repo *PostsRepository) GetLiked(id string) []model.Post {
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
				for k := range userPosts[j].LIKES {
					if userPosts[j].LIKES[k].UserID.String() == id {
						posts = append(posts, userPosts[j])
					}
				}
				break
			}
		}
	}
	return posts

}



func(repo *PostsRepository) GetDisliked(id string) []model.Post {
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
				for k := range userPosts[j].DISLIKES {
					if userPosts[j].DISLIKES[k].UserID.String() == id {
						posts = append(posts, userPosts[j])
					}
				}
				break
			}
		}
	}
	return posts

}
func (repo *PostsRepository) Delete (deletePost *dto.DeletePostDto) bool {
	var userPosts []model.Post
	var newPosts []model.Post
	result, err := repo.Database.Get(deletePost.OWNERID.String()).Result()
	if err!=nil {
		fmt.Println("error")
		fmt.Println(err)
		return false
	}
	bytes := []byte(result)
	json.Unmarshal(bytes, &userPosts)
	for i := range userPosts {
		if userPosts[i].ID != deletePost.POSTID {
			newPosts = append(newPosts, userPosts[i])
		}
	}
	err = repo.Database.Del(deletePost.OWNERID.String()).Err()
	json, _ := json.Marshal(newPosts)
	err = repo.Database.Set(deletePost.OWNERID.String(), json, 0).Err()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func (repo *PostsRepository) DeleteCampaign (deletePost *dto.DeletePostDto) bool {
	var campaigns []model.Campaign
	var newCampaigns []model.Campaign
	result, err := repo.Database.Get(deletePost.OWNERID.String() + "_campaign").Result()
	if err!=nil {
		fmt.Println("error")
		fmt.Println(err)
		return false
	}
	bytes := []byte(result)
	json.Unmarshal(bytes, &campaigns)
	for i := range campaigns {
		if campaigns[i].ID != deletePost.POSTID {
			newCampaigns = append(newCampaigns, campaigns[i])
		}
	}
	err = repo.Database.Del(deletePost.OWNERID.String() + "_campaign").Err()
	json, _ := json.Marshal(newCampaigns)
	err = repo.Database.Set(deletePost.OWNERID.String() + "_campaign", json, 0).Err()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func (repo *PostsRepository) DeleteCampaignReq (deleteReq *dto.DeleteCampaignReq) bool {
	var campaignReqs []dto.CampaignRequestDto
	var newReqs []dto.CampaignRequestDto
	result, err := repo.Database.Get(deleteReq.OWNERID.String() + "_campaignReq").Result()
	if err!=nil {
		fmt.Println("error")
		fmt.Println(err)
		return false
	}
	bytes := []byte(result)
	json.Unmarshal(bytes, &campaignReqs)
	for i := range campaignReqs {
		if campaignReqs[i].ID != deleteReq.CAMPAIGNREQID {
			newReqs = append(newReqs, campaignReqs[i])
		}
	}
	err = repo.Database.Del(deleteReq.OWNERID.String() + "_campaignReq").Err()
	json, _ := json.Marshal(newReqs)
	err = repo.Database.Set(deleteReq.OWNERID.String() + "_campaignReq", json, 0).Err()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}


func(repo *PostsRepository) GetReported(ids []dto.UserId) ([]model.Post) {
	var reportedPosts []model.Post
	var userPosts []model.Post
	for i := range  ids {
		result, err := repo.Database.Get(ids[i].Id).Result()
		if err != nil {
			fmt.Println(err)
			continue
		}
		bytes := []byte(result)
		json.Unmarshal(bytes, &userPosts)
		for j:= range userPosts {
			if len(userPosts[j].REPORTS) > 0 {
				reportedPosts = append(reportedPosts, userPosts[j])
			}
		}
	}
	return reportedPosts
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

func (repo *PostsRepository) LeaveComment(postId uuid.UUID, ownerId uuid.UUID, comment *model.Comment) error {
	var posts []model.Post
	result, _ :=  repo.Database.Get(ownerId.String()).Result()
	bytes := []byte(result)
	err := json.Unmarshal(bytes, &posts)
	if err != nil {
		return err
	}

	for i := range posts {
		if posts[i].ID == postId {
			posts[i].COMMENTS = append(posts[i].COMMENTS, *comment)
			break
		}
	}
	jsonPosts, _ := json.Marshal(posts)
	newErr := repo.Database.Set(ownerId.String(), jsonPosts, 0).Err()
	if newErr != nil {
		return newErr
	}
	return err
}

func (repo *PostsRepository) GetByIds(userid string, postid string) interface{} {
	var posts []model.Post
	var post model.Post
	result, _ :=  repo.Database.Get(userid).Result()
	bytes := []byte(result)
	json.Unmarshal(bytes, &posts)
	for i := range posts {
		if posts[i].ID.String() == postid {
			post = posts[i]
			break
		}
	}
	return post
}
func (repo *PostsRepository) GetCampaignsByIds(userid string, campaignid string) model.Campaign {
	var tempCampaigns []model.Campaign
	var tempCampaign model.Campaign
	result1, _ :=  repo.Database.Get(userid+ "_campaignTemp").Result()
	bytes1 := []byte(result1)
	json.Unmarshal(bytes1, &tempCampaigns)
	for i := range tempCampaigns {
		if tempCampaigns[i].ID.String() == campaignid {
			tempCampaign = tempCampaigns[i]
			break
		}
	}
	var campaings []model.Campaign
	var campaing model.Campaign
	result, _ :=  repo.Database.Get(userid+ "_campaign").Result()
	bytes := []byte(result)
	json.Unmarshal(bytes, &campaings)
	for i := range campaings {
		if campaings[i].ID.String() == campaignid {
			campaing = campaings[i]
			break
		}
	}
	if campaing.ID == tempCampaign.ID {
		campaing = tempCampaign
	}

	return campaing
}
func (repo *PostsRepository) GetCampaignsByInfluencerIds(userid string, campaignid string) model.Campaign {
	var campaings []model.Campaign
	var campaing model.Campaign
	result, _ :=  repo.Database.Get(userid+ "_campaignInf").Result()
	bytes := []byte(result)
	json.Unmarshal(bytes, &campaings)
	for i := range campaings {
		if campaings[i].ID.String() == campaignid {
			campaing = campaings[i]
			break
		}
	}


	return campaing
}





func (repo *PostsRepository) GetByUserId(userid string) interface{} {
	var posts []model.Post
	result, _ :=  repo.Database.Get(userid).Result()
	bytes := []byte(result)
	json.Unmarshal(bytes, &posts)
	return posts
}

func (repo *PostsRepository) SavePost(post *model.SavedPost) error {
	results, err := repo.Database.Get(post.USERID.String() + "_archive").Result()
	var archivedPosts []model.SavedPost
	if err != nil {
		archivedPosts = append(archivedPosts, *post)
	} else {
		bytes := []byte(results)
		err = json.Unmarshal(bytes, &archivedPosts)
		flag := true
		for i := range archivedPosts {
			if archivedPosts[i].POSTID == post.POSTID {
				flag = false
			}
		}
		if flag {
			archivedPosts = append(archivedPosts, *post)
		} else {
			return errors.New("Post already archived")
		}
	}
	jsonPosts, _ := json.Marshal(archivedPosts)
	newErr := repo.Database.Set(post.USERID.String() + "_archive", jsonPosts, 0).Err()
	return newErr
}

func (repo *PostsRepository) GetAllArchived(id string) []model.SavedPost {
	var archivedPosts []model.SavedPost
	results, err := repo.Database.Get(id + "_archive").Result()
	fmt.Println(results)
	if err == nil {
		bytes := []byte(results)
		err = json.Unmarshal(bytes, &archivedPosts)
	}
	return archivedPosts
}

func (repo *PostsRepository) EditArchived(post model.SavedPost) error {
	var archivedPosts []model.SavedPost
	result, _ :=  repo.Database.Get(post.USERID.String() + "_archive").Result()
	bytes := []byte(result)
	err := json.Unmarshal(bytes, &archivedPosts)
	if err != nil {
		return err
	}

	for i := range archivedPosts {
		if archivedPosts[i].POSTID == post.POSTID {
			archivedPosts[i].COLLECTION = post.COLLECTION
			break
		}
	}
	jsonPosts, _ := json.Marshal(archivedPosts)
	newErr := repo.Database.Set(post.USERID.String() + "_archive", jsonPosts, 0).Err()
	if newErr != nil {
		return newErr
	}
	return err
}