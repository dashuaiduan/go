package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
)

var MongoDB *mongo.Client

type MongoOpt struct {
	Addr     string
	UserName string
	Password string
}

func NewMongo(mo MongoOpt) error {
	// Set Client options
	clientOptions := options.Client().ApplyURI(mo.Addr)
	if mo.UserName != "" {
		clientOptions.Auth = &options.Credential{Username: mo.UserName, Password: mo.Password}
	}
	// Connect to MongoDB
	var err error
	MongoDB, err = mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return err
	}
	// Check the connection
	if err = MongoDB.Ping(context.TODO(), nil); err != nil {
		return err
	}
	fmt.Println("mongoDB run success!")

	return nil
}

func TestName(t *testing.T) {
	var mongoOpt = MongoOpt{
		Addr:     "mongodb://129.211.102.208:27017",
		UserName: "worthcloud",
		Password: "worthcloud2022",
	}

	//初始化mongo
	if err := NewMongo(mongoOpt); err != nil {
		panic(fmt.Sprintf("初始化mongo客户端 err: %v", err))
	}
	database := "worthcloud"
	ThingReportCool := MongoDB.Database(database).Collection(ThingProperty{}.TableName())
	_, err := ThingReportCool.InsertOne(context.TODO(), map[string]interface{}{"key": "666", "value": "ffdfd"})
	fmt.Println(err)
	//	其他示例  查看modle_thing_property model文件
}
