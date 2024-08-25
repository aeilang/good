-- name: GetJobById :one
select href, company_image, company_name, title, city, fulltime, job_type, description, requirement, price_down, price_up, version
from jobs
where id = $1
and is_deleted = false;

-- name: CreateJob :exec
insert into jobs (
    id, href, company_name, company_image, title, keyword, city, 
    fulltime, job_type, description, requirement, price_down, price_up
) values (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
);

-- name: UpdateJobById :exec
update jobs
set href = $1, company_image = $2, company_name = $3, title = $4, keyword = $5, city = $6, 
fulltime = $7, job_type = $8, description = $9, requirement = $10,
price_down = $11, price_up = $12, version = version+1
where id = $13 and version = $14;

-- name: DeleteJobById :exec
update jobs set
is_deleted = true
where id = $1;

-- name: GetAllJobs :many
select id, company_name,company_image, title, city, fulltime, job_type, keyword, price_down, price_up,
COUNT(*) OVER() as total
from jobs
where is_deleted = false
AND (price_down >= @price_down::integer or @price_down = 0)
AND (price_up <= @price_up::integer or @price_up = 0)
AND (to_tsvector('simple', title) @@ plainto_tsquery('simple', @title::text) or @title = '')
AND (to_tsvector('simple', company_name) @@ plainto_tsquery('simple', @companyName::text) or @companyName = '')
AND (city = @city::text or @city = '')
AND (job_type = @jobType::text or @jobType = '')
ORDER BY @oderBy::text, id ASC
LIMIT $1 OFFSET $2;
 
