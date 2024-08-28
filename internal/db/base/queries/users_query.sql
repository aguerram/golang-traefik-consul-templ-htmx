-- name: GetAllUsers :many
SELECT * FROM users
WHERE users.created_at > @start_date AND users.created_at < @end_date
LIMIT @max_elements OFFSET @start_offset;
