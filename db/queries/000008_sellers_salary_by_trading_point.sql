SELECT id, salary
FROM seller;

SELECT id, salary
FROM seller
WHERE trading_point_type = ?;

SELECT id, salary
FROM seller
WHERE (trading_point_id, trading_point_type) = (?, ?);