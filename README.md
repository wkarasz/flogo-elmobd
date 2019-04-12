# flogo-elmobd
flogo-elmobd is a library of Flogo activities designed for interfacing with a car's OBD-II system using ELM327 based devices.  The underlying libraries is based on the work from github.com/rzetterberg/elmobd.<br>
<br>
An ELM327 device is a requirement to use this library - this may be a OBDII USB-dongle or connection via Bluetooth.<br>
<br>
The library consists of pre-cast functions for the most common OBD commands and a direct device command function for everything else not yet implemented.
<br>
* Direct device command
  * [activity/elmdirect](activity/elmdirect/README.md)
* Pre-cast functions
  * [activity/elmcoolanttemperature](activity/elmcoolanttemperature/README.md)
  * [activity/elmdistsincedtcclear](activity/elmdistsincedtcclear/README.md)
  * [activity/elmengineload](activity/elmengineload/README.md)
  * [activity/elmenginerpm](activity/elmenginerpm/README.md)
  * [activity/elmfuel](activity/elmfuel/README.md)
  * [activity/elmfuelpressure](activity/elmfuelpressure/README.md)
  * [activity/elmgetdevice](activity/elmgetdevice/README.md)
  * [activity/elmintakeairtemperature](activity/elmintakeairtemperature/README.md)
  * [activity/elmlongfueltrim1](activity/elmlongfueltrim1/README.md)
  * [activity/elmlongfueltrim2](activity/elmlongfueltrim2/README.md)
  * [activity/elmmafairflowrate](activity/elmmafairflowrate/README.md)
  * [activity/elmmonitorstatus](activity/elmmonitorstatus/README.md)
  * [activity/elmobdstandards](activity/elmobdstandards/README.md)
  * [activity/elmruntimesincestart](activity/elmruntimesincestart/README.md)
  * [activity/elmshortfueltrim1](activity/elmshortfueltrim1/README.md)
  * [activity/elmshortfueltrim2](activity/elmshortfueltrim2/README.md)
  * [activity/elmthrottleposition](activity/elmthrottleposition/README.md)
  * [activity/elmtimingadvance](activity/elmtimingadvance/README.md)
  * [activity/elmvehiclespeed](activity/elmvehiclespeed/README.md)
