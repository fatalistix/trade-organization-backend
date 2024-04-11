ALTER TABLE receipt_with_accounting
    RENAME COLUMN receipting_point_with_accounting_type TO customer_accounting_point_type ;

ALTER TABLE receipt_with_accounting
    RENAME COLUMN receipting_point_with_accounting_id TO customer_accounting_point_id;

ALTER TABLE receipt_with_accounting
    RENAME TO receipt;
