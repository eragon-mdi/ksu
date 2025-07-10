package sqlrepo

const TABLE_NAME = `task`

const QUERY_SAVE_TASK_RETURN = `
INSERT INTO ` + TABLE_NAME + `
	(id, result, status, created_at, duration, started_at)
VALUES
	($1, $2, $3, $4, $5, $6)
RETURNING
	id, result, status, created_at, duration, started_at
`

const QUERY_DELETE_TASK = `
DELETE 
FROM ` + TABLE_NAME + `
WHERE id = $1
`

const QUERY_GET_TASK_RESULT = `
SELECT 
	status, created_at, duration
FROM ` + TABLE_NAME + `
WHERE
	id = $1
`

const QUERY_GET_TASK_STATUS = `
SELECT 
	result
FROM ` + TABLE_NAME + `
WHERE
	id = $1
`

const QUERY_UPDATE_TASK = `
UPDATE ` + TABLE_NAME + `
SET
	status = $2, duration = $3, started_at = $4, result = $5
WHERE
	id = $1
`

const QUERY_GET_ALL_TASKS = `
SELECT 
	id, result, status, created_at, duration, started_at
FROM ` + TABLE_NAME + `
`
