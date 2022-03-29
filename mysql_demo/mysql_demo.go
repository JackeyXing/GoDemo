package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	//数据库连接
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/xxj")
	checkErr(err)

	//插入数据
	// db.Prepare()函数用来返回准备要执行的sql操作，然后返回准备完毕的执行状态。
	stmt, err := db.Prepare("INSERT userinfo SET username=?, departname=?,created=?")
	checkErr(err)

	res, err := stmt.Exec("JackeyXing", "研发部门", "2022-3-29")
	checkErr(err)
	// LastInsertId 返回数据库为响应命令而生成的整数。 通常，这将来自插入新行时的“自动增量”列。
	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
	//更新数据
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)
	// stmt.Exec()函数用来执行stmt准备好的SQL语句
	res, err = stmt.Exec("JackeyXingUpdate", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	//查询数据
	// db.Query()函数用来直接执行Sql返回Rows结果。
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	//删除数据
	stmt, err = db.Prepare("delete from userinfo where username=?")
	checkErr(err)

	res, err = stmt.Exec("JackeyXingUpdate")
	checkErr(err)

	// RowsAffected 返回受更新、插入或删除影响的行数。 并非每个数据库或数据库驱动程序都支持这一点。
	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	defer db.Close()

}
