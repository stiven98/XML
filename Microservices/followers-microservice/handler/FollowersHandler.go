package handler

import (
	"encoding/json"
	"fmt"
	"followers-microservice/model"
	"followers-microservice/model/dto"
	"followers-microservice/service"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

type FollowersHandler struct {
	FollowersService *service.FollowersService
}

func (h FollowersHandler) GetFollowers(writer http.ResponseWriter, request *http.Request) {
	tokens := strings.Split(request.URL.Path, "/")
	ID := tokens[int(len(tokens))-1]

	// Check User exist in Profile service and return bad request if don't
	fmt.Println(ID)

	//retVal := h.FollowersService.UserExists(ID)

	//if retVal == nil {
	//	writer.WriteHeader(http.StatusBadRequest)
	//	return
	//}

	keys := h.FollowersService.GetFollowers(ID)

	if keys == nil {
		keys = []string{}
	}

	// retVal contains ids of followers

	var followers model.Followers
	followers.KEYS = keys

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	bytes, err := json.Marshal(followers)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.Write(bytes)
}

func (h FollowersHandler) GetFollowing(writer http.ResponseWriter, request *http.Request) {
	tokens := strings.Split(request.URL.Path, "/")
	ID := tokens[int(len(tokens))-1]

	// Check User exist in Profile service and return bad request if don't
	fmt.Println(ID)

	//retVal := h.FollowersService.UserExists(ID)
	//
	//if retVal == nil {
	//	writer.WriteHeader(http.StatusBadRequest)
	//	return
	//}

	keys := h.FollowersService.GetFollowing(ID)

	if keys == nil {
		keys = []string{}

		//	writer.WriteHeader(http.StatusBadRequest)
	//	return
	}

	var followers model.Followers
	followers.KEYS = keys

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	bytes, err := json.Marshal(followers)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.Write(bytes)
}

func (h FollowersHandler) Follow(writer http.ResponseWriter, request *http.Request) {
	tokens := strings.Split(request.URL.Path, "/")
	userID := tokens[int(len(tokens))-2]
	targetID := tokens[int(len(tokens))-1]
	retVal := h.FollowersService.UserExists(userID)
	fmt.Println(retVal)
	fmt.Println("posotji")
	if retVal == nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	retVal = h.FollowersService.UserExists(targetID)
	fmt.Println(retVal)
	fmt.Println("postoji")
	if retVal == nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	retVal = h.FollowersService.CheckRelationship(userID, targetID)
	fmt.Println(retVal)
	if retVal == true {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	response, err := http.Get("http://localhost:8085/users/isPublic/" + targetID)
	if err != nil {
		fmt.Println("Error in communication with profile microservice!")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	var userDTO dto.IsUserPublicDTO
	err = json.NewDecoder(response.Body).Decode(&userDTO)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	if userDTO.IsPublic == false {
		e := h.FollowersService.Request(userID, targetID)
		if e == nil {
			writer.WriteHeader(http.StatusOK)
			return
		}
	} else {
		e := h.FollowersService.Follow(userID, targetID)
		if e == nil {
			writer.WriteHeader(http.StatusOK)
			return
		}
	}

	writer.WriteHeader(http.StatusOK)

}

func (h FollowersHandler) IsFollowing(writer http.ResponseWriter, request *http.Request) {
	tokens := strings.Split(request.URL.Path, "/")
	userID := tokens[int(len(tokens))-2]
	targetID := tokens[int(len(tokens))-1]
	retVal:= h.FollowersService.CheckRelationship(userID, targetID)

	if retVal == true {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println(retVal)
	//retVal = h.FollowersService.UserExists(targetID)
	//
	//if retVal == nil {
	//	writer.WriteHeader(http.StatusBadRequest)
	//	return
	//}
	writer.WriteHeader(http.StatusOK)
}

func (h FollowersHandler) AddNode(writer http.ResponseWriter, request *http.Request) {

	vars :=mux.Vars(request)
	ID := vars["id"]
	fmt.Println(ID)
	if !IsValidUUID(ID) {
		fmt.Println("ID is invalid!")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	res := h.FollowersService.UserExists(ID)

	if res != nil {
		fmt.Println("User with ID already exists!")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	_ , err := h.FollowersService.AddNode(ID)
	if err != nil {
		fmt.Println("Error with base!")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusOK)

}

func (h FollowersHandler) GetRequest(writer http.ResponseWriter, request *http.Request) {
	tokens := strings.Split(request.URL.Path, "/")
	ID := tokens[int(len(tokens))-1]
	if !IsValidUUID(ID){
		fmt.Println("Invalid ID!")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	retVal := h.FollowersService.UserExists(ID)

	if retVal == nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	keys := h.FollowersService.GetRequests(ID)

	fmt.Println(len(keys))
	if keys == nil {
		keys = []string{}

		//	fmt.Println("DJOLE")
	//	writer.WriteHeader(http.StatusBadRequest)
	//	return
	}

	var followers model.Followers
	followers.KEYS = keys

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	bytes, err := json.Marshal(followers)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.Write(bytes)

}

func (h FollowersHandler) Unfollow(writer http.ResponseWriter, request *http.Request) {
	tokens := strings.Split(request.URL.Path, "/")
	userID := tokens[int(len(tokens))-2]
	targetID := tokens[int(len(tokens))-1]
	retVal := h.FollowersService.UserExists(userID)

	if retVal == nil {
		fmt.Println("User doesn't exists!")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	retVal = h.FollowersService.UserExists(targetID)

	if retVal == nil {
		fmt.Println("User doesn't exists!")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}


	e := h.FollowersService.Unfollow(userID, targetID)
	if e == nil {
		_, _ = http.Post("http://localhost:8087/users/unsubscribe/" + userID + "/" + targetID, "application/json", nil)
		writer.WriteHeader(http.StatusOK)
		return
	}
	writer.WriteHeader(http.StatusBadRequest)

}

func (h FollowersHandler) AcceptRequest(writer http.ResponseWriter, request *http.Request) {
	tokens := strings.Split(request.URL.Path, "/")
	userID := tokens[int(len(tokens))-2]
	targetID := tokens[int(len(tokens))-1]
	retVal := h.FollowersService.UserExists(userID)

	if retVal == nil {
		fmt.Println("User doesn't exists!")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	retVal = h.FollowersService.UserExists(targetID)

	if retVal == nil {
		fmt.Println("User doesn't exists!")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	retVal = h.FollowersService.CheckRelationship(userID, targetID)
	fmt.Println(retVal)
	if retVal != true {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	e := h.FollowersService.Unfollow(userID, targetID)
	if e != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	e = h.FollowersService.Follow(userID, targetID)
	if e != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusOK)


}

func (h FollowersHandler) CheckFollowing(writer http.ResponseWriter, request *http.Request) {
	tokens := strings.Split(request.URL.Path, "/")
	userID := tokens[int(len(tokens))-2]
	targetID := tokens[int(len(tokens))-1]
	retVal:= h.FollowersService.CheckRelationship(userID, targetID)

	followingDto := dto.CheckFollowingDTO{IsFollowing: retVal}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	bytes, err := json.Marshal(followingDto)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.Write(bytes)
}


