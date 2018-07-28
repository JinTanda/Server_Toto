create table classes(
  id int not null auto_increment primary key,
  name varchar(255),
  teacher varchar(255),
  room varchar(255),
  class_time varchar(255),
  university varchar(255),
  document varchar(255),
  created datetime
);
