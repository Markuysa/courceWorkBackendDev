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
	set deadline = $1        ,
	    participant_id = $2  ,
	    description = $3     ,
	    status = $4          ,
	    category = $5        ,
	    priority = $6        
	where id = $7
`
	queryGetStatusList = `
	select id,
			description
	from lists.status_list
`
	queryGetPriorityList = `
	select id,
			description
	from lists.priority_list
	
`
	queryGetCategoriesList = `
	select id,
			description
	from lists.category_list
`
	queryAddComment = `
	update tasks.task
	set comments = jsonb_concat(comments, $1::jsonb)
	where id = $2
`
)
