-- +migrate Up
create table nagasu.tasks
(
    id          bigint auto_increment,
    mission_id  int                     not null,
    name        varchar(255)            not null,
    description varchar(1024)           not null,
    created_at  timestamp default NOW() not null,
    updated_at  timestamp default NOW() not null,
    constraint tasks_pk
        primary key (id),
    constraint tasks_projects_fk
        FOREIGN KEY (mission_id) REFERENCES nagasu.missions (id)
);
-- +migrate Down
drop table nagasu.tasks;