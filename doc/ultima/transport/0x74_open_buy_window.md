# (0x74) OPEN_BUY_WINDOW

Sends shop inventory information to the client. Variable length.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x74`
`0x01` | `USHORT` | `blockSize` | Total packet size
`0x03` | `UINT` | `vendorID` | Vendor identifier
`0x07` | `BYTE` | `numItems` | Number of items in inventory
`--` | `UINT` | `price` | Item cost (repeats per item)
`--` | `BYTE` | `length` | Length of the item description text
`--` | `CHAR[]` | `name` | Item description text (repeats per item)

## Notes

- This packet must be preceded by a Describe Contents packet (0x3C) with container ID `(vendorID | 0x40000000)`, followed by an Open Container packet (0x24) using `vendorID` with model number `0x0030`.
