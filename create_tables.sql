drop table if exists CookBookEntry;
drop table if exists Ingredient;
drop table if exists Recipe;


create table if not exists Ingredient (
	id int NOT NULL AUTO_INCREMENT,
	ingredient_name varchar(30) NOT NULL unique,
    ingredient_category varchar(30) NOT NULL,
    primary key ( id )
);
	
create table if not exists Recipe (
	id int NOT NULL AUTO_INCREMENT,
    recipe_name varchar(50) NOT NULL unique,
    recipe_description varchar(200) NOT NULL,
    primary key ( id )
);

create table CookBookEntry (
	recipe_id int NOT NULL,
    ingredient_id int NOT NULL,
    ingredient_amount varchar(30),
    foreign key ( recipe_id ) REFERENCES Recipe(id),
    foreign key ( ingredient_id ) REFERENCES Ingredient(id),
    primary key ( recipe_id, ingredient_id )
);

/* Starting Ingredients */
insert into ingredient (ingredient_name, ingredient_category) values ('Cheese', 'Dairy');
insert into ingredient (ingredient_name, ingredient_category) values ('Bread', 'Grain');
insert into ingredient (ingredient_name, ingredient_category) values ('Butter', 'Dairy');
insert into ingredient (ingredient_name, ingredient_category) values ('Ribeye', 'Meat');
insert into ingredient (ingredient_name, ingredient_category) values ('Thyme', 'Herb');
insert into ingredient (ingredient_name, ingredient_category) values ('Can of Whole Tomatoes', 'Fruit');
insert into ingredient (ingredient_name, ingredient_category) values ('Ginger', 'Spice');
insert into ingredient (ingredient_name, ingredient_category) values ('Extra Virgin Olive Oil', 'Oil');
insert into ingredient (ingredient_name, ingredient_category) values ('Medium Onion', 'Vegetable');
insert into ingredient (ingredient_name, ingredient_category) values ('Garlic', 'Vegetable');
insert into ingredient (ingredient_name, ingredient_category) values ('Smoked Paprika', 'Spice');
insert into ingredient (ingredient_name, ingredient_category) values ('Spinach', 'Vegetable');
insert into ingredient (ingredient_name, ingredient_category)  values ('Chickpeas', 'Beans');
insert into ingredient (ingredient_name, ingredient_category) values ('Bay Leaves', 'Spice');
insert into ingredient (ingredient_name, ingredient_category) values ('Soy Sauce', 'Sauce');
insert into ingredient (ingredient_name, ingredient_category) values ('Sherry Vinegar', 'Vinegar');

/* Starting Recipes */
insert into recipe (recipe_name, recipe_description) 
	values ('Grilled Cheese', 'Grilled cheese so delicious you\'ll want to eat it');
insert into recipe (recipe_name, recipe_description) 
	values ('Pan Steak', 'Delicious pan seared steak');
insert into recipe (recipe_name, recipe_description) 
	values ('Quick Chickpea and Spinach Stew', 'A delicious vegetarian stew that will last for days');

/* Entries for Grilled Cheese */
insert into CookBookEntry (recipe_id, ingredient_id) values (1, 1);
insert into CookBookEntry (recipe_id, ingredient_id) values (1, 2);
insert into CookBookEntry (recipe_id, ingredient_id) values (1, 3);

/* Entries for Pan Steak */
insert into CookBookEntry (recipe_id, ingredient_id) values (2, 4);
insert into CookBookEntry (recipe_id, ingredient_id) values (2, 5);
insert into CookBookEntry (recipe_id, ingredient_id) values (2, 3);

/* Entries for Quick Chickpea and Spinach Stew */
insert into CookBookEntry (recipe_id, ingredient_id) values (3, 6);
insert into CookBookEntry (recipe_id, ingredient_id) values (3, 7);
insert into CookBookEntry (recipe_id, ingredient_id) values (3, 8);
insert into CookBookEntry (recipe_id, ingredient_id) values (3, 9);
insert into CookBookEntry (recipe_id, ingredient_id) values (3, 10);
insert into CookBookEntry (recipe_id, ingredient_id) values (3, 11);
insert into CookBookEntry (recipe_id, ingredient_id) values (3, 12);
insert into CookBookEntry (recipe_id, ingredient_id) values (3, 13);
insert into CookBookEntry (recipe_id, ingredient_id) values (3, 14);
insert into CookBookEntry (recipe_id, ingredient_id) values (3, 15);
insert into CookBookEntry (recipe_id, ingredient_id) values (3, 16);
    
-- select recipe_id, ingredient_id, recipe_name, ingredient_name
-- 	from cookbookentry 
--     inner join recipe on recipe_id = recipe.id 
--     inner join ingredient on ingredient_id = ingredient.id;
