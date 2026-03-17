# (0x3B) BUY_ITEMS

This is sent by the client to buy items from a vendor. When sent by the server with numItems = 0, it removes the buy gump. Variable length.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x3B`
`0x01` | `USHORT` | `blockSize` | Total packet size in bytes
`0x03` | `UINT` | `vendorID` | Vendor serial identifier
`0x07` | `BYTE` | `numItems` | Number of items to buy; 0 removes the buy gump

Each item entry repeats for every item being purchased:

Offset (per item) | Type | Name | Description
--- | --- | --- | ---
`+0x00` | `BYTE` | `layer` | Item layer, usually `0x1A`
`+0x01` | `UINT` | `itemID` | Item serial from the `0x3C` container packet
`+0x05` | `USHORT` | `amount` | Quantity to purchase
