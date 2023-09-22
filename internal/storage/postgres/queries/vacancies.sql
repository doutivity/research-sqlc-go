-- name: VacancyNew :exec
INSERT INTO vacancies (company_id, title, created_by, created_at)
VALUES (@company_id, @title, @created_by, @created_at);

-- name: VacancyByCompanies :many
SELECT v.id    AS vacancy_id,
       v.title AS vacancy_title,
       c.id    AS company_id,
       c.alias AS company_alias,
       c.name  AS company_name
FROM vacancies v
         INNER JOIN companies c ON (v.company_id = c.id)
WHERE c.id = ANY (@companies::BIGINT[]);
