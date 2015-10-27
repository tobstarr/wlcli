package main

type completeTasks struct {
	IDs []int `cli:"arg required"`
}

func (r *completeTasks) Run() error {
	cl, err := loadClient()
	if err != nil {
		return err
	}
	for _, id := range r.IDs {
		if err := cl.CompleteTask(id); err != nil {
			return err
		}
	}
	return nil
}
