WITH r AS (
    SELECT id, type
    FROM receipt
    WHERE created_at BETWEEN ? AND ?
), t as (
    SELECT receipt_id, receipt_type
    FROM product_receipt
    WHERE receipt_type = 'receipt_with_accounting'
      AND product_id = ?
      AND (receipt_id, receipt_type) IN (r)
)
SELECT customer_id
FROM receipt_with_accounting
WHERE (id, type) IN (t)
GROUP BY customer_id;