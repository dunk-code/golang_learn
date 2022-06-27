package main

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"strings"
)

type User struct {
	ID       int    `db:"id"`
	UserName string `db:"username"`
	NickName string `db:"nickname"'`
	Picture  string `db:"picture"`
}

func main() {
	database, err := sqlx.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/entry_task")
	if err != nil {
		panic(err)
	}
	var sqlStr = "SELECT id,username,nickname,picture FROM t_user WHERE id in ( ? )"
	ids := []int{1, 2, 3, 4}
	query, args, err := sqlx.In(sqlStr, ids)
	if err != nil {
		panic(err)
	}
	fmt.Println(query)
	fmt.Println(args)
	rows, err := database.Query(query, args...)
	defer rows.Close()
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var id int
		var username, nickname, picture string
		err := rows.Scan(&id, &username, &nickname, &picture)
		user := User{UserName: username, ID: id, NickName: nickname, Picture: picture}
		if err != nil {
			panic(err)
		}
		fmt.Println(user)
	}
	var person []User
	ctx := context.Background()
	err = database.SelectContext(ctx, &person, query, args...)
	if err != nil {
		panic(err)
	}
	fmt.Println(person)
	strs := []string{"detail", "read"}
	join := strings.Join(strs, ".")
	fmt.Println(join)
}
