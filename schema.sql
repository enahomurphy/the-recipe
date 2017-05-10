CREATE TABLE users (
  id int(11) AUTO_INCREMENT NOT NULL,
  firts_name varchar(255),
  last_name varchar(255),
  email varchar(255),
  username varchar(255),
  profile_pic text,
  PRIMARY KEY (id)
)

CREATE TABLE categories (
  id int(11) AUTO_INCREMENT NOT NULL,
  title varchar(255),
  description text(300)
  PRIMARY KEY (id)
)

CREATE TABLE ingredients (
  id int(11) NOT NULL AUTO_INCREMENT
  name varchar(255),
  quantity varchar(50),
  unit varchar(20)
  PRIMARY KEY id
)

CREATE TABLE recipes (
  id int(11) NOT NUll AUTO_INCREMENT,
  name varchar(255),
  description text(300),
  PRIMARY KEY (id)
 
)

ALTER TABLE recipes
			ADD FOREIGN KEY (categoryID) REFERENCES categories(id);

ALTER TABLE recipes
			ADD FOREIGN KEY (ingredientID) REFERENCES ingredients(id);