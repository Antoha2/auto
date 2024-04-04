CREATE TABLE IF NOT EXISTS autos (
    id serial PRIMARY KEY,
    regnum varchar(255) NOT NULL,
    mark varchar(255) NOT NULL,
    model varchar(255)NOT NULL,
    owner varchar(255) 
)