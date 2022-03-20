package sdf

import (
	"database/sql"
	"fmt"
	"testing"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
	err error
)
type Wp_usermeta struct {
	Umeta_id int `json:"umeta_id"`
	User_id int `json:"user_id"`
	Meta_key string `json:"meta_key"`
	Meta_value string `json:"meta_value"`
}

func painc_err(err error) {
	if err != nil {
		panic(err)
	}
}
func init() {
	// 格式：username:password@链接协议(ip:port)/db?参数
	db, err = sql.Open("mysql", "root:fstxlab@2017!@tcp(123.57.174.21:3306)/horizon?charset=utf8")
	painc_err(err)
}
//查询单条记录
func TestName(t *testing.T) {
	stmt, err := db.Prepare("select * from wp_usermeta where umeta_id = ?")
	defer stmt.Close()  
	painc_err(err)
	row:= stmt.QueryRow(12)
	// 以上为使用预处理方式 查询 。也可以使用非预处理方式查询
	//row := db.QueryRow("select * from wp_usermeta where umeta_id = 12")
	var usermeta Wp_usermeta
	err = row.Scan(&usermeta.Umeta_id, &usermeta.User_id, &usermeta.Meta_key, &usermeta.Meta_value)
	painc_err(err)
	fmt.Println(usermeta)
}
//查询多条记录
func TestName1(t *testing.T) {
	stmt, err := db.Prepare("select * from wp_usermeta where umeta_id < ?")
	defer stmt.Close()
	painc_err(err)
	res, err := stmt.Query(12)
	defer res.Close()
	painc_err(err)
	// 以上为使用预处理方式 查询 。也可以使用非预处理方式查询
	//res, err = db.Query("select * from wp_usermeta where umeta_id < 12")
	for res.Next() {
		var usermeta Wp_usermeta
		err = res.Scan(&usermeta.Umeta_id, &usermeta.User_id, &usermeta.Meta_key, &usermeta.Meta_value)
		painc_err(err)
		fmt.Println(usermeta)
	}
}
//	写入 更新 删除数据 一致 只是sql 不一样
func TestName3(t *testing.T) {
	stmt, err := db.Prepare("insert into wp_usermeta(user_id,meta_key,meta_value) value (?,?,?)")
	painc_err(err)
	defer stmt.Close()
	res, err := stmt.Exec(2,"test3","test4")
	painc_err(err)
	row_num, err := res.RowsAffected() // 受影响的行数
	painc_err(err)
	id,err := res.LastInsertId()	// 最后写入的id
	painc_err(err)
	fmt.Println(row_num,id)
}
// 事务
func Test44(t *testing.T) {
	tx, err := db.Begin()
	err = tx.Commit()
	err = tx.Rollback()
	fmt.Println(err)
}