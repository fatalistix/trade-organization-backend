CREATE TABLE receipt (
    id SERIAL UNIQUE NOT NULL,
    type TEXT NOT NULL CHECK ( type IN ('receipt_with_accounting', 'receipt_without_accounting') ),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    seller_id INTEGER NOT NULL,
    PRIMARY KEY (id, type),
    FOREIGN KEY (seller_id) REFERENCES seller(id)
);