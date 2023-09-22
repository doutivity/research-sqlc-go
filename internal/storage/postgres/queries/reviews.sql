-- name: ReviewNew :one
INSERT INTO reviews (parent_id, company_id, content, created_by, created_at)
VALUES (@parent_id, @company_id, @content, @created_by, @created_at)
RETURNING id;

-- name: ReviewsRootByCompany :many
SELECT id, parent_id, company_id, content, created_by, created_at
FROM reviews
WHERE company_id = @company_id
  AND parent_id IS NULL;

-- name: ReviewsNested :many
WITH RECURSIVE nested AS (
    SELECT r.*
    FROM reviews r
    WHERE r.id = @id::BIGINT
        UNION ALL
    SELECT r.*
    FROM reviews r
        INNER JOIN nested n ON (r.parent_id = n.id)
)
SELECT id, parent_id, company_id, content, created_by, created_at
FROM nested n
WHERE n.company_id = @company_id::BIGINT;
