create table sec_kill_goods
(
    id          bigint auto_increment
        primary key comment '主键ID',
    sku_code    varchar(32)                         not null comment '商品sku',
    stock       int                                 not null comment '商品库存',
    create_time timestamp default CURRENT_TIMESTAMP not null comment '创建时间',
    update_time timestamp default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP not null comment '更新时间',
    state       tinyint                             not null comment '数据状态'
) engine = innodb comment '秒杀商品表';