CREATE TABLE application (
    id SERIAL UNIQUE NOT NULL,
    trading_point_id INTEGER NOT NULL,
    trading_point_type TEXT NOT NULL CHECK ( trading_point_type IN ('department_store', 'store', 'kiosk', 'tray') ),
    PRIMARY KEY (id),
    FOREIGN KEY (trading_point_id, trading_point_type) REFERENCES trading_point(id, type)
);