# SKILLS.MUL

Skill name definitions. Accessed via `SKILLS.IDX`.

Reference: [Heptazane — Section 3.14](https://uo.stratics.com/heptazane/fileformats.shtml#3.14)

## File structure

Each skill entry starts at the position given by `SKILLS.IDX`.

### Skill entry

Offset | Size | Type | Name | Description
--- | --- | --- | --- | ---
`0x00` | `1` | `UBYTE` | `Button` | `1` = skill has an action button; `0` = no button
`0x01` | var | `CHAR[]` | `Name` | Skill name, null-terminated string
