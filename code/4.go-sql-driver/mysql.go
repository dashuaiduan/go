package sdf

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// go-sql-driver 最大的问题  scan 是按照顺序 解析的  一旦数据库结构改变  将会全部错乱  代码直接报错
var Db *sql.DB

func init() {
	database, err := sql.Open("mysql", "root:123456@tcp(42.194.133.108:3306)/bbs")
	//checkErr(err)
	if err != nil {
		panic("连接数据库错误")
	}
	Db = database
}

func Exec(sqlString string, parames ...interface{}) (sql.Result, error) {
	stmt, err := Db.Prepare(sqlString)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(parames...)
	// 判断修改数据 成功失败 应该以err 是不是为空 为判断依据， 返回res 用于获取其他数据
	return res, err
}

//普通 query 函数封装  limit page 不用的话 需要传递 0,0
func Query(sqlString string, limit int, page int, parames ...interface{}) (*sql.Rows, error) {
	stmt, err := Db.Prepare(sqlString)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var rows *sql.Rows
	if limit != 0 && page != 0 {
		parames = append(parames, (page-1)*limit, limit)
	}
	rows, err = stmt.Query(parames...)
	if err != nil {
		panic(err)
	}
	return rows, err
}

func QueryNoParame(sqlString string) (*sql.Rows, error) {
	rows, err := Db.Query(sqlString)
	//defer rows.Close()   函数调用方还需要用到该资源 此处不能关闭，不然外面获取不到值  所以必须在调用处 defer close
	return rows, err
}

//查询一行记录
func QueryRow(sqlString string, parames ...interface{}) (*sql.Row, error) {
	stmt, err := Db.Prepare(sqlString)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var row *sql.Row
	row = stmt.QueryRow(parames...)
	return row, err
}
