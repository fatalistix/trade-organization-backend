SELECT
    product.id,
    product.name,
    product.description,
    product_trading_point.quantity,
    product_trading_point.price
FROM product_trading_point
    INNER JOIN product
        ON product.id = product_trading_point.product_id
WHERE (product_trading_point.trading_point_id, product_trading_point.trading_point_type) = (?, ?);