CREATE TABLE transfer (
    id SERIAL UNIQUE NOT NULL,
    from_trading_point_id INTEGER NOT NULL,
    from_trading_point_type trading_point_type NOT NULL,
    to_trading_point_id INTEGER NOT NULL,
    to_trading_point_type trading_point_type NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (from_trading_point_id, from_trading_point_type) REFERENCES trading_point(id, type),
    FOREIGN KEY (to_trading_point_id, to_trading_point_type) REFERENCES trading_point(id, type)
);