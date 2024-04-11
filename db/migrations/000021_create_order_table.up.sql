CREATE TABLE "order" (
    id SERIAL UNIQUE NOT NULL,
    supplier_id INTEGER NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (supplier_id) REFERENCES supplier(id)
);