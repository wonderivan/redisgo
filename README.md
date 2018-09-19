# Redis存储访问说明文档
## API使用说明
包名为`store`,引用路径`rotoava/sdk.redis`,依赖`github.com/chasex/redis-go-cluster`。
很多命令可以从这里找到参考 [Redis 命令](http://www.redis.cn/commands.html)。

- store.NewPool()
  - 结构定义: `func NewPool(host, username, password string) *Pool`
  - 参数`host`: 主机信息，可选，例如:`127.0.0.1:6379`
  - 说明: 生成`Pool`对象，用于后续的各项数据操作

- store.Pool{}
  - 说明: Store对象池，通过NewPool创建。对于不使用的Store对象会在一定时间内释放
  - 结构定义:
  ```go
  type Pool struct {
    ClusterOp
    OpString
    OpHash
    OpList
  }
  ```

  - 成员`redis.ClusterOp`:
     - 说明: 用于操作所有的redis通用命令和返回结果的对象，OpString/OpHash/OpList都是基于此对象执行命令
  - 成员`store.OpString`:
     - 说明: 对String类型的数据操作的对象
  - 成员`store.OpHash`:
     - 说明: 对Hash类型的数据操作的对象
  - 成员`store.OpList`:
     - 说明: 对List类型的数据操作的对象		
  - 成员方法`Close()`:
     - 说明: 关闭由NewPool创建的Store对象池
  - 成员方法`Exist()`:
     - 参数`key`: 用于查询的key
     - 说明: 用于查询的key是否存在与redis存储中
  - 成员方法`BeginPackage()`:
     - 说明:启动pipeline功能
     - 返回值`batch`: 操作pipeline的句柄
  - 成员方法`EndPackage()`:
     - 参数`batch`: 操作pipeline的句柄
     - 说明: 执行pipeline，并返回操作结果
     - 返回值`[]interface`: 操作pipeline的返回结果
	
----------
- store.OpString{}
  - 说明: 操作List的接口，存储列表。见 [Redis Lists数据类型](http://www.redis.cn/commands.html#string)
  - 结构定义:
	```go
	type OpString struct {
	  GetRaw(key interface{}) ([]byte, error) 
	  Get(key interface{}) (string, error) 
	  Set(key, data interface{}) error 
	  Increase(key interface{}) (int64, error) 
	  IncreaseBy(key interface{}, increment int64) (int64, error) 
	  GetRange(key interface{}, start, end int64) (string, error) 
	  Append(key interface{}, str string) (int64, error) 
	  Length(key interface{}) (int64, error) 
	  Remove(key interface{}) error 
	}
	```

  - 成员方法`GetRaw()`:
     - 参数`key`: 用于查询的key
     - 返回值`[]byte`: 获取该key的原始数据
     - 返回值`error`: 如果发生错误，此值为错误的信息，否则为nil
  	 - 说明: 获取指定key的值。见 [Redis GET命令](http://www.redis.cn/commands/get.html)
  
  - 成员方法`Get()`:
     - 参数`key`: 用于查询的key
     - 返回值`string`: 获取该key的值
     - 返回值`error`: 如果发生错误，此值为错误的信息，否则为nil
  	 - 说明: 获取指定key的值。见 [Redis GET命令](http://www.redis.cn/commands/get.html)
  - 成员方法`Set()`:
     - 参数`key`: 用于查询的key
     - 参数`data`: 要设置的value值
     - 返回值`string`: 获取该key的值
     - 返回值`error`: 如果发生错误，此值为错误的信息，否则为nil
  	 - 说明: 获取指定位置的元素。见 [Redis SET命令](http://www.redis.cn/commands/set.html)
  - 成员方法`Increase()`:
     - 参数`key`: 用于查询的key
     - 返回值`int64`: 执行递增操作后key对应的值
     - 返回值`error`: 如果发生错误，此值为错误的信息，否则为nil
  	 - 说明: 对存储在指定key的数值执行原子的加1操作 见 [Redis INCR命令](http://www.redis.cn/commands/incr.html)
  - 成员方法`IncreaseBy()`:
     - 参数`key`: 用于查询的key
     - 参数`increment`: 执行值增加操作的大小
     - 返回值`int64`: 执行增加操作后key对应的值
     - 返回值`error`: 如果发生错误，此值为错误的信息，否则为nil
  	 - 说明: 将key对应的数字加increment 见 [Redis INCR命令](http://www.redis.cn/commands/incrby.html)	
  - 成员方法`GetRange()`:
     - 参数`key`: 用于查询的key
     - 参数`start`: 取子串的起始位置
     - 参数`end`: 取子串的结束位置
     - 返回值`string`: 返回key对应的字符串value的子串
     - 返回值`error`: 如果发生错误，此值为错误的信息，否则为nil
  	 - 说明: 返回key对应的字符串value的子串，这个子串是由start和end位移决定的（两者都在string内）<br> 见 [Redis GETRANGE命令](http://www.redis.cn/commands/getrange.html)	
  - 成员方法`Append()`:
     - 参数`key`: 用于查询的key
     - 参数`str`: 要追加到的原有value值后的字符串
     - 返回值`int64`: 返回append后字符串值（value）的长度
     - 返回值`error`: 如果发生错误，此值为错误的信息，否则为nil
  	 - 说明: 获取指定位置的元素。见 [Redis APPEND命令](http://www.redis.cn/commands/append.html)
  - 成员方法`Length()`:
     - 参数`key`: 用于查询的key
     - 返回值`int64` 返回key的string类型value的长度
     - 返回值`error`: 如果发生错误，此值为错误的信息，否则为nil
  	 - 说明:  删除指定的key。见 [Redis STRLEN命令](http://www.redis.cn/commands/strlen.html)	
  - 成员方法`Remove()`:
     - 参数`key`: 用于查询的key
     - 返回值`error`: 如果发生错误，此值为错误的信息，否则为nil
  	 - 说明:  删除指定的key。见 [Redis DEL命令](http://www.redis.cn/commands/del.html)	
	
	
----------

- store.OpHash{}
  - 说明: 操作Hash的接口，存储键值对。见 [Redis Hashs数据类型](http://www.redis.cn/commands.html#hash)
  - 结构定义:
	```go
    type KV [2]interface{}
    type OpHash struct {
      Get(key, name interface{}) ([]byte,error)
      GetRange(key interface{}, names... interface{}) ([]KV,error)
      GetAll(key interface{}) ([]KV,error)
      Set(name, value interface{}) error
      SetRange(key interface{}, datus... KV) error
      Exist(key, name interface{}) (bool,error)
      Keys(key interface{}) ([][]byte,error)
      Remove(key interface{}, names... interface{}) (bool,error)
      RemoveAll(key interface{}) error
    }
	```
  - 成员方法`Get()`:
     - 参数`key`: 用于查询的key
     - 参数`name`: 键名
     - 返回值`[]byte`: 该名称的元素值，name不存在返回nil
     - 返回值`error`: 如果发生错误，此值为错误的信息，否则为nil
  	 - 说明: 获取指定名称的元素值。见 [Redis HGET命令](http://www.redis.cn/commands/hget.html)
  - 成员方法`GetRange()`:
     - 参数`key`: 用于查询的key
     - 参数`names`: 键名集合
     - 返回值`[]KV`: 键-值的集合。"[2]interface{}"数组，[0]为键，[1]值。name不存在的[1]为nil
     - 返回值`error`: 如果发生错误，此值为错误的信息，否则为nil
  	 - 说明: 获取指定名称的元素值。见 [Redis HMGET命令](http://www.redis.cn/commands/hmget.html)
  - 成员方法`GetAll()`:
     - 参数`key`: 用于查询的key
     - 返回值`[]KV`: 键-值的集合。"[2]interface{}"数组，[0]为键，[1]值
     - 返回值`error`: 如果发生错误，此值为错误的信息，否则为nil
  	 - 说明: 获取所有的键-值。见 [Redis HGETALL命令](http://www.redis.cn/commands/hgetall.html)
  - 成员方法`Set()`:
     - 参数`key`: 用于查询的key
     - 参数`name`: 键名
     - 参数`value`: 值
     - 返回值`error`: 如果发生错误，此值为错误的信息，否则为nil
  	 - 说明: 设置键-值。见 [Redis HSET命令](http://www.redis.cn/commands/hset.html)
  - 成员方法`SetRange()`:
     - 参数`key`: 用于查询的key
     - 参数`[]KV`: 键-值的集合
     - 返回值`error`: 如果发生错误，此值为错误的信息，否则为nil
  	 - 说明: 批量设置键-值。见 [Redis HMSET命令](http://www.redis.cn/commands/hmset.html)
  - 成员方法`Exist()`:
     - 参数`key`: 用于查询的key
     - 参数`name`: 键名
     - 返回值`bool`: 指定的'name'存在则为true，否则为false
     - 返回值`error`: 如果发生错误，此值为错误的信息，否则为nil
  	 - 说明: 判断键是否存在。见 [Redis HEXISTS命令](http://www.redis.cn/commands/hexists.html)
  - 成员方法`Keys()`:
     - 参数`key`: 用于查询的key
     - 返回值`[][]byte`: 所有的键集合
     - 返回值`error`: 如果发生错误，此值为错误的信息，否则为nil
  	 - 说明: 获取获取所有的键。见 [Redis HKEYS命令](http://www.redis.cn/commands/hkeys.html)
  - 成员方法`Remove()`:
     - 参数`key`: 用于查询的key
     - 参数`names`: 键名集合
     - 返回值`bool`: 如果移除成功的数量和给定names的数量不相等则返回false,否则返回true
     - 返回值`error`: 如果发生错误，此值为错误的信息，否则为nil
  	 - 说明: 移除元素。见 [Redis HDEL命令](http://www.redis.cn/commands/hdel.html)
  - 成员方法`RemoveAll()`:
     - 参数`key`: 用于查询的key
     - 返回值`error`: 如果发生错误，此值为错误的信息，否则为nil
  	 - 说明:  删除所有元素。见 [Redis DEL命令](http://www.redis.cn/commands/del.html)

----------
	
- store.OpList{}
  - 说明: 操作List的接口，存储列表。见 [Redis Lists数据类型](http://www.redis.cn/commands.html#list)
  - 结构定义:
	```go
    type OpList struct {
      Get(key interface{},pos int64) ([]byte,error)
      GetRange(key interface{},pos int64, count int64) ([][]byte,error)
      PopHead(key interface{}) ([]byte,error)
      PopTail(key interface{}) ([]byte,error)
      Count(key interface{}) (int64,error)
      Add(key interface{}, datus... interface{}) error
      Insert(key interface{}, pos int64, datus... interface{}) error
      Replace(key interface{}, pos int64, datus... interface{}) error
      Remove(key interface{}, poses... int64) error
      RemoveAll(key interface{}) error
    }
	```
  - 成员方法`Get()`:
     - 参数`key`: 用于查询的key
     - 参数`pos`: 从0开始的元素的位置索引，负数表示从最后一个位置算起。例如：0表示第一个元素，1表示第二个元素，-1表示最后一个元素，-2表示倒数第二个元素
     - 返回值`[]byte`: 该位置的数据
     - 返回值`error`: 如果发生错误，此值为错误的信息，否则为nil
  	 - 说明: 获取指定位置的元素。见 [Redis LINDEX命令](http://www.redis.cn/commands/lindex.html)
  - 成员方法`GetRange()`:
     - 参数`key`: 用于查询的key
     - 参数`pos`: 从0开始的元素的位置索引，负数表示从尾位置算起。例如：0表示第一个元素，1表示第二个元素，-1表示最后一个元素，-2表示倒数第二个元素
     - 参数`count`: 元素的数量，-1表示到末尾
     - 返回值`[][]byte`: 获取的元素
     - 返回值`error`: 如果发生错误，此值为错误的信息，否则为nil
  	 - 说明: 据获取指定范围的元素。见 [Redis LRANGE命令](http://www.redis.cn/commands/lrange.html)<br>
  	   假定list数据为：[A0,B1,C2,D3,E4,F5,G6]
         - GetRange(1,2) = [B1,C2]                    -- 正向下标取数据
         - GetRange(4,5) = [E4,F5,G6]                 -- 正向下标取数据, 超出范围的忽略
         - GetRange(1,-1) = [B1,C2,D3,E4,F5,G6]       -- 正向下标取之后的所有数据
         - GetRange(-3,2) = [E4,F5]                   -- 反向下标取数据
         - GetRange(-3,-1) = [E4,F5,G6]               -- 反向下标取之后的所有数
  - 成员方法`PopHead()`:
     - 参数`key`: 用于查询的key
     - 返回值`[]byte`: 元素的数据。无数据则返回nil
     - 返回值`error`: 如果发生错误，此值为错误的信息，否则为nil
  	 - 说明: 移除并且返回第一个元素。见 [Redis LPOP命令](http://www.redis.cn/commands/lpop.html)
  - 成员方法`PopTail()`:
     - 参数`key`: 用于查询的key
     - 返回值`[]byte`: 元素的数据。无数据则返回nil
     - 返回值`error`: 如果发生错误，此值为错误的信息，否则为nil
  	 - 说明: 移除并且返回最后一个元素。见 [Redis RPOP命令](http://www.redis.cn/commands/rpop.html)
  - 成员方法`Count()`:
     - 参数`key`: 用于查询的key
     - 返回值`int64`: 元素数量
     - 返回值`error`: 如果发生错误，此值为错误的信息，否则为nil
  	 - 说明: 获取元素的数量。见 [Redis LLEN命令](http://www.redis.cn/commands/llen.html)
  - 成员方法`Add()`:
     - 参数`key`: 用于查询的key
     - 参数`datus`: 添加的元素集合。为可变参数，可以为任意多了
     - 返回值`error`: 如果发生错误，此值为错误的信息，否则为nil
  	 - 说明: 向集合末尾添加元素。见 [Redis RPUSH命令](http://www.redis.cn/commands/rpush.html)
  - 成员方法`Insert()`:
     - 参数`key`: 用于查询的key
     - 参数`pos`: 从0开始的元素插入的位置索引，负数表示从尾位置算起。例如：0表示第一个元素，1表示第二个元素，-1表示最后一个元素，-2表示倒数第二个元素
     - 参数`datus`: 添加的元素集合。为可变参数，可以为任意多了
     - 返回值`error`: 如果发生错误，此值为错误的信息，否则为nil
  	 - 说明: 向集合指定位置开始插入元素。*首先会根据`pos`位置找到元素，然后再使用`LINSERT`命令向元素的"BEFORE"插入数据，因此如果列表中有重复的数据，可能插入的位置不是`pos`指定的*。见 [Redis LINSERT命令](http://www.redis.cn/commands/linsert.html)
  - 成员方法`Replace()`:
     - 参数`pos`: 从0开始的元素插入的位置索引，负数表示从最后一个位置算起(累加`pos`,比如替换-2,-1,0,1,2)。例如：0表示第一个元素，1表示第二个元素，-1表示最后一个元素，-2表示倒数第二个元素。
     - 参数`datus`: 添加的元素集合。为可变参数，可以为任意多了
     - 返回值`error`: 如果发生错误，此值为错误的信息，否则为nil
  	 - 说明: 从指定位置开始替换内容，替换内容的数组长度不应该超过List的容量，否则只会替换前面存在的位置的元素。见 [Redis LSET命令](http://www.redis.cn/commands/lset.html)
  - 成员方法`Remove()`:
     - 参数`key`: 用于查询的key
     - 参数`poses`: 从0开始的元素插入的位置索引集合，负数表示从尾位置算起。为可变参数，可以为任意多了。例如：0表示第一个元素，1表示第二个元素，-1表示最后一个元素，-2表示倒数第二个元素
     - 返回值`error`: 如果发生错误，此值为错误的信息，否则为nil
  	 - 说明: 删除指定位置的元素。redis不支持直接指定索引删除元素，内部的做法是将指定位置的元素用`LSET`命令设置为"\_\_deleted\_\_"值，再使用`LREM`删除这些元素。见 [Redis LSET命令](http://www.redis.cn/commands/lset.html)， [Redis LREM命令](http://www.redis.cn/commands/lrem.html)
  - 成员方法`RemoveAll()`:
     - 参数`key`: 用于操作的key
     - 返回值`error`: 如果发生错误，此值为错误的信息，否则为nil
  	 - 说明: 删除所有元素。见 [Redis DEL命令](http://www.redis.cn/commands/del.html)

----------

- redis.Cluster
  - 说明: 操作其他不常用的数据接口。见 [Redis 数据类型](http://www.redis.cn/commands.html)

----------
## 设计目的
能够为所有的访问redis存储的go服务提供访问方法，可以在此基础上针对业务做进一步的结构封装


##类图
<div align=center>
redis client api structure
<img src="images/store.png" />
</div>
