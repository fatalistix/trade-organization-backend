CREATE TYPE hall_container_type AS ENUM(
    'section',
    'shop'
);

CREATE TABLE hall_container (
    id SERIAL UNIQUE NOT NULL,
    type hall_container_type NOT NULL,
    PRIMARY KEY (id, type)
);