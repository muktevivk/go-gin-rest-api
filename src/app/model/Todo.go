package model
// Todo is structure to hold the tasks
type Todo struct {
	ID string `json:"task_id"`
	TaskName string `json:"task_name"`
	Category string `json:"category"`
	UserName string `json:"user_name"`
	IsCompleted bool `json:"is_completed"`
}

//TableName is function to specifiy the tablename
func (b *Todo) TableName() string {
	return "todo"
}