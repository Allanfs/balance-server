create database allanfs;

create schema balance;

create table balance.entry (
    id bigserial primary key,
    name varchar(255) not null,
    amount numeric(10,2) not null,
    type varchar(10) not null,
    external_info varchar(255),
    created_at timestamp not null default now()
);
