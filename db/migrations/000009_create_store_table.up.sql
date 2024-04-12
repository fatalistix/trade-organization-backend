CREATE TABLE store (
    id SERIAL UNIQUE NOT NULL,
    type TEXT NOT NULL CHECK ( type = 'store' ),
    PRIMARY KEY (id, type),
    FOREIGN KEY (id, type) REFERENCES trading_point(id, type),
    FOREIGN KEY (id, type) REFERENCES hall_container(id, type),
    FOREIGN KEY (id, type) REFERENCES receipting_point_with_accounting(id, type)
);