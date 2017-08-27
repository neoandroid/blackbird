// @APIVersion 1.0.0
// @Title Blackbird API
// @Description Backend service for PA systems (formerly SoundSwitchWS)
// @Contact neoandroid@kaledoniah.net
// @TermsOfServiceUrl http://github.com/neoandroid/blackbird/
// @License GNU General Public License v3.0
// @LicenseUrl https://www.gnu.org/licenses/gpl-3.0.en.html
package routers

import (
	"github.com/neoandroid/blackbird/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/arduino",
			beego.NSInclude(
				&controllers.ArduinoController{},
			),
		),

		beego.NSNamespace("/group",
			beego.NSInclude(
				&controllers.GroupController{},
			),
		),

		beego.NSNamespace("/groupMap",
			beego.NSInclude(
				&controllers.GroupMapController{},
			),
		),

		beego.NSNamespace("/location",
			beego.NSInclude(
				&controllers.LocationController{},
			),
		),

		beego.NSNamespace("/music",
			beego.NSInclude(
				&controllers.MusicController{},
			),
		),

		beego.NSNamespace("/schedule",
			beego.NSInclude(
				&controllers.ScheduleController{},
			),
		),

		beego.NSNamespace("/speaker",
			beego.NSInclude(
				&controllers.SpeakerController{},
			),
		),

		beego.NSNamespace("/status",
			beego.NSInclude(
				&controllers.StatusController{},
			),
		),

		beego.NSNamespace("/stream",
			beego.NSInclude(
				&controllers.StreamController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
