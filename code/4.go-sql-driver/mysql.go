package sdf

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

// go-sql-driver 最大的问题  scan 是按照顺序 解析的  一旦数据库结构改变  将会全部错乱  代码直接报错
//  不封装mysql 的事务   query 相关的stmt必须在commit 或者rollback之后才能关闭  封装并不能defer close 意义不大  用原生db写吧

var Db *sql.DB
var err error

func init() {
	Db, err = sql.Open("mysql", "root:123456@tcp(42.194.133.108:3306)/bbs")
	Db.SetMaxOpenConns(99)                                 // 最大连接数
	Db.SetMaxIdleConns(90)                                 //最大空闲连接数
	Db.SetConnMaxLifetime(time.Duration(time.Second * 60)) // 连接超时世间
	Db.SetConnMaxIdleTime(time.Duration(time.Second * 60)) // 连接空闲持续时间

	//checkErr(err)
}

func Exec(sqlString string, parames ...interface{}) (sql.Result, error) {
	var stmt *sql.Stmt
	stmt, err := Db.Prepare(sqlString)
	if err != nil {
		return nil, err
	}
	defer stmt.Close() //busy buffer  bad connection 错误的始作俑者  必须在tx.commit 之后再关闭
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
	defer stmt.Close() //busy buffer  bad connection 错误的始作俑者  必须在tx.commit 之后再关闭

	if limit != 0 && page != 0 {
		parames = append(parames, (page-1)*limit, limit)
	}
	rows, err := stmt.Query(parames...)
	if err != nil {
		panic(err)
	}
	return rows, err
}

func QueryNoParame(sqlString string) (*sql.Rows, error) {
	rows, err := Db.Query(sqlString)
	if err != nil {
		return nil, err
	}
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
	row := stmt.QueryRow(parames...)
	return row, nil
}
