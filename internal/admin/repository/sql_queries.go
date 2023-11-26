package repository

const (
	queryAddTask = `
		insert into tasks.task(
		                       category,
		                       deadline,
		                       status,
		                       priority,
		                       creator_id,
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
		ts.id,
		ct.description,
		extract(epoch from ts.deadline)::bigint as deadline,
		st.description,
		pr.description,
		creator_id,
		ts.description,
		participant_id
	from tasks.task ts
	left join lists.status_list st on st.id = ts.status
	left join lists.category_list ct on ct.id = ts.category
	left join lists.priority_list pr on pr.id = ts.priority
	where
		participant_id is null or participant_id = $1
	limit $2 offset $3
`
)
