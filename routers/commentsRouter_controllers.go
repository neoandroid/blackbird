package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:ArduinoController"] = append(beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:ArduinoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:ArduinoController"] = append(beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:ArduinoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:ArduinoController"] = append(beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:ArduinoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:ArduinoController"] = append(beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:ArduinoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:ArduinoController"] = append(beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:ArduinoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:GroupController"] = append(beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:GroupController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:GroupController"] = append(beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:GroupController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:GroupController"] = append(beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:GroupController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:GroupController"] = append(beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:GroupController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:GroupController"] = append(beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:GroupController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:GroupMapController"] = append(beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:GroupMapController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:GroupMapController"] = append(beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:GroupMapController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:GroupMapController"] = append(beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:GroupMapController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:GroupMapController"] = append(beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:GroupMapController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:GroupMapController"] = append(beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:GroupMapController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:LocationController"] = append(beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:LocationController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:LocationController"] = append(beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:LocationController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:LocationController"] = append(beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:LocationController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:LocationController"] = append(beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:LocationController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:LocationController"] = append(beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:LocationController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:MusicController"] = append(beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:MusicController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:MusicController"] = append(beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:MusicController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:MusicController"] = append(beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:MusicController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:MusicController"] = append(beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:MusicController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:MusicController"] = append(beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:MusicController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:ScheduleController"] = append(beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:ScheduleController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:ScheduleController"] = append(beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:ScheduleController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:ScheduleController"] = append(beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:ScheduleController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:ScheduleController"] = append(beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:ScheduleController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:ScheduleController"] = append(beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:ScheduleController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:SpeakerController"] = append(beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:SpeakerController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:SpeakerController"] = append(beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:SpeakerController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:SpeakerController"] = append(beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:SpeakerController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:SpeakerController"] = append(beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:SpeakerController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:SpeakerController"] = append(beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:SpeakerController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:StatusController"] = append(beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:StatusController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:StatusController"] = append(beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:StatusController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:StatusController"] = append(beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:StatusController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:StatusController"] = append(beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:StatusController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:StatusController"] = append(beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:StatusController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:StreamController"] = append(beego.GlobalControllerRouter["github.com/neoandroid/blackbird/controllers:StreamController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/:id`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
