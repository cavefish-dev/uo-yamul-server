# (0xB1) Generic Gump Trigger

This is sent when the user responds to a gump.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0xB1`
`0x01` | `USHORT` | `size` | Packet size
`0x03` | `UINT` | `player` | Player serial
`0x07` | `UINT` | `serial` | Gump serial
`0x0B` | `UINT` | `button` | The serial of the button that was pressed
`0x0F` | `UINT` | `switches` | The number of switches (checkboxes/radios) in the loop

Repeated for each selected switch (loops `switches` times):

Offset | Type | Name | Description
--- | --- | --- | ---
`—` | `UINT` | `switchid` | The switch's serial

Following the switches loop:

Offset | Type | Name | Description
--- | --- | --- | ---
`—` | `UINT` | `entries` | The number of text entry responses

Repeated for each text entry (loops `entries` times):

Offset | Type | Name | Description
--- | --- | --- | ---
`—` | `USHORT` | `number` | Text entry serial
`—` | `USHORT` | `length` | The length of the text
`—` | `UNI[]` | `text` | Text entry content
