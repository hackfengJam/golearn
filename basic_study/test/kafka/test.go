package main

func c() {
	// ast := assert.New(t)
	// req := &sarama.CreateTopicsRequest{
	// 	TopicDetails: map[string]*sarama.TopicDetail{
	// 		"topic": {
	// 			NumPartitions:     -1,
	// 			ReplicationFactor: -1,
	// 			ReplicaAssignment: map[int32][]int32{
	// 				0: []int32{0, 1, 2},
	// 			},
	// 		},
	// 	},
	// 	Timeout: 100 * time.Millisecond,
	// }
	// config := sarama.NewConfig()
	// config.Version = sarama.V1_0_0_0
	// admin, err := sarama.NewClusterAdmin([]string{"127.0.0.1:9092"}, nil)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// err = admin.CreateTopic("hk4e_e20201218answer_app_deploy", &sarama.TopicDetail{
	// 	NumPartitions:     1,
	// 	ReplicationFactor: 1}, false)
	// ast.Nil(err)

}

// func Test(t *testing.T) {
// 	ast := assert.New(t)
// 	// req := &sarama.CreateTopicsRequest{
// 	// 	TopicDetails: map[string]*sarama.TopicDetail{
// 	// 		"topic": {
// 	// 			NumPartitions:     -1,
// 	// 			ReplicationFactor: -1,
// 	// 			ReplicaAssignment: map[int32][]int32{
// 	// 				0: []int32{0, 1, 2},
// 	// 			},
// 	// 		},
// 	// 	},
// 	// 	Timeout: 100 * time.Millisecond,
// 	// }
// 	config := sarama.NewConfig()
// 	config.Version = sarama.V1_0_0_0
// 	admin, err := sarama.NewClusterAdmin([]string{"127.0.0.1:9092"}, config)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	err = admin.CreateTopic("hk4e_e20201218answer_app_deploy", &sarama.TopicDetail{
// 		NumPartitions:     1,
// 		ReplicationFactor: 1,}, false)
// 	ast.Nil(err)
// }

func main() {
}
