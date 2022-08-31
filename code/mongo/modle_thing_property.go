package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ThingReport *mongo.Collection

type ThingProperty struct {
	Key        string      `bson:"key" json:"key"`
	Value      interface{} `bson:"value" json:"value"`
	ProductKey string      `bson:"product_key" json:"product_key"`
	DeviceId   string      `bson:"device_id" json:"device_id"`
	//数据类型 1 设备上报  2 指令下发
	DataType  int   `bson:"data_type" json:"data_type"`
	Timestamp int64 `bson:"timestamp" json:"timestamp"`
}

//TableName
func (tpl ThingProperty) TableName() string {
	return "c_thing_property"
}

func (tpl *ThingProperty) List(pageNum, pageSize int64, m map[string]interface{}) (list []ThingProperty, count int64, err error) {
	findOptions := &options.FindOptions{}
	findOptions.SetSort(map[string]interface{}{"timestamp": -1})
	findOptions.SetLimit(pageSize)
	findOptions.SetSkip(pageSize * (pageNum - 1))
	cursor, err := ThingReport.Find(context.Background(), m, findOptions)

	defer cursor.Close(context.Background())
	if err != nil {
		return list, count, err
	}
	count, err = ThingReport.CountDocuments(context.Background(), m)
	if err != nil {
		return list, count, err
	}
	for cursor.Next(context.Background()) {
		var thingReportLog ThingProperty
		if err = cursor.Decode(&thingReportLog); err != nil {
			return list, count, err
		}

		list = append(list, thingReportLog)
	}
	return list, count, nil
}

//根据设备 时间  统计物模型各属性总次数
func (tpl *ThingProperty) GetDayReport(deviceId string, startTime, endTime int64) (listDay, listTotal []map[string]interface{}, err error) {

	match := bson.M{"device_id": deviceId, "timestamp": bson.M{"$gte": startTime, "$lte": endTime}}
	//28800000 是东八区的偏移时间
	dateBson := bson.M{"$dateToString": bson.M{"format": "%Y-%m-%d", "date": bson.M{"$toDate": bson.M{"$add": bson.A{28800000, bson.M{"$multiply": bson.A{"$timestamp", 1}}}}}}}
	pipeDay := []bson.M{
		{"$match": match},
		{"$project": bson.M{"date": dateBson, "key": 1}},
		{"$group": bson.M{"_id": bson.M{"date": "$date", "key": "$key"}, "count": bson.M{"$sum": 1}}},
		{"$sort": bson.M{"count": -1}},
		{"$limit": 4},
		{"$project": bson.M{"_id": 0, "count": 1, "key": "$_id.key", "date": "$_id.date"}},
	}
	pipeTotal := []bson.M{
		{"$match": match},
		{"$project": bson.M{"date": dateBson, "key": 1}},
		{"$group": bson.M{"_id": bson.M{"key": "$key"}, "count": bson.M{"$sum": 1}}},
		{"$sort": bson.M{"count": -1}},
		{"$limit": 4},
		{"$project": bson.M{"_id": 0, "count": 1, "key": "$_id.key", "date": "$_id.date"}},
	}

	cursorDay, err := ThingReport.Aggregate(context.TODO(), pipeDay)
	if err != nil {
		return
	}
	cursorTotal, err := ThingReport.Aggregate(context.TODO(), pipeTotal)
	if err != nil {
		return
	}
	for cursorDay.Next(context.TODO()) {
		var p1 map[string]interface{}
		if err = cursorDay.Decode(&p1); err != nil {
			return
		}
		listDay = append(listDay, p1)
	}
	for cursorTotal.Next(context.TODO()) {
		var p2 map[string]interface{}
		if err = cursorTotal.Decode(&p2); err != nil {
			return
		}
		listTotal = append(listTotal, p2)
	}

	return
}

type DayReportCount struct {
	Count    int    `bson:"count"`
	Key      string `bson:"key"`
	DateTime string `bson:"date_time"`
}

func (tpl *ThingProperty) CountDocuments(filter map[string]interface{}) (total int64, err error) {
	total, err = ThingReport.CountDocuments(context.Background(), filter)
	return
}
