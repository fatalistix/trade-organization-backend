CREATE TABLE section (
    id SERIAL UNIQUE NOT NULL,
    type TEXT NOT NULL CHECK ( type = 'section' ),
    department_store_id INTEGER NOT NULL,
    department_store_type TEXT NOT NULL CHECK ( department_store_type = 'department_store' ),
    PRIMARY KEY (id, type),
    FOREIGN KEY (id, type) REFERENCES hall_container(id, type),
    FOREIGN KEY (department_store_id, department_store_type) REFERENCES department_store(id, type)
)