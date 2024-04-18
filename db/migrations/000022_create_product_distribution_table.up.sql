CREATE TABLE product_distribution (
    id SERIAL UNIQUE NOT NULL,
    distribution_id INTEGER NOT NULL,
    product_id INTEGER NOT NULL,
    quantity INTEGER NOT NULL CHECK ( quantity >= 0 ),
    price NUMERIC(18, 2) NOT NULL CHECK ( price >= 0 ),
    PRIMARY KEY (id),
    FOREIGN KEY (distribution_id) REFERENCES distribution(id),
    FOREIGN KEY (product_id) REFERENCES product(id)
);