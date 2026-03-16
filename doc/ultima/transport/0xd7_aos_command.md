# (0xD7) AOS Command

Age of Shadows (AOS) command packet used to perform various actions, mostly related to AOS features such as house customization and special moves.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0xD7`
`0x01` | `USHORT` | `packetsize` | Total packet length
`0x03` | `UINT` | `serial` | Target object serial
`0x07` | `USHORT` | `subcommand` | Specific action type
`0x09` | `BYTE` | `unknown` | Reserved/unknown field

## Subcommands

Subcommand | Description
--- | ---
`0x0002` | House Customization — Backup
`0x0003` | House Customization — Restore
`0x0004` | House Customization — Commit
`0x0005` | House Customization — Destroy Item
`0x0006` | House Customization — Place Item
`0x000C` | House Customization — Exit
`0x000D` | House Customization — Place Multi (Stairs)
`0x000E` | House Customization — Synch
`0x0010` | House Customization — Clear
`0x0012` | House Customization — Switch Floors
`0x0019` | Special Moves — Activate/Deactivate
`0x001A` | House Customization — Revert

No individual field tables are documented for these subcommands beyond the base packet structure.
