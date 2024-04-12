CREATE TABLE receipt (
    id SERIAL UNIQUE NOT NULL,
    type TEXT NOT NULL CHECK ( type IN ('receipt_with_accounting', 'receipt_without_accounting') ),
    PRIMARY KEY (id, type)
);