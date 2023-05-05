create database allanfs;

create schema balance;

create table balance.entry (
    id bigserial primary key,
    name varchar(255) not null,
    amount numeric(10,2) not null,
    entry_type varchar(10) not null,
    external_info varchar(255),
    created_at timestamp not null default now()
);

COMMENT ON COLUMN balance.entry.external_info IS 'user''s correlation info to track this entry';
COMMENT ON COLUMN balance.entry.entry_type IS 'enum: C to Credit; D to Debit. This enum is maintained by application';
COMMENT ON COLUMN balance.entry.amount IS 'entry amount value';
COMMENT ON COLUMN balance.entry."name" IS 'entry name';

CREATE INDEX entry_external_info_idx ON balance.entry (external_info);
