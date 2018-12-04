create table ingredients (
    id INT PRIMARY KEY DEFAULT nextval('ingredients_seq'),
    description varchar(50) not null,
    quantity decimal,
    -- No important for now
    measure_unit varchar(10),
    recipe_id integer references recipes(id)
);
