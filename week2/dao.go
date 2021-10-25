package main

// 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，
// 是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

// 答：个人认为应该将 sql.ErrNoRows 抛给上层，但是同时需要加上对应的上下文信息
import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

type User struct {
	Id   uint64 `json:"Id"`
	Name string `json: "name"`
}

func main() {
	res, err := GetUserById(4)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}

func GetUserById(id uint64) (user *User, err error) {
	user = &User{Id: id}
	db, err := sql.Open("mysql", "root:@(127.0.0.1:3306)/test")
	if err != nil {
		panic("数据库链接失败")
	}

	err = db.Ping()
	if err != nil {
		panic("未找到数据库")
	}
	defer db.Close()

	sqlStr := "select id, name from user where id=?"
	err = db.QueryRow(sqlStr, id).Scan(&user.Id, &user.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.Wrap(err, fmt.Sprintf("find user null,user id: %v", id))
		} else {
			return nil, errors.Wrap(err, fmt.Sprintf("query faild when find user,user id: %v", id))
		}
	}
	return user, nil
}
