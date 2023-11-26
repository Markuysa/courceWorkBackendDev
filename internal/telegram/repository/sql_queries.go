package repository

const (
	queryLinkTelegram = `
	update "user".user	
	set tg_chat = $1
	where id = $2
`
	queryGetUsersTask = `
	SELECT u.tg_chat AS chat_id, t.description, t.deadline
	FROM "user"."user" u
	JOIN tasks.task t ON u.id = t.participant_id
	WHERE u.tg_chat IS NOT NULL
`
)
