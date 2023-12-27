# oversold
超卖问题解决方案+源码实现

## 环境准备

### SQL

```sql
create table product
(
    id     int unsigned auto_increment primary key,
    sku_id int unsigned default 0 not null comment '商品sku_id',
    num    int default 0 not null comment '商品库存数量'
) comment '产品表';

create table orders
(
    id         int unsigned auto_increment primary key,
    product_id int unsigned default 0 not null comment '商品id',
    user_id    int unsigned default 0 not null comment '用户id'
) comment '订单表';

insert into product(sku_id, num) value(132, 10)
```

### MySQL
```shell
docker run -itd --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root mysql
```

### ab并发测试

1. 准备 data.json 文件
```json
{
  "user_id": 666,
  "sku_id": 132,
  "num": 1
}
```
2. 启动 ab 并发测试
```shell
./ab.exe -n 1000 -c 100 -T application/json -p data.json http://localhost:9090/oversold
```
参数解释：
> -n 请求总数  
> -c 并发数  
> -T 发送请求的Content-Type  
> -p 发送请求的Body
