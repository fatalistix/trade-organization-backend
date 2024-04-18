WITH r AS (
    SELECT id, type
    FROM receipt
    WHERE created_at BETWEEN ? AND ?
)
SELECT receipt_id, receipt_type, quantity
FROM product_receipt
WHERE product_id = ?
  AND (receipt_id, receipt_type) IN (r);

WITH r AS (
    SELECT id, type
    FROM receipt
    WHERE created_at BETWEEN ? AND ?
), p AS (
    SELECT product_receipt.receipt_id, product_receipt.receipt_type, product_receipt.quantity
    FROM product_receipt
    WHERE product_id = ?
      AND (receipt_id, receipt_type) IN (r)
)
SELECT receipt_with_accounting.id, receipt_with_accounting.type, p.quantity
FROM receipt_with_accounting
    INNER JOIN p
        ON (p.receipt_id, p.receipt_type) = (receipt_with_accounting.id, receipt_with_accounting.type)
WHERE receipt_with_accounting.receipting_point_with_accounting_type = ?
UNION ALL
SELECT receipt_with_accounting.id, receipt_with_accounting.type, p.quantity
FROM receipt_with_accounting
    INNER JOIN p
        ON (p.receipt_id, p.receipt_type) = (receipt_with_accounting.id, receipt_with_accounting.type)
WHERE receipt_with_accounting.receipting_point_with_accounting_type = ?;

WITH r AS (
    SELECT id, type
    FROM receipt
    WHERE created_at BETWEEN ? AND ?
), p AS (
    SELECT product_receipt.receipt_id, product_receipt.receipt_type, product_receipt.quantity
    FROM product_receipt
    WHERE product_id = ?
      AND (receipt_id, receipt_type) IN (r)
)
SELECT receipt_with_accounting.id, receipt_with_accounting.type, p.quantity
FROM receipt_with_accounting
    INNER JOIN p
        ON (p.receipt_id, p.receipt_type) = (receipt_with_accounting.id, receipt_with_accounting.type)
WHERE (receipt_with_accounting.receipting_point_with_accounting_id, receipt_with_accounting.receipting_point_with_accounting_type) = (?, ?)
UNION ALL
SELECT receipt_without_accounting.id, receipt_without_accounting.type, p.quantity
FROM receipt_without_accounting
    INNER JOIN p
        ON (p.receipt_id, p.receipt_type) = (receipt_without_accounting.id, receipt_without_accounting.type)
WHERE (receipt_without_accounting.receipting_point_without_accounting_id, receipt_without_accounting.receipting_point_without_accounting_type) = (?, ?)