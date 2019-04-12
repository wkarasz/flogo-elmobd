# 	ELM Direct - Activity

Communicate with car's OBD-II system using ELM327 based devices.
Send raw/direct ELM commands to ELM327 chipset.

## Installation

```console
flogo install github.com/wkarasz/flogo-elmobd/activity/elmdirect
```

Link for flogo web:
```console
https://github.com/wkarasz/flogo-elmobd/activity/elmdirect
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
    },
    { "name": "directCmd",
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
| directCmd        | The raw command supported by the ELM chipset; e.g. AT@1
https://www.elmelectronics.com/wp-content/uploads/2017/01/AT_Command_Table.pdf|

# Outputs
| Output           | Description    |
|:-----------------|:---------------|
| result           | The result will contain a string response to the command or will contain an error message |
