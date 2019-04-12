# 	ELM OBD Standards - Activity

Communicate with car's OBD-II system using ELM327 based devices.
OBDStandards represents a command that checks the OBD standards the vehicle conforms to as a single decimal value:
1       OBD-II as defined by the CARB
2       OBD as defined by the EPA
3       OBD and OBD-II
4       OBD-I
5       Not OBD compliant
6       EOBD (Europe)
7       EOBD and OBD-II
8       EOBD and OBD
9       EOBD, OBD and OBD II
10      JOBD (Japan)
11      JOBD and OBD II
12      JOBD and EOBD
13      JOBD, EOBD, and OBD II
14      Reserved
15      Reserved
16      Reserved
17      Engine Manufacturer Diagnostics (EMD)
18      Engine Manufacturer Diagnostics Enhanced (EMD+)
19      Heavy Duty On-Board Diagnostics (Child/Partial) (HD OBD-C)
20      Heavy Duty On-Board Diagnostics (HD OBD)
21      World Wide Harmonized OBD (WWH OBD)
22      Reserved
23      Heavy Duty Euro OBD Stage I without NOx control (HD EOBD-I)
24      Heavy Duty Euro OBD Stage I with NOx control (HD EOBD-I N)
25      Heavy Duty Euro OBD Stage II without NOx control (HD EOBD-II)
26      Heavy Duty Euro OBD Stage II with NOx control (HD EOBD-II N)
27      Reserved
28      Brazil OBD Phase 1 (OBDBr-1)
29      Brazil OBD Phase 2 (OBDBr-2)
30      Korean OBD (KOBD)
31      India OBD I (IOBD I)
32      India OBD II (IOBD II)
33      Heavy Duty Euro OBD Stage VI (HD EOBD-IV)
34-250  Reserved
251-255 Not available for assignment (SAE J1939 special meaning)

## Installation
Command for Flogo CLI:
```console
flogo install github.com/wkarasz/flogo-elmobd/activity/elmobdstandards
```

Link for Flogo Web UI:
```console
https://github.com/wkarasz/flogo-elmobd/activity/elmobdstandards
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
