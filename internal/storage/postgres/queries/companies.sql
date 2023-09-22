-- name: CompanyNewAndGetID :one
INSERT INTO companies (alias, name, created_by, created_at)
VALUES (@alias, @name, @created_by, @created_at)
RETURNING id;
