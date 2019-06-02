package common

import (
	"encoding/json"
	"strings"
)

// 定时任务
type Job struct {
	Name     string `json:"name"`     // 任务名
	Command  string `json:"command"`  // shell 命令
	CronExpr string `json:"cronExpr"` // cron 表达式
}

// HTTP 接口响应
type Response struct {
	Errno int         `json:"errno"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
}
type JobEvent struct {
	EventType int // SAVE, DELETE
	Job       *Job
}

// 应答方法
func BuildResponse(errno int, msg string, data interface{}) (resp []byte, err error) {
	// 1. 定义一个 Response
	var (
		response Response
	)

	response.Errno = errno
	response.Msg = msg
	response.Data = data

	// 2. 序列化 json
	resp, err = json.Marshal(response)

	return
}

// 反序列化 Job
func UnpackJob(value []byte) (ret *Job, err error) {
	var (
		job *Job
	)
	job = &Job{}

	if err = json.Unmarshal(value, job); err != nil {
		return
	}
	ret = job
	return
}

// 从etcd 的 key 提取任务名
// /cron/jobs/job10 抹掉 /cron/jobs
func ExtractJobName(jobKey string) (string) {
	return strings.TrimPrefix(jobKey, JOB_SAVE_DIR)
}

// 任务变化事件有两种 1. 更新任务 2. 删除任务
func BuildJobEvent(eventType int, job *Job) (jobEvent *JobEvent) {
	return &JobEvent{
		EventType: eventType,
		Job:       job,
	}
}
