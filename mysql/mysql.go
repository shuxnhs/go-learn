package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"	// 引入，虽然不使用但是需要使用他的初始化
	"github.com/jmoiron/sqlx"
)

type Person struct {
	Name string `db:"name"`
	Age int `db:"age"`

}


func main()  {
	// 连接数据库
	db, err := sqlx.Open("mysql", "root:password@tcp(localhost:3306)/hxh_test")
	if err == nil{
		fmt.Println("连接成功")
		/**-----------开始数据库的操作新增----------------**/
		result,e := db.Exec("insert into person(name,age) values (?,?);", "hxh", 22)
		if e == nil{
			Rows, _ := result.RowsAffected()
			Id, _ := result.LastInsertId()
			fmt.Println("添加成功", )
			fmt.Println("影响的行数", Rows)
			fmt.Println("插入的id：", Id)
		}else {
			fmt.Println("添加失败，err：",e)
		}

		/**-----------开始数据库的操作查询----------------**/
		var personContainer []Person	// 先创建个切片容器来存储
		e = db.Select(&personContainer, "select name, age from person where id = ?", 1)
		if e == nil{
			fmt.Println("查询成功", personContainer)
		}else {
			fmt.Println("查询失败，原因：",e)
		}

		/**-----------开始数据库的操作更新----------------**/
		result1, e1 := db.Exec("update person set age = ? where id = ?", 20,1)
		if e1 == nil{
			Rows, _ := result1.RowsAffected()
			Id, _ := result1.LastInsertId()
			fmt.Println("更新成功")
			fmt.Println("影响的行数", Rows)
			fmt.Println("插入的id：", Id)
		}else {
			fmt.Println("更新失败，err：",e1)
		}

		/**-----------开始数据库的操作删除----------------**/
		result2, e2 := db.Exec("delete from person where id = (select max(id) from person)")
		if e2 == nil{
			Rows, _ := result2.RowsAffected()
			Id, _ := result2.LastInsertId()
			fmt.Println("删除成功")
			fmt.Println("影响的行数", Rows)
			fmt.Println("删除的id：", Id)
		}else {
			fmt.Println("删除失败，err：",e2)
		}


	}else {
		fmt.Println("连接失败，err：", err)
	}
	defer func() {
		_ = db.Close()
	}()


}
