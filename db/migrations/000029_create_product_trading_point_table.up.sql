CREATE TABLE product_trading_point (
    id SERIAL UNIQUE NOT NULL,
    trading_point_id INTEGER NOT NULL,
    trading_point_type TEXT NOT NULL CHECK ( trading_point_type IN ('department_store', 'store', 'tray', 'kiosk') ),
    product_id INTEGER NOT NULL,
    quantity INTEGER NOT NULL CHECK ( quantity >= 0 ),
    price NUMERIC(18, 2) NOT NULL CHECK ( price >= 0 ),
    PRIMARY KEY (id),
    FOREIGN KEY (trading_point_id, trading_point_type) REFERENCES trading_point(id, type),
    FOREIGN KEY (product_id) REFERENCES product(id)
);