ALTER TYPE customer_accounting_point_type
    RENAME TO receipting_point_with_accounting_type;

ALTER TABLE customer_accounting_point
    RENAME TO receipting_point_with_accounting;