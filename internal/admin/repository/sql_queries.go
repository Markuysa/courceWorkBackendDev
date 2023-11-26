package repository

const (
	queryAddTask = `
		insert into tasks.task(
		                       category,
		                       deadline,
		                       status,
		                       priority,
		                       creator_id.
		                       description,
		                       participant_id
		)values (
		         $1,$2,$3,$4,$5,$6,$7
		)
	`
	queryAssignTask = `
	update tasks.task
	set participant_id = $1
	where id = $2
`
	queryGetUsersTasks = `
		select
			id,
			category,
			deadline,
			status,
			priority,
			creator_id.
			description,
			participant_id
		from tasks.task
		where 
			participant_id is null or participant_id = $1
		limit $2 offset $3

`
)
