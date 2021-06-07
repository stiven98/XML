package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type FollowersRepository struct {
	Driver *neo4j.Driver
}

func (r FollowersRepository) UserExist(id string) interface{} {

	session := (*r.Driver).NewSession(neo4j.SessionConfig{})
	defer session.Close()
	responses, err := session.Run("MATCH (u:User {id: $id}) return u", map[string]interface{}{
		"id": id,
	})

	if err != nil {
		fmt.Println("Query error!")
		return nil
	}

	response, err := responses.Single()
	if err != nil {
		return nil
	}
	return response
}

func (r FollowersRepository) GetFollowers(id string) interface{} {
	session := (*r.Driver).NewSession(neo4j.SessionConfig{})
	defer session.Close()

	responses, err := session.Run("MATCH (u:User) -[:FOLLOWING] ->(user:User{id: $id}) return u.id", map[string]interface{}{
		"id": id,
	})

	if err != nil {
		fmt.Println("Query error!")
		return nil
	}
	//fmt.Println(responses.Collect())
	retVal, err := responses.Collect()

	var followersIds [] uuid.UUID

	fmt.Println(len(retVal))
	fmt.Println(retVal[0].Values[0])

	for i := range retVal {
		s, ok := retVal[i].Values[0].(string)
		if ok {
			followersIds = append(followersIds, uuid.MustParse(s))
			fmt.Println(followersIds)
		}
	}

	return followersIds
}

func (r FollowersRepository) GetFollowing(id string) interface{} {
	session := (*r.Driver).NewSession(neo4j.SessionConfig{})
	defer session.Close()

	responses, err := session.Run("MATCH (u:User) <- [:FOLLOWING] - (user:User{id: $id}) return u.id", map[string]interface{}{
		"id": id,
	})

	if err != nil {
		fmt.Println("Query error!")
		return nil
	}
	retVal, err := responses.Collect()

	var followersIds [] uuid.UUID

	for i := range retVal {
		s, ok := retVal[i].Values[0].(string)
		if ok {
			followersIds = append(followersIds, uuid.MustParse(s))
			fmt.Println(followersIds)
		}
	}

	return followersIds
}

func (r FollowersRepository) CheckRelationship(userID string, targetID string) bool {
	session := (*r.Driver).NewSession(neo4j.SessionConfig{})
	defer session.Close()

	fmt.Println(userID)
	fmt.Println(targetID)

	responses, err := session.Run("MATCH (userA:User{id: $userID}) - [r] -> (userB:User{id: $targetID}) return count(type(r))", map[string]interface{}{
		"userID": userID,
		"targetID": targetID,
	})

	if err != nil {
		fmt.Println("Query error!")
		return false
	}



	if responses.Next() {
		fmt.Println(responses.Record().Values[0].(int64))
		if responses.Record().Values[0].(int64) > 0 {
			return true
		} else {
			return false
		}
	}

	return false

}

func (r FollowersRepository) Follow(userID string, targetID string) error{
	session := (*r.Driver).NewSession(neo4j.SessionConfig{})
	defer session.Close()

	fmt.Println(userID)
	fmt.Println(targetID)


	_, err := session.Run("MATCH  (a:User),  (b:User)WHERE a.id = $idA AND b.id = $idB CREATE (a)<-[r:FOLLOWING]-(b) RETURN type(r)", map[string]interface{}{
		"idB" : userID,
		"idA" : targetID,
	})

	if err != nil {
		fmt.Println("Query error!")
		return err
	}
	return nil
}

func (r FollowersRepository) AddNode(id string) (interface{}, error) {
	session := (*r.Driver).NewSession(neo4j.SessionConfig{})
	defer session.Close()

	fmt.Println(id)

	res, err := session.Run("CREATE (u:User{id: $id}) return u", map[string]interface{}{
		"id": id,
	})

	if err != nil {
		fmt.Println("Query error!")
		return nil, err
	}

	response, err := res.Single()
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (r FollowersRepository) Request(userID string, targetID string) interface{} {
	session := (*r.Driver).NewSession(neo4j.SessionConfig{})
	defer session.Close()

	fmt.Println(userID)
	fmt.Println(targetID)


	_, err := session.Run("MATCH  (a:User), (b:User) WHERE a.id = $idA AND b.id = $idB CREATE (a) -[r:REQUEST]-> (b) RETURN type(r)", map[string]interface{}{
		"idB" : targetID,
		"idA" : userID,
	})

	if err != nil {
		fmt.Println("Query error!")
		return err
	}
	return nil
}

func (r FollowersRepository) GetRequests(id string) interface{} {
	session := (*r.Driver).NewSession(neo4j.SessionConfig{})
	defer session.Close()

	responses, err := session.Run("MATCH (u:User) - [:REQUEST] -> (user:User{id: $id}) return u.id", map[string]interface{}{
		"id": id,
	})

	if err != nil {
		fmt.Println("Query error!")
		return nil
	}
	//fmt.Println(responses.Collect())
	retVal, err := responses.Collect()

	var followersIds [] uuid.UUID

	for i := range retVal {
		s, ok := retVal[i].Values[0].(string)
		if ok {
			followersIds = append(followersIds, uuid.MustParse(s))
			fmt.Println(followersIds)
		}
	}

	return followersIds


}

func (r FollowersRepository) Unfollow(userID string, targetID string) interface{} {
	session := (*r.Driver).NewSession(neo4j.SessionConfig{})
	defer session.Close()

	fmt.Println(userID)
	fmt.Println(targetID)


	_, err := session.Run("MATCH  (a:User {id: $idA}) <- [r] - (b:User {id: $idB}) DELETE r", map[string]interface{}{
		"idB" : userID,
		"idA" : targetID,
	})

	if err != nil {
		fmt.Println("Query error!")
		return err
	}
	return nil

}



