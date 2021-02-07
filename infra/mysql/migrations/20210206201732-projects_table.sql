-- +migrate Up
create table nagasu.projects (
    id int auto_increment,
    name varchar(255) not null,
    created_at timestamp default NOW() not null,
    updated_at timestamp default NOW() not null,
    constraint projects_pk
        primary key (id)
);

-- +migrate Down
DROP TABLE nagasu.projects;
