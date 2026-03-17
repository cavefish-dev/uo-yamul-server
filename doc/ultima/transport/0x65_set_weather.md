# (0x65) SET_WEATHER

Controls weather effects displayed to the client, including rain, snow, and storms with configurable intensity.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x65`
`0x01` | `BYTE` | `type` | Weather type (see values below)
`0x02` | `BYTE` | `num` | Count of active weather effects on screen
`0x03` | `BYTE` | `temperature` | Temperature value (currently unused)

**Total size:** 4 bytes

## Weather type values

Value | Effect
--- | ---
`0x00` | "It starts to rain"
`0x01` | "A fierce storm approaches."
`0x02` | "It begins to snow"
`0x03` | "A storm is brewing."
`0xFE` | Temperature adjustment (no visual effect)
`0xFF` | None (disables sound effects)

## Notes

- Maximum 70 weather effects displayable simultaneously.
- Messages only appear when weather initiates.
- Weather automatically ceases after 6 minutes without updates.
- Teleportation can reset weather display.
- Combining snow and rain requires sending the rain packet first, then the snow packet with cumulative effect counts.
- The temperature field has no current functional impact.
