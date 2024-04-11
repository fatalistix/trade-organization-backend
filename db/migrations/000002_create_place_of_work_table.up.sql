CREATE TYPE place_of_work_type AS ENUM(
    'hall',
    'kiosk',
    'tray'
);

CREATE TABLE place_of_work (
    id SERIAL UNIQUE NOT NULL,
    type place_of_work_type NOT NULL,
    PRIMARY KEY (id, type)
);