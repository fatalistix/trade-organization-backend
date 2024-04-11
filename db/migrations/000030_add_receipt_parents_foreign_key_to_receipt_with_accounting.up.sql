ALTER TABLE product_receipt
    DROP CONSTRAINT product_receipt_receipt_id_fkey;

DROP TABLE receipt_with_accounting;

CREATE TABLE receipt_with_accounting (
    id SERIAL UNIQUE NOT NULL,
    receipt_id INTEGER NOT NULL,
    receipt_type receipt_type NOT NULL,
    customer_id INTEGER NOT NULL,
    receipting_point_with_accounting_id INTEGER NOT NULL,
    receipting_point_with_accounting_type receipting_point_with_accounting_type NOT NULL,
    PRIMARY KEY (id),
    UNIQUE (receipt_id, receipt_type),
    FOREIGN KEY (receipt_id, receipt_type) REFERENCES receipt(id, type),
    FOREIGN KEY (customer_id) REFERENCES customer(id),
    FOREIGN KEY (receipting_point_with_accounting_id, receipting_point_with_accounting_type) REFERENCES receipting_point_with_accounting(id, type)
);

ALTER TABLE product_receipt
    ADD FOREIGN KEY (receipt_id) REFERENCES receipt_with_accounting(id);