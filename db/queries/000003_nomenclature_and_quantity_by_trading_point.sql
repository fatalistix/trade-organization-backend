SELECT product.id, product.name, product.description, product_trading_point.quantity, product_trading_point.price
FROM product_trading_point
    LEFT JOIN product
        ON product_trading_point.trading_point_id = ?;