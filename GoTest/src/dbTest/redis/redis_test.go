package redist_test

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {
	logger := log.Logger{}

	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	ping := client.Ping(context.Background())
	if ping.Err() != nil {
		logger.Println("ping error")
	}

	// watch自带事务
	err := client.Watch(context.Background(), func(tx *redis.Tx) error {
		res := client.HSetNX(context.Background(), "name", "1", "han")
		return res.Err()
	})
	if err != nil {
		logger.Println("ping error")
		return
	}
	all := client.HGetAll(context.Background(), "name")
	fmt.Println(all)

	client.HSet(context.Background(), "name", "han", "tuo")
	res, err := client.Get(context.Background(), "aeqfqwefurfqrg").Uint64()
	fmt.Println("===============================", err, res)
	if err != redis.Nil {
		return
	}

	err = client.Set(context.Background(), "123", "123", 0).Err()
	fmt.Println(err)
	err = client.SetNX(context.Background(), "123", "321", 0).Err()
	fmt.Println("++++++++++++++++++++++++++", err, client.Get(context.Background(), "123").String())

	fmt.Println(fmt.Sprintf("1123_"+"%d", 123))

}

func TestChan1(t *testing.T) {

	c := make(chan int, 12)
	for i := 0; i < 12; i++ {
		c <- i
	}
	close(c)

	for i := 0; i < 13; i++ {
		t, ok := <-c
		fmt.Println(len(c), t, ok)
	}
}

func TestChan(t *testing.T) {
	c := make(chan int, 5)
	for i := 0; i < 5; i++ {
		c <- i
	}

	var wg sync.WaitGroup
	wg.Add(runtime.NumCPU() * 34)
	for i := 0; i < runtime.NumCPU()*34; i++ {
		go func() {
			defer func() {
				wg.Done()
			}()
			for {
				if len(c) < 1 {
					return
				}
				time.Sleep(time.Second)
				v, ok := <-c
				if ok && len(c) == 0 {
					close(c)
				}

				fmt.Println(ok)

				fmt.Println(v)
			}
		}()
	}
	wg.Wait()
	fmt.Println("---------")
}
func TestSlice(t *testing.T) {
	var a []*int
	var b *int
	*b = 10
	a = append(a, b)

}

func TestABC(t *testing.T) {
	var a = []int64{2, 3, 1}
	for x := range a {
		fmt.Println(x)
	}
}

var mutex = &sync.Mutex{}

func TestDefer(t *testing.T) {
	b := []int{1, 2, 3, 4, 5}
	var a =make(map[string][]int)

	a["123"]=b
	a["213213213"]=b
	fmt.Println(a["123"])

}

//func aaa(){
//	defer mutex.
//
//}
