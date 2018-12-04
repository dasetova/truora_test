create table categories (
	id INT PRIMARY KEY DEFAULT nextval('categories_seq'),
	name varchar(50) unique not null
);
