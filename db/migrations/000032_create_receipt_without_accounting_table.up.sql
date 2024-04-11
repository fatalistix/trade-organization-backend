CREATE TABLE receipt_without_accounting (
    id SERIAL UNIQUE NOT NULL,
    receipt_id INTEGER NOT NULL,
    receipt_type receipt_type NOT NULL,
    receipting_point_without_accounting_id INTEGER NOT NULL,
    receipting_point_without_accounting_type receipting_point_without_accounting_type NOT NULL,
    PRIMARY KEY (id),
    UNIQUE (receipt_id, receipt_type),
    FOREIGN KEY (receipt_id, receipt_type) REFERENCES receipt(id, type),
    FOREIGN KEY (receipting_point_without_accounting_id, receipting_point_without_accounting_type) REFERENCES receipting_point_without_accounting(id, type)
);