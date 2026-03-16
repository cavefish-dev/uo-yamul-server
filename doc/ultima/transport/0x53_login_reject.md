# (0x53) LOGIN_REJECT

This is sent by the server to display a variety of pre-defined rejection messages, most related to logging in to a shard. Fixed size of 2 bytes.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x53`
`0x01` | `BYTE` | `msgid` | Message identifier (see values below)

## Message codes

Value | Description
--- | ---
`0x00` | Incorrect password
`0x01` | This character does not exist any more
`0x02` | This character already exists
`0x03` | Could not attach to game server
`0x04` | Could not attach to game server
`0x05` | Another character is already logged in
`0x06` | Synchronization error
`0x07` | Idle too long
`0x08` | Could not attach to game server
`0x09` | Character transfer in progress
