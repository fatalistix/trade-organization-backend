CREATE TABLE product_transfer (
    id SERIAL UNIQUE NOT NULL,
    transfer_id INTEGER NOT NULL,
    product_id INTEGER NOT NULL,
    quantity INTEGER NOT NULL CHECK ( quantity > 0 ),
    PRIMARY KEY (id),
    FOREIGN KEY (transfer_id) REFERENCES transfer(id),
    FOREIGN KEY (product_id) REFERENCES product(id)
);