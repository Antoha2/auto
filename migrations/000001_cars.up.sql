CREATE TABLE IF NOT EXISTS cars (
    id serial PRIMARY KEY,
    regnum varchar(255) NOT NULL,
    mark varchar(255) NOT NULL,
    model varchar(255) NOT NULL,
    owners varchar(255) NOT NULL
)