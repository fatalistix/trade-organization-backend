CREATE TABLE department_store (
    id SERIAL UNIQUE NOT NULL,
    type TEXT NOT NULL CHECK ( type = 'department_store' ),
    PRIMARY KEY (id, type),
    FOREIGN KEY (id, type) REFERENCES trading_point(id, type),
    FOREIGN KEY (id, type) REFERENCES receipting_point_with_accounting(id, type)
);