DROP TABLE store;

CREATE TABLE store (
    id SERIAL UNIQUE NOT NULL,
    trading_point_id INTEGER NOT NULL,
    trading_point_type trading_point_type NOT NULL,
    hall_container_id INTEGER NOT NULL,
    hall_container_type hall_container_type NOT NULL,
    customer_accounting_point_id INTEGER NOT NULL,
    customer_accounting_point_type customer_accounting_point_type NOT NULL,
    PRIMARY KEY (id),
    UNIQUE (trading_point_id, trading_point_type),
    UNIQUE (hall_container_id, hall_container_type),
    UNIQUE (customer_accounting_point_id, customer_accounting_point_type),
    FOREIGN KEY (trading_point_id, trading_point_type) REFERENCES trading_point(id, type),
    FOREIGN KEY (hall_container_id, hall_container_type) REFERENCES hall_container(id, type),
    FOREIGN KEY (customer_accounting_point_id, customer_accounting_point_type) REFERENCES customer_accounting_point(id, type)
);

ALTER TABLE section
    DROP CONSTRAINT section_department_store_id_fkey;

DROP TABLE department_store;

CREATE TABLE department_store (
    id SERIAL UNIQUE NOT NULL,
    trading_point_id INTEGER NOT NULL,
    trading_point_type trading_point_type NOT NULL,
    customer_accounting_point_id INTEGER NOT NULL,
    customer_accounting_point_type customer_accounting_point_type NOT NULL,
    PRIMARY KEY (id),
    UNIQUE (trading_point_id, trading_point_type),
    UNIQUE (customer_accounting_point_id, customer_accounting_point_type),
    FOREIGN KEY (trading_point_id, trading_point_type) REFERENCES trading_point(id, type),
    FOREIGN KEY (customer_accounting_point_id, customer_accounting_point_type) REFERENCES customer_accounting_point(id, type)
);

ALTER TABLE section
    ADD FOREIGN KEY (department_store_id) REFERENCES department_store(id);