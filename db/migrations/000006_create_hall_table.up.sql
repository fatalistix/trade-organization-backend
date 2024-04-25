CREATE TABLE hall (
    id SERIAL UNIQUE NOT NULL,
    type TEXT NOT NULL CHECK ( type = 'hall' ),
    hall_container_id INTEGER NOT NULL,
    hall_container_type TEXT NOT NULL CHECK ( hall_container_type IN ('section', 'store') ),
    trading_point_id INTEGER NOT NULL,
    trading_point_type TEXT NOT NULL CHECK ( trading_point_type IN ('department_store', 'store') ),
    PRIMARY KEY (id, type),
    FOREIGN KEY (id, type) REFERENCES place_of_work(id, type),
    FOREIGN KEY (hall_container_id, hall_container_type) REFERENCES hall_container(id, type)
);