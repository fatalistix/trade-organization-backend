ALTER TABLE receipting_point_with_accounting
    RENAME TO customer_accounting_point;

ALTER TYPE receipting_point_with_accounting_type
    RENAME TO customer_accounting_point_type;