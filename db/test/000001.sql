START TRANSACTION ISOLATION LEVEL SERIALIZABLE;

INSERT INTO trading_point (id, type, area_plot, rental_charge, counter_count, address) VALUES (0, 'department_store', 10, 10, 10, 'address');

INSERT INTO receipting_point_with_accounting (id, type) VALUES (0, 'department_store');

INSERT INTO department_store (id, type) VALUES (0, 'department_store');

COMMIT TRANSACTION;