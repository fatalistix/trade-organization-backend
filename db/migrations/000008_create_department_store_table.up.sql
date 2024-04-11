CREATE TABLE department_store (
    id SERIAL UNIQUE NOT NULL,
    trading_point_id INTEGER NOT NULL,
    trading_point_type trading_point_type NOT NULL,
    PRIMARY KEY (id),
    UNIQUE (trading_point_id, trading_point_type),
    FOREIGN KEY (trading_point_id, trading_point_type) REFERENCES trading_point(id, type)
);