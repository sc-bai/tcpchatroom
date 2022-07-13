package rediscache

// redis 采用hash 把clientuser 存储起来
import (
	"chatroom/comm"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
)

var DBclient *redis.Client

var g_userkey string = "chat_users"

type RedisDb struct {
}

func InitDb() error {
	DBclient = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	s, err := DBclient.Ping().Result()
	fmt.Printf("s: %v\n", s)
	return err
}

func (p *RedisDb) PutUser(user comm.LoginMessage) error {

	b, _ := json.Marshal(user)
	s, err := DBclient.HSet(g_userkey, user.UserName, string(b)).Result()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}
	fmt.Printf("[server]: DBclient.HMSet ret string: %v\n", s)
	return nil
}

func (p *RedisDb) DelUser(username string) {
	DBclient.HDel(g_userkey, username).Result()
}

func (p *RedisDb) FindUser(username string) comm.LoginMessage {
	var ret comm.LoginMessage

	s, err2 := DBclient.HGet(g_userkey, username).Result()
	if err2 != nil {
		fmt.Printf("[server] redis hget error: %v\n", err2)
		return ret
	}

	err := json.Unmarshal([]byte(s), &ret)
	if err != nil {
		fmt.Println("redis: finduser json.Unmarshal error,", err.Error())
		return ret
	}

	return ret
}

func (p *RedisDb) ListUserOnline() []string {
	s, err := DBclient.HKeys(g_userkey).Result()
	if err != nil {
		fmt.Printf("ListUserOnline: %v\n", err)
		return s
	}

	return s
}
