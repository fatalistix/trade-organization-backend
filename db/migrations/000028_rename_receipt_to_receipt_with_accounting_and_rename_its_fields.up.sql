ALTER TABLE receipt
    RENAME TO receipt_with_accounting;

ALTER TABLE receipt_with_accounting
    RENAME COLUMN customer_accounting_point_id TO receipting_point_with_accounting_id;

ALTER TABLE receipt_with_accounting
    RENAME COLUMN customer_accounting_point_type TO receipting_point_with_accounting_type;