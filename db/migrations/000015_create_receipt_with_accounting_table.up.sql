CREATE TABLE receipt_with_accounting (
    id SERIAL UNIQUE NOT NULL,
    type TEXT NOT NULL CHECK ( type = 'receipt_with_accounting' ),
    customer_id INTEGER NOT NULL,
    receipting_point_with_accounting_id INTEGER NOT NULL,
    receipting_point_with_accounting_type TEXT NOT NULL CHECK ( receipting_point_with_accounting_type IN ('department_store', 'store') ),
    PRIMARY KEY (id, type),
    FOREIGN KEY (id, type) REFERENCES receipt(id, type),
    FOREIGN KEY (customer_id) REFERENCES customer(id),
    FOREIGN KEY (receipting_point_with_accounting_id, receipting_point_with_accounting_type) REFERENCES receipting_point_with_accounting(id, type)
);