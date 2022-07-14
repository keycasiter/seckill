create table sec_kill_order
(
    id          bigint auto_increment
        primary key,
    user_id     varchar(32)                         not null,
    sku_code    varchar(32)                         not null,
    sku_num     int                                 not null,
    create_time timestamp default CURRENT_TIMESTAMP not null,
    update_time timestamp default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP not null,
    state       tinyint                             not null,
    order_state tinyint                             not null
);