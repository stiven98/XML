package handler

import (
	"fmt"
	"followers-microservice/service"
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

	retVal := h.FollowersService.UserExists(ID)

	if retVal == nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	retVal = h.FollowersService.GetFollowers(ID)

	if retVal == nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	// retVal contains ids of followers


	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")

}

func (h FollowersHandler) GetFollowing(writer http.ResponseWriter, request *http.Request) {
	tokens := strings.Split(request.URL.Path, "/")
	ID := tokens[int(len(tokens))-1]

	// Check User exist in Profile service and return bad request if don't
	fmt.Println(ID)

	retVal := h.FollowersService.UserExists(ID)

	if retVal == nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	retVal = h.FollowersService.GetFollowing(ID)

	if retVal == nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	// retVal contains ids of followers


	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
}

func (h FollowersHandler) Follow(writer http.ResponseWriter, request *http.Request) {
	tokens := strings.Split(request.URL.Path, "/")
	userID := tokens[int(len(tokens))-2]
	targetID := tokens[int(len(tokens))-1]
	retVal := h.FollowersService.UserExists(userID)

	if retVal == nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	retVal = h.FollowersService.UserExists(targetID)

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

	e := h.FollowersService.Follow(userID, targetID)
	if e == nil {
		writer.WriteHeader(http.StatusOK)
		return
	}


	writer.WriteHeader(http.StatusBadRequest)


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
