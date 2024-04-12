CREATE TABLE place_of_work (
    id SERIAL UNIQUE NOT NULL,
    type TEXT NOT NULL CHECK ( type IN ('hall', 'kiosk', 'tray') ),
    PRIMARY KEY (id, type)
);