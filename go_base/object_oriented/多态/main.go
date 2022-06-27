package main

import "fmt"

// User 定义结构体
type User struct {
	ID       int
	Username string
}

// UserDao 定义接口
type UserDao interface {
	getUserById(int) User
	getUserByName(interface{}) interface{}
}

type UserDaoImpl struct {
}

func (udi *UserDaoImpl) getUserById(id int) User {
	return User{ID: id, Username: "username"}
}

func (udi *UserDaoImpl) getUserByName(name interface{}) interface{} {
	username := name.(string)
	return User{ID: 0, Username: username}
}

func main() {
	var userDao UserDao
	userDao = &UserDaoImpl{}
	user := userDao.getUserById(1)
	// user := ans.(User)
	fmt.Printf("%#v\n", user)
}
