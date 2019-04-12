# 	ELM Distance Since DTC Clear - Activity

Communicate with car's OBD-II system using ELM327 based devices.
DistSinceDTCClear represents a command that checks distance (kilometers) since last DTC clear
Distance accumulated since DTCs were cleared (via an external test equipment or possibly, a battery disconnect).  If greater than 65,535 km have occured, CLR_DIST shall remain at 65,535 km and not wrap to zero.


## Installation
Command for Flogo CLI:
```console
flogo install github.com/wkarasz/flogo-elmobd/activity/elmdistsincedtcclear
```

Link for Flogo Web UI:
```console
https://github.com/wkarasz/flogo-elmobd/activity/elmdistsincedtcclear
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
