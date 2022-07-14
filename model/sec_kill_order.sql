create table sec_kill_order
(
    id          bigint auto_increment
        primary key comment '主键ID',
    order_no    varchar(32)                                                     not null comment '订单号',
    user_id     varchar(32)                                                     not null comment '用户ID',
    sku_code    varchar(32)                                                     not null comment '商品sku',
    sku_num     int                                                             not null comment '购买数量',
    create_time timestamp default CURRENT_TIMESTAMP                             not null comment '创建时间',
    update_time timestamp default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP not null comment '更新时间',
    state       tinyint                                                         not null comment '数据状态',
    order_state tinyint                                                         not null comment '订单状态'
) engine = innodb comment '秒杀订单表';