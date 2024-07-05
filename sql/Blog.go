package sql

import (
	"fmt"
	"time"
)

type Blog struct {
	Id         int       // id
	Title      string    // 标题
	Pic        string    // 展示图片
	Content    string    // 内容
	Type       string    // 分类
	CreateTime time.Time // 创建时间
	UpdateTime time.Time // 更新时间
	ClickNum   int       // 点击数
	Status     int       // 当前状态
}

// GetById 根据ID获取 ✅
func GetById(id int) (Blog, error) {
	var blog Blog
	var ct, ut string
	db, db_err := Db()
	if db == nil {
		return blog, db_err
	}

	command := "SELECT id, title, pic, content, type, create_time, update_time, click_num, status FROM blog WHERE id = ?"
	err := db.QueryRow(command, id).Scan(&blog.Id, &blog.Title, &blog.Pic, &blog.Content, &blog.Type, &ct, &ut, &blog.ClickNum, &blog.Status)

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

// GetByType 根据分类获取
func GetByType(_type int) ([]Blog, error) {
	var blogs []Blog

	db, db_err := Db()
	if db == nil {
		return nil, db_err
	}

	command := "SELECT id, title, pic, content, type, create_time, update_time, click_num, status FROM blog WHERE type = ?"
	rows, err := db.Query(command, _type)
	//err := db.Query(command, _type).Scan(&blog.Id, &blog.Title, &blog.Pic, &blog.Content, &blog.Type, &ct, &ut, &blog.ClickNum, &blog.Status)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var blog Blog
		var ct, ut string

		// 转化完是默认的 RFC3339 格式（"2024-04-29T18:43:40Z"）
		// T 表示日期和时间的分隔符，Z 表示 UTC 时间（即零时区）
		ctime, _ := time.Parse("2006-01-02 15:04:05", ct)
		utime, _ := time.Parse("2006-01-02 15:04:05", ut)

		//如果想自定义格式的字符串可以使用time.Format("2006-01-02 15:04:05")
		blog.CreateTime = ctime
		blog.UpdateTime = utime
		fmt.Println(blog)

		blogs = append(blogs, blog)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return blogs, nil
}

// Find 需要查tags和blog的链接表
func Find(title string, bt int64, et int64, content string, tag string, _type int, status int, limit, offset int) ([]*Blog, error) {
	// 构建 SQL 查询语句
	query := "SELECT id, title, pic, content, type, create_time, update_time, click_num, status FROM blogs WHERE 1=1"
	var args []interface{}

	// 添加条件
	if title != "" {
		query += " AND title LIKE CONCAT('%', ?, '%')"
		args = append(args, title)
	}
	if bt > 0 {
		begin := time.Unix(bt, 0)
		query += " AND created_at >= ?"
		args = append(args, begin)
	}
	if et > 0 {
		query += " AND created_at <= ?"
		args = append(args, et)
	}
	if content != "" {
		query += " AND MATCH (content) AGAINST (?)"
		args = append(args, content)
	}
	if tag != "" {
		query += " AND tag = ?"
		args = append(args, tag)
	}
	if _type > 0 {
		query += " AND type = ?"
		args = append(args, _type)
	}
	if status != 0 {
		query += " AND status = ?"
		args = append(args, status)
	}

	// 添加分页限制
	query += " LIMIT ? OFFSET ?"
	args = append(args, limit, offset)

	// 执行查询
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var blogs []*Blog
	for rows.Next() {
		var blog Blog
		var ct, ut string
		err := rows.Scan(&blog.Id, &blog.Title, &blog.Pic, &blog.Content, &blog.Type, &ct, &ut, &blog.ClickNum, &blog.Status)
		if err != nil {
			return nil, err
		}

		ctime, err := time.Parse("2006-01-02 15:04:05", ct)
		utime, err := time.Parse("2006-01-02 15:04:05", ut)

		blog.CreateTime = ctime
		blog.UpdateTime = utime
		blogs = append(blogs, &blog)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return blogs, nil
}

// ListPage 获取博客列表分页
// page 页数
// size 每页个数
func ListPage(page, size int) ([]*Blog, error) {
	offset := (page - 1) * size
	command := "SELECT id, title, pic, content, type, create_time, update_time, click_num, status FROM blog LIMIT ? OFFSET ?"
	rows, err := db.Query(command, size, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var blogs []*Blog
	for rows.Next() {
		var blog Blog
		var ct, ut string
		err := rows.Scan(&blog.Id, &blog.Title, &blog.Pic, &blog.Content, &blog.Type, &ct, &ut, &blog.ClickNum, &blog.Status)
		if err != nil {
			return nil, err
		}

		ctime, err := time.Parse("2006-01-02 15:04:05", ct)
		utime, err := time.Parse("2006-01-02 15:04:05", ut)

		blog.CreateTime = ctime
		blog.UpdateTime = utime
		blogs = append(blogs, &blog)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return blogs, nil
}

// Add 添加
func Add(blog Blog) (int64, error) {
	colums := ""
	values := ""
	var args []interface{}

	// 添加条件 title, pic, content, type, click_num, status
	if blog.Title != "" {
		colums += "title, "
		values += "?, "
		args = append(args, blog.Title)
	}

	if blog.Pic != "" {
		colums += "pic, "
		values += "?, "
		args = append(args, blog.Pic)
	}

	if blog.Content != "" {
		colums += "content, "
		values += "?, "
		args = append(args, blog.Content)
	}

	if blog.Type != "" {
		colums += "type, "
		values += "?, "
		args = append(args, blog.Type)
	}

	colums += "status"
	values += "0"

	command := "INSERT INTO blog (" + colums + ") values (" + values + ")"

	// !使用...运算符展开args
	result, err := db.Exec(command, args...)
	if err != nil {
		return 0, err
	}

	var newID int64
	newID, err = result.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err: %v\n", err)
		return 0, err
	}
	return newID, nil
}
