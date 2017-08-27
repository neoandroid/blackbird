package tasks

import (
	//"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
	"github.com/neoandroid/blackbird/models"
)

func init() {
	cleanSchedule := toolbox.NewTask("cleanSchedule", "5 */5 * * * *", func() error {
		// This task will run every 5 minutes
		beego.Info("Running task: clean Schedule")

		schedules, err := models.GetAllSchedule(map[string]string{"name": "LiveStream"}, nil, nil, nil, 0, 1000)
		if err != nil {
			return err
		}
		for _, schedule := range schedules {
			beego.Debug("Clean schedule id: ", schedule.(models.Schedule).Id)
			// TODO: Verify this isn't a currently running stream/schedule
			err = models.DeleteSchedule(schedule.(models.Schedule).Id)
			if err != nil {
				return err
			}
		}

		return nil
	})

	toolbox.AddTask("cleanSchedule", cleanSchedule)
	toolbox.StartTask()
	defer toolbox.StopTask()
}
