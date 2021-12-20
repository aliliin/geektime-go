#### 1.总结几种 socket 粘包的解包方式：fix length/delimiter based/length field based frame decoder。尝试举例其应用

* 粘包问题是指当发送两条消息时，比如发送了 ABC 和 DEF，但另一端接收到的却是 ABCD，像这种一次性读取了两条数据的情况就叫做粘包（正常情况应该是一条一条读取的）,正确读取 ABC 和 DEF 两条信息。
* 当发送的消息是 ABC 时，另一端却接收到的是 AB 和 C 两条信息，像这种情况就叫做半包。

##### 为什么会有粘包和半包？

* 这是因为 TCP 是面向连接的传输协议，TCP 传输的数据是以流的形式，而流数据是没有明确的开始结尾边界，所以 TCP 也没办法判断哪一段流属于一个消息。

* 造成粘包的主要原因
  1. 发送方每次写入数据 < 套接字（Socket）缓冲区大小
  2. 接收方读取套接字（Socket）缓冲区数据不够及时。
* 造成半包的主要原因
  1. 发送方每次写入数据 > 套接字（Socket）缓冲区大小
  2. 发送的数据大于协议的 MTU (Maximum Transmission Unit，最大传输单元)，因此必须拆包。

##### fix length

*  [点击查看实现代码](https://github.com/aliliin/geektime-go/tree/main/week9/fix_length)

*执行程序在 client server 目录*

* 每次发送固定缓冲区大小数据.客户端和服务器约定每次发送请求的大小.例如客户端发送1024个字节，服务器接受1024个字节。  
  这样虽然可以解决粘包的问题，但是如果发送的数据小于1024个字节，就会导致数据内存冗余和浪费;且如果发送请求大于1024字节，会出现半包的问题，也就是数据接收的不完整。

##### delimiter based

*  [点击查看实现代码](https://github.com/aliliin/geektime-go/tree/main/week9/delimiter_based)

​		*执行程序在 client  server 目录*

* 基于定界符来判断是不是一个请求（例如结尾'\n'). 客户端发送过来的数据，每次以 \n 结束，服务器每接受到一个 \n 就以此作为一个请求.然后对其拆分后的头部部分与前一个包的剩余部分进行合并，这样就得到了一个完整的包。这种方式的缺点在于如果数据量过大，查找定界符会消耗一些性能

##### length field based frame decoder

*  [点击查看实现代码](https://github.com/aliliin/geektime-go/tree/main/week9/frame_decoder)

​		*执行程序在 client  server 目录*

* 在 TCP 协议头里面写入每次发送请求的长度。 客户端在协议头里面带入数据长度，服务器在接收到请求后，根据协议头里面的数据长度来决定接受多少数据,只有在读取到足够长度的消息之后才算是读到了一个完整的消息。之后会按照参数指定的包长度偏移量数据对接收到的数据进行解码，从而得到目标消息体数据。

#### goim 协议的解码器
*  [点击查看实现代码](https://github.com/aliliin/geektime-go/tree/main/week9/goim_decode)