CREATE TABLE product (
    id SERIAL UNIQUE NOT NULL,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    PRIMARY KEY (id),
    UNIQUE (name),
    UNIQUE (description)
);