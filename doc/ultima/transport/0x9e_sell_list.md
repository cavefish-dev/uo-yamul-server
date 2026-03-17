# (0x9E) Sell List

Sent by the server to open the sell window, showing items the vendor will buy from the player. Each entry includes the item's serial, graphic, hue, quantity, price, and name. Variable length.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x9E`
`0x01` | `USHORT` | `size` | Total packet length in bytes
`0x03` | `UINT` | `vendorUID` | Serial of the vendor NPC
`0x07` | `USHORT` | `itemCount` | Number of item entries that follow
`0x09` | `UINT` | `itemUID` | Serial of first item
`0x0D` | `USHORT` | `graphic` | Graphic ID of first item
`0x0F` | `USHORT` | `hue` | Hue of first item
`0x11` | `USHORT` | `qty` | Quantity of first item
`0x13` | `USHORT` | `price` | Price the vendor will pay per unit
`0x15` | `USHORT` | `nameLen` | Length of the item name string
`0x17` | `CHAR[nameLen]` | `name` | Item name (not null-terminated)
`...` | `...` | `...` | Additional item entries repeating from `itemUID`
