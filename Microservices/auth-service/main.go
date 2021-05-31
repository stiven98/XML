package main

import (
	"auth-service/handler"
	"auth-service/model"
	"auth-service/repository"
	"auth-service/service"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"time"
)

func initDB() *gorm.DB {
	hostName := os.Getenv("HOST_NAME")
	host := "localhost"
	if len(hostName) != 0 {
		host = hostName
	}
	user := "postgres"
	password := "root"
	dbname := "auth_service"
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
	if database != nil {
		database.Migrator().DropTable(&model.LoginDetails{})

		database.AutoMigrate(&model.LoginDetails{})

		loginDetailsList := [] model.LoginDetails {
			{
				ID:          uuid.New(),
				Email:       "aca@gmail.com",
				Password:    "aca12345",
			}, {
				ID:          uuid.New(),
				Email:       "jovan@gmail.com",
				Password:    "jovan123",
			}, {
				ID:          uuid.New(),
				Email:       "djordjije@gmail.com",
				Password:    "djordjije123",
			}, {
				ID:          uuid.New(),
				Email:       "aleksandar@gmail.com",
				Password:    "aleksandar123",
			},
		}
		for i := range loginDetailsList {
			database.Create(&loginDetailsList[i])
		}
	}
	return database
}

func initRepo(database *gorm.DB) *repository.LoginDetailsRepository {
	return &repository.LoginDetailsRepository{Database: database}
}



func initServices(loginDetailsRepository *repository.LoginDetailsRepository) *service.LoginDetailsService {
	return &service.LoginDetailsService{LoginDetailsRepository: loginDetailsRepository}

}



func initHandler(loginDetailsService *service.LoginDetailsService) *handler.LoginDetailsHandler {
	return &handler.LoginDetailsHandler{Service: loginDetailsService}
}

func handleFunc(LoginDetailsHandler *handler.LoginDetailsHandler) {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/login/create", LoginDetailsHandler.Create).Methods("POST")
	router.HandleFunc("/login/update", LoginDetailsHandler.Update).Methods("PUT")
	router.HandleFunc("/login/getAll", LoginDetailsHandler.GetAll).Methods("GET")
	router.HandleFunc("/login/get/{email}", LoginDetailsHandler.GetByEmail).Methods("GET")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8088"), router))
}

func main() {
	database := initDB()
	loginDetailsRepo := initRepo(database)
	loginDetailsService := initServices(loginDetailsRepo)
	loginDetailsHandler := initHandler(loginDetailsService)
	handleFunc(loginDetailsHandler)
}
