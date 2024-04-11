ALTER TABLE store
    RENAME COLUMN receipting_point_with_accounting_type TO customer_accounting_point_type;

ALTER TABLE store
    RENAME COLUMN receipting_point_with_accounting_id TO customer_accounting_point_id ;

ALTER TABLE department_store
    RENAME COLUMN receipting_point_with_accounting_type TO customer_accounting_point_type;

ALTER TABLE department_store
    RENAME COLUMN receipting_point_with_accounting_id TO customer_accounting_point_id;
