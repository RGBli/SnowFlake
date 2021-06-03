# SnowFlake
---
### 1.介绍
SnowFlake 算法用于生成64位整数类型的 ID，能够保证按时间戳递增，能够作为分布式系统的主键。

<br/>

### 2.用法
``` go
import (
    "fmt"
    "github.com/RGBli/snowflake"
)

func main() {
	// 3 表示数据中心 ID，2表示机器 ID，5表示用5位来表示机器 ID，最后一个参数表示起始时间
	// 其中机器 ID 的位数不能超过10，不然会引发 panic
	sf := NewSnowFlake(3, 2, 5, time.Date(2021, 6, 3, 0, 0, 0, 0, time.Local).Unix())
	fmt.Println(sf.NextId())
}
```

<br/>

### 3.参考文章
https://www.jianshu.com/p/b1124283fc43
