CREATE TABLE salary_history (
    id SERIAL UNIQUE NOT NULL,
    seller_id INTEGER NOT NULL,
    amount NUMERIC(18, 2) NOT NULL CHECK ( amount >= 0 ),
    payment_time TIMESTAMP NOT NULL DEFAULT NOW(),
    PRIMARY KEY (id),
    FOREIGN KEY (seller_id) REFERENCES seller(id)
)