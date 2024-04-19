CREATE TABLE "order" (
    id SERIAL UNIQUE NOT NULL,
    supplier_id INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    PRIMARY KEY (id),
    FOREIGN KEY (supplier_id) REFERENCES supplier(id)
);