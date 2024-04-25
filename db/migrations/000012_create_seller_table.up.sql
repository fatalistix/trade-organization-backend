CREATE TYPE seller_status AS ENUM(
    'working',
    'not_working'
);

CREATE TABLE seller (
    id SERIAL UNIQUE NOT NULL,
    status seller_status NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    middle_name TEXT NOT NULL,
    birth_date DATE NOT NULL,
    salary NUMERIC(18, 2) NOT NULL CHECK ( salary >= 0 ),
    phone_number TEXT NOT NULL,
    place_of_work_id INTEGER NOT NULL,
    place_of_work_type TEXT NOT NULL CHECK ( place_of_work_type IN ('hall', 'tray', 'kiosk') ),
    trading_point_id INTEGER NOT NULL,
    trading_point_type TEXT NOT NULL CHECK ( trading_point_type IN ('department_store', 'store', 'tray', 'kiosk') ),
    PRIMARY KEY (id),
    UNIQUE (phone_number),
    FOREIGN KEY (place_of_work_id, place_of_work_type) REFERENCES place_of_work(id, type),
    FOREIGN KEY (trading_point_id, trading_point_type) REFERENCES trading_point(id, type)
);