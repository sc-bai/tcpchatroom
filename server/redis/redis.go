package rediscache

// redis 采用hash 把clientuser 存储起来
import (
	"chatroom/comm"
	"errors"
	"fmt"
	"reflect"

	"github.com/go-redis/redis"
)

type RedisDb struct {
	DBclient *redis.Client
}

func InitDb() (*redis.Client, error) {
	db := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	s, err := db.Ping().Result()
	fmt.Printf("s: %v\n", s)
	return db, err
}

func (p *RedisDb) PutUser(user comm.LoginMessage, clientip string) error {

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
	_, err = p.DBclient.HSet(user.UserId, "clientip", user.UserPasswd).Result()
	if err != nil {
		fmt.Printf("sb hset clientip error 2: %v\n", err)
		return err
	}
	*/

	userinfo := make(map[string]interface{})
	userinfo["userid"] = user.UserId
	userinfo["username"] = user.UserName
	userinfo["userpasswd"] = user.UserPasswd

	s, err := p.DBclient.HMSet(clientip, userinfo).Result()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}
	fmt.Printf(" DBclient.HMSet ret string: %v\n", s)
	return nil
}

func (p *RedisDb) DelUser(clientip string) error {

	// 不用差了 直接删
	/* 	i, err := p.DBclient.HMGet(clientip, "userid", "username", "userpasswd").Result()
	   	if err != redis.Nil || i == nil {
	   		fmt.Printf("err: %v\n", err)
	   		return err // 不存在
	   	} */
	_, err := p.DBclient.HDel(clientip).Result()
	if err != redis.Nil {
		fmt.Printf("DBclient.HDel err: %v\n", err)
		return err
	}
	return redis.Nil
}

func (p *RedisDb) FindUser(clientip, username string) (comm.LoginMessage, error) {
	var ret comm.LoginMessage

	i, err := p.DBclient.HMGet(clientip, "userid", "username", "userpasswd").Result()
	if err != nil {
		fmt.Printf("DBclient.HMGet: %v\n", err)
		return ret, nil
	}

	if i[0] == nil {
		return ret, nil
	}

	fmt.Printf("DBclient.HMGet result: %v\n", i)
	if len(i) != 3 {
		return ret, errors.New("hmget len not 3")
	}

	fmt.Println(" dbclient finduser success..")
	ret.UserId = reflect.TypeOf(i[0]).String()
	ret.UserName = reflect.TypeOf(i[1]).String()
	ret.UserPasswd = reflect.TypeOf(i[2]).String()

	return ret, redis.Nil
}
