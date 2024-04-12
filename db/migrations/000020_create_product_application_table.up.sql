CREATE TABLE product_application (
    id SERIAL UNIQUE NOT NULL,
    application_id INTEGER NOT NULL,
    product_id INTEGER NOT NULL,
    quantity INTEGER NOT NULL CHECK ( quantity > 0 ),
    PRIMARY KEY (id),
    FOREIGN KEY (application_id) REFERENCES application(id),
    FOREIGN KEY (product_id) REFERENCES product(id)
);