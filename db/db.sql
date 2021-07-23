create table users
(
    id        int auto_increment,
    user_name varchar(255) not null,
    email     varchar(255) not null,
    nick_name varchar(255) not null,
    password  varchar(255) not null,
    constraint user_id_index
        unique (id)
);

alter table users
    add primary key (id);

create table documents
(
    id            int auto_increment
        primary key,
    user_id       int          null,
    title         varchar(255) null,
    content       text         not null,
    modified_time datetime     null,
    create_time   datetime     null,
    constraint documents_users_id_fk
        foreign key (user_id) references users (id)
);

