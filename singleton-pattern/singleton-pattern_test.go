package singleton_pattern

import (
	"fmt"
	"sync"
	"testing"
)

func Test_SingletonPattern(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(100)
	GetInstance().SetValue(999)
	for i := 0; i < 100; i++ {
		go func() {
			fmt.Println(GetInstance().GetValue())
			wg.Done()
		}()
	}
	wg.Wait()
	t.Log("test singleton pass")
}
