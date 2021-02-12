-- +migrate Up
CREATE TABLE nagasu.task_status
(
    id           bigint                  not null auto_increment,
    mission_id   int                     not null,
    name         varchar(255)            not null,
    status_order int                     not null,
    created_at   timestamp default NOW() not null,
    updated_at   timestamp default NOW() not null,
    constraint task_status_pk
        primary key (id),
    constraint task_status_mission_fk
        foreign key (mission_id) references nagasu.missions (id)
);
-- +migrate Down
DROP TABLE nagasu.task_status;
