WITH r AS (
    SELECT id, type, seller_id
    FROM receipt
    WHERE created_at BETWEEN ? AND ?
), s AS (
    SELECT SUM(quantity * price) AS total
    FROM product_receipt
    WHERE (receipt_id, receipt_type) IN (r.id, r.type)
), c AS (
    SELECT COUNT(*) AS total
    FROM r
    GROUP BY seller_id;
)
SELECT s.total / c.total;

WITH a AS (
    SELECT
        id,
        type
--         receipting_point_with_accounting_id AS point_id,
--         receipting_point_with_accounting_type AS point_type
    FROM receipt_with_accounting
    WHERE receipting_point_with_accounting_type = ?
    UNION ALL
    SELECT
        id,
        type
--         receipting_point_without_accounting_id AS point_id,
--         receipting_point_without_accounting_type AS point_type
    FROM receipt_without_accounting
    WHERE receipting_point_without_accounting_type = ?
), r AS (
    SELECT id, type, seller_id
    FROM receipt
    WHERE created_at BETWEEN ? AND ?
      AND (id, type) IN (a)
), s AS (
    SELECT SUM(quantity * price) AS total
    FROM product_receipt
    WHERE (receipt_id, receipt_type) IN (r.id, r.type)
), c AS (
    SELECT COUNT(*) AS total
    FROM r
    GROUP BY seller_id
)
SELECT s.total / c.total;
