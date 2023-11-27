SELECT
user_id AS '用户编号',
org_id AS '机构编号'
FROM users_org
WHERE user_id = %s
;
