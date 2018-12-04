create table recipes (
	id INT PRIMARY KEY DEFAULT nextval('recipes_seq'),
	name varchar(50) unique not null,
    description varchar(200),
    category_id integer references categories(id),
    indications text
);
