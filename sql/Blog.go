package sql

type Blog struct {
	id   int
	title  string
	pic  string
	content string
	tags   string
	type   string
	cre
	name string
}

func GetById(id int) (Blog, error) {
	var blog Blog
	db, db_err := Db()
	if db == nil {
		return nil, db_err
	}

	command := "SELECT * FROM blog WHERE id = ?"

	err := db.QueryRow(command, id).Scan(&blog.id, &blog.age, &blog.name)

	if err != nil {
		return nil, err
	}

	return nil
}
