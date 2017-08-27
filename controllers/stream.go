package controllers

import (
	"encoding/json"
	"sort"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"github.com/neoandroid/blackbird/models"
)

// StreamController operations for Stream
type StreamController struct {
	beego.Controller
}

// URLMapping ...
func (c *StreamController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// Post ...
// @Title Post
// @Description Play stream
// @Param	body		body 	[]models.Group	true		"body for Group content"
// @Success 201 {int} models.AudioPlayerResponse
// @Failure 403 body is empty
// @router /:id [post]
func (c *StreamController) Post() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	groups := make([]models.Group, 0)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &groups)
	if err != nil {
		beego.Critical(err)
		c.Data["json"] = err.Error()
	} else {
		music, err := models.GetMusicById(id)
		if err != nil {
			beego.Critical(err)
			c.Data["json"] = err.Error()
		} else {
			if status, err := c.PlayStream(music, groups, -1); err == nil {
				c.Ctx.Output.SetStatus(201)
				c.Data["json"] = status
			} else {
				c.Data["json"] = err.Error()
			}
		}
	}
	c.ServeJSON()
}

func (c *StreamController) PlayStream(music *models.Music, groups []models.Group, scheduleId int) (*models.AudioPlayerResponse, error) {
	models.CleanStatus()
	if !c.isAvailable() {
		return &models.AudioPlayerResponse{Status: "System is busy"}, nil
	}

	if scheduleId == -1 {
		beego.Info("Setup LiveStream status")
		err := c.setupLiveStatus(music, groups)
		if err != nil {
			return &models.AudioPlayerResponse{Status: "Error setting LiveStream status"}, nil
		}
	} else {
		// Get date in long format
		now := time.Now()
		date, err := strconv.Atoi(now.Format("200601021504"))
		if err != nil {
			beego.Critical(err)
			return &models.AudioPlayerResponse{Status: "Error setting scheduled task status"}, nil
		}
		longDate := int64(date)
		// Register status
		_, err = models.AddStatus(&models.Status{ScheduleId: &models.Schedule{Id: scheduleId}, Status: 1, FiredOn: longDate})
		if err != nil {
			beego.Critical(err)
			return &models.AudioPlayerResponse{Status: "Error setting scheduled task status"}, nil
		}
	}

	beego.Info("Start Arduino configuration")
	switches, err := c.getSwitches(groups)
	if err != nil {
		return &models.AudioPlayerResponse{Status: "Error getting Arduino switches list"}, nil
	}

	err = c.configArduinos(switches)
	if err != nil {
		return &models.AudioPlayerResponse{Status: "Error setting Arduino relays"}, nil
	}

	beego.Info("Start playing stream")
	var ap models.AudioPlayer
	apStatus, err := ap.Play(music.Filename)
	if err != nil {
		beego.Critical(err)
		return &models.AudioPlayerResponse{Status: "Error playing stream"}, nil
	}
	return apStatus, nil
}

func (c *StreamController) getSwitches(groups []models.Group) (models.ArduinoSwitchList, error) {
	var groupMaps []interface{}
	for _, group := range groups {
		gms, err := models.GetAllGroupMap(map[string]string{"group_id": strconv.Itoa(group.Id)}, nil, nil, nil, 0, 1000)
		if err != nil {
			beego.Critical(err)
			return nil, err
		}
		groupMaps = append(groupMaps, gms...)
	}
	var locations []models.Location
	for _, gm := range groupMaps {
		beego.Debug("Group map id", gm.(models.GroupMap).Id)
		beego.Debug("Group map group_id", gm.(models.GroupMap).GroupId.Id)
		beego.Debug("Group map location_id", gm.(models.GroupMap).LocationId.Id)
		locations = append(locations, *gm.(models.GroupMap).LocationId)
	}

	o := orm.NewOrm()
	var switches models.ArduinoSwitchList
	for _, loc := range locations {
		_, err := o.LoadRelated(&loc, "SpeakerId")
		if err != nil {
			beego.Critical(err)
			return nil, err
		}
		//beego.Debug("Push switch with ArduinoId", loc.SpeakerId.ArduinoId.Id, "and relay", loc.SpeakerId.Relay)
		switches = append(switches, models.ArduinoSwitch{loc.SpeakerId.ArduinoId.Id, loc.SpeakerId.Relay})
	}
	sort.Sort(switches)
	for _, value := range switches {
		beego.Debug("Push switch with ArduinoId", value.ArduinoId, "and relay", value.Relay)
	}

	return switches, nil
}

func (c *StreamController) configArduinos(switches models.ArduinoSwitchList) error {
	if len(switches) > 0 {
		arduino, err := models.GetArduinoById(switches[0].ArduinoId)
		if err != nil {
			beego.Critical(err)
		}
		for _, speakerSwitch := range switches {
			if arduino.Id != speakerSwitch.ArduinoId {
				err = arduino.ApplyState()
				if err != nil {
					beego.Critical(err)
				}
				arduino, err = models.GetArduinoById(speakerSwitch.ArduinoId)
				if err != nil {
					beego.Critical(err)
				}
			}
			arduino.SetActiveRelay(speakerSwitch.Relay)
		}
		arduino.ApplyState()
		if err != nil {
			beego.Critical(err)
		}
	}
	return nil
}

func (c *StreamController) isAvailable() bool {
	status, err := models.GetAllStatus(nil, nil, nil, nil, 0, 10)
	if err != nil || len(status) > 0 {
		return false
	}
	return true
}

func (c *StreamController) setupLiveStatus(music *models.Music, groups []models.Group) error {
	now := time.Now()
	liveSchedule := models.Schedule{
		Name: "LiveStream",
		Description: "Automatically generated schedule for Live stream",
		Hour: now.Hour(),
		Minute: now.Minute(),
		DayOfMonth: now.Day(),
		Month: int(now.Month()),
		DayOfWeek: int(now.Weekday()),
		Year: now.Year(),
		Length: 45,
		MusicId: music,
		GroupId: &groups[0],
	}
	_, err := models.AddSchedule(&liveSchedule)
	if err != nil {
		beego.Critical(err)
		return err
	}

	date, err := strconv.Atoi(now.Format("20060102150405"))
	_, err = models.AddStatus(&models.Status{ScheduleId: &liveSchedule, Status: 1, FiredOn: int64(date)})
	if err != nil {
		beego.Critical(err)
		return err
	}
	return nil
}
