# ws 使用ws 在送礼业务中来 提高并发
```
1 使用mysql 和 redis 连接池 减少client 连接时间
2 携程执行socket 的收发消息，将收发消息放入channel，
3 启动另一个携程来 循环处理channel消息，提高系统吞吐量
```
