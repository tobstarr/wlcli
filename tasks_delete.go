package main

type deleteTasks struct {
	IDs []int `cli:"arg required"`
}

func (r *deleteTasks) Run() error {
	cl, err := loadClient()
	if err != nil {
		return err
	}
	for _, i := range r.IDs {
		task, err := cl.Task(i)
		if err != nil {
			return err
		}

		if err := cl.DeleteTask(task.ID, task.Revision); err != nil {
			return err
		}
	}
	return nil
}
