# (0x7C) OBJECT_PICKER

Can be used to display either a grey menu with options or the old-fashioned object picker. Variable length.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x7C`
`0x01` | `USHORT` | `size` | Total size of the packet
`0x03` | `UINT` | `dialog` | The dialog's serial
`0x07` | `USHORT` | `menu` | The menu's serial
`0x09` | `BYTE` | `len_title` | Length of the title string
`0x0A` | `CHAR[]` | `title` | Title of the dialog
`--` | `BYTE` | `options` | Number of options / items

## Options loop

Repeated for each option.

Offset | Type | Name | Description
--- | --- | --- | ---
`--` | `USHORT` | `artwork` | Item's artwork number; use `0x0000` for an option menu
`--` | `USHORT` | `hue` | Item's hue
`--` | `BYTE` | `length` | Length of the option text
`--` | `CHAR[]` | `text` | Text to display for the option
