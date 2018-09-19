/**
 * 实现Pool redis数据库操作基本初始化需求
 * wangruiming create 2017.10.26
 */
package store

import (
	"log"
	"testing"
)

var redisAddr []string = []string{"192.168.14.75:6379", "192.168.14.76:6379",
	"192.168.14.77:6379",
	"192.168.14.75:6380",
	"192.168.14.76:6380",
	"192.168.14.77:6380"}

func TestNewPool(t *testing.T) {
	ctx, err := NewPool(&RedisOpt{HostAddrs: redisAddr})
	if err != nil {
		log.Fatalf("NewPool error: %s", err.Error())
	}
	defer ctx.Close()
}

func TestExists(t *testing.T) {
	ctx, err := NewPool(&RedisOpt{HostAddrs: redisAddr})
	if err != nil {
		log.Fatalf("NewPool error: %s", err.Error())
	}
	defer ctx.Close()
	f, err := ctx.Exist("foo")
	if err != nil {
		log.Fatalf("ctx.Exist error: %s", err.Error())
	}
	log.Printf(`%s ctx.Exist:%v`, "foo", f)
}

func TestPipeLine(t *testing.T) {

	ctx, err := NewPool(&RedisOpt{HostAddrs: redisAddr})
	if err != nil {
		log.Fatalf("NewPool error: %s", err.Error())
	}
	defer ctx.Close()

	pipe := ctx.BeginPackage()
	pipe.Put("set", "a", "123")
	pipe.Put("set", "b", "234")
	pipe.Put("set", "c", "456")
	rtl, err := ctx.EndPackage(pipe)
	if err != nil {
		log.Fatalf("1 ctx.EndPackage error: %s", err.Error())
	}
	log.Printf(`1 ctx.EndPackage:%v`, rtl)

	pipe = ctx.BeginPackage()

	pipe.Put("get", "a")
	pipe.Put("get", "b")
	pipe.Put("get", "c")
	rtl, err = ctx.EndPackage(pipe)
	if err != nil {
		log.Fatalf("2 ctx.EndPackage error: %s", err.Error())
	}
	log.Printf(`2 ctx.EndPackage:%v`, rtl)

	pipe = ctx.BeginPackage()

	pipe.Put("del", "a")
	pipe.Put("del", "b")
	pipe.Put("del", "c")
	rtl, err = ctx.EndPackage(pipe)
	if err != nil {
		log.Fatalf("3 ctx.EndPackage error: %s", err.Error())
	}
	log.Printf(`3 ctx.EndPackage:%v`, rtl)
}
