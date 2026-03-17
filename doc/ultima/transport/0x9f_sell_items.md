# (0x9F) Sell Items

Sent by the client to confirm selling selected items to a vendor. Contains the vendor's serial and the list of items with quantities the player wishes to sell. Variable length.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x9F`
`0x01` | `USHORT` | `size` | Total packet length in bytes
`0x03` | `UINT` | `vendorUID` | Serial of the vendor NPC
`0x07` | `USHORT` | `itemCount` | Number of items being sold
`0x09` | `UINT` | `itemUID` | Serial of first item to sell
`0x0D` | `USHORT` | `amount` | Quantity of first item to sell
`...` | `...` | `...` | Additional `itemUID(UINT) + amount(USHORT)` pairs
