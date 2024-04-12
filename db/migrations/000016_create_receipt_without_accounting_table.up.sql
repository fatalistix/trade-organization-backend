CREATE TABLE receipt_without_accounting (
    id SERIAL UNIQUE NOT NULL,
    type TEXT NOT NULL CHECK ( type = 'receipt_without_accounting' ),
    receipting_point_without_accounting_id INTEGER NOT NULL,
    receipting_point_without_accounting_type TEXT NOT NULL CHECK ( receipting_point_without_accounting_type IN ('kiosk', 'tray') ),
    PRIMARY KEY (id, type),
    FOREIGN KEY (id, type) REFERENCES receipt(id, type),
    FOREIGN KEY (receipting_point_without_accounting_id, receipting_point_without_accounting_type) REFERENCES receipting_point_without_accounting(id, type)
);