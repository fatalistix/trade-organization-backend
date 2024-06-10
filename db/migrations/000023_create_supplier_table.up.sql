CREATE TYPE supplier_status AS ENUM(
    'available',
    'not_available'
);

CREATE TABLE supplier (
    id SERIAL UNIQUE NOT NULL,
    name TEXT UNIQUE NOT NULL,
    type supplier_status NOT NULL DEFAULT 'available',
    PRIMARY KEY (id)
);