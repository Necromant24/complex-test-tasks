WITH ranked_requests AS (
    SELECT 
        ur."userId",
        ur."productId",
        ur.created,
        ROW_NUMBER() OVER (PARTITION BY ur."userId" ORDER BY ur.created DESC) AS rank
    FROM "users-requests" ur
)
SELECT 
    u.id AS user_id,
    u.name AS user_name,
    p.id AS product_id,
    p.name AS product_name,
    rr.created AS order_date
FROM ranked_requests rr
INNER JOIN users u ON rr."userId" = u.id
INNER JOIN products p ON rr."productId" = p.id
WHERE rr.rank <= 10
ORDER BY u.id, rr.created DESC;
