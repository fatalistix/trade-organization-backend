CREATE TABLE distribution (
    id SERIAL UNIQUE NOT NULL,
    trading_point_id INTEGER NOT NULL,
    trading_point_type TEXT NOT NULL CHECK ( trading_point_type IN ('department_store', 'store', 'kiosk', 'tray') ),
    order_id INTEGER NOT NULL,
    product_id INTEGER NOT NULL,
    quantity INTEGER NOT NULL CHECK ( quantity >= 0 ),
    PRIMARY KEY (id),
    FOREIGN KEY (trading_point_id, trading_point_type) REFERENCES trading_point(id, type),
    FOREIGN KEY (order_id) REFERENCES "order"(id),
    FOREIGN KEY (product_id) REFERENCES product(id)
)