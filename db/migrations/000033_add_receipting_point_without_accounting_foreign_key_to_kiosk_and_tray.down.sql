DROP TABLE tray;

CREATE TABLE tray (
    id SERIAL UNIQUE NOT NULL,
    trading_point_id INTEGER NOT NULL,
    trading_point_type trading_point_type NOT NULL,
    place_of_work_id INTEGER NOT NULL,
    place_of_work_type place_of_work_type NOT NULL,
    PRIMARY KEY (id),
    UNIQUE (trading_point_id, trading_point_type),
    UNIQUE (place_of_work_id, place_of_work_type),
    FOREIGN KEY (trading_point_id, trading_point_type) REFERENCES trading_point(id, type),
    FOREIGN KEY (place_of_work_id, place_of_work_type) REFERENCES place_of_work(id, type)
);

DROP TABLE kiosk;

CREATE TABLE kiosk (
    id SERIAL UNIQUE NOT NULL,
    trading_point_id INTEGER NOT NULL,
    trading_point_type trading_point_type NOT NULL,
    place_of_work_id INTEGER NOT NULL,
    place_of_work_type place_of_work_type NOT NULL,
    PRIMARY KEY (id),
    UNIQUE (trading_point_id, trading_point_type),
    UNIQUE (place_of_work_id, place_of_work_type),
    FOREIGN KEY (trading_point_id, trading_point_type) REFERENCES trading_point(id, type),
    FOREIGN KEY (place_of_work_id, place_of_work_type) REFERENCES place_of_work(id, type)
);
