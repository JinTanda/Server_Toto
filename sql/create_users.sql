create table users(
  id int not null auto_increment primary key,
  university varchar(255),
  name varchar(255),
  email varchar(255) unique,
  password char(32),
  department varchar(255),
  major varchar(255),
  created datetime,
  confirmed boolean default false
);
