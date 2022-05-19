insert into users(email, password, is_admin, birth_date, username, profile_img, status)
values ('bond.annyka@gmail.icloud', 'qwertyuiop1', false, '2000-12-12', 'arkasha1', 'static/img/1.jpg', 'follow your heart');
insert into users(email, password, is_admin, birth_date, username, profile_img, status)
values ('bond.annyka@gmail.com', 'qwertyuiop2', false, '2000-12-12', 'arkasha2', 'static/img/1.jpg', 'my friends are my estate');
insert into users(email, password, is_admin, birth_date, username, profile_img, status)
values ('annyka@gmail.com', 'qwertyuiop3', false, '2000-12-12', 'arkasha3', 'static/img/1.jpg', 'live without regretsd');

insert into message(from_id, to_id, text, date_time)
values (1, 2, 'hello', '2021-11-01T21:30:48');
insert into message(from_id, to_id, text, date_time)
values (2, 3, 'hello', '2021-11-01T21:30:48');

insert into pin(description, likes_count, dislikes_count, author_id, pin_link)
values ('amazon', 0, 0, 1, 'static/img/6.jpg');
insert into pin(description, likes_count, dislikes_count, author_id, pin_link)
values ('shampoo', 0, 0, 2, 'static/img/3.jpg');
insert into pin(description, likes_count, dislikes_count, author_id, pin_link)
values ('family', 0, 0, 3, 'static/img/8.jpg');

insert into saved_pins(pin_id, user_id)
values (1,1);
insert into saved_pins(pin_id, user_id)
values (2,1);

insert into comment(is_deleted, pin_id, text, author_id, date_time)
values (false, 1, ';)', 2, '2021-09-30T22:13:43');
insert into comment(is_deleted, pin_id, text, author_id, date_time)
values (false, 2, ';(', 2, '2021-09-30T22:13:43');