create table recipes (
	id serial primary key,
	name varchar(50) unique not null,
    description varchar(200),
    category_id integer references categories(id),
    indications text
);
