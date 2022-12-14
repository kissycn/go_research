package chapter9_concurrent

import (
	"fmt"
	"sync/atomic"
	"testing"
)

var (
	// 序列号
	seq int64
)

// 序列号生成器
func GenID() int64 {
	// 尝试原子的增加序列号
	return atomic.AddInt64(&seq, 1)
}

func TestAtomic(t *testing.T) {
	//生成10个并发序列号
	for i := 0; i < 10; i++ {
		go GenID()
	}
	fmt.Println(GenID())
}
