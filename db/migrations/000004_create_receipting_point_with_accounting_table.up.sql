CREATE TABLE receipting_point_with_accounting (
    id SERIAL UNIQUE NOT NULL,
    type TEXT NOT NULL CHECK ( type IN ('department_store', 'store') ),
    PRIMARY KEY (id, type)
);