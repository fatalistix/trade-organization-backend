WITH sum_square AS (
    SELECT SUM(area_plot)
    FROM trading_point
    WHERE type = ?
), receipts_with_trading_point_type AS (
    SELECT id, type
    FROM receipt_with_accounting
    WHERE receipting_point_with_accounting_type = ?
    UNION ALL
    SELECT id, type
    FROM receipt_without_accounting
    WHERE receipting_point_without_accounting_type = ?
), sum_quantity AS (
    SELECT SUM(quantity)
    FROM product_receipt
    WHERE (receipt_id, receipt_type) IN (receipts_with_trading_point_type)
)
SELECT sum_quantity / sum_square;

WITH halls_count AS (
    SELECT COUNT(*)
    FROM hall
    WHERE trading_point_type = ?
), receipts_with_trading_point_type AS (
    SELECT id, type
    FROM receipt_with_accounting
    WHERE receipting_point_with_accounting_type = ?
    UNION ALL
    SELECT id, type
    FROM receipt_without_accounting
    WHERE receipting_point_without_accounting_type = ?
), sum_quantity AS (
    SELECT SUM(quantity)
    FROM product_receipt
    WHERE (receipt_id, receipt_type) IN (receipts_with_trading_point_type)
)
SELECT sum_quantity / halls_count;

WITH counter_count AS (
    SELECT SUM(num_of_counter)
    FROM trading_point
    WHERE type = ?
), receipts_with_trading_point_type AS (
    SELECT id, type
    FROM receipt_with_accounting
    WHERE receipting_point_with_accounting_type = ?
    UNION ALL
    SELECT id, type
    FROM receipt_without_accounting
    WHERE receipting_point_without_accounting_type = ?
), sum_quantity AS (
    SELECT SUM(quantity)
    FROM product_receipt
    WHERE (receipt_id, receipt_type) IN (receipts_with_trading_point_type)
)
SELECT sum_quantity / counter_count;