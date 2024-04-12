CREATE TABLE receipting_point_without_accounting (
    id SERIAL UNIQUE NOT NULL,
    type TEXT NOT NULL CHECK ( type IN ('tray', 'kiosk') ),
    PRIMARY KEY (id, type)
);