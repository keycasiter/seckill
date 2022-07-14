# 工程结构
- conf
  
  配置文件

- model
    - po 
      
      数据库持久化对象
      
    - vo
    
      业务对象封装
  
- dal

    数据持久层，使用GormV2与MySQL交互

- logic
    
    业务逻辑层，封装复杂业务逻辑

- method
    
    方法实现层，

- mq 

# 组件
- ORM https://github.com/go-gorm/gorm
- Redis https://github.com/go-redis/redis
- RocketMQ https://github.com/apache/rocketmq-client-go
- UUID https://github.com/satori/go.uuid
- Log https://github.com/ian-kent/go-log/log