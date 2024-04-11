CREATE TYPE receipting_point_without_accounting_type AS ENUM(
    'kiosk',
    'tray'
);

CREATE TABLE receipting_point_without_accounting (
    id SERIAL UNIQUE NOT NULL,
    type receipting_point_without_accounting_type NOT NULL,
    PRIMARY KEY (id, type)
);