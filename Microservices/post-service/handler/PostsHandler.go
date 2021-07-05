package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"post_service/model"
	"post_service/model/dto"
	"post_service/service"
	"strings"
	"time"
)

type PostsHandler struct {
	Service *service.PostsService
}


func (handler *PostsHandler) Create(w http.ResponseWriter, r *http.Request) {
	var post model.Post
	fmt.Println(json.NewDecoder(r.Body).Decode(&post))
	err := json.NewDecoder(r.Body).Decode(&post)
	post.ID = uuid.New()
	fmt.Println("*************************************\n\n\n\n")
	fmt.Println(post.ID.String())
	fmt.Println("\n\n\n\n*************************************")
	post.TIMESTAMP = time.Now()
	post.COMMENTS = []model.Comment{}
	post.LIKES = []model.Like{}
	post.DISLIKES = []model.Dislike{}
	post.REPORTS = []model.ReportedBy{}
	err = handler.Service.Create(&post)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}
	rest, err := http.Get("http://localhost:8088/users/getFollowers/" + post.USERID.String())
	//rest, err := http.Get("https://mocki.io/v1/84324533-ee57-4eb2-8042-3f5845dcc41b")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var followersDto dto.FollowersDto
	err = json.NewDecoder(rest.Body).Decode(&followersDto)
	fmt.Println(followersDto.KEYS)
	err = handler.Service.AddPostToFeed(followersDto.KEYS, &post)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	respSubs, errSubs  := http.Get("http://localhost:8087/users/subscribers/" + post.USERID.String())

	if errSubs == nil {
		var subsIds [] string
		err = json.NewDecoder(respSubs.Body).Decode(&subsIds)
		if err == nil {
			for i := range subsIds {
				notReq := dto.NotificationRequestDto{USERID: post.USERID.String(), NOTIFYID: post.ID.String(), NOTIFYUSERID: subsIds[i], TYPEOFNOTIFY: "post"}
				payloadBuf := new(bytes.Buffer)
				json.NewEncoder(payloadBuf).Encode(notReq)
				_, _ = http.Post("http://localhost:8085/notify/create", "application/json", payloadBuf)
			}
		}
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
func (handler *PostsHandler) CreateCampaign(w http.ResponseWriter, r *http.Request) {
	var campaign model.Campaign
	fmt.Println(json.NewDecoder(r.Body).Decode(&campaign))
	err := json.NewDecoder(r.Body).Decode(&campaign)
	fmt.Println(err)
	campaign.ID = uuid.New()
	campaign.INFLUENCERS = []uuid.UUID{}
	campaign.COMMENTS = []model.Comment{}
	campaign.LIKES = []model.Like{}
	campaign.DISLIKES = []model.Dislike{}
	for i := range campaign.ADS {
		campaign.ADS[i].ID = uuid.New()
	}
	fmt.Println(campaign)
	err = handler.Service.CreateCampaign(&campaign)
	fmt.Println("handler")
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *PostsHandler) UpdateCampaign(w http.ResponseWriter, r *http.Request) {
	var campaign model.Campaign
	fmt.Println(json.NewDecoder(r.Body).Decode(&campaign))
	err := json.NewDecoder(r.Body).Decode(&campaign)
	fmt.Println(err)

	fmt.Println(campaign)
	existingCampaign := handler.Service.GetCampaignsByIds(campaign.USERID.String(), campaign.ID.String())
	newCampaign := existingCampaign

	newCampaign.DESCRIPTION = campaign.DESCRIPTION
	newCampaign.STARTDAY = campaign.STARTDAY
	newCampaign.ENDDAY = campaign.ENDDAY
	newCampaign.TIMESTOPLACE = campaign.TIMESTOPLACE
	newCampaign.WHENTOPLACE = campaign.WHENTOPLACE
	newCampaign.SHOWUNDER18 = campaign.SHOWUNDER18
	newCampaign.SHOW18TO24 = campaign.SHOW18TO24
	newCampaign.SHOW24TO35 = campaign.SHOW24TO35
	newCampaign.SHOWOVER35 = campaign.SHOWOVER35
	newCampaign.SHOWTOMEN = campaign.SHOWTOMEN
	newCampaign.SHOWTOWOMEN = campaign.SHOWTOWOMEN
	dto := dto.DeletePostDto{
		OWNERID: existingCampaign.USERID,
		POSTID:  existingCampaign.ID,
	}
	handler.Service.DeleteCampaign(&dto)
	err = handler.Service.CreateCampaign(&newCampaign)
	err1 := handler.Service.CreateTemporaryCampaign(&existingCampaign)

	fmt.Println("handler")
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}
	if err1 != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}
	//rest, err := http.Get("http://localhost:8088/users/getFollowers/" + post.USERID.String())
	//rest, err := http.Get("https://mocki.io/v1/84324533-ee57-4eb2-8042-3f5845dcc41b")

	//if err != nil {
	//	w.WriteHeader(http.StatusBadRequest)
	//	return
	//}
	//
	//var dto dto.FollowersDto
	//err = json.NewDecoder(rest.Body).Decode(&dto)
	//fmt.Println(dto.KEYS)
	//err = handler.Service.AddPostToFeed(dto.KEYS, &post)
	//if err != nil {
	//	fmt.Println(err)
	//	w.WriteHeader(http.StatusBadRequest)
	//}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}





func (postsHandler *PostsHandler) UploadFile(w http.ResponseWriter, r *http.Request) {
	var postItems[] model.PostItem
	// Maximum upload of 10 MB files
	r.ParseMultipartForm(32 << 20) // 32MB is the default used by FormFile
	fhs := r.MultipartForm.File["files"]
	var i int
	i = 0
	for _, fh := range fhs {
		i = i + 1
		f, err := fh.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		fmt.Printf("Uploaded File: %+v\n", fh.Filename)
		fmt.Printf("File Size: %+v\n", fh.Size)
		fmt.Printf("MIME Header: %+v\n", fh.Header)
		fileName := strings.Split(fh.Filename, ".")
		var filePath string
		var resourceName string
		if(len(fileName) >= 2){
			resourceName = uuid.NewString() +  "." + fileName[1]
			filePath = filepath.Join("user_posts", resourceName)
		}else{
			filePath = filepath.Join("user_posts", fh.Filename)
		}
		dst, err := os.Create(filePath)
		defer dst.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Copy the uploaded file to the created file on the filesystem
		if _, err := io.Copy(dst, f); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var postItem model.PostItem
		postItem.PATH = resourceName
		postItem.ID = uuid.New()
		postItems = append(postItems, postItem)
		defer f.Close()

	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	renderJSON(w, &postItems)
}

func (handler *PostsHandler) GetByKey(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	fmt.Println(vars["key"])
	posts :=handler.Service.GetByKey(vars["key"])
	renderJSON(w, &posts)
}

func (handler *PostsHandler) LikePost(w http.ResponseWriter, r *http.Request){
	var likeReq dto.LikeDto
	err := json.NewDecoder(r.Body).Decode(&likeReq)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	err = handler.Service.LikePost(likeReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *PostsHandler) DislikePost(w http.ResponseWriter, r *http.Request){
	var dislikeReq dto.LikeDto
	err := json.NewDecoder(r.Body).Decode(&dislikeReq)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	err = handler.Service.DislikePost(dislikeReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
func (handler *PostsHandler) ReportPost(w http.ResponseWriter, r *http.Request){
	var reportReq dto.ReportDto
	err := json.NewDecoder(r.Body).Decode(&reportReq)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	fmt.Println(reportReq.POSTID)
	fmt.Println(reportReq.USERID)
	err = handler.Service.ReportPost(reportReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}


func (handler *PostsHandler) GetLiked(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	fmt.Println(vars["id"])
	post :=handler.Service.GetLiked(vars["id"])
	renderJSON(w, &post)
}

func (handler *PostsHandler) GetCampaigns(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	fmt.Println(vars["id"])
	post :=handler.Service.GetCampaigns(vars["id"])
	renderJSON(w, &post)
}

func (handler *PostsHandler) GetTemporaryCampaigns(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	fmt.Println(vars["id"])
	post :=handler.Service.GetTemporaryCampaigns(vars["id"])
	renderJSON(w, &post)
}
func (handler *PostsHandler) GetDisliked(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	fmt.Println(vars["id"])
	post :=handler.Service.GetDisliked(vars["id"])
	renderJSON(w, &post)
}

func (handler *PostsHandler) DeletePost(w http.ResponseWriter, r *http.Request) {

	var deletePost dto.DeletePostDto
	err := json.NewDecoder(r.Body).Decode(&deletePost)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	ret := handler.Service.Delete(&deletePost)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	renderJSON(w, &ret)
}


func (handler *PostsHandler) DeleteCampaign(w http.ResponseWriter, r *http.Request) {

	var deletePost dto.DeletePostDto
	err := json.NewDecoder(r.Body).Decode(&deletePost)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	ret := handler.Service.DeleteCampaign(&deletePost)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	renderJSON(w, &ret)
}






func (handler *PostsHandler) GetReported(w http.ResponseWriter, r *http.Request){
	rest, err := http.Get("http://localhost:8085/getIds")

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}


	var ids []dto.UserId
	err = json.NewDecoder(rest.Body).Decode(&ids)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(ids)

	reportedPosts := handler.Service.GetReported(ids)
	renderJSON(w, &reportedPosts)

}



func (handler *PostsHandler) GetFeed(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	fmt.Println(vars["id"])
	var posts []model.Post
	var firstFilter []model.Post
	var finalFilter []model.Post
	posts = handler.Service.GetFeed(vars["id"])

	// Filtrira postove mutovanih korisnika

	respMut, _  := http.Get("http://localhost:8087/users/muted/" + vars["id"])
	var mutedIds [] string
	errJson := json.NewDecoder(respMut.Body).Decode(&mutedIds)
	if errJson == nil {
		for i := range posts {
			isNotByMutedUser := true
			for j := range mutedIds {
				if posts[i].USERID.String() == mutedIds[j] {
					isNotByMutedUser = false
				}
			}
			if isNotByMutedUser {
				firstFilter = append(firstFilter, posts[i])
			}
		}
	} else {
		firstFilter = posts
	}

	// Filtrira postove mutovanih korisnika

	respBlock, _  := http.Get("http://localhost:8087/users/blocked/" + vars["id"])

	var blockedIds [] string
	errJson2 := json.NewDecoder(respBlock.Body).Decode(&blockedIds)
	if errJson2 == nil {
		for i := range firstFilter {
			isNotByBlockedUser := true
			for j := range blockedIds {
				if firstFilter[i].USERID.String() == blockedIds[j] {
					isNotByBlockedUser = false
				}
			}
			if isNotByBlockedUser {
				finalFilter = append(finalFilter, firstFilter[i])
			}
		}
	} else {
		finalFilter = firstFilter
	}

	renderJSON(w, &finalFilter)
}

func (handler *PostsHandler) GetPublic(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	rest, err := http.Get("http://localhost:8085/users/public-ids")
	var finalFilter []model.Post
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var dto dto.KeyValueListDto
	err = json.NewDecoder(rest.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(dto.KEYS)
	posts := handler.Service.GetPublic(dto.KEYS)
	if vars["id"] == "" {
		renderJSON(w, &posts)
		return
	}
	respBlock, _  := http.Get("http://localhost:8087/users/blocked/" + vars["id"])

	var blockedIds [] string
	errJson2 := json.NewDecoder(respBlock.Body).Decode(&blockedIds)
	if errJson2 == nil {
		for i := range posts {
			isNotByBlockedUser := true
			for j := range blockedIds {
				if finalFilter[i].USERID.String() == blockedIds[j] {
					isNotByBlockedUser = false
				}
			}
			if isNotByBlockedUser {
				finalFilter = append(finalFilter, posts[i])
			}
		}
	} else {
		finalFilter = posts
	}
	renderJSON(w, &finalFilter)
	return
}

func (handler *PostsHandler) GetAllTagsPublic(w http.ResponseWriter, r *http.Request){
	//public users ids
	rest, err := http.Get("http://localhost:8085/users/public-ids")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var dtoIds dto.KeyValueListDto
	err = json.NewDecoder(rest.Body).Decode(&dtoIds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(dtoIds.KEYS)
	posts := handler.Service.GetPublic(dtoIds.KEYS)

	var tags dto.KeyValueListDto

	for i := range posts {
		tags.KEYS = append(tags.KEYS, posts[i].HASHTAG)
	}

	renderJSON(w, &tags)
}

func (handler *PostsHandler) GetAllTagsSignedIn(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	fmt.Println(vars["id"])

	//public users ids
	rest, err := http.Get("http://localhost:8085/users/public-ids")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var dtoIds dto.KeyValueListDto
	err = json.NewDecoder(rest.Body).Decode(&dtoIds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(dtoIds.KEYS)


	restFollowers, _ := http.Get("http://localhost:8088/users/getFollowers/" + vars["id"])
	var dtoFollowers dto.FollowersDto
	err = json.NewDecoder(restFollowers.Body).Decode(&dtoFollowers)
	fmt.Println(dtoFollowers.KEYS)

	var idsList []string

	idsList = append(idsList, dtoIds.KEYS...)
	idsList = append(idsList, dtoFollowers.KEYS...)

	posts := handler.Service.GetPublic(idsList)

	var tags dto.KeyValueListDto

	for i := range posts {
		tags.KEYS = append(tags.KEYS, posts[i].HASHTAG)
	}

	renderJSON(w, &tags)
}

func (handler *PostsHandler) GetAllLocationsPublic(w http.ResponseWriter, r *http.Request){
	//public users ids
	rest, err := http.Get("http://localhost:8085/users/public-ids")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var dtoIds dto.KeyValueListDto
	err = json.NewDecoder(rest.Body).Decode(&dtoIds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(dtoIds.KEYS)
	posts := handler.Service.GetPublic(dtoIds.KEYS)

	var locations dto.KeyValueListDto

	for i := range posts {
		locations.KEYS = append(locations.KEYS, posts[i].LOCATION)
	}

	renderJSON(w, &locations)
}

func (handler *PostsHandler) GetAllLocationsSignedIn(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	fmt.Println(vars["id"])

	//public users ids
	rest, err := http.Get("http://localhost:8085/users/public-ids")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var dtoIds dto.KeyValueListDto
	err = json.NewDecoder(rest.Body).Decode(&dtoIds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(dtoIds.KEYS)


	restFollowers, _ := http.Get("http://localhost:8088/users/getFollowers/" + vars["id"])
	var dtoFollowers dto.FollowersDto
	err = json.NewDecoder(restFollowers.Body).Decode(&dtoFollowers)
	fmt.Println(dtoFollowers.KEYS)

	var idsList []string

	idsList = append(idsList, dtoIds.KEYS...)
	idsList = append(idsList, dtoFollowers.KEYS...)

	posts := handler.Service.GetPublic(idsList)

	var locations dto.KeyValueListDto

	for i := range posts {
		locations.KEYS = append(locations.KEYS, posts[i].LOCATION)
	}

	renderJSON(w, &locations)
}

func (handler *PostsHandler) LeaveComment(writer http.ResponseWriter, request *http.Request) {
	var commentDto dto.CommentDto
	err := json.NewDecoder(request.Body).Decode(&commentDto)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	comment := new(model.Comment)
	comment.USERID = commentDto.USERID
	comment.ID = uuid.New()
	comment.TIMESTAMP = time.Now()
	comment.VALUE = commentDto.COMMENT
	err = handler.Service.LeaveComment(commentDto.POSTID, commentDto.OWNERID, comment)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	return
}

func (handler *PostsHandler) GetByIds(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars["userid"])
	fmt.Println(vars["postid"])
	post :=handler.Service.GetByIds(vars["userid"], vars["postid"])
	if post == nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	renderJSON(w, &post)
}
func (handler *PostsHandler) GetCampaignsByIds(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars["userid"])
	fmt.Println(vars["campaignid"])
	campaigns :=handler.Service.GetCampaignsByIds(vars["userid"], vars["campaignid"])
	if len(campaigns.ADS) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	renderJSON(w, &campaigns)
}

func (handler *PostsHandler) GetByUserId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	posts :=handler.Service.GetByUserId(vars["userid"])
	if posts == nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	renderJSON(w, &posts)
}

func (handler *PostsHandler) Save(writer http.ResponseWriter, request *http.Request) {
	var savePostDto dto.SavePostDto
	err := json.NewDecoder(request.Body).Decode(&savePostDto)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	archivedPost := new(model.SavedPost)
	archivedPost.USERID = savePostDto.USERID
	archivedPost.OWNERID = savePostDto.OWNERID
	archivedPost.POSTID = savePostDto.POSTID
	collection := model.PostCollection{NAME: "Default"}
	archivedPost.COLLECTION = collection
	err = handler.Service.SavePost(archivedPost)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	return
}

func (handler *PostsHandler) GetAllArchived(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars["id"])
	var archivedPosts []model.SavedPost
	archivedPosts = handler.Service.GetAllArchived(vars["id"])
	renderJSON(w, &archivedPosts)
}

func (handler *PostsHandler) EditArchived(writer http.ResponseWriter, request *http.Request) {
	var editPost model.SavedPost
	err := json.NewDecoder(request.Body).Decode(&editPost)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.EditArchived(editPost)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	return
}