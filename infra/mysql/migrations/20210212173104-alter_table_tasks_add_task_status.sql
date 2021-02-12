-- +migrate Up
alter table nagasu.tasks
    add column current_status bigint not null;
alter table nagasu.tasks
    add constraint tasks_task_status_fk
        foreign key (current_status) references nagasu.task_status (id);

-- +migrate Down
alter table nagasu.tasks
    drop constraint tasks_task_status_fk;
alter table nagasu.tasks
    drop column current_status;