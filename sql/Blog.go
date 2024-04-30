package sql

import (
	"fmt"
	"time"
)

type Blog struct {
	Id         int
	Title      string
	Pic        string
	Content    string
	Tags       string
	Type       string
	CreateTime time.Time
	UpdateTime time.Time
	ClickNum   int
	Status     int
}

func GetById(id int) (Blog, error) {
	var blog Blog
	var ct, ut string
	db, db_err := Db()
	if db == nil {
		return blog, db_err
	}

	command := "SELECT id, title, pic, content, tags, type, create_time, update_time, click_num, status FROM blog WHERE id = ?"
	err := db.QueryRow(command, id).Scan(&blog.Id, &blog.Title, &blog.Pic, &blog.Content, &blog.Tags, &blog.Type, &ct, &ut, &blog.ClickNum, &blog.Status)

	if err != nil {
		fmt.Println(err)
		return blog, err
	}

	// 转化完是默认的 RFC3339 格式（"2024-04-29T18:43:40Z"）
	// T 表示日期和时间的分隔符，Z 表示 UTC 时间（即零时区）
	ctime, err := time.Parse("2006-01-02 15:04:05", ct)
	utime, err := time.Parse("2006-01-02 15:04:05", ut)

	//如果想自定义格式的字符串可以使用time.Format("2006-01-02 15:04:05")
	blog.CreateTime = ctime
	blog.UpdateTime = utime
	fmt.Println(blog)

	return blog, nil
}
