# 	ELM Intake Manifold Pressure- Activity

Communicate with car's OBD-II system using ELM327 based devices.
IntakeManifoldPressure represents a command that checks the intake manifold pressure in kPa
Min: 0
Max: 255

## Installation
Command for Flogo CLI:
```console
flogo install github.com/wkarasz/flogo-elmobd/activity/elmintakemanifoldpressure
```

Link for Flogo Web UI:
```console
https://github.com/wkarasz/flogo-elmobd/activity/elmintakemanifoldpressure
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
