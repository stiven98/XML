package main

import (
	"fmt"
	"github.com/google/uuid"
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
	host := "db"
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
	if database != nil {
		database.Migrator().DropTable(&model.SystemUser{})
		database.Migrator().DropTable(&model.Administrator{})
		database.Migrator().DropTable(&model.User{})
		database.Migrator().DropTable(&model.Agent{})

		database.AutoMigrate(&model.SystemUser{})
		database.AutoMigrate(&model.Administrator{})
		database.AutoMigrate(&model.User{})
		database.AutoMigrate(&model.Agent{})

		systemUsers := [] model.SystemUser {
			{
				ID:          uuid.UUID{},
				FirstName:   "Aca",
				LastName:    "Simic",
				Username:    "acasimic",
				Email:       "aca@gmail.com",
				Password:    "aca12345",
				Gender:      model.MALE,
				TypeOfUser: model.ADMIN,
				DateOfBirth: time.Time{}.AddDate(1998, 10, 1),
			}, {
				ID:          uuid.UUID{},
				FirstName:   "Jovan",
				LastName:    "Bosnic",
				Username:    "jovanbosnic",
				Email:       "jovan@gmail.com",
				Password:    "jovan123",
				Gender:      model.MALE,
				TypeOfUser: model.ADMIN,
				DateOfBirth: time.Time{}.AddDate(1998, 7, 31),
			}, {
				ID:          uuid.UUID{},
				FirstName:   "Djordjije",
				LastName:    "Kundacina",
				Username:    "djordjijekundacina",
				Email:       "djordjije@gmail.com",
				Password:    "djordjije123",
				Gender: 	 model.MALE,
				TypeOfUser: model.USER,
				DateOfBirth: time.Time{}.AddDate(1998, 9, 10),
			}, {
				ID:          uuid.UUID{},
				FirstName:   "Aleksandar",
				LastName:    "Stevanovic",
				Username:    "aleksandarstevanovic",
				Email:       "aleksandar@gmail.com",
				Password:    "aleksandar123",
				Gender:      model.MALE,
				TypeOfUser: model.USER,
				DateOfBirth: time.Time{}.AddDate(1998, 10, 10),
			},
		}

		administrators := []model.Administrator {
			{
				UserID: systemUsers[0].ID,
				SystemUser: systemUsers[0],
			}, {
				UserID: systemUsers[1].ID,
				SystemUser: systemUsers[1],
			},
		}

		users := [] model.User {
			{
				UserID:      systemUsers[2].ID,
				SystemUser:  systemUsers[2],
				IsPublic:    true,
				PhoneNumber: "00381-44-44-44-221",
				WebSite:     "www.qndaa.com",
				Biography:   "Zanimljiv decak!",
				AllowedTags: true,
				IsBlocked:   false,
			}, {
				UserID:      systemUsers[3].ID,
				SystemUser:  systemUsers[3],
				IsPublic:    false,
				PhoneNumber: "00381-22-22-22-3333",
				WebSite:     "www.stiven.com",
				Biography:   "Berem jagode!",
				AllowedTags: true,
				IsBlocked:   false,
			},
		}

		for i := range administrators {
			database.Create(&administrators[i])
		}
		for i := range users {
			database.Create(&users[i])
		}
	}
	return database
}

func initRepo(database *gorm.DB) (*repository.SystemUsersRepository,
								  *repository.AdministratorsRepository,
								  *repository.UsersRepository,
								  *repository.AgentsRepository) {

	return &repository.SystemUsersRepository{Database: database}, &repository.AdministratorsRepository{Database: database},
																  &repository.UsersRepository{Database: database},
																  &repository.AgentsRepository{Database: database}
}



func initServices(systemUsersRepo *repository.SystemUsersRepository, administratorsRepo *repository.AdministratorsRepository,
																	 usersRepo *repository.UsersRepository,
																	 agentsRepo *repository.AgentsRepository) (*service.SystemUsersService,
																	                                           *service.AdministratorsService,
																	                                           *service.UsersService,
																	                                           *service.AgentsService){

	return &service.SystemUsersService{Repo: systemUsersRepo}, &service.AdministratorsService{AdministratorRepo: administratorsRepo, SystemUserRepo: systemUsersRepo},
	                                                           &service.UsersService{UsersRepo: usersRepo, SystemUserRepo: systemUsersRepo},
	                                                           &service.AgentsService{SystemUserRepo: systemUsersRepo, AgentsRepo: agentsRepo}
}



func initHandler(SystemUsersService *service.SystemUsersService, administratorsService *service.AdministratorsService,
															     usersService *service.UsersService,
															     agentsService *service.AgentsService) (*handler.SystemUsersHandler,
															     										*handler.AdministratorsHandler,
															     										*handler.UsersHandler,
															     										*handler.AgentsHandler) {
	return &handler.SystemUsersHandler{Service: SystemUsersService}, &handler.AdministratorsHandler{Service: administratorsService},
		&handler.UsersHandler{Service: usersService}, &handler.AgentsHandler{Service: agentsService}
}



func handleFunc(SystemUsersHandler *handler.SystemUsersHandler, administratorsHandler *handler.AdministratorsHandler,
	usersHandler *handler.UsersHandler,agentsHandler *handler.AgentsHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/sysusers/create", SystemUsersHandler.Create).Methods("POST")
	router.HandleFunc("/sysusers/update", SystemUsersHandler.Update).Methods("PUT")
	router.HandleFunc("/sysusers/getAll",SystemUsersHandler.GetAll).Methods("GET")
	router.HandleFunc("/administrators/update",  administratorsHandler.Update).Methods("PUT")
	router.HandleFunc("/administrators/create",  administratorsHandler.Create).Methods("POST")
	router.HandleFunc("/administrators/getAll",  administratorsHandler.GetAll).Methods("GET")
	router.HandleFunc("/users/update",  usersHandler.Update).Methods("PUT")
	router.HandleFunc("/users/create",  usersHandler.Create).Methods("POST")
	router.HandleFunc("/users/getAll",  usersHandler.GetAll).Methods("GET")
	router.HandleFunc("/users/changeWhetherIsPublic", usersHandler.ChangeWhetherIsPublic).Methods("POST")
	router.HandleFunc("/users/changeAllowedTags", usersHandler.ChangeAllowedTags).Methods("POST")
	router.HandleFunc("/agents/update",  agentsHandler.Update).Methods("PUT")
	router.HandleFunc("/agents/create",  agentsHandler.Create).Methods("POST")
	router.HandleFunc("/agents/getAll",  agentsHandler.GetAll).Methods("GET")


	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8085"), router))
}

func main() {
	database := initDB()
	sysusersRepo, administratorsRepo, usersRepo, agentsRepo := initRepo(database)
	systemUsersService, administratorsService, usersService, agentsService := initServices(sysusersRepo, administratorsRepo, usersRepo, agentsRepo)
	systemUsersHandler, administratorsHandler, usersHandler, agentsHandler := initHandler(systemUsersService, administratorsService, usersService, agentsService)
	handleFunc(systemUsersHandler, administratorsHandler, usersHandler, agentsHandler)
}

