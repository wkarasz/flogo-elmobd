# flogo-elmobd
flogo-elmobd is a library of Flogo activities designed for interfacing with a car's OBD-II system using ELM327 based devices.<br>
<br>
An ELM327 device is a requirement to use this library - this may be a OBDII USB-dongle or connection via Bluetooth.<br>
<br>
The library consists of pre-cast functions for the most common OBD commands and a direct device command function for everything else not yet implemented.
<br>
* Direct device command
  * activity/elmdirect
* Pre-cast functions
  * activity/elmcoolanttemperature
  * activity/elmdistsincedtcclear
  * activity/elmengineload
  * activity/elmenginerpm
  * activity/elmfuel
  * activity/elmfuelpressure
  * activity/elmgetdevice
  * activity/elmintakeairtemperature
  * activity/elmlongfueltrim1
  * activity/elmlongfueltrim2
  * activity/elmmafairflowrate
  * activity/elmmonitorstatus
  * activity/elmobdstandards
  * activity/elmruntimesincestart
  * activity/elmshortfueltrim1
  * activity/elmshortfueltrim2
  * activity/elmthrottleposition
  * activity/elmtimingadvance
  * activity/elmvehiclespeed
