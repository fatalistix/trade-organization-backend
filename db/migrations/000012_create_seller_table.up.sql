CREATE TYPE seller_status AS ENUM(
    'working',
    'not_working'
);

CREATE TABLE seller (
    id SERIAL UNIQUE NOT NULL,
    status seller_status NOT NULL DEFAULT 'working',
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    middle_name TEXT NOT NULL,
    birth_date DATE NOT NULL,
    salary NUMERIC(18, 2) NOT NULL CHECK ( salary >= 0 ),
    phone_number TEXT NOT NULL,
    place_of_work_id INTEGER DEFAULT NULL,
    place_of_work_type TEXT CHECK ( place_of_work_type IN ('hall', 'tray', 'kiosk', NULL) ) DEFAULT NULL,
    trading_point_id INTEGER DEFAULT NULL,
    trading_point_type TEXT CHECK ( trading_point_type IN ('department_store', 'store', 'tray', 'kiosk', NULL) ) DEFAULT NULL,
    CHECK (
        (place_of_work_id IS NULL AND place_of_work_type IS NULL AND trading_point_id IS NULL AND trading_point_type IS NULL) OR
        (place_of_work_id IS NOT NULL AND place_of_work_type IS NOT NULL AND trading_point_id IS NOT NULL AND trading_point_type IS NOT NULL)),
    PRIMARY KEY (id),
    UNIQUE (phone_number),
    FOREIGN KEY (place_of_work_id, place_of_work_type) REFERENCES place_of_work(id, type),
    FOREIGN KEY (trading_point_id, trading_point_type) REFERENCES trading_point(id, type)
);