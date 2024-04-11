DROP TABLE product_receipt;

CREATE TABLE product_receipt (
    id SERIAL UNIQUE NOT NULL,
    receipt_id INTEGER NOT NULL,
    product_id INTEGER NOT NULL,
    quantity INTEGER NOT NULL CHECK ( quantity > 0 ),
    price NUMERIC(18, 2) NOT NULL CHECK ( price > 0 ),
    PRIMARY KEY (id),
    FOREIGN KEY (receipt_id) REFERENCES receipt(id),
    FOREIGN KEY (product_id) REFERENCES product(id)
);