CREATE TABLE hall_container (
    id SERIAL UNIQUE NOT NULL,
    type TEXT NOT NULL CHECK ( type in ('section', 'store') ),
    PRIMARY KEY (id, type)
);