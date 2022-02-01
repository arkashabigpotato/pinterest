select * from users;
select * from message;
select * from pin;
select * from saved_pins;
select * from comment;

select is_admin from users where is_admin=false;
select * from users where email like 'bond%';
select count(*) from pin;