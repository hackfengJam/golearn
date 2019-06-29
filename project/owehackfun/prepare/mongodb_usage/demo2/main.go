package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// 任务的执行时间点
type TimePoint struct {
	StartTime int64 `bson:"startTime"`
	EndTime   int64 `bson:"endTime"`
}

// 一条日志
type LogRecord struct {
	JobName   string    `bson:"jobName"`   // 任务名
	Command   string    `bson:"command"`   // shell 命令
	Err       string    `bson:"err"`       // 脚本错误
	Content   string    `bson:"content"`   // 脚本输出
	TimePoint TimePoint `bson:"timePoint"` // 执行时间点
}

func main() {
	var (
		client     *mongo.Client
		err        error
		database   *mongo.Database
		collection *mongo.Collection
		record     *LogRecord
		result     *mongo.InsertOneResult
	)
	// 1. 建立连接
	/*
		if client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017")); err != nil {
			fmt.Println(err)
			return
		}
		ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	*/
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	if client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017")); err != nil {
		fmt.Println(err)
		return
	}
	// 2. 选择数据库
	//database = client.Database("cron")
	database = client.Database("test")

	// 3. 选择表 my_collection
	//collection = database.Collection("log")
	collection = database.Collection("my_collection")

	// 4. 插入记录(bson)
	record = &LogRecord{
		JobName:   "job 10",
		Command:   "echo hello",
		Err:       "",
		Content:   "hello",
		TimePoint: TimePoint{StartTime: time.Now().Unix(), EndTime: time.Now().Unix() + 10},
	}

	if result, err = collection.InsertOne(context.TODO(), record); err != nil {
		fmt.Println(err)
		return
	}

	// id: 默认生成一个全局唯一ID，ObjectID： 12 字节的二进制
	// TODO
	docId := result.InsertedID
	fmt.Println("自增ID：", docId)
}
