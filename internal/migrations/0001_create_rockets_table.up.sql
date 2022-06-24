CREATE TABLE IF NOT EXIST rockets(
    id serial NOT NULL PRIMARY KEY,
    type VARCHAR (50),
    name VARCHAR (50)
);