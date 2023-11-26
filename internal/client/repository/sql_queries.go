package repository

const (
	queryGetTasksList = `
	select id,
		   category,
		   deadline,
		   status,
		   comments,
		   priority,
		   creator_id,
		   participant_id,
		   description
	from tasks.task
	where participant_id = $1;
`
	queryUpdateTask = `
	update tasks.task
	set deadline = $1
	where id = $2
`
	queryLinkTelegram = `
	update "user".user	
	set tg_chat = $1
	where id = $2
`
)
