package util

import (
	"testing"
)

func TestEncrypt(t *testing.T) {
	_, err := Encrypt("12345")
	if err != nil {
		t.Error(err)
	}

	t.Log("TestEncrypt test pass")
}

func BenchmarkEncrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Encrypt("12345")
	}
}

func BenchmarkGenShortIdTimeConsuming(b *testing.B) {
	b.StopTimer() //调用该函数停止压力测试的时间计数

	_, err := Encrypt("12345")
	if err != nil {
		b.Error(err)
	}

	b.StartTimer() //重新开始时间

	for i := 0; i < b.N; i++ {
		Encrypt("12345")
	}
}
