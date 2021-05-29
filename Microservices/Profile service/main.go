package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"profileservice/handler"
	"profileservice/model"
	"profileservice/repository"
	"profileservice/service"
	"time"
)

func initDB() *gorm.DB {
	host := "localhost"
	user := "postgres"
	password := "root"
	dbname := "users_service"
	dbport := "5432"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host, user, password, dbname, dbport)

	var database *gorm.DB

	log.Println("Connecting to database...")
	for 1 == 1 {
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			time.Sleep(2 * time.Second)
			log.Println("Reconnecting to database...")
			continue
		}

		database = db
		log.Println("Connected to database")
		break
	}

	database.Migrator().DropTable(&model.SystemUser{})
	database.AutoMigrate(&model.SystemUser{})


	return database

}

func initRepo(database *gorm.DB) (*repository.UsersRepository) {
	return &repository.UsersRepository{Database: database}
}

func initServices(usersRepo *repository.UsersRepository) (*service.UsersService) {
	return &service.UsersService{Repo: usersRepo}
}

func initHandler(usersService *service.UsersService) (*handler.UsersHandler) {
	return &handler.UsersHandler{Service: usersService}
}

func handleFunc(usersHandler *handler.UsersHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/users/create", usersHandler.Create).Methods("POST")
	router.HandleFunc("/users/update", usersHandler.Update).Methods("PUT")
	router.HandleFunc("/users/getAll",usersHandler.GetAll).Methods("GET")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8085"), router))
}

func main() {
	database := initDB()
	usersRepo := initRepo(database)
	usersService := initServices(usersRepo)
	usersHandler := initHandler(usersService)
	handleFunc(usersHandler)
}

