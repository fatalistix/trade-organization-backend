CREATE TABLE store (
    id SERIAL UNIQUE NOT NULL,
    trading_point_id INTEGER NOT NULL,
    trading_point_type trading_point_type NOT NULL,
    hall_container_id INTEGER NOT NULL,
    hall_container_type hall_container_type NOT NULL,
    PRIMARY KEY (id),
    UNIQUE (trading_point_id, trading_point_type),
    UNIQUE (hall_container_id, hall_container_type),
    FOREIGN KEY (trading_point_id, trading_point_type) REFERENCES trading_point(id, type),
    FOREIGN KEY (hall_container_id, hall_container_type) REFERENCES hall_container(id, type)
);