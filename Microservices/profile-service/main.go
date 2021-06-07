package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"profileservice/handler"
	"profileservice/model"
	"profileservice/repository"
	"profileservice/service"
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
				ID:          uuid.MustParse("69b0597e-4a63-49e5-ae40-5b159ada82b9"),
				FirstName:   "Aca",
				LastName:    "Simic",
				Username:    "acasimic",
				Email:       "aca@gmail.com",
				Password:    "$2y$10$szTo3OrMpAUX0kIvWHh0seRntNn/GG6zBWIRnK.DJ7y.zItJRLYO2",
				Gender:      model.MALE,
				TypeOfUser: model.ADMIN,
				DateOfBirth: time.Time{}.AddDate(1998, 10, 1),
			}, {
				ID:          uuid.MustParse("965208b9-287b-4da5-b772-73df5e74ebbc"),
				FirstName:   "Jovan",
				LastName:    "Bosnic",
				Username:    "jovanbosnic",
				Email:       "jovan@gmail.com",
				Password:    "$2y$10$szTo3OrMpAUX0kIvWHh0seRntNn/GG6zBWIRnK.DJ7y.zItJRLYO2",
				Gender:      model.MALE,
				TypeOfUser: model.USER,
				DateOfBirth: time.Time{}.AddDate(1998, 7, 31),
			}, {
				ID:          uuid.MustParse("4579daae-1567-42d5-a25c-1a3818077c84"),
				FirstName:   "Djordjije",
				LastName:    "Kundacina",
				Username:    "djordjije",
				Email:       "djordjije@gmail.com",
				Password:    "$2y$10$szTo3OrMpAUX0kIvWHh0seRntNn/GG6zBWIRnK.DJ7y.zItJRLYO2",
				Gender: 	 model.MALE,
				TypeOfUser: model.USER,
				DateOfBirth: time.Time{}.AddDate(1998, 9, 10),
			}, {
				ID:          uuid.MustParse("5cb65bc8-6130-4436-a1f9-ad4778f112bc"),
				FirstName:   "Aleksandar",
				LastName:    "Stevanovic",
				Username:    "aleksandar",
				Email:       "aleksandar@gmail.com",
				Password:    "$2y$10$szTo3OrMpAUX0kIvWHh0seRntNn/GG6zBWIRnK.DJ7y.zItJRLYO2",
				Gender:      model.MALE,
				TypeOfUser: model.USER,
				DateOfBirth: time.Time{}.AddDate(1998, 10, 10),
			}, {
				ID:          uuid.MustParse("708b65de-fb77-4934-bfd0-d14161a74905"),
				FirstName:   "Marko",
				LastName:    "Markovic",
				Username:    "marko",
				Email:       "marko@gmail.com",
				Password:    "$2y$10$szTo3OrMpAUX0kIvWHh0seRntNn/GG6zBWIRnK.DJ7y.zItJRLYO2",
				Gender:      model.MALE,
				TypeOfUser: model.USER,
				DateOfBirth: time.Time{}.AddDate(1998, 1, 10),
			}, {
				ID:          uuid.MustParse("0cf8a7ff-7bb5-48f0-a834-7b07eb306f90"),
				FirstName:   "Janko",
				LastName:    "Jankovic",
				Username:    "janko_98",
				Email:       "janko@gmail.com",
				Password:    "$2y$10$szTo3OrMpAUX0kIvWHh0seRntNn/GG6zBWIRnK.DJ7y.zItJRLYO2",
				Gender:      model.MALE,
				TypeOfUser: model.USER,
				DateOfBirth: time.Time{}.AddDate(1994, 10, 10),
			}, {
				ID:          uuid.MustParse("be71d1da-0749-480f-a563-dcc35a14e542"),
				FirstName:   "Dejan",
				LastName:    "Dejanovic",
				Username:    "deki_99",
				Email:       "dejan@gmail.com",
				Password:    "$2y$10$szTo3OrMpAUX0kIvWHh0seRntNn/GG6zBWIRnK.DJ7y.zItJRLYO2",
				Gender:      model.MALE,
				TypeOfUser: model.USER,
				DateOfBirth: time.Time{}.AddDate(1991, 10, 10),
			},{
				ID:          uuid.MustParse("d3ea863d-350e-44f2-bd6e-809aa7100476"),
				FirstName:   "Milica",
				LastName:    "Milicevic",
				Username:    "milica00",
				Email:       "milica@gmail.com",
				Password:    "$2y$10$szTo3OrMpAUX0kIvWHh0seRntNn/GG6zBWIRnK.DJ7y.zItJRLYO2",
				Gender:      model.MALE,
				TypeOfUser: model.USER,
				DateOfBirth: time.Time{}.AddDate(1992, 12, 10),
			},

		}

		administrators := [] model.Administrator {
			{
				UserID: systemUsers[0].ID,
				SystemUser: systemUsers[0],
			},
		}

		users := [] model.User {
			{
				UserID:      systemUsers[1].ID,
				SystemUser:  systemUsers[1],
				IsPublic:    true,
				PhoneNumber: "00381-44-44-44-111",
				WebSite:     "www.org.com",
				Biography:   "Nema je!",
				AllowedTags: true,
				IsBlocked:   false,
			},
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
				IsPublic:    true,
				PhoneNumber: "00381-22-22-22-333",
				WebSite:     "www.stiven.com",
				Biography:   "Berem jagode!",
				AllowedTags: true,
				IsBlocked:   false,
			},{
				UserID:      systemUsers[4].ID,
				SystemUser:  systemUsers[4],
				IsPublic:    false,
				PhoneNumber: "00381-22-22-22-123",
				WebSite:     "www.google.com",
				Biography:   "!",
				AllowedTags: true,
				IsBlocked:   false,
			},{
				UserID:      systemUsers[5].ID,
				SystemUser:  systemUsers[5],
				IsPublic:    false,
				PhoneNumber: "00381-22-22-22-555",
				WebSite:     "www.yas.com",
				Biography:   "Berem jagode!",
				AllowedTags: true,
				IsBlocked:   false,
			},{
				UserID:      systemUsers[6].ID,
				SystemUser:  systemUsers[6],
				IsPublic:    false,
				PhoneNumber: "00381-22-22-22-666",
				WebSite:     "www.stiven.com",
				Biography:   "",
				AllowedTags: true,
				IsBlocked:   false,
			},{
				UserID:      systemUsers[7].ID,
				SystemUser:  systemUsers[7],
				IsPublic:    true,
				PhoneNumber: "00381-33-22-22-777",
				WebSite:     "www.whynot.com",
				Biography:   "",
				AllowedTags: true,
				IsBlocked:   false,
			},
		}

		for i := range administrators {
			fmt.Println(administrators[i])
			database.Create(&administrators[i])
		}
		for i := range users {
			fmt.Println(users[i])
			database.Create(&users[i])
			//_, err := http.Post("http://localhost:8088/users/addNode/" + users[i].UserID.String(), "", nil)
			//if err != nil {
			//	fmt.Println("Pokrenite neo4j i followers microservice!")
			//	continue
			//}
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
	router.HandleFunc("/sysusers/getAllUsernames",SystemUsersHandler.GetAllUsernames).Methods("GET")
	router.HandleFunc("/sysusers/getUserId/{username}",SystemUsersHandler.GetUserId).Methods("GET")
	router.HandleFunc("/sysusers/getById/{id}",  SystemUsersHandler.GetById).Methods("GET")
	router.HandleFunc("/administrators/update",  administratorsHandler.Update).Methods("PUT")
	router.HandleFunc("/administrators/create",  administratorsHandler.Create).Methods("POST")
	router.HandleFunc("/administrators/getAll",  administratorsHandler.GetAll).Methods("GET")
	router.HandleFunc("/users/update",  usersHandler.Update).Methods("PUT")
	router.HandleFunc("/users/create",  usersHandler.Create).Methods("POST")
	router.HandleFunc("/users/getAll",  usersHandler.GetAll).Methods("GET")
	router.HandleFunc("/users/getById/{id}",  usersHandler.GetById).Methods("GET")
	router.HandleFunc("/users/changeWhetherIsPublic", usersHandler.ChangeWhetherIsPublic).Methods("POST")
	router.HandleFunc("/users/changeAllowedTags", usersHandler.ChangeAllowedTags).Methods("POST")
	router.HandleFunc("/agents/update",  agentsHandler.Update).Methods("PUT")
	router.HandleFunc("/agents/create",  agentsHandler.Create).Methods("POST")
	router.HandleFunc("/agents/getAll",  agentsHandler.GetAll).Methods("GET")
	router.HandleFunc("/users/isPublic/{id}", usersHandler.IsPublic).Methods("GET")

	headers := handlers.AllowedHeaders([] string{"Content-Type"})
	methods := handlers.AllowedMethods([] string{"GET", "POST", "PUT"})
	origins := handlers.AllowedOrigins([] string{"*"})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8085"), handlers.CORS(headers, methods, origins) (router)))
}

func main() {
	fmt.Println(uuid.New())
	database := initDB()
	sysusersRepo, administratorsRepo, usersRepo, agentsRepo := initRepo(database)
	systemUsersService, administratorsService, usersService, agentsService := initServices(sysusersRepo, administratorsRepo, usersRepo, agentsRepo)
	systemUsersHandler, administratorsHandler, usersHandler, agentsHandler := initHandler(systemUsersService, administratorsService, usersService, agentsService)
	handleFunc(systemUsersHandler, administratorsHandler, usersHandler, agentsHandler)
}

