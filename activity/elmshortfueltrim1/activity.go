package elmshortfueltrim1

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
        "github.com/wkarasz/elmobd"
        "flag"
)

// log is the default logger which we'll use to log
var log = logger.GetLogger("activity-elm-ShortFuelTrim1")

// String to hold the pointer for serial flag object
var serialPathP string

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
        log.Infof("Device path capture [%s]", device)

        if flag.Lookup("serial") == nil {
                flag.StringVar(
                        &serialPathP,
                        "serial",
                        device,
			"Path to the serial device to use",
                )
        }
        flag.Parse()

        dev, err := elmobd.NewTestDevice(serialPathP, false)

        if err != nil {
                log.Infof("Failed to create new device [%s]", err)
                return
        }

        elmObj, err := dev.RunOBDCommand(elmobd.NewShortFuelTrim1())

	if err != nil {
                log.Infof("Failed to Short Fuel Trim 1 [%s]", err)
                return
        }

        log.Infof("Short Fuel Trim 1 is [%s]", elmObj.ValueAsLit())
        context.SetOutput("result", elmObj.ValueAsLit())
        dev = nil

	return true, nil
}
