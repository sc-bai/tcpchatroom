package rediscache

// redis 采用hash 把clientuser 存储起来
import (
	"chatroom/comm"
	"fmt"
	"reflect"

	"github.com/go-redis/redis"
)

var DBclient *redis.Client

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

	// 单个设置
	/* _, err := p.DBclient.HSet(user.UserId, "username", user.UserName).Result()
	if err != nil {
		fmt.Printf("sb hset username error: %v\n", err)
		return err
	}
	_, err = p.DBclient.HSet(user.UserId, "userpasswd", user.UserPasswd).Result()
	if err != nil {
		fmt.Printf("sb hset userpasswd error: %v\n", err)
		return err
	}
	*/

	userinfo := make(map[string]interface{})
	userinfo["userid"] = user.UserId
	userinfo["username"] = user.UserName
	userinfo["userpasswd"] = user.UserPasswd

	s, err := DBclient.HMSet(user.UserName, userinfo).Result()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}
	fmt.Printf("[server]: DBclient.HMSet ret string: %v\n", s)
	return nil
}

func (p *RedisDb) DelUser(username string) {

	// 不用差了 直接删
	/* 	i, err := p.DBclient.HMGet(username, "userid", "username", "userpasswd").Result()
	   	if err != redis.Nil || i == nil {
	   		fmt.Printf("err: %v\n", err)
	   		return err // 不存在
	   	} */
	DBclient.HDel(username).Result()
}

func (p *RedisDb) FindUser(username string) comm.LoginMessage {
	var ret comm.LoginMessage

	i, err := DBclient.HMGet(username, "userid", "username", "userpasswd").Result()
	if err != nil {
		fmt.Printf("DBclient.HMGet: %v\n", err)

		if err == redis.Nil {
			fmt.Println("HmGet redis.nil")
		}
		return ret
	}

	if i[0] == nil {
		return ret
	}

	fmt.Printf("[server]: DBclient.HMGet result: %v\n", i)
	if len(i) != 3 {
		return ret
	}

	fmt.Println(" dbclient finduser success..")
	ret.UserId = reflect.ValueOf(i[0]).String()
	ret.UserName = reflect.ValueOf(i[1]).String()
	ret.UserPasswd = reflect.ValueOf(i[2]).String()

	return ret
}
