SELECT 
    U.id,
    U.username,
    U.email,
    CASE 
        WHEN SU.system_id IS NOT NULL THEN TRUE
        ELSE FALSE
    END AS registered
FROM users U
LEFT JOIN systems_users SU ON U.id = SU.user_id AND SU.system_id = 1;


SELECT 
		P.id,
		P.name,
		CASE WHEN SU.system_id IS NOT NULL THEN TRUE ELSE FALSE END AS registered
FROM permissions P
LEFT JOIN systems_users_permissions SU ON P.id = SU.permission_id AND SU.system_id = 1 AND SU.user_id = 1
INNER JOIN roles R ON R.id = P.role_id 
WHERE R.id = 1;