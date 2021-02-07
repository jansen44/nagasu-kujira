-- +migrate Up
create table projects (
    id int auto_increment,
    name varchar(255) not null,
    constraint projects_pk
        primary key (id)
);

-- +migrate Down
DROP TABLE projects;
