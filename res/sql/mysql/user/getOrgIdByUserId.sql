SELECT
user_id AS 'UID',
org_id AS '机构编号'
FROM users_org
WHERE user_id = ?
;
