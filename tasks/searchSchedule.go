package tasks

import (
	"sort"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
	"github.com/neoandroid/blackbird/controllers"
	"github.com/neoandroid/blackbird/models"
)

func init() {
	searchSchedule := toolbox.NewTask("searchSchedule", "0/15 * * * * *", func() error {
		// This task will run every 15 seconds
		beego.Info("Running task: Search Schedule")

		// Get all schedules ordered by date
		scheds, err := models.GetAllSchedule(nil, nil, nil, nil, 0, 1000)
		if err != nil {
			beego.Critical(err)
			return err
		}
		var schedList models.ScheduleList
		for _, schedule := range scheds {
			schedList = append(schedList, schedule.(models.Schedule))
		}
		sort.Sort(schedList)

		// Get date in long format
		now := time.Now()
		date, err := strconv.Atoi(now.Format("200601021504"))
		if err != nil {
			beego.Critical(err)
			return err
		}
		longDate := int64(date)
		beego.Info("Current date:", date)

		// Get first schedule to play
		var currentSchedule *models.Schedule
		for _, schedule := range schedList {
			if schedule.GetComposeDate() < longDate {
				continue
			} else {
				currentSchedule = &schedule
				break
			}
		}

		if currentSchedule != nil {
			beego.Info("Next scheduled job:", currentSchedule.GetComposeDate())
		} else {
			beego.Info("There is no scheduled job")
			return nil
		}


		// Play schedule
		if currentSchedule.GetComposeDate() == longDate {
			// Check if we're already playing this schedule.
			status, err := models.GetAllStatus(map[string]string{"schedule_id": strconv.Itoa(currentSchedule.Id)}, nil, nil, nil, 0, 1000)
			if err != nil {
				beego.Critical(err)
				return err
			}
			if len(status) > 0 {
				return nil
			}

			// Get Music
			music, err := models.GetMusicById(currentSchedule.MusicId.Id)
			if err != nil {
				beego.Critical(err)
				return err
			}

			// Get Group
			group, err := models.GetGroupById(currentSchedule.GroupId.Id)
			if err != nil {
				beego.Critical(err)
				return err
			}

			var stream controllers.StreamController
			playerStatus, err := stream.PlayStream(music, []models.Group{*group}, currentSchedule.Id)
			beego.Info("Player status", playerStatus)
		}

		return nil
	})

	toolbox.AddTask("searchSchedule", searchSchedule)
	toolbox.StartTask()
	defer toolbox.StopTask()
}

