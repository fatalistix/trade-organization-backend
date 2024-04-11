CREATE TYPE customer_accounting_point_type AS ENUM(
    'department_store',
    'store'
);

CREATE TABLE customer_accounting_point (
    id SERIAL UNIQUE NOT NULL,
    type customer_accounting_point_type NOT NULL,
    PRIMARY KEY (id, type)
);