create table app_user
(
    id         varchar(36) primary key,
    username   varchar(36) not null,
    password   varchar(36) not null,
    first_name varchar(36) not null,
    last_name  varchar(36) not null,
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp
)

create
unique index app_user_username_uindex
	on app_user (username);
