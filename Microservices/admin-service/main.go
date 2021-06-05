package main


import (
	"admin-service/handler"
	"admin-service/model"
	"admin-service/repository"
	"admin-service/service"
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

		database.AutoMigrate(&model.Administrator{})
		database.AutoMigrate(&model.AgentRegistrationRequest{})
		database.AutoMigrate(&model.AccountVerificationRequest{})

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
		for i := range administrators {
			database.Create(&administrators[i])
		}
		for i := range accountVerificationRequests {
			database.Create(&accountVerificationRequests[i])
		}
		for i := range agentRegistrationRequests {
			database.Create(&agentRegistrationRequests[i])
		}

	}
	return database
}

func initRepo(database *gorm.DB) (*repository.AdministratorsRepository,
	                              *repository.AccountVerificationsRepository,
	                              *repository.AgentRegistrationsRepository, ) {

	return &repository.AdministratorsRepository{Database: database}, &repository.AccountVerificationsRepository{Database: database},
		&repository.AgentRegistrationsRepository{Database: database}
}



func initServices(adminsRepo *repository.AdministratorsRepository, accountVerificationsRepo *repository.AccountVerificationsRepository,
	agentRegistrationRepo *repository.AgentRegistrationsRepository) (*service.AdministratorsService,
																	 *service.AccountVerificationsService,
																	 *service.AgentRegistrationsService){

	return &service.AdministratorsService{AdministratorRepo: adminsRepo}, &service.AccountVerificationsService{AccountVerificationsRepo: accountVerificationsRepo},
		&service.AgentRegistrationsService{AgentRegistrationsRepo: agentRegistrationRepo}
}


func initHandler(adminsService *service.AdministratorsService, accountVerificationService *service.AccountVerificationsService,
	agentRegistrationsService *service.AgentRegistrationsService) ( *handler.AdministratorsHandler,
																	*handler.AccountVerificationsHandler,
																	*handler.AgentRegistrationsHandler) {
	return &handler.AdministratorsHandler{Service: adminsService}, &handler.AccountVerificationsHandler{Service: accountVerificationService},
		&handler.AgentRegistrationsHandler{Service: agentRegistrationsService}
}



func handleFunc(administratorsHandler *handler.AdministratorsHandler, accountVerificationHadnler *handler.AccountVerificationsHandler,
	agentRegistrationsHandler *handler.AgentRegistrationsHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/administrators/update",  administratorsHandler.Update).Methods("PUT")
	router.HandleFunc("/administrators/create",  administratorsHandler.Create).Methods("POST")
	router.HandleFunc("/administrators/getAll",  administratorsHandler.GetAll).Methods("GET")
	router.HandleFunc("/accountVerification/update",  accountVerificationHadnler.Update).Methods("PUT")
	router.HandleFunc("/accountVerification/create",  accountVerificationHadnler.Create).Methods("POST")
	router.HandleFunc("/accountVerification/getAll",  accountVerificationHadnler.GetAll).Methods("GET")
	router.HandleFunc("/agentRegistration/update",  agentRegistrationsHandler.Update).Methods("PUT")
	router.HandleFunc("/agentRegistration/create",  agentRegistrationsHandler.Create).Methods("POST")
	router.HandleFunc("/agentRegistration/getAll",  agentRegistrationsHandler.GetAll).Methods("GET")


	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8089"), router))
}

func main() {
	database := initDB()
	adminsRepo, accountVerificationRepo, agentRegistrationRepo := initRepo(database)
	adminsService, accountVerifciationService, agentRegistrationService := initServices(adminsRepo, accountVerificationRepo, agentRegistrationRepo)
	adminsHandler, accountVerificationHandler, agentRegistrationHanlder := initHandler(adminsService, accountVerifciationService, agentRegistrationService)
	handleFunc(adminsHandler, accountVerificationHandler, agentRegistrationHanlder)
}

