SELECT
id AS '用户编号',
mobile_no AS '手机号'
FROM users
WHERE id = %s
;
