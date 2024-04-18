WITH r AS (
    SELECT receipt_id, receipt_type, SUM(quantity * price) AS total_receipt_sum
    FROM product_receipt
    WHERE receipt_type = 'receipt_with_accounting'
    GROUP BY (receipt_id, receipt_type)
    ORDER BY total_receipt_sum DESC
)
SELECT receipt_with_accounting.customer_id, SUM(r.total_receipt_sum) AS total_customer_sum
FROM r
    INNER JOIN receipt_with_accounting
        ON (r.receipt_id, r.receipt_type) = (id, type)
GROUP BY receipt_with_accounting.customer_id;
