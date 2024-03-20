package model

import (
	"encoding/json"
	"fmt"

	"github.com/28267/chatroom/common/message"

	_ "github.com/go-redis/redis"
	"github.com/gomodule/redigo/redis"
)

var (
	MyUserDao *UserDao
)

type UserDao struct {
	pool *redis.Pool
}

func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

func (userDao *UserDao) GetUserById(conn redis.Conn, id int) (user *User, err error) {

	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		if err == redis.ErrNil {
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	user = &User{}
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(res),user)-错误,err=", err)
	}
	return

}

func (userDao *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	conn := userDao.pool.Get()
	defer conn.Close()
	user, err = userDao.GetUserById(conn, userId)
	if err != nil {
		err = ERROR_USER_NOTEXISTS
		return
	}
	// if err == nil {
	// 	return
	// } else {
	// 	err = ERROR_USER_NOTEXISTS
	// }
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}

func (userDao *UserDao) Register(user *message.User) (err error) {
	conn := userDao.pool.Get()
	defer conn.Close()
	_, err = userDao.GetUserById(conn, user.UserId)
	if err == nil {
		err = ERROR_USER_EXISTS //用户已存在,无需注册
		return
	}
	data, err := json.Marshal(user)
	if err != nil {
		return
	}
	_, err = conn.Do("HSet", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("保存注册用户错误err=", err)
		return
	}
	return
}
