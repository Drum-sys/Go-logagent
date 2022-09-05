create table tbl_app_info
(
    app_id       int primary key auto_increment,
    app_name     varchar(1024) not null,
    app_type     varchar(64) not null,
    create_time  TIMESTAMP default current_timestamp,
    develop_path varchar(256)  not null
);

create table tbl_app_ip
(
    app_id int,
    ip     varchar(64),
    Key    app_id_ip_index (app_id, ip)
);

create table tbl_log_info
(
    log_id      int auto_increment primary key,
    app_id      varchar(1024) not null,
    log_path    varchar(64)   not null,
    create_time TIMESTAMP default current_timestamp,
    status      int       default 1
);