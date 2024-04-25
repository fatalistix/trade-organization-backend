-- Only list of ids and quantity for each
WITH r AS (
    SELECT receipt_id, receipt_type, quantity
    FROM product_receipt
    WHERE product_id = ?
      AND receipt_type = 'receipt_with_accounting'
), c AS (
    SELECT id, type
    FROM receipt
    WHERE created_at
        BETWEEN ? AND ?
      AND type = 'receipt_with_accounting'
), t AS (
    SELECT c.id, c.type, r.quantity
    FROM c
        INNER JOIN r
            ON (c.id, c.type) = (r.receipt_id, r.receipt_type)
)
SELECT receipt_with_accounting.customer_id, SUM(t.quantity) as total
FROM receipt_with_accounting
    INNER JOIN t
        ON (receipt_with_accounting.id, receipt_with_accounting.type) = (t.id, t.type)
GROUP BY receipt_with_accounting.customer_id
HAVING SUM(t.quantity) > ?;