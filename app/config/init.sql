create table users (
	id serial not null unique,
	first_name varchar(64),
	last_name varchar(64),
	primary key(id)
);

insert into users(first_name, last_name)
values
	("Developer", "Developer"),
	("Admin", "Admin");