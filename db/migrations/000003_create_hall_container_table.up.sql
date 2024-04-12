CREATE TABLE hall_container (
    id SERIAL UNIQUE NOT NULL,
    type TEXT NOT NULL CHECK ( type in ('section', 'shop') ),
    PRIMARY KEY (id, type)
);