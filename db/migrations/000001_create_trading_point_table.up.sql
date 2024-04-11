CREATE TYPE trading_point_type AS ENUM(
    'department_store',
    'store',
    'kiosk',
    'tray'
);

CREATE TABLE trading_point (
    id SERIAL UNIQUE NOT NULL,
    type trading_point_type NOT NULL,
    area_plot FLOAT NOT NULL CHECK ( area_plot >= 0 ),
    rental_charge NUMERIC(18, 2) NOT NULL CHECK ( rental_charge >= 0 ),
    num_of_counter INTEGER NOT NULL CHECK ( num_of_counter >= 0 ),
    PRIMARY KEY (id, type)
);

CREATE TABLE utility_service (
    id SERIAL UNIQUE NOT NULL,
    payment_day DATE NOT NULL DEFAULT CURRENT_DATE,
    amount NUMERIC(18, 2) NOT NULL CHECK ( amount >= 0 ),
    trading_point_id INTEGER NOT NULL,
    trading_point_type trading_point_type NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (trading_point_id, trading_point_type) REFERENCES trading_point(id, type)
);