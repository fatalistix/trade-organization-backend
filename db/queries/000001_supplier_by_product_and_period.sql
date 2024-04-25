-- Only list of ids and quantity for each
WITH p AS (
    SELECT order_id, quantity
    FROM product_order
    WHERE product_id = ?
), o AS (
    SELECT id, supplier_id
    FROM "order"
    WHERE status = 'completed'
      AND completed_at BETWEEN ? AND ?
)
SELECT o.supplier_id, SUM(p.quantity) AS total
FROM o
    INNER JOIN p
        ON o.id = p.order_id
GROUP BY o.supplier_id
HAVING SUM(p.quantity) > ?;