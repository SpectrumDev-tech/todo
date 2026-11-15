-- +goose UP
create table tasks(
    id integer primary key autoincrement,
    title text not null ,
    description text,
    due_time text not null,
    priority text not null default 'normal',
    completed boolean not null default 0,
    created_at text not null
);

create index idx_tasks_completed on tasks(completed);
create index idx_tasks_due_time on tasks(due_time);

-- +goose DOWN
drop table tasks;
drop index idx_tasks_completed;
drop index idx_tasks_due_time;