# flogo-elmobd
flogo-elmobd is a library of Flogo activities designed for interfacing with a car's OBD-II system using ELM327 based devices.  The underlying libraries is based on the work from github.com/rzetterberg/elmobd.<br>
<br>
An ELM327 device is a requirement to use this library - this may be a OBDII USB-dongle or connection via Bluetooth.<br>
<br>
The library consists of pre-cast functions for the most common OBD commands and a direct device command function for everything else not yet implemented.
<br>
* Direct device command
  * [activity/elmdirect](tree/master/activity/elmdirect)
* Pre-cast functions
  * [activity/elmcoolanttemperature](github.com/wkarasz/flogo-elmobd/tree/master/activity/elmcoolanttemperature)
  * [activity/elmdistsincedtcclear](github.com/wkarasz/flogo-elmobd/activity/elmdistsincedtcclear)
  * [activity/elmengineload](github.com/wkarasz/flogo-elmobd/activity/elmengineload)
  * [activity/elmenginerpm](github.com/wkarasz/flogo-elmobd/activity/elmenginerpm)
  * [activity/elmfuel](github.com/wkarasz/flogo-elmobd/activity/elmfuel)
  * [activity/elmfuelpressure](github.com/wkarasz/flogo-elmobd/activity/elmfuelpressure)
  * [activity/elmgetdevice](github.com/wkarasz/flogo-elmobd/activity/elmgetdevice)
  * [activity/elmintakeairtemperature](github.com/wkarasz/flogo-elmobd/activity/elmintakeairtemperature)
  * [activity/elmlongfueltrim1](github.com/wkarasz/flogo-elmobd/activity/elmlongfueltrim1)
  * [activity/elmlongfueltrim2](github.com/wkarasz/flogo-elmobd/activity/elmlongfueltrim2)
  * [activity/elmmafairflowrate](github.com/wkarasz/flogo-elmobd/activity/elmmafairflowrate)
  * [activity/elmmonitorstatus](github.com/wkarasz/flogo-elmobd/activity/elmmonitorstatus)
  * [activity/elmobdstandards](github.com/wkarasz/flogo-elmobd/activity/elmobdstandards)
  * [activity/elmruntimesincestart](github.com/wkarasz/flogo-elmobd/activity/elmruntimesincestart)
  * [activity/elmshortfueltrim1](github.com/wkarasz/flogo-elmobd/activity/elmshortfueltrim1)
  * [activity/elmshortfueltrim2](github.com/wkarasz/flogo-elmobd/activity/elmshortfueltrim2)
  * [activity/elmthrottleposition](github.com/wkarasz/flogo-elmobd/activity/elmthrottleposition)
  * [activity/elmtimingadvance](github.com/wkarasz/flogo-elmobd/activity/elmtimingadvance)
  * [activity/elmvehiclespeed](github.com/wkarasz/flogo-elmobd/activity/elmvehiclespeed)
