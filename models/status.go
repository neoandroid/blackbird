package models

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Status struct {
	Id         int       `orm:"column(id);auto"`
	ScheduleId *Schedule `orm:"column(schedule_id);rel(fk)"`
	Status     int8      `orm:"column(status)"`
	FiredOn    int64     `orm:"column(fired_on)"`
}

func (t *Status) TableName() string {
	return "status"
}

func init() {
	orm.RegisterModel(new(Status))
}

// AddStatus insert a new Status into database and returns
// last inserted Id on success.
func AddStatus(m *Status) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetStatusById retrieves Status by Id. Returns error if
// Id doesn't exist
func GetStatusById(id int) (v *Status, err error) {
	o := orm.NewOrm()
	v = &Status{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllStatus retrieves all Status matches certain condition. Returns empty list if
// no records exist
func GetAllStatus(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Status))
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

	var l []Status
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

// UpdateStatus updates Status by Id and returns error if
// the record to be updated doesn't exist
func UpdateStatusById(m *Status) (err error) {
	o := orm.NewOrm()
	v := Status{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteStatus deletes Status by Id and returns error if
// the record to be deleted doesn't exist
func DeleteStatus(id int) (err error) {
	o := orm.NewOrm()
	v := Status{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Status{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func CleanStatus() (err error) {
	// Get date in long format
	now := time.Now()
	date, err := strconv.Atoi(now.Format("200601021504"))
	if err != nil {
		beego.Critical(err)
		return err
	}
	longDate := int64(date)

	status, err := GetAllStatus(nil, nil, nil, nil, 0, 10)
	if err != nil {
		return err
	}
	for _, value := range status {
		schedule, err := GetScheduleById(value.(Status).ScheduleId.Id)
		if err != nil {
			return err
		}
		finishDate := value.(Status).FiredOn + int64(schedule.GetLengthInMinutes())
		beego.Debug("Finish date", finishDate)
		beego.Debug("Long date", longDate)
		if finishDate <= longDate {
			err = DeleteStatus(value.(Status).Id)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
