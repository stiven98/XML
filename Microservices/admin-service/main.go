package main


import (
	"admin-service/handler"
	"admin-service/model"
	"admin-service/repository"
	"admin-service/service"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/handlers"
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
	dbname := "admins_service"
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
		database.Migrator().DropTable(&model.Administrator{})
		database.Migrator().DropTable(&model.AgentRegistrationRequest{})
		database.Migrator().DropTable(&model.AccountVerificationRequest{})
		database.Migrator().DropTable(&model.VerificationRequest{})

		database.AutoMigrate(&model.Administrator{})
		database.AutoMigrate(&model.AgentRegistrationRequest{})
		database.AutoMigrate(&model.AccountVerificationRequest{})
		database.AutoMigrate(&model.VerificationRequest{})

		administrators := [] model.Administrator{
			{
				ID:          uuid.UUID{},
				FIRSTNAME:   "Aca",
				LASTNAME:    "Simic",
				USERNAME:    "acasimic",
				EMAIL:       "aca@gmail.com",
				PASSWORD:    "aca12345",
				GENDER:      model.MALE,
				TYPEOFUSER:  model.ADMIN,
				DATEOFBIRTH: time.Time{}.AddDate(1998, 10, 1),
			},
		}
		accountVerificationRequests := []model.AccountVerificationRequest {
			{
				ID:        uuid.UUID{},
				FIRSTNAME: "Aca",
				LASTNAME:  "Simic",
				PHOTOPATH: "putanja do slike",
				CATEGORY:  model.INFLUENCER,
				STATUS:    model.SUBMITTED,
			},
		}
		agentRegistrationRequests := []model.AgentRegistrationRequest {
			{
				ID:     uuid.UUID{},
				LINK:   "link",
				USER:   "neki user ne znam sta hocete ovde da bude",
				STATUS: model.SUBMITTED,
			},
		}

		verificationRequests := [] model.VerificationRequest {
			{
				ID:           uuid.MustParse("f6b7b1b1-49e7-476b-b895-88baf888a0f6"),
				UserID:       uuid.MustParse(string("4579daae-1567-42d5-a25c-1a3818077c84")),
				DocumentPath: "stiven.jpeg",
				Status: model.SUBMITTED,
			},
		}

		for i := range administrators {
			database.Create(&administrators[i])
		}
		for i := range accountVerificationRequests {
			database.Create(&accountVerificationRequests[i])
		}
		for i := range agentRegistrationRequests {
			database.Create(&agentRegistrationRequests[i])
		}

		for i := range verificationRequests {
			database.Create(&verificationRequests[i])
		}

	}
	return database
}

func initRepo(database *gorm.DB) *repository.VerificationRequestRepository {
	return &repository.VerificationRequestRepository{Database: database}
}

func initServices(verificationRequestRepository *repository.VerificationRequestRepository) *service.VerificationRequestService {
	return &service.VerificationRequestService{VerificationRequestRepository: verificationRequestRepository}
}


func initHandler(verificationRequestService *service.VerificationRequestService) *handler.VerificationRequestHandler {
	return &handler.VerificationRequestHandler{VerificationRequestService: verificationRequestService}
}



func handleFunc(verificationRequestHandler *handler.VerificationRequestHandler) {
	router := mux.NewRouter().StrictSlash(true)

	//router.HandleFunc("/administrators/update",  administratorsHandler.Update).Methods("PUT")
	//router.HandleFunc("/administrators/create",  administratorsHandler.Create).Methods("POST")
	//router.HandleFunc("/administrators/getAll",  administratorsHandler.GetAll).Methods("GET")
	//router.HandleFunc("/accountVerification/update",  accountVerificationHadnler.Update).Methods("PUT")
	//router.HandleFunc("/accountVerification/create",  accountVerificationHadnler.Create).Methods("POST")
	//router.HandleFunc("/accountVerification/getAll",  accountVerificationHadnler.GetAll).Methods("GET")
	//router.HandleFunc("/agentRegistration/update",  agentRegistrationsHandler.Update).Methods("PUT")
	//router.HandleFunc("/agentRegistration/create",  agentRegistrationsHandler.Create).Methods("POST")
	//router.HandleFunc("/agentRegistration/getAll",  agentRegistrationsHandler.GetAll).Methods("GET")
	router.HandleFunc("/verificationRequest/getAll", verificationRequestHandler.GetAll).Methods("GET")
	router.HandleFunc("/verificationRequest/{id}", verificationRequestHandler.CreateVerificationRequest).Methods("POST")
	router.HandleFunc("/verificationRequest/accept/{id}", verificationRequestHandler.Accept).Methods("PUT")
	router.HandleFunc("/verificationRequest/decline/{id}", verificationRequestHandler.Decline).Methods("PUT")
	router.Handle("/images/{rest}", http.StripPrefix("/images/", http.FileServer(http.Dir("./documents/"))))

	methods := handlers.AllowedMethods([] string{"GET", "POST", "PUT"})
	origins := handlers.AllowedOrigins([] string{"*"})
	headers := handlers.AllowedHeaders([] string{"Content-Type", "Authorization"})



	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8089"), handlers.CORS(headers, methods, origins) (router)))
}

func main() {
	database := initDB()
	verificationRequestRepository := initRepo(database)
	verificationRequestService := initServices(verificationRequestRepository)
	verificationRequestHandler := initHandler(verificationRequestService)
	handleFunc(verificationRequestHandler)
}

