package main

import (
	"database/sql"
	"fmt"
	"_github.com/go-sql-driver/mysql"
)

var db *sql.DB


func main (){
	caidan :="root:@tcp(127.0.0.1:3306)/user?charset=utf8"

	var err error

	db,err =sql.Open("mysql",caidan)

	if err !=nil {
		fmt.Println("错误")
	}
	stmt,err := db.Prepare("insert into `caidan`(cainame,caimoney)values(?,?)")
	if err!=nil{
		fmt.Println("预处理失败:",err)
		return
	}
	result,err := stmt.Exec("小面","6元")
	if err!=nil{
		fmt.Println("执行预处理失败:",err)
		return
	}else{
		rows,_ := result.RowsAffected()
		fmt.Println("执行成功,影响行数",rows,"行" )
	}


	result, err = stmt.Exec("饭团", "7元")
	if err!=nil{
		fmt.Println("执行预处理失败:",err)
		return
	}else{
		rows,_ := result.RowsAffected()
		fmt.Println("执行成功,影响行数",rows,"行" )
	}

	result, err = stmt.Exec("香菇滑鸡", "12元")
	if err!=nil{
		fmt.Println("执行预处理失败:",err)
		return
	}else{
		rows,_ := result.RowsAffected()
		fmt.Println("执行成功,影响行数",rows,"行" )
	}

	rows,err:=db.Query("select  from `cainame`")

	defer rows.Close()
	if err !=nil{
		fmt.Println("获取数据失败")
		return
	}
	err =db.Ping()
	if err !=nil{
		fmt.Println("打开菜单出错")
	}


	go da()

}

func da(){
	var cainame string
	fmt.Println("输入的菜名")
	fmt.Scan("%s",cainame)
	fmt.Printf("a=%s",cainame)
}