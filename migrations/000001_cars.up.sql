CREATE TABLE IF NOT EXISTS cars (
    id serial PRIMARY KEY,
    regnum varchar(255) NOT NULL,
    mark varchar(255) NOT NULL,
    model varchar(255) NOT NULL,
    year int NOT NULL,
    name varchar(255) NOT NULL,
    surname varchar(255) NOT NULL,
    patronymic varchar(255)
)