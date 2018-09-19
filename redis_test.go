/**
 * 对garyburd/redigo的测试
 * 对chasex/redis-go-cluster的测试
 * wangruiming create 2017.10.26
 */
package store_test

import (
	"fmt"
	"testing"
	"time"

	redis1 "github.com/chasex/redis-go-cluster"
	"github.com/garyburd/redigo/redis"
)

func TestGoRedis(t *testing.T) {
	conn, err := redis.DialTimeout("tcp", "10.15.63.246:6379", 0, 1*time.Second, 1*time.Second)
	if nil != err {
		t.Fatal("DialTimeout", err)
	}
	defer conn.Close()
	conn.Do("AUTH", "password")

	_, err = conn.Do("DBSIZE")
	if nil != err {
		t.Error("Do DBSIZE", err)
	}

	_, err = conn.Do("SET", "user:user0", 123)
	if nil != err {
		t.Error("Do SET", err)
	}
	_, err = conn.Do("SET", "user:user1", 456)
	if nil != err {
		t.Error("Do SET2", err)
	}
	_, err = conn.Do("APPEND", "user:user0", 87)
	if nil != err {
		t.Error("Do APPEND", err)
	}
	user0, err := redis.Int(conn.Do("GET", "user:user0"))
	if nil != err {
		t.Error("Int", err)
	}
	if user0 != 12387 {
		t.Error("User0 %d not equal 12387", user0)
	}
	user1, err := redis.Int(conn.Do("GET", "user:user1"))
	if nil != err {
		t.Error("Int", err)
	}
	if user1 != 456 {
		t.Error("User1 %d not equal 456", user1)
	}
}

func TestGoRedisCluster(t *testing.T) {
	cluster, err := redis1.NewCluster(
		&redis1.Options{ //
			StartNodes:   []string{"192.168.14.75:6379", "192.168.14.76:6379", "192.168.14.77:6379"},
			ConnTimeout:  50 * time.Millisecond,
			ReadTimeout:  50 * time.Millisecond,
			WriteTimeout: 50 * time.Millisecond,
			KeepAlive:    16,
			AliveTime:    60 * time.Second,
		})

	if err != nil {
		t.Error(err)
	}
	defer cluster.Close()

	_, err = cluster.Do("set", "{user000}.name", "Joel")
	_, err = cluster.Do("set", "{user000}.age", "26")
	_, err = cluster.Do("set", "{user000}.country", "China")

	name, err := redis1.String(cluster.Do("get", "{user000}.name"))
	if err != nil {
		t.Error(err)
	}
	age, err := redis1.Int(cluster.Do("get", "{user000}.age"))
	if err != nil {
		t.Error(err)
	}
	country, err := redis1.String(cluster.Do("get", "{user000}.country"))
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("name: %s, age: %d, country: %s\n", name, age, country)

	_, err = cluster.Do("del", "{user000}.name")
	_, err = cluster.Do("del", "{user000}.age")
	_, err = cluster.Do("del", "{user000}.country")
	return
}
