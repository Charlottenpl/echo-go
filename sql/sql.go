package sql

import (
	"database/sql"
	"fmt"
	//执行driver.go文件中的init(),向"database/sql"注册一个mysql的驱动
	_ "github.com/go-sql-driver/mysql"
)

var (
	// 定义一个全局对象db
	db *sql.DB
	//定义数据库连接的相关参数值
	//连接数据库的用户名
	userName string = "root"
	//连接数据库的密码
	password string = "admin"
	//连接数据库的地址
	ipAddress string = "127.0.0.1"
	//连接数据库的端口号
	port int = 3306
	//连接数据库的具体数据库名称
	dbName string = "go_test"
	//连接数据库的编码格式
	charset string = "utf8"
)

func Init() {
	dsn := "root:1234@tcp(117.50.187.91:3306)/echo?charset=utf8"
	//Open打开一个driverName指定的数据库，dataSourceName指定数据源
	//不会校验用户名和密码是否正确，只会对dsn的格式进行检测
	Db, err := sql.Open("mysql", dsn)
	db = Db
	if err != nil { //dsn格式不正确的时候会报错
		fmt.Printf("打开数据库失败,err:%v\n", err)
		return
	}
	//尝试连接数据库，Ping方法可检查数据源名称是否合法,账号密码是否正确。
	err = db.Ping()
	if err != nil {
		fmt.Printf("连接数据库失败,err:%v\n", err)
		return
	}
	fmt.Println("连接数据库成功！")
}
