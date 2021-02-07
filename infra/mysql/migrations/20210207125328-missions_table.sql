-- +migrate Up
create table nagasu.missions
(
    id          int auto_increment,
    projects_id int                     not null,
    name        varchar(255)            not null,
    created_at  timestamp default NOW() not null,
    updated_at  timestamp default NOW() not null,
    constraint missions_pk
        primary key (id),
    constraint missions_projects_fk
        FOREIGN KEY (projects_id) REFERENCES nagasu.projects(id)
);
-- +migrate Down
drop table nagasu.missions;