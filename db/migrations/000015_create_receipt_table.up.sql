CREATE TABLE receipt (
    id SERIAL UNIQUE NOT NULL,
    customer_id INTEGER NOT NULL,
    customer_accounting_point_id INTEGER NOT NULL,
    customer_accounting_point_type customer_accounting_point_type NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (customer_id) REFERENCES customer(id),
    FOREIGN KEY (customer_accounting_point_id, customer_accounting_point_type) REFERENCES customer_accounting_point(id, type)
);