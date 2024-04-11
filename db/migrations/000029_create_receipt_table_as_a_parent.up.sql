CREATE TYPE receipt_type AS ENUM(
    'receipt_with_accounting',
    'receipt_without_accounting'
);

CREATE TABLE receipt (
    id SERIAL UNIQUE NOT NULL,
    type receipt_type NOT NULL,
    PRIMARY KEY (id, type)
);