# (0xBF) Generic Command

General-purpose packet with subcommand routing for various server-to-client and client-to-server communications.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0xBF`
`0x01` | `USHORT` | `length` | Total packet length
`0x03` | `USHORT` | `subcommand` | Identifies the specific operation
`0x05` | `BYTE[]` | `submessage` | Variable-length subcommand payload

## Subcommand 0x0001 — Initialize Fast Walk Prevention

Server message. Establishes an anti-cheat stack of movement keys on the client. Total packet size: 29 bytes.

Offset | Type | Name | Description
--- | --- | --- | ---
`0x05` | `UINT` | `key1` | First stack element (top)
`0x09` | `UINT` | `key2` | Second element
`0x0D` | `UINT` | `key3` | Third element
`0x11` | `UINT` | `key4` | Fourth element
`0x15` | `UINT` | `key5` | Fifth element
`0x19` | `UINT` | `key6` | Bottom stack element

## Subcommand 0x0002 — Add Key to Fast Walk Stack

Server message. Pushes a new validation key onto the top of the fast-walk prevention stack. Total packet size: 9 bytes.

Offset | Type | Name | Description
--- | --- | --- | ---
`0x05` | `UINT` | `newkey` | Key added to the top of the stack

## Subcommand 0x0005 — Unknown

Client message. Purpose unknown. Total packet size: 13 bytes.

Offset | Type | Name | Description
--- | --- | --- | ---
`0x05` | `UINT` | `unknown1` | Value: `0x00000320`
`0x09` | `UINT` | `unknown2` | Value: `0x00000005`

## Subcommand 0x0006 — Party System

Contains nested subsubcommands for party management.

Offset | Type | Name | Description
--- | --- | --- | ---
`0x05` | `BYTE` | `subsubcommand` | Party action identifier

### Subsubcommand 0x01 — Add Party Member

Client message. Requests adding a party member. If id is 0, the client enters targeting mode.

Offset | Type | Name | Description
--- | --- | --- | ---
`0x06` | `UINT` | `id` | Target player serial (0 = activate targeting)
`0x0A` | `BYTE` | `numMembers` | Total party member count
`—` | `UINT[]` | `member_ids` | Serial for each party member (repeats numMembers times)

### Subsubcommand 0x02 — Remove Party Member

No detailed field structure documented.

### Subsubcommand 0x03 — Send Message to Single Member

No detailed field structure documented.

### Subsubcommand 0x04 — Broadcast Message to All Members

No detailed field structure documented.

### Subsubcommand 0x06 — Grant Looting Permissions

No detailed field structure documented.

### Subsubcommand 0x08 — Accept Party Invitation

No detailed field structure documented.

### Subsubcommand 0x09 — Decline Party Invitation

No detailed field structure documented.

## Subcommand 0x0008 — Set Cursor Hue

No detailed field structure documented.

## Subcommand 0x000B — Client Language Specification

No detailed field structure documented.

## Subcommand 0x000C — Closed Status Gump Notification

No detailed field structure documented.
