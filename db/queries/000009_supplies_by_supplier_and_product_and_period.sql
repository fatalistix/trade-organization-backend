WITH o AS (
    SELECT id
    FROM "order"
    WHERE supplier_id = ?
      AND status = 'completed'
      AND completed_at BETWEEN ? AND ?
)
SELECT o.id, product_order.quantity, product_order.price
FROM product_order
    INNER JOIN o
        ON o.id = product_order.order_id;
