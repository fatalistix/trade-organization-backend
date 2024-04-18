SELECT trading_point_id, trading_point_type, price, quantity
FROM product_trading_point
WHERE product_id = ?;

SELECT trading_point_id, trading_point_type, price, quantity
FROM product_trading_point
WHERE product_id = ?
  AND trading_point_type = ?;

SELECT trading_point_id, trading_point_type, price, quantity
FROM product_trading_point
WHERE product_id = ?
  AND trading_point_id = ?
  AND trading_point_type = ?;