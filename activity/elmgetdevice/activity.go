package elmgetdevice

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/wkarasz/elmobd"
	"flag"
)

// log is the default logger which we'll use to log
var log = logger.GetLogger("activity-elm-getdevice")

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error)  {

	// do eval
	device := context.GetInput("devicePath").(string)
	log.Infof("Device path capture: [%s]", device)
	serialPath := flag.String(
		"serial",
		device,
		"Path to the serial device to use",
	)

	flag.Parse()
	
	dev, err := elmobd.NewTestDevice(*serialPath, false)

	if err != nil {
		log.Infof("Failed to create new device: [%s]", err)
		return
	}

	version, err := dev.GetVersion()

	if err != nil {
		log.Infof("Failed to get version: [%s]", err)
		return
	}

	log.Infof("Device has version [%s]", version)
	context.SetOutput("result", "Device has version "+version)
	dev = nil
	return true, nil
}
