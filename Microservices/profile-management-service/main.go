package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"profile-management-service/handler"
	"profile-management-service/model"
	"profile-management-service/repository"
	"profile-management-service/service"
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
	dbname := "management_profile_service"
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
		database.Migrator().DropTable(&model.CloseFriends{})
		database.Migrator().DropTable(&model.BlockedUsers{})
		database.Migrator().DropTable(&model.MutedUsers{})
		database.Migrator().DropTable(&model.SubscribeAcc{})


		database.AutoMigrate(&model.CloseFriends{})
		database.AutoMigrate(&model.BlockedUsers{})
		database.AutoMigrate(&model.MutedUsers{})
		database.AutoMigrate(&model.SubscribeAcc{})

	
		//blockedUsersInit := [] model.BlockedUsers {
		//	{
		//		BlockedByID: uuid.MustParse("9da543e6-0484-4dce-9cef-68ab8a291826"),
		//		BlockedID:   uuid.MustParse("7c0d42ad-aedf-47c6-93d3-fc53dcc57099"),
		//	}, {
		//		BlockedByID: uuid.MustParse("9da543e6-0484-4dce-9cef-68ab8a291826"),
		//		BlockedID:   uuid.MustParse("a8c4752e-c09b-11eb-8529-0242ac130003"),
		//	}, {
		//		BlockedByID: uuid.MustParse("9da543e6-0484-4dce-9cef-68ab8a291826"),
		//		BlockedID:   uuid.MustParse("ad79e1bc-c09b-11eb-8529-0242ac130003"),
		//	},
		//
		//}
		//
		//closeFriendsInit := [] model.CloseFriends {
		//	{
		//		UserID:   uuid.MustParse("9da543e6-0484-4dce-9cef-68ab8a291826"),
		//		FriendID: uuid.MustParse("e748f162-c09b-11eb-8529-0242ac130003"),
		//	}, {
		//		UserID:   uuid.MustParse("9da543e6-0484-4dce-9cef-68ab8a291826"),
		//		FriendID: uuid.MustParse("eb1eafc0-c09b-11eb-8529-0242ac130003"),
		//	}, {
		//		UserID:   uuid.MustParse("9da543e6-0484-4dce-9cef-68ab8a291826"),
		//		FriendID: uuid.MustParse("ed77d88c-c09b-11eb-8529-0242ac130003"),
		//	},
		//
		//}
		//
		//mutedUsersInit := [] model.MutedUsers {
		//	{
		//		MutedByID: uuid.MustParse("9da543e6-0484-4dce-9cef-68ab8a291826"),
		//		MutedID:   uuid.MustParse("23b899e0-c09c-11eb-8529-0242ac130003"),
		//	}, {
		//		MutedByID: uuid.MustParse("1e35dc53-3743-4737-8463-4400418149ca"),
		//		MutedID:   uuid.MustParse("89d660eb-9a79-4f11-97f1-0a7476af35ab"),
		//	},
		//}
		//
		//
		//for i := range blockedUsersInit {
		//	database.Create(&blockedUsersInit[i])
		//}
		//
		//for i:= range closeFriendsInit {
		//	database.Create(&closeFriendsInit[i])
		//}
		//
		//for i := range mutedUsersInit {
		//	database.Create(&mutedUsersInit[i])
		//}
	}


	return database
}

func initRepo(database *gorm.DB) (*repository.BlockedUsersRepository, *repository.CloseFriendsRepository, *repository.MutedUsersRepository, *repository.SubscribeAccRepository) {
	return &repository.BlockedUsersRepository{DataBase: database}, &repository.CloseFriendsRepository{Database: database}, &repository.MutedUsersRepository{DataBase: database},
	&repository.SubscribeAccRepository{DataBase:database}
}

func initServices(blockedUsersRepository *repository.BlockedUsersRepository, closeFriendRepository *repository.CloseFriendsRepository,
	mutedUsersRepository * repository.MutedUsersRepository, subscribeAccRepository *repository.SubscribeAccRepository) (*service.BlockedUsersService,
	*service.CloseFriendsService, *service.MutedUsersService, *service.SubscribeAccService){
	return &service.BlockedUsersService{BlockedUsersRepository: blockedUsersRepository}, &service.CloseFriendsService{CloseFriendsService: closeFriendRepository},
	&service.MutedUsersService{MutedUsersRepository: mutedUsersRepository}, &service.SubscribeAccService{SubscriberAccRepository: subscribeAccRepository}
}

func initHandler(blockedUserService *service.BlockedUsersService, closeFriendService *service.CloseFriendsService,
	mutedUsersService *service.MutedUsersService, subscribeAccService *service.SubscribeAccService) (*handler.BlockedUsersHandler,
	*handler.CloseFriendHandler, *handler.MutedUsersHandler, *handler.SubscribeAccHandler) {
	return &handler.BlockedUsersHandler{BlockedUsersService: blockedUserService},
	&handler.CloseFriendHandler{CloseFriendService: closeFriendService},
	&handler.MutedUsersHandler{MutedUsersService: mutedUsersService},
	&handler.SubscribeAccHandler{SubscriberAccService: subscribeAccService}
}

func handleFunc(blockedUsersHandler *handler.BlockedUsersHandler,
	closeFriendsHandler *handler.CloseFriendHandler,
	mutedUsersHandler *handler.MutedUsersHandler, subscribeAccHandler *handler.SubscribeAccHandler) {



	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/users/blocked/{id}", blockedUsersHandler.GetAllBlockedBy).Methods("GET")
	router.HandleFunc("/users/block/{blockedById}/{blockedId}", blockedUsersHandler.BlockUserByUser).Methods("POST")
	router.HandleFunc("/users/unblock/{blockedById}/{blockedId}", blockedUsersHandler.UnBlockUserByUser).Methods("POST")
	router.HandleFunc("/users/isBlocked/{blockedById}/{blockedId}", blockedUsersHandler.IsBlocked).Methods("GET")
	router.HandleFunc("/users/isMuted/{mutedById}/{mutedId}", mutedUsersHandler.IsMuted).Methods("GET")
	router.HandleFunc("/users/mute/{mutedById}/{mutedId}", mutedUsersHandler.MutedUserByUser).Methods("POST")
	router.HandleFunc("/users/unmute/{mutedById}/{mutedId}", mutedUsersHandler.UnMutedUserByUser).Methods("POST")
	router.HandleFunc("/users/muted/{id}", mutedUsersHandler.GetAllMutedBy).Methods("GET")
	router.HandleFunc("/users/subscribe/{subscribedById}/{subscribedId}", subscribeAccHandler.Subscribe).Methods("POST")
	router.HandleFunc("/users/unsubscribe/{subscribedById}/{subscribedId}", subscribeAccHandler.UnSubscribe).Methods("POST")
	router.HandleFunc("/users/subscribers/{id}", subscribeAccHandler.GetAllSubscribers).Methods("GET")
	router.HandleFunc("/users/isSubscribed/{subscribedById}/{subscribedId}", subscribeAccHandler.IsSubscribed).Methods("GET")
	router.HandleFunc("/users/addCloseFriend/{userId}/{friendId}", closeFriendsHandler.AddCloseFriend).Methods("POST")
	router.HandleFunc("/users/removeCloseFriend/{userId}/{friendId}", closeFriendsHandler.RemoveCloseFriend).Methods("POST")
	router.HandleFunc("/users/closeFriend/{id}", closeFriendsHandler.GetAllCloseFriend).Methods("GET")
	//router.HandleFunc("/users/getAll",SystemUsersHandler.GetAll).Methods("GET")
	//router.HandleFunc("/administrators/update",  administratorsHandler.Update).Methods("PUT")


	headers := handlers.AllowedHeaders([] string{"Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([] string{"GET", "POST", "PUT"})
	origins := handlers.AllowedOrigins([] string{"*"})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8087"),handlers.CORS(headers, methods, origins)  (router)))
}

func main() {
	database := initDB()
	blockedUsersRepository, closeFriendRepository, mutedUsersRepository, subscribeAccRepository := initRepo(database)
	blockedUsersService, closeFriendService, mutedUsersService, subscribeAccService := initServices(blockedUsersRepository, closeFriendRepository, mutedUsersRepository,subscribeAccRepository)
	blockedUsersHandler, closeFriendHandler, mutedUsersHandler, subscribeAccHandler := initHandler(blockedUsersService, closeFriendService, mutedUsersService,subscribeAccService)
	handleFunc(blockedUsersHandler, closeFriendHandler, mutedUsersHandler,subscribeAccHandler)
}
