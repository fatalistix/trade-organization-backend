ALTER TABLE department_store
    RENAME COLUMN customer_accounting_point_id TO receipting_point_with_accounting_id;

ALTER TABLE department_store
    RENAME COLUMN customer_accounting_point_type TO receipting_point_with_accounting_type;

ALTER TABLE store
    RENAME COLUMN customer_accounting_point_id TO receipting_point_with_accounting_id;

ALTER TABLE store
    RENAME COLUMN customer_accounting_point_type TO receipting_point_with_accounting_type;
