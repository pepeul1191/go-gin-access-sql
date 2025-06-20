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