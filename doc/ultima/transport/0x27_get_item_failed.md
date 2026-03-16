# (0x27) Get Item Failed

Sent by the server to deny a player's request to pick up an item. 2 bytes.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x27`
`0x01` | `BYTE` | `reason` | Rejection reason code (see values below)

## Reason codes

Code | Description
--- | ---
`0x00` | Display "You cannot pick that up."
`0x01` | Display "That is too far away."
`0x02` | Display "That is out of sight."
`0x03` | Display "That item does not belong to you. You will have to steal it."
`0x04` | Display "You are already holding an item."
`0x05` | Destroy the item
`0x06` | No message
