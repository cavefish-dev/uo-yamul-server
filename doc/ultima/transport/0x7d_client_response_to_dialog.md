# (0x7D) CLIENT_RESPONSE_TO_DIALOG

Client response to an Object Picker dialog (0x7C).

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x7D`
`0x01` | `UINT` | `dialogId` | Echoed back from the `0x7C` packet
`0x05` | `USHORT` | `menuId` | Echoed back from the `0x7C` packet
`0x07` | `USHORT` | `cindex` | 1-based index of the selected choice
`0x09` | `USHORT` | `modelNo` | Model number of the selected choice
`0x0B` | `BYTE[2]` | `unknown1` | Fixed padding: `0x00` `0x00`

**Total size:** 13 bytes
