package sql

import "time"

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
	db, db_err := Db()
	if db == nil {
		return blog, db_err
	}

	command := "SELECT id, title, pic, content, tags, type, create_time, update_time, click_num, status FROM blog WHERE id = ?"

	err := db.QueryRow(command, id).Scan(&blog.Id, &blog.Title, &blog.Pic, &blog.Content, &blog.Tags, &blog.Type, &blog.CreateTime, &blog.UpdateTime, &blog.ClickNum, &blog.Status)

	if err != nil {
		return blog, err
	}

	return blog, nil
}
