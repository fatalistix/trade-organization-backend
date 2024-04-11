CREATE TABLE hall (
    id SERIAL UNIQUE NOT NULL,
    hall_container_id INTEGER NOT NULL,
    hall_container_type hall_container_type NOT NULL,
    place_of_work_id INTEGER NOT NULL,
    place_of_work_type place_of_work_type NOT NULL,
    PRIMARY KEY (id),
    UNIQUE (place_of_work_id, place_of_work_type),
    FOREIGN KEY (hall_container_id, hall_container_type) REFERENCES hall_container(id, type),
    FOREIGN KEY (place_of_work_id, place_of_work_type) REFERENCES place_of_work(id, type)
);