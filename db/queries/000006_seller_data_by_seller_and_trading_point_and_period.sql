WITH r AS (
    SELECT id, type
    FROM receipt
    WHERE seller_id = ?
      AND created_at BETWEEN ? AND ?
)
SELECT SUM(quantity * price) AS total
FROM product_receipt
WHERE (receipt_id, receipt_type) IN (r);