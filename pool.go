package store

import (
	"fmt"
	"sync"
	"time"

	redis "github.com/chasex/redis-go-cluster"
)

var doOnceInit = new(sync.Once)

type Pool struct {
	ClusterOp *redis.Cluster
	OpString  *OpString
	OpHash    *OpHash
	OpList    *OpList //IM消息
}

// 获取配置的连接参数
type RedisOpt struct {
	HostAddrs       []string
	Connections     int //连接数 默认10
	ReadTimeOut     int //读超时时间(默认10秒)
	WriteTimeOut    int //写超时时间(默认10秒)
	ConntectTimeOut int //连接超时时间(默认3秒)
	AliveTime       int //每个连接的空闲时间
	EnableMgo       bool
	EnableKafka     bool
}

const (
	default_ConnTimeout  = 10
	default_ReadTimeout  = 10
	default_WriteTimeout = 10
	default_KeepAlive    = 1
	default_AliveTime    = 60
)

func NewPool(opt *RedisOpt) (*Pool, error) {

	if 0 > len(opt.HostAddrs) {
		return nil, fmt.Errorf("NewPool need set storage address in function param or etc/config.json")
	}
	if opt.ConntectTimeOut == 0 {
		opt.ConntectTimeOut = default_ConnTimeout
	}
	if opt.ReadTimeOut == 0 {
		opt.ReadTimeOut = default_ReadTimeout
	}
	if opt.WriteTimeOut == 0 {
		opt.WriteTimeOut = default_WriteTimeout
	}
	if opt.Connections == 0 {
		opt.Connections = default_KeepAlive
	}
	if opt.AliveTime == 0 {
		opt.AliveTime = default_AliveTime
	}
	cluster, err := redis.NewCluster(
		&redis.Options{
			StartNodes:   opt.HostAddrs,
			ConnTimeout:  time.Duration(opt.ConntectTimeOut) * time.Second,
			ReadTimeout:  time.Duration(opt.ReadTimeOut) * time.Second,
			WriteTimeout: time.Duration(opt.WriteTimeOut) * time.Second,
			KeepAlive:    opt.Connections,
			AliveTime:    time.Duration(opt.AliveTime) * time.Second,
		})
	if err != nil {
		return nil, fmt.Errorf("redis NewCluster error: %s", err.Error())
	}
	return &Pool{ClusterOp: cluster, OpString: &OpString{cluster},
		OpHash: &OpHash{cluster}, OpList: &OpList{cluster}}, nil
}

func (this *Pool) Close() {
	this.ClusterOp.Close()
}

func (this *Pool) BeginPackage() *redis.Batch {
	return this.ClusterOp.NewBatch()
}

func (this *Pool) EndPackage(batch *redis.Batch) ([]interface{}, error) {
	return this.ClusterOp.RunBatch(batch)
}

// 判断key是否存在
func (this *Pool) Exist(key string) (bool, error) {
	data, err := this.ClusterOp.Do("EXISTS", key)
	if nil != err {
		return false, err
	}
	return data.(int64) == 1, nil
}
