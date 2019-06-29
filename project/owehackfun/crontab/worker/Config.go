package worker

import (
	"encoding/json"
	"io/ioutil"
)

// 程序配置
type Config struct {
	EtcdEndpoints         []string `json:"etcdEndpoints"`
	EtcdDialTmeOut        int      `json:"etcdDialTimeOut"`
	MongodbUri            string   `json:"mongodbUri"`
	MongodbConnectTimeOut int      `json:"mongodbConnectTimeOut"`
	JobLogBatchSize       int      `json:"jobLogBatchSize"`
	JobLogCommitTimeout   int      `json:"jobLogCommitTimeout"`
}

var (
	// 单例
	G_config *Config
)

// 加载配置
func InitConfig(filename string) (err error) {
	var (
		content []byte
		conf    Config
	)

	// 1. 读取配置文件
	if content, err = ioutil.ReadFile(filename); err != nil {
		return
	}

	// 2. JSON 反序列化
	if err = json.Unmarshal(content, &conf); err != nil {
		return
	}

	// 3. 赋值单例
	G_config = &conf

	return
}
