-- auto-generated definition
create table users
(
    id         bigint unsigned not null
        primary key,
    username   int unsigned    not null comment '用户名称',
    password   varchar(255)    not null comment '密码',
    created_at timestamp       not null,
    updated_at timestamp       not null,
    deleted_at timestamp       null,
    constraint username
        unique (username)
)
    collate = utf8mb4_unicode_ci;

create index deleted_at
    on users (deleted_at);

