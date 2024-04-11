CREATE TABLE customer (
    id SERIAL UNIQUE NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    phone_number TEXT NOT NULL,
    birth_date DATE NOT NULL,
    PRIMARY KEY (id),
    UNIQUE (phone_number)
);