package tasks

var InsertTaskQuery = `
	INSERT taskSchema::Tasks {
		title := <str>$title,
		description := <str>$description,
		due_date := <cal::local_date>$due_date
	};
`
var GetTaskQuery = `
	SELECT taskSchema::Tasks { 
		title, 
		description, 
		due_date 
	} FILTER .id = <uuid>$id;
`

var GetAllTasksQuery = `
	SELECT taskSchema::Tasks {
		id,
		title,
		description,
		due_date,
		created_at
	};
`

var UpdateTasksQuery = `
	UPDATE taskSchema::Tasks
	FILTER .id = <uuid>$id
	SET { 
		title := <str>$title,
		description := <str>$description,
		due_date := <cal::local_date>$due_date
	};
`

var DeleteTasksQuery = `
	DELETE taskSchema::Tasks
	FILTER .id = <uuid>$id ;
`
