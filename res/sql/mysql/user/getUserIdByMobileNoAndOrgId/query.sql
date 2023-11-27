SELECT
O.user_id AS '用户编号',
O.org_id AS '机构编号',
U.mobile_no AS '手机号'
FROM users AS U
LEFT JOIN users_org AS O
ON U.id = O.user_id
WHERE U.mobile_no = %s
AND O.org_id = %s
;
