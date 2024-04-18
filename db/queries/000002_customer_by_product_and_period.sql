SELECT id, first_name, last_name, phone_number
FROM customer
WHERE id IN
      (SELECT receipt_with_accounting.customer_id
       FROM receipt_with_accounting
           LEFT JOIN customer
               ON receipt_with_accounting.customer_id = customer.id
       WHERE (receipt_with_accounting.id, receipt_with_accounting.type) IN
             (SELECT receipt_id, receipt_type
              FROM product_receipt
              WHERE (receipt_id, receipt_type) IN
                    (SELECT id, type
                     FROM receipt
                     WHERE type = 'receipt_with_accounting'
                       AND created_at BETWEEN ? AND ?)
                AND product_id = ?
              GROUP BY (receipt_id, receipt_type))
       GROUP BY (receipt_with_accounting.customer_id));