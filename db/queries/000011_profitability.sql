WITH receipts_in_interval AS (
    SELECT id, type
    FROM receipt
    WHERE created_at BETWEEN ? AND ?
), receipts_with_trading_point_type AS (
    SELECT id, type
    FROM receipt_with_accounting
    WHERE (receipting_point_with_accounting_id, receipting_point_with_accounting_type) = (?, ?)
      AND (id, type) IN (receipts_in_interval)
    UNION ALL
    SELECT id, type
    FROM receipt_without_accounting
    WHERE (receipting_point_without_accounting_id, receipting_point_without_accounting_type) = (?, ?)
      AND (id, type) IN (receipts_in_interval)
), sell_volume AS (
    SELECT SUM(quantity)
    FROM product_receipt
    WHERE (receipt_id, receipt_type) IN (receipts_with_trading_point_type)
), sum_rental AS (
    SELECT SUM(amount)
    FROM rental_charge_history
    WHERE (trading_point_id, trading_point_type) = (?, ?)
      AND payment_day BETWEEN ? AND ?
), sum_utility_service AS (
    SELECT SUM(amount)
    FROM utility_service
    WHERE (trading_point_id, trading_point_type) = (?, ?)
      AND payment_day BETWEEN ? AND ?
), sellers_at_trading_point AS (
    SELECT id
    FROM seller
    WHERE (trading_point_id, trading_point_type) = (?, ?)
), sum_salary AS (
    SELECT SUM(amount)
    FROM salary_history
    WHERE payment_time BETWEEN ? AND ?
      AND seller_id IN (sellers_at_trading_point)
)
SELECT sell_volume / (sum_rental + sum_utility_service + sum_salary);