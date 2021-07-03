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
	"shop-backend/handler"
	"shop-backend/model"
	"shop-backend/repository"
	"shop-backend/service"
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
	dbname := "shop_db"
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
		database.Migrator().DropTable(&model.User{})
		database.Migrator().DropTable(&model.Product{})

		database.AutoMigrate(&model.User{})
		database.AutoMigrate(&model.Product{})


		users := [] model.User {
			{
				UserID:     uuid.MustParse("3d202f62-0d31-4f36-a965-0e4792eb32f3"),
				FirstName: "Ranko",
				LastName:  "Rankovic",
				Username:  "ranko",
				Email:     "ranko@gmail.com",
				Password:  "$2y$10$szTo3OrMpAUX0kIvWHh0seRntNn/GG6zBWIRnK.DJ7y.zItJRLYO2",
			}, {
				UserID:    uuid.MustParse("007339c0-58be-4ba5-b1a1-9ef8ea2b89dc"),
				FirstName: "Marija",
				LastName:  "Markovic",
				Username:  "marija",
				Email:     "marija@gmail.com",
				Password:  "$2y$10$szTo3OrMpAUX0kIvWHh0seRntNn/GG6zBWIRnK.DJ7y.zItJRLYO2",
			}, {
				UserID:     uuid.MustParse("aee5cdda-2937-4d35-adf3-47cf8a19a239"),
				FirstName: "Isidora",
				LastName:  "Popovic",
				Username:  "isi",
				Email:     "isi@gmail.com",
				Password:  "$2y$10$szTo3OrMpAUX0kIvWHh0seRntNn/GG6zBWIRnK.DJ7y.zItJRLYO2",
			},

		}

		products := [] model.Product {
			{
				ID:          uuid.New(),
				User:        users[0],
				Price:       100.0,
				PicturePath: "chair.jpeg",
				Quantity:    400,
				Name:        "Stolica",
				Deleted: 	false,
			}, {
				ID:          uuid.New(),
				User:        users[0],
				Name:        "Sto",
				Price:       200.00,
				PicturePath: "table.jpeg",
				Quantity:    100,
				Deleted: 	false,

			}, {
				ID:          uuid.New(),
				User:        users[1],
				Name:        "Computer HP",
				Price:       1000,
				PicturePath: "computer.jpeg",
				Quantity:    10,
				Deleted: 	false,

			}, {
				ID:          uuid.New(),
				User:        users[2],
				Name:        "Cup",
				Price:       10,
				PicturePath: "cup.jpeg",
				Quantity:    150,
				Deleted: 	false,

			}, {
				ID:          uuid.New(),
				User:        users[0],
				Name:        "Rucni sat",
				Price:       350,
				PicturePath: "watch.jpeg",
				Quantity:    11,
				Deleted: 	false,

			},



		}

		for i := range users {
			database.Create(&users[i])
		}

		for i := range products {
			database.Create(&products[i])
		}
	}




		//for i := range administrators {
		//	//fmt.Println(administrators[i])
		//	database.Create(&administrators[i])
		//}


	return database
}

func initRepo(database *gorm.DB) (*repository.UsersRepository, *repository.ProductsRepository) {
	return &repository.UsersRepository{Database: database}, &repository.ProductsRepository{Database: database}
}


func initServices(usersRepository *repository.UsersRepository, productsRepository *repository.ProductsRepository) (*service.UsersService, *service.ProductsService){

	return &service.UsersService{UsersRepository: usersRepository}, &service.ProductsService{ProductsRepository: productsRepository}
}



func initHandler(userService *service.UsersService, productsService *service.ProductsService) (*handler.UsersHandler, *handler.ProductsHandler) {
	return &handler.UsersHandler{UsersService: userService}, &handler.ProductsHandler{ProductsService: productsService}
}



func handleFunc(usersHandler *handler.UsersHandler, productsHandler *handler.ProductsHandler) {
	router := mux.NewRouter().StrictSlash(true)


	router.HandleFunc("/users/create", usersHandler.Create).Methods("POST")
	router.HandleFunc("/users/login", usersHandler.Login).Methods("POST")
	router.HandleFunc("/products/user/{id}", productsHandler.GetProductByUser).Methods("GET")
	router.HandleFunc("/products/get/{id}", productsHandler.GetProductById).Methods("GET")
	router.HandleFunc("/products/update", productsHandler.Update).Methods("PUT")
	router.HandleFunc("/products/delete/{id}", productsHandler.Delete).Methods("DELETE")
	router.HandleFunc("/products/all", productsHandler.GetAllProducts).Methods("GET")
	router.HandleFunc("/products/create", productsHandler.Create).Methods("POST")
	router.HandleFunc("/images/upload", productsHandler.ImageUpload).Methods("POST")
	router.Handle("/images/{rest}",
		http.StripPrefix("/images/", http.FileServer(http.Dir("./products_images/"))))

	headers := handlers.AllowedHeaders([] string{"Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([] string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([] string{"*"})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8081"), handlers.CORS(headers, methods, origins) (router)))
}

func main() {
	database := initDB()
	usersRepository, productsRepository := initRepo(database)
	usersService, productsService := initServices(usersRepository, productsRepository)
	usersHandler, productsHandler := initHandler(usersService, productsService)
	handleFunc(usersHandler, productsHandler)
}

