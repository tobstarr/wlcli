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
		if err := cl.DeleteTask(i); err != nil {
			return err
		}
	}
	return nil
}
