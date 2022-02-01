truncate table users restart identity cascade;
truncate table message restart identity cascade;
truncate table pin restart identity cascade;
truncate table saved_pins restart identity cascade;
truncate table comment restart identity cascade;

drop table if exists users cascade;
drop table if exists message cascade;
drop table if exists pin cascade;
drop table if exists saved_pins cascade;
drop table if exists comment cascade;
