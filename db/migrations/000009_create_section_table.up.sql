CREATE TABLE section (
    id SERIAL UNIQUE NOT NULL,
    department_store_id INTEGER NOT NULL,
    hall_container_id INTEGER NOT NULL,
    hall_container_type hall_container_type NOT NULL,
    PRIMARY KEY (id),
    UNIQUE (hall_container_id, hall_container_type),
    FOREIGN KEY (department_store_id) REFERENCES department_store(id),
    FOREIGN KEY (hall_container_id, hall_container_type) REFERENCES hall_container(id, type)
)