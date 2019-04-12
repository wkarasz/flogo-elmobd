# 	ELM Throttle Position - Activity

Communicate with car's OBD-II system using ELM327 based devices.
ThrottlePosition represents a command that checks the throttle position in percent
Min: 0.0
Max: 1.0

## Installation
Command for Flogo CLI:
```console
flogo install github.com/wkarasz/flogo-elmobd/activity/elmthrottleposition
```

Link for Flogo Web UI:
```console
https://github.com/wkarasz/flogo-elmobd/activity/elmthrottleposition
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
