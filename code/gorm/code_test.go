package gorm

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"log"
	"os"
	"runtime"
	"strconv"
	"testing"
	"time"
)

var DB *gorm.DB

//自定义 时间 替换 create update time 的默认 time.time 解决读取数据时，读出来的时间是没有格式化的time问题
type LocalTime time.Time

func (t *LocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02 15:04:05"))), nil
}
func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	tlt := time.Time(t)
	//判断给定时间是否和默认零时间的时间戳相同
	if tlt.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return tlt, nil
}
func (t *LocalTime) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = LocalTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

type sqlWrite struct {
	l *log.Logger
}

//New 初始化 Logger
func (s *sqlWrite) New(f io.Writer, prefix string, flag int) {
	s.l = log.New(f, prefix, flag) // io writer
	return
}

//Printf 实现gorm打印接口
func (s *sqlWrite) Printf(format string, a ...interface{}) {
	_, file, line, ok := runtime.Caller(6)
	if ok {
		format = file + ":" + strconv.FormatInt(int64(line), 10) + "\n" + format
	}
	s.l.Output(2, fmt.Sprintf(format, a...))
	return
}

func init() {

	dsn := "root:djf136169@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//enable Gorm mysql log
	f, err := os.OpenFile("gorm.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("could not create mysql gorm log file", err)
		return
	}
	//初始化
	mySqlWrite := &sqlWrite{}
	mySqlWrite.New(f, "\r\n", log.LstdFlags)
	newLogger := logger.New(
		mySqlWrite,
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)
	db.Logger = newLogger
	DB = db
}

type Product struct {
	Model
	Code  string
	Price uint
}
type Model struct {
	ID        uint `gorm:"primarykey" json:"id"`
	CreatedAt LocalTime
	UpdatedAt LocalTime      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func Test1(t *testing.T) {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情

	// 迁移 schema
	DB.AutoMigrate(&Product{})

	// Create
	DB.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	//product.ID = 3
	DB.First(&product, 1)                 // 根据整型主键查找
	DB.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录
	DB.First(&product)
	//Update - 将 product 的 price 更新为 200
	DB.Model(&product).Update("Price", 200)
	// Update - 更新多个字段
	DB.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	DB.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 product
	DB.Delete(&product)
	jsons, errs := product.CreatedAt.MarshalJSON()
	fmt.Println(string(jsons), errs)
}

type User struct {
	ID           uint
	Name         string
	Email        string
	Age          uint8
	Birthday     time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// 查询
func Test2(t *testing.T) {
	DB.AutoMigrate(&User{})
	//user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}

	//result := DB.Create(&user) // 通过数据的指针来创建
	//
	//fmt.Println(user.ID)             // 返回插入数据的主键
	//fmt.Println(result.Error)        // 返回 error
	//fmt.Println(result.RowsAffected) // 返回插入记录的条数
	//var user *User
	//DB.First(&user)
	////虽然user 数据完整 每个字段都有数据  但是select 之后 只会取select的字段 进行插入
	////INSERT INTO `users` (`name`,`age`,`created_at`,`updated_at`) VALUES ('Jinzhu',18,'2022-07-30 23:57:57.514','2022-07-30 23:57:57.514')
	//DB.Select("Name", "Age", "CreatedAt").Create(&user)
	//// 与select 相反的是omit 忽略字段 其他字段全部插入
	//DB.Omit("Name", "Age", "CreatedAt").Create(&user)
	//
	//批量写入
	//INSERT INTO `users` (`name`,`email`,`age`,`birthday`,`member_number`,`activated_at`,`created_at`,`updated_at`) VALUES ('jinzhu1','',0,'0000-00-00 00:00:00',NULL,NULL,'2022-07-31 00:16:58.292','2022-07-31 00:16:58.292'),('jinzhu2','',0,'0000-00-00 00:00:00',NULL,NULL,'2022-07-31 00:16:58.292','2022-07-31 00:16:58.292'),('jinzhu3','',0,'0000-00-00 00:00:00',NULL,NULL,'2022-07-31 00:16:58.292','2022-07-31 00:16:58.292')
	var users = []User{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
	DB.Create(&users)

	for _, user := range users {
		fmt.Println(user.ID) // 1,2,3
	}

}
