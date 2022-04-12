insert into users(email, password, is_admin, birth_date, username, profile_img, status)
values ('bond.annyka@gmail.icloud', 'qwertyuiop1', false, '2000-12-12', 'arkasha1', 'static/1.jpg', 'qwertyuiop');
insert into users(email, password, is_admin, birth_date, username, profile_img, status)
values ('bond.annyka@gmail.com', 'qwertyuiop2', false, '2000-12-12', 'arkasha2', 'static/1.jpg', 'qwertyuiop');
insert into users(email, password, is_admin, birth_date, username, profile_img, status)
values ('annyka@gmail.com', 'qwertyuiop3', false, '2000-12-12', 'arkasha3', 'static/1.jpg', 'qwertyuiop');

insert into message(from_id, to_id, text, date_time)
values (1, 2, 'hello', '2021-11-01T21:30:48');
insert into message(from_id, to_id, text, date_time)
values (2, 3, 'hello', '2021-11-01T21:30:48');

insert into pin(description, likes_count, dislikes_count, author_id, pin_link)
values ('poiuytrewq', 0, 0, 1, 'asdfghjkl');
insert into pin(description, likes_count, dislikes_count, author_id, pin_link)
values ('poiuytrewq', 0, 0, 2, 'asdfghjkl');
insert into pin(description, likes_count, dislikes_count, author_id, pin_link)
values ('poiuytrewq', 0, 0, 3, 'asdfghjkl');

insert into saved_pins(pin_id, user_id)
values (1,1);
insert into saved_pins(pin_id, user_id)
values (2,1);

insert into comment(is_deleted, pin_id, text, author_id, date_time)
values (false, 1, 'zxcvbnm', 2, '2021-09-30T22:13:43');
insert into comment(is_deleted, pin_id, text, author_id, date_time)
values (false, 2, 'zxcvbnm', 2, '2021-09-30T22:13:43');