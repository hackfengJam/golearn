package main

import (
	"encoding/base64"
	"encoding/json"
	"reflect"
	"sync"
	"time"

	"github.com/google/uuid"
)

// TaskMessage is celery-compatible message
type TaskMessage struct {
	ID      string                 `json:"id"`
	Task    string                 `json:"task"`
	Args    []interface{}          `json:"args"`
	Kwargs  map[string]interface{} `json:"kwargs"`
	Retries int                    `json:"retries"`
	ETA     *string                `json:"eta"`
}

func (tm *TaskMessage) reset() {
	tm.ID = uuid.Must(uuid.New(), nil).String()
	tm.Task = ""
	tm.Args = nil
	tm.Kwargs = nil
}

var taskMessagePool = sync.Pool{
	New: func() interface{} {
		eta := time.Now().Format(time.RFC3339)
		return &TaskMessage{
			ID:      uuid.Must(uuid.New(), nil).String(),
			Retries: 0,
			Kwargs:  nil,
			ETA:     &eta,
		}
	},
}

func getTaskMessage(task string) *TaskMessage {
	msg := taskMessagePool.Get().(*TaskMessage)
	msg.Task = task
	msg.Args = make([]interface{}, 0)
	msg.Kwargs = make(map[string]interface{})
	msg.ETA = nil
	return msg
}

func releaseTaskMessage(v *TaskMessage) {
	v.reset()
	taskMessagePool.Put(v)
}

// DecodeTaskMessage decodes base64 encrypted body and return TaskMessage object
func DecodeTaskMessage(encodedBody string) (*TaskMessage, error) {
	body, err := base64.StdEncoding.DecodeString(encodedBody)
	if err != nil {
		return nil, err
	}
	message := taskMessagePool.Get().(*TaskMessage)
	err = json.Unmarshal(body, message)
	if err != nil {
		return nil, err
	}
	return message, nil
}

// Encode returns base64 json encoded string
func (tm *TaskMessage) Encode() (string, error) {
	jsonData, err := json.Marshal(tm)
	if err != nil {
		return "", err
	}
	encodedData := base64.StdEncoding.EncodeToString(jsonData)
	return encodedData, err
}

// ResultMessage is return message received from broker
type ResultMessage struct {
	ID        string        `json:"task_id"`
	Status    string        `json:"status"`
	Traceback interface{}   `json:"traceback"`
	Result    interface{}   `json:"result"`
	Children  []interface{} `json:"children"`
}

func (rm *ResultMessage) reset() {
	rm.Result = nil
}

var resultMessagePool = sync.Pool{
	New: func() interface{} {
		return &ResultMessage{
			Status:    "SUCCESS",
			Traceback: nil,
			Children:  nil,
		}
	},
}

func getResultMessage(val interface{}) *ResultMessage {
	msg := resultMessagePool.Get().(*ResultMessage)
	msg.Result = val
	return msg
}

func getReflectionResultMessage(val *reflect.Value) *ResultMessage {
	msg := resultMessagePool.Get().(*ResultMessage)
	// msg.Result = GetRealValue(val)
	msg.Result = val
	return msg
}

func releaseResultMessage(v *ResultMessage) {
	v.reset()
	resultMessagePool.Put(v)
}

// GetRealValue returns real value of reflect.Value
// Required for JSON Marshalling
func GetRealValue(val *reflect.Value) interface{} {
	if val == nil {
		return nil
	}
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return val.Int()
	case reflect.String:
		return val.String()
	case reflect.Bool:
		return val.Bool()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return val.Uint()
	case reflect.Float32, reflect.Float64:
		return val.Float()
	default:
		return val.Interface()
	}
}
