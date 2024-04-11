CREATE TABLE seller (
    id SERIAL UNIQUE NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    middle_name TEXT NOT NULL,
    birth_date DATE NOT NULL,
    salary NUMERIC(18, 2) NOT NULL CHECK ( salary >= 0 ),
    phone_number TEXT NOT NULL,
    place_of_work_id INTEGER NOT NULL,
    place_of_work_type place_of_work_type NOT NULL,
    PRIMARY KEY (id),
    UNIQUE (phone_number),
    FOREIGN KEY (place_of_work_id, place_of_work_type) REFERENCES place_of_work(id, type)
);