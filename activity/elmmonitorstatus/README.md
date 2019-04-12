# 	ELM Monitor Status - Activity

Communicate with car's OBD-II system using ELM327 based devices.
MonitorStatus represents a command that checks the status since DTCs were cleared last time.  This includes the MIL status and the amount of DTCs.

## Installation
Command for Flogo CLI:
```console
flogo install github.com/wkarasz/flogo-elmobd/activity/elmmonitorstatus
```

Link for Flogo Web UI:
```console
https://github.com/wkarasz/flogo-elmobd/activity/elmmonitorstatus
```

## Schema
Inputs and Outputs:
```json
{
  "inputs":[
    {
      "name": "devicePath",
      "type": "string",
      "required": true
    }
  ],
  "outputs": [
    {
      "name": "result",
      "type": "string"
    }
  ]
}
```
## Inputs
| Input            | Description    |
|:-----------------|:---------------|
| devicePath       | The path to the ELM device on the host system; e.g. /dev/ttyUSB0 |

# Outputs
| Output           | Description    |
|:-----------------|:---------------|
| result           | The result will contain a string response of the command or will contain an error message |
