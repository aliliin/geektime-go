### 1.使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。

进入已经安装好的 redis 服务器中，执行 `redis-benchmark -d 10 -t get,set` 命令，支持的参数

- -t 表示需要执行的命令，如set、get
- -c 表示客户端连接数
- -d 表示单条数据大小，单位 byte
- -n 表示测试包数量
- -r 表示使用随机key数量

#### 结果如下

##### GET

| 数据大小 | 执行次数和所需的时间                      | 每秒请求次数                  |
| :------: | ----------------------------------------- | ----------------------------- |
|  10byte  | 100000 requests completed in 0.72 seconds | 118063.76 requests per second |
|  20byte  | 100000 requests completed in 0.73 seconds | 136986.30 requests per second |
|  50byte  | 100000 requests completed in 0.74 seconds | 135501.36 requests per second |
| 100byte  | 100000 requests completed in 0.73 seconds | 137931.03 requests per second |
| 200byte  | 100000 requests completed in 0.77 seconds | 136986.30 requests per second |
| 1024byte | 100000 requests completed in 0.73 seconds | 136239.78 requests per second |
| 5210byte | 100000 requests completed in 0.74 seconds | 135501.36 requests per second |

##### SET

| 数据大小 |           执行次数和所需的时间            | 每秒请求次数                  |
| :------: | :---------------------------------------: | ----------------------------- |
|  10byte  | 100000 requests completed in 0.72 seconds | 138696.25 requests per second |
|  20byte  | 100000 requests completed in 1.01 seconds | 98814.23 requests per second  |
|  50byte  | 100000 requests completed in 0.74 seconds | 134952.77 requests per second |
| 100byte  | 100000 requests completed in 0.73 seconds | 137174.22 requests per second |
| 200byte  | 100000 requests completed in 0.73 seconds | 137174.22 requests per second |
| 1024byte | 100000 requests completed in 1.19 seconds | 84033.61 requests per second  |
| 5210byte | 100000 requests completed in 1.16 seconds | 85836.91 requests per second  |



### 2.写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息 , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。

代码已经在 此目录的 `main.go` 文件中了，具体结果参看 `reports` 目录下的 csv 文件。

相同长度的 `value` 在写入数量越多情况下，平均每个 `value` 占用内存更多。

![memory](/Users/aliliin/sites/golang/geektime-go/week8/memory.png)

