CREATE TABLE transfer (
    id SERIAL UNIQUE NOT NULL,
    from_trading_point_id INTEGER NOT NULL,
    from_trading_point_type TEXT NOT NULL CHECK ( to_trading_point_type IN ('department_store', 'store', 'kiosk', 'tray') ),
    to_trading_point_id INTEGER NOT NULL,
    to_trading_point_type TEXT NOT NULL CHECK ( to_trading_point_type IN ('department_store', 'store', 'kiosk', 'tray') ),
    PRIMARY KEY (id),
    FOREIGN KEY (from_trading_point_id, from_trading_point_type) REFERENCES trading_point(id, type),
    FOREIGN KEY (to_trading_point_id, to_trading_point_type) REFERENCES trading_point(id, type)
);