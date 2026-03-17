# MOBTYPES.MUL

Mobile (NPC/creature) type definitions used to select behavior and animation set for each body ID.

Reference: [SphereServer — uo_files/CUOMobtypes.h](https://github.com/Sphereserver/Source-X/blob/master/src/game/uo_files/CUOMobtypes.h)

## File structure

The file is a flat array of 4-byte records indexed by body ID. Record count = `FileSize / 4`.

### MobType record (4 bytes)

Source: `CUOMobtypesRec` (`m_type`, `m_flags`)

Offset | Size | Type | Name | Description
--- | --- | --- | --- | ---
`0x00` | `2` | `UWORD` | `Type` | Entity type enum (see table below)
`0x02` | `2` | `UWORD` | `Flags` | Behavior/animation flags

## Entity type values

| Value | Name | Description |
| --- | --- | --- |
| `0` | `Monster` | Standard land monster |
| `1` | `SeaMonster` | Aquatic creature |
| `2` | `Animal` | Non-hostile animal |
| `3` | `Human` | Human/humanoid body |
| `4` | `Equipment` | Wearable equipment body |

## Usage

The body ID used here is the same value stored in `AnimID` in `TILEDATA.MUL` static records. Servers use the type to determine which animation set and combat/movement behaviors apply to a given mobile.
