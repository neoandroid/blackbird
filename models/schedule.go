package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Schedule struct {
	Id          int       `orm:"column(id);auto"`
	Name        string    `orm:"column(name);size(50)"`
	Description string    `orm:"column(description);size(255)"`
	Hour        int       `orm:"column(hour)"`
	Minute      int       `orm:"column(minute)"`
	DayOfMonth  int       `orm:"column(day_of_month)"`
	Month       int       `orm:"column(month)"`
	DayOfWeek   int       `orm:"column(day_of_week)"`
	Year        int       `orm:"column(year)"`
	Length      int       `orm:"column(length)"`
	MusicId     *Music    `orm:"column(music_id);rel(fk)"`
	GroupId     *Group    `orm:"column(group_id);rel(fk)"`
}

func (t *Schedule) TableName() string {
	return "schedule"
}

func init() {
	orm.RegisterModel(new(Schedule))
}

// AddSchedule insert a new Schedule into database and returns
// last inserted Id on success.
func AddSchedule(m *Schedule) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetScheduleById retrieves Schedule by Id. Returns error if
// Id doesn't exist
func GetScheduleById(id int) (v *Schedule, err error) {
	o := orm.NewOrm()
	v = &Schedule{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSchedule retrieves all Schedule matches certain condition. Returns empty list if
// no records exist
func GetAllSchedule(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Schedule))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Schedule
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateSchedule updates Schedule by Id and returns error if
// the record to be updated doesn't exist
func UpdateScheduleById(m *Schedule) (err error) {
	o := orm.NewOrm()
	v := Schedule{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSchedule deletes Schedule by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSchedule(id int) (err error) {
	o := orm.NewOrm()
	v := Schedule{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Schedule{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func (s *Schedule) GetComposeDate() int64 {
	now := time.Now()
	year := s.Year
	month := s.Month
	dayOfWeek := s.DayOfWeek
	dayOfMonth := s.DayOfMonth
	hour := s.Hour
	minute := s.Minute
	if year == -1 {
		year = now.Year()
	}
	if month == -1 {
		month = int(now.Month())
	}
	if dayOfMonth == -1 {
		dayOfMonth = now.Day()
	}
	if hour == -1 {
		hour = now.Hour()
	}
	if minute == -1 {
		minute = now.Minute()
	}
	if dayOfWeek == -1 {
		dayOfWeek = int(now.Weekday())
	} else if dayOfWeek != int(now.Weekday()) {
		date, _ := time.Parse("2006-01-02T15:04", fmt.Sprintf("%d-%02d-%02dT%02d:%02d", year, month, dayOfMonth, hour, minute))
		daysAhead := dayOfWeek - int(date.Weekday())
		if daysAhead < 0 {
			daysAhead += 7
		}
		date = date.AddDate(0, 0, daysAhead)
		return int64(date.Year())*int64(100000000) + int64(date.Month())*int64(1000000) + int64(date.Day())*int64(10000) + int64(date.Hour())*int64(100) + int64(date.Minute())
	}
	return int64(year)*int64(100000000) + int64(month)*int64(1000000) + int64(dayOfMonth)*int64(10000) + int64(hour)*int64(100) + int64(minute)

}

func (s *Schedule) GetLengthInMinutes() int {
	remainder := s.Length % 60
	if remainder > 0 {
		return s.Length / 60 + 1
	} else {
		return s.Length / 60
	}
}

type ScheduleList []Schedule

func (sl ScheduleList) Len() int {
	return len(sl)
}

func (sl ScheduleList) Less(i, j int) bool {
	return sl[i].GetComposeDate() < sl[j].GetComposeDate()
}

func (sl ScheduleList) Swap(i, j int) {
	sl[i], sl[j] = sl[j], sl[i]
}
