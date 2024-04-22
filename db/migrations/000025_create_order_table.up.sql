CREATE TYPE order_status AS ENUM(
    'completed',
    'canceled',
    'in_progress'
);

CREATE TABLE "order" (
    id SERIAL UNIQUE NOT NULL,
    supplier_id INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    status order_status NOT NULL DEFAULT 'in_progress',
    completed_at TIMESTAMP DEFAULT NULL,
    canceled_at TIMESTAMP DEFAULT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (supplier_id) REFERENCES supplier(id)
);