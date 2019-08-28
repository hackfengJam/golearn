package main

import (
	"testing"
)

func BenchmarkQueuedSchedulerSpider(b *testing.B)  {
	for i:=0;i<b.N; i++{
		QueuedSchedulerSpider()
	}
}


func BenchmarkSimpleSchedulerSpider(b *testing.B)  {
	//for i:=0;i<b.N; i++{
	for i:=0;i<3; i++{
		SimpleSchedulerSpider()
	}
}


func BenchmarkSimpleEngineSpider(b *testing.B)  {
	for i:=0;i<b.N; i++{
		SimpleEngineSpider()
	}
}

