CREATE TABLE kiosk (
    id SERIAL UNIQUE NOT NULL,
    type TEXT NOT NULL CHECK ( type = 'kiosk' ),
    PRIMARY KEY (id, type),
    FOREIGN KEY (id, type) REFERENCES trading_point(id, type),
    FOREIGN KEY (id, type) REFERENCES place_of_work(id, type),
    FOREIGN KEY (id, type) REFERENCES receipting_point_without_accounting(id, type)
);