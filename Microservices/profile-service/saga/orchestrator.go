package saga

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"log"
)


const (
	ProfileChannel    string = "ProfileChannel"
	FollowerChannel    string = "FollowerChannel"
	ReplyChannel    string = "ReplyChannel"
	ServiceProfile    string = "Profile"
	ServiceFollower    string = "Follower"
	ActionStart     string = "Start"
	ActionDone      string = "DoneMsg"
	ActionError     string = "ErrorMsg"
	ActionRollback  string = "RollbackMsg"
)

type Orchestrator struct {
	c *redis.Client
	r *redis.PubSub
}

func NewOrchestrator() *Orchestrator {
	var err error
	// create client and ping redis
	client := redis.NewClient(&redis.Options{Addr: "localhost:6379", Password: "", DB: 0})
	if _, err = client.Ping().Result(); err != nil {
		log.Fatalf("error creating redis client %s", err)
	}

	// initialize and start the orchestrator in the background
	o := &Orchestrator{
		c: client,
		r: client.Subscribe(ProfileChannel, FollowerChannel, ReplyChannel),
	}

	return o
}

func (o Orchestrator) Start() {
	var err error
	if _, err = o.r.Receive(); err != nil {
		log.Fatalf("error setting up redis %s \n", err)
	}
	ch := o.r.Channel()
	defer func() { _ = o.r.Close() }()

	log.Println("starting the redis client")
	for {
		select {
		case msg := <-ch:
			m := Message{}
			if err = json.Unmarshal([]byte(msg.Payload), &m); err != nil {
				log.Println(err)
				// continue to skip bad messages
				continue
			}

			// only process the messages on ReplyChannel
			switch msg.Channel {
			case ReplyChannel:
				// if there is any error, just rollback
				if m.Action != ActionDone {
					o.Rollback(m)
					continue
				}

				// else start the next stage
				switch m.Service {
				case ServiceFollower:
					o.Next(FollowerChannel, ServiceFollower, m)
				case ServiceProfile:
					o.Next(ProfileChannel, ServiceProfile, m)
				}
			}
		}
	}
}

func (o Orchestrator) Next(channel, service string, message Message) {
	fmt.Println("jhbhjkhkjhkhkgk")
	var err error
	message.Action = ActionStart
	message.Service = service
	if err = o.c.Publish(channel, message).Err(); err != nil {
		log.Printf("error publishing start-message to %s channel", channel)
		log.Fatal(err)
	}
	log.Printf("start message published to channel :%s", channel)
}

func (o Orchestrator) Rollback(m Message) {
	var err error
	var channel string
	switch m.Service {
	case ServiceFollower:
		channel = FollowerChannel
	case ServiceProfile:
		channel = ProfileChannel
	}
	m.Action = ActionRollback
	if err = o.c.Publish(channel, m).Err(); err != nil {
		log.Printf("error publishing rollback message to  channel")
	}
}