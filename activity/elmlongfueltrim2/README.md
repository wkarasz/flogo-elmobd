# 	ELM Long Fuel Trim 2 - Activity

Communicate with car's OBD-II system using ELM327 based devices.
Long Fuel Trim 2 represents a command that checks the long term fuel trim for bank 2 in percent
Min: -1
Max: .9922

## Installation
Command for Flogo CLI:
```console
flogo install github.com/wkarasz/flogo-elmobd/activity/elmlongfueltrim2
```

Link for Flogo Web UI:
```console
https://github.com/wkarasz/flogo-elmobd/activity/elmlongfueltrim2
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
