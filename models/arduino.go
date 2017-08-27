package models

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Arduino struct {
	Id          int        `orm:"column(id);auto"`
	Ipv4        string     `orm:"column(ipv4);size(15)"`
	Ipv6        string     `orm:"column(ipv6);size(50)"`
	Port        int        `orm:"column(port)"`
	RelayCount  int        `orm:"column(relay_count)"`
	Location    string     `orm:"column(location);size(255)"`
	Status      int8       `orm:"column(status)"`
	conn        net.Conn   `json:"-"`
	relaysState [21]string `json:"-"`
}

func (t *Arduino) TableName() string {
	return "arduino"
}

func init() {
	orm.RegisterModel(new(Arduino))
}

// AddArduino insert a new Arduino into database and returns
// last inserted Id on success.
func AddArduino(m *Arduino) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetArduinoById retrieves Arduino by Id. Returns error if
// Id doesn't exist
func GetArduinoById(id int) (v *Arduino, err error) {
	o := orm.NewOrm()
	v = &Arduino{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllArduino retrieves all Arduino matches certain condition. Returns empty list if
// no records exist
func GetAllArduino(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Arduino))
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

	var l []Arduino
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

// UpdateArduino updates Arduino by Id and returns error if
// the record to be updated doesn't exist
func UpdateArduinoById(m *Arduino) (err error) {
	o := orm.NewOrm()
	v := Arduino{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteArduino deletes Arduino by Id and returns error if
// the record to be deleted doesn't exist
func DeleteArduino(id int) (err error) {
	o := orm.NewOrm()
	v := Arduino{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Arduino{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func TurnOnAllArduinos() error {
	beego.Info("Turnning on all arduinos")
	arduinos, err := GetAllArduino(map[string]string{"status": "1" }, nil, nil, nil, 0, 10)
	if err != nil {
		beego.Debug(err)
		return err
	}
	for _, ardu := range arduinos {
		arduino := ardu.(Arduino)
		arduino.SetAllOn()
		err := arduino.ApplyState()
		if err!= nil {
			beego.Debug(err)
		}
		beego.Info("Operation ON completed for arduino id", arduino.Id)
	}
	return nil
}

func (a *Arduino) ApplyState() error {
	beego.Debug(a.relaysState)
	if a.Status == 1 {
		beego.Debug("Connect to arduino id", a.Id)
		err := a.connect()
		if err != nil {
			beego.Critical(err)
		}
		command := strings.Join(a.relaysState[:], "")
		beego.Debug("Send command", command)
		a.sendLine(command)
		retVal, err := a.readLine()
		if err != nil {
			beego.Critical(err)
		}
		beego.Debug("Arduino result", retVal)
		retry := 1
		reg, _ := regexp.Compile("OK.*")
		for reg.MatchString(retVal) == false && retry < 5 {
			beego.Debug("Retrying send command...", retry)
			a.sendLine(command)
			beego.Debug("Read confirmation command")
			retVal, err := a.readLine()
			if err != nil {
				beego.Critical(err)
			}
			beego.Debug("Arduino result", retVal)
			retry++
		}
		if retVal == "" {
			a.disconnect()
			return errors.New("Error sending command to Arduino")
		}
		a.disconnect()
		return nil
	}
	return nil
}

// Not used
func (a *Arduino) changeRelay(pos int, HL string) error {
	if a.conn != nil {
		a.sendLine("S")
		state, err := a.readLine()
		stateValues := strings.Fields(state)
		cmd := "D"
		for i := 1; i < 21; i++ {
			if i == pos {
				cmd = cmd + HL
			} else {
				cmd = cmd + stateValues[i]
			}
			beego.Debug("Command values:", cmd)
		}
		a.sendLine(cmd)
		result, err := a.readLine()
		if len(result) == 0 {
			beego.Critical(err)
			return err
		}
		return nil
	}
	return errors.New("Error: Connection not established")
}

func (a *Arduino) connect() error {
	if a.conn != nil {
		beego.Critical("Already connected to IP:", a.Ipv4)
		return errors.New("Error: Already connected to IP " + a.Ipv4)
	}
	var err error
	a.conn, err = net.DialTimeout("tcp", a.Ipv4 + ":" + strconv.Itoa(a.Port), time.Duration(3)*time.Second)
	if err != nil {
		// TODO: handle error
		beego.Critical(err)
		return err
	}
	return nil
}

func (a *Arduino) disconnect() error {
	if a.conn != nil {
		err := a.conn.Close()
		if err != nil {
			beego.Critical(err)
			return err
		}
	}
	return nil
}

func (a *Arduino) readLine() (string, error) {
	if a.conn != nil {
		line, err := bufio.NewReader(a.conn).ReadString('\n')
		if err != nil {
			beego.Critical(err)
			return "", err
		}
		return line, nil
	}
	// No connection established
	return "", errors.New("Error: Connection not established")
}

func (a *Arduino) sendLine(message string) error {
	if a.conn != nil {
		fmt.Fprintf(a.conn, message + "\r\n")
	}
	return nil
}

func (a *Arduino) SetActiveRelay(relay int) {
	if a.relaysState[0] == "" {
		a.relaysState = [21]string{"D","L","L","L","L","L","L","L","L","L","L","L","L","L","L","L","L","L","L","L","L"}
	}
	a.relaysState[relay] = "H"
}

func (a *Arduino) SetAllOn() {
	a.relaysState = [21]string{"D","H","H","H","H","H","H","H","H","H","H","H","H","H","H","H","H","H","H","H","H"}
}


type ArduinoSwitch struct {
	ArduinoId int
	Relay     int
}

type ArduinoSwitchList []ArduinoSwitch

func (as ArduinoSwitchList) Len() int {
	return len(as)
}

func (as ArduinoSwitchList) Less(i, j int) bool {
	return as[i].ArduinoId < as[j].ArduinoId
}

func (as ArduinoSwitchList) Swap(i, j int) {
	as[i], as[j] = as[j], as[i]
}

