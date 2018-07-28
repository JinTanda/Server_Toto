# usersへのinsert例
insert into users (id, university, name, email, password, department, major, created, confirmed)
  values(0, '名工大', 'fukui@gmail.com', 1234, '工学部', '情報工学', '2018', NOW(), false);

# classesへのisert例
insert into classes (name, teacher, room, time, university, document, created)
values('知能処理', '世木', '2002', '月1', '名工大', 'powerpoint', now());
