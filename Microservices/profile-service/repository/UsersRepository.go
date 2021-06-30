package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	_ "github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
	"profileservice/model"
	"profileservice/model/Dto"
)

type UsersRepository struct {
	Database *gorm.DB
}

func (repo *UsersRepository) Update(user *model.User) error {
	fmt.Println(user)
	result := repo.Database.Model(&model.User{}).Where("user_id = ?", user.UserID).Updates(map[string]interface{}{
				"IsPublic": user.IsPublic,
				"AllowedTags": user.AllowedTags,
				"IsBlocked": user.IsBlocked,
				"AcceptMessagesFromNotFollowProfile" : user.AcceptMessagesFromNotFollowProfile,
				"SystemUser": user.SystemUser,
				"PhoneNumber": user.PhoneNumber,
				"WebSite": user.WebSite,
				"Biography": user.Biography,
				"NotifyPosts" : user.NotifyPosts,
				"NotifyMessages": user.NotifyMessages,
				"NotifyStory": user.NotifyStory,
				"NotifyComments": user.NotifyComments,
		})
	return result.Error
}
func(repo *UsersRepository) GetAll() []model.User{
	var users []model.User
	repo.Database.Preload("SystemUser").Find(&users)
	return users
}

func(repo *UsersRepository) GetAllPublic() []model.User{
	var users []model.User
	repo.Database.Preload("SystemUser").Where("is_public = 'true'").Find(&users)
	return users
}


func (repo *UsersRepository) Create(user *model.User) error {
	result := repo.Database.Create(user)
	var dto = Dto.CreateUserDTO{
		ID:       user.UserID,
		USERNAME: user.SystemUser.Username,
		PASSWORD: user.SystemUser.Password,
		ACTIVE:   true,
		ROLE:     "ROLE_SYSTEM_USER",
	}
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(dto)
	payloadBuf1 := new(bytes.Buffer)
	json.NewEncoder(payloadBuf1).Encode(dto.ID)
	_, err := http.Post("http://localhost:8080/api/createUser","application/json", payloadBuf)
	_, err1 := http.Post("http://localhost:8088/users/addNode/" + dto.ID.String(),"application/json", payloadBuf1)

	if err != nil {
		fmt.Println(err)
		return err
	}
	if err1 != nil {
		fmt.Println(err1)
		return err1
	}
	fmt.Println(result.RowsAffected)
	return result.Error
}
func (repo *UsersRepository) ChangeWhetherIsPublic(dto *Dto.ChangeWhetherIsPublicDto) error {
	result := repo.Database.Model(model.User{}).Where("user_id = ?", dto.USERID).UpdateColumn("is_public", dto.FLAG)
	return result.Error
}

func (repo *UsersRepository) ChangeAllowedTags(dto *Dto.ChangeAllowedTagsDto) error {
	result := repo.Database.Model(model.User{}).Where("user_id = ?", dto.USERID).UpdateColumn("allowed_tags", dto.FLAG)
	return result.Error
}

func (repo *UsersRepository) GetById(id string) (model.User, error) {
	var user model.User
	response := repo.Database.Preload("SystemUser").Find(&user, "user_id = ?", id)
	return user, response.Error
}
func (repo *UsersRepository) GetIds() ([]Dto.UserId, error) {
	var ids []Dto.UserId
	var users []model.User
	response:= repo.Database.Preload("SystemUser").Find(&users)

	for i := range users {
		var id Dto.UserId
		id.Id = users[i].UserID.String()
		ids = append(ids, id)
	}

	return ids, response.Error
}