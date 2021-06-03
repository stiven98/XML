package main

import (
	"fmt"
	"followers-microservice/handler"
	"followers-microservice/model"
	"followers-microservice/repository"
	"followers-microservice/service"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"log"
	"net/http"
	"time"
)

func initDB() *neo4j.Driver {
	dbUri := "neo4j://localhost:7687"
	var driver neo4j.Driver
	for 1 == 1 {
		d, err := neo4j.NewDriver(dbUri, neo4j.BasicAuth("neo4j", "root", ""))
		if err != nil {
			time.Sleep(2 * time.Second)
			log.Println("Reconnecting to database...")
			continue
		}
		driver = d
		log.Println("Connected to database")
		break
	}

	//defer driver.Close()

	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	//fmt.Println(session)


	err, _ := session.WriteTransaction(insertUsers)

	if err != nil {
		return nil
	}

	return &driver

}

func insertUsers(tx neo4j.Transaction) (interface{},  error) {
	tx.Run("MATCH ()-[r:FOLLOWING]->() DELETE r", map[string]interface{}{})
	tx.Run("MATCH (u) REMOVE u:User", map[string]interface{}{})

	users := [] model.User {
		{
			ID: uuid.MustParse("b94a0d50-c3c0-11eb-8529-0242ac130004"),
		},{
			ID: uuid.MustParse("5d9faac6-c3c7-11eb-8529-0242ac130004"),
		},{
			ID: uuid.MustParse("e8cf21de-c3f5-11eb-8529-0242ac130003"),
		},{
			ID: uuid.MustParse("ec34c504-c3f5-11eb-8529-0242ac130003"),
		},{
			ID: uuid.MustParse("ef99cdf2-c3f5-11eb-8529-0242ac130003"),
		},

	}

	for i := range users {
		_ , err := tx.Run("CREATE (u:User {id: $id}) return u.id", map[string]interface{}{
			"id": users[i].ID.String(), //string(users[i].ID)
		})
		if err != nil {
			return nil, err
		}

	}

	_, err := tx.Run("MATCH  (a:User),  (b:User)WHERE a.id = $idA AND b.id = $idB CREATE (a)-[r:FOLLOWING]->(b) RETURN type(r)", map[string]interface{}{
		"idB" : "b94a0d50-c3c0-11eb-8529-0242ac130004",
		"idA" : "5d9faac6-c3c7-11eb-8529-0242ac130004",
	})

	_, err = tx.Run("MATCH  (a:User),  (b:User)WHERE a.id = $idA AND b.id = $idB CREATE (a)-[r:FOLLOWING]->(b) RETURN type(r)", map[string]interface{}{
		"idB" : "b94a0d50-c3c0-11eb-8529-0242ac130004",
		"idA" : "e8cf21de-c3f5-11eb-8529-0242ac130003",
	})

	_, err = tx.Run("MATCH  (a:User),  (b:User)WHERE a.id = $idA AND b.id = $idB CREATE (a)-[r:FOLLOWING]->(b) RETURN type(r)", map[string]interface{}{
		"idB" : "b94a0d50-c3c0-11eb-8529-0242ac130004",
		"idA" : "ef99cdf2-c3f5-11eb-8529-0242ac130003",
	})

	_, err = tx.Run("MATCH  (a:User),  (b:User)WHERE a.id = $idA AND b.id = $idB CREATE (a)-[r:FOLLOWING]->(b) RETURN type(r)", map[string]interface{}{
		"idA" : "b94a0d50-c3c0-11eb-8529-0242ac130004",
		"idB" : "ef99cdf2-c3f5-11eb-8529-0242ac130003",
	})

	_, err = tx.Run("MATCH  (a:User),  (b:User)WHERE a.id = $idA AND b.id = $idB CREATE (a)-[r:FOLLOWING]->(b) RETURN type(r)", map[string]interface{}{
		"idA" : "ec34c504-c3f5-11eb-8529-0242ac130003",
		"idB" : "ef99cdf2-c3f5-11eb-8529-0242ac130003",
	})




	if err != nil {
		return nil, err
	}

	return  nil, nil
}


func initRepo(driver *neo4j.Driver) *repository.FollowersRepository {
	return &repository.FollowersRepository{Driver: driver}
}

func initService(followersRepository *repository.FollowersRepository) *service.FollowersService {
	return &service.FollowersService{FollowersRepository: followersRepository}
}

func initHandler(followersService *service.FollowersService) *handler.FollowersHandler  {
	return &handler.FollowersHandler{FollowersService: followersService}
}

func handlerFunc(followersHandler *handler.FollowersHandler)  {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/users/getFollowers/{id}", followersHandler.GetFollowers).Methods("GET")
	router.HandleFunc("/users/getFollowing/{id}", followersHandler.GetFollowing).Methods("GET")
	router.HandleFunc("/users/follow/{idUser}/{idTarget}", followersHandler.Follow).Methods("POST")
	router.HandleFunc("/users/isFollowing/{idUser}/{idTarget}", followersHandler.IsFollowing).Methods("GET")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8088"), router))
}


func main() {
	driver := initDB()
	followersRepository := initRepo(driver)
	followersService := initService(followersRepository)
	followersHandler := initHandler(followersService)
	handlerFunc(followersHandler)

}

// TO DO:
// getFollowers
// getFollowing
// isFollowing/{idUser}/{idTarget}


// ** Nemojte ovo ispod jos brisati moze posluziti za kasnije zbog baze! **

	// Neo4j 4.0, defaults to no TLS therefore use bolt:// or neo4j://
	// Neo4j 3.5, defaults to self-signed certificates, TLS on, therefore use bolt+ssc:// or neo4j+ssc://
	//dbUri := "neo4j://localhost:7687"
	//driver, err := neo4j.NewDriver(dbUri, neo4j.BasicAuth("neo4j", "root", ""))
	//if err != nil {
	//	fmt.Println("Greska")
	//	panic(err)
	//}
	//// Handle driver lifetime based on your application lifetime requirements  driver's lifetime is usually
	//// bound by the application lifetime, which usually implies one driver instance per application
	//defer driver.Close()
	//err = insertItem(driver)
	//if err != nil {
	//	fmt.Println("Greska")
	//	panic(err)
	//}
	////fmt.Printf("%v\n", item)


//func insertItem(driver neo4j.Driver) error {
//	// Sessions are short-lived, cheap to create and NOT thread safe. Typically create one or more sessions
//	// per request in your web application. Make sure to call Close on the session when done.
//	// For multi-database support, set sessionConfig.DatabaseName to requested database
//	// Session config will default to write mode, if only reads are to be used configure session for
//	// read mode.
//	session := driver.NewSession(neo4j.SessionConfig{})
//	defer session.Close()
//	_, err := session.WriteTransaction(createItemFn)
//	if err != nil {
//		return  err
//	}
//	return  nil
//}
//
//func createItemFn(tx neo4j.Transaction) (interface{}, error) {
//
//
//
//		if err != nil {
//			return nil, err
//		}
//		record, err := records.Single()
//		if err != nil {
//			return nil, err
//		}
//		fmt.Println(record.Values)
//
//
//
//
//	//records, err := tx.Run("CREATE (u:User { id: $id}) RETURN u.id", map[string]interface{}{
//	//	"id": uuid.MustParse("b94a0d50-c3c0-11eb-8529-0242ac130003"),
//	//},{
//	//	"id": uuid.MustParse("b94a0d50-c3c0-11eb-8529-0242ac130003"),
//	//})
//	// In face of driver native errors, make sure to return them directly.
//	// Depending on the error, the driver may try to execute the function again.
//
//
//	// You can also retrieve values by name, with e.g. `id, found := record.Get("n.id")`
//	//return &Item{
//	//	Id:   record.Values[0].(int64),
//	//	Name: record.Values[1].(string),
//	//}, nil
//	return nil, nil
//}


