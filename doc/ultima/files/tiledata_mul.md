# TILEDATA.MUL

Tile attribute definitions for all land and static tiles. Contains 512 land tile groups followed by static tile groups.

Reference: [Heptazane — Section 3.19](https://uo.stratics.com/heptazane/fileformats.shtml#3.19) · [SphereServer — uo_files/CUOTiledata.h](https://github.com/Sphereserver/Source-X/blob/master/src/game/uo_files/CUOTiledata.h)

## File structure

Two sections: land tile groups first, then static tile groups. Each group holds 32 tile records preceded by a 4-byte unknown header.

> **High Seas note:** Pre-HS clients use the structures below. HS+ clients expanded both structs. Check file size to detect the variant (see HS variant sections).

### Land tile group (836 bytes)

Offset | Size | Type | Name | Description
--- | --- | --- | --- | ---
`0x00` | `4` | `DWORD` | `Header` | Unknown; always present
`0x04` | `26*32` | `LandTile[32]` | `Tiles` | 32 land tile records

### Land tile record (26 bytes)

Offset | Size | Type | Name | Description
--- | --- | --- | --- | ---
`0x00` | `4` | `DWORD` | `Flags` | Tile property flags (see flags table)
`0x04` | `2` | `UWORD` | `TexID` | Texture map ID (`0` = no texture)
`0x06` | `20` | `CHAR[20]` | `Name` | Tile name (null-padded)

### Land tile record — High Seas variant (30 bytes)

Source: `CUOTerrainTypeRec_HS`

Offset | Size | Type | Name | Description
--- | --- | --- | --- | ---
`0x00` | `4` | `DWORD` | `Flags` | Tile property flags (see flags table)
`0x04` | `2` | `UWORD` | `TexID` | Texture map ID (`0` = no texture)
`0x06` | `4` | `DWORD` | `Unknown` | Added in HS expansion; purpose unknown
`0x0A` | `20` | `CHAR[20]` | `Name` | Tile name (null-padded)

---

### Static tile group (1,188 bytes)

Offset | Size | Type | Name | Description
--- | --- | --- | --- | ---
`0x00` | `4` | `DWORD` | `Header` | Unknown; always present
`0x04` | `37*32` | `StaticTile[32]` | `Tiles` | 32 static tile records

### Static tile record (37 bytes)

Offset | Size | Type | Name | Description
--- | --- | --- | --- | ---
`0x00` | `4` | `DWORD` | `Flags` | Item property flags (see flags table)
`0x04` | `1` | `UBYTE` | `Weight` | Item weight; `255` = immovable
`0x05` | `1` | `UBYTE` | `Quality` | Layer or light ID depending on tile type
`0x06` | `2` | `UWORD` | `Unknown` | Purpose unknown
`0x08` | `1` | `UBYTE` | `Unknown1` | Purpose unknown
`0x09` | `1` | `UBYTE` | `Quantity` | Class or armor value
`0x0A` | `2` | `UWORD` | `AnimID` | Body ID; add 50,000 or 60,000 for gump index
`0x0C` | `1` | `UBYTE` | `Unknown2` | Purpose unknown
`0x0D` | `1` | `UBYTE` | `Hue` | Colored light value
`0x0E` | `2` | `UWORD` | `Unknown3` | Purpose unknown
`0x10` | `1` | `UBYTE` | `Height` | Z-height or container capacity
`0x11` | `20` | `CHAR[20]` | `Name` | Tile name (null-padded)

### Static tile record — High Seas variant (41 bytes)

Source: `CUOItemTypeRec_HS`. `Flags` expands from `DWORD` (4 bytes) to `QWORD` (8 bytes); all subsequent offsets shift by 4.

Offset | Size | Type | Name | Description
--- | --- | --- | --- | ---
`0x00` | `8` | `QWORD` | `Flags` | Item property flags (64-bit in HS+)
`0x08` | `1` | `UBYTE` | `Weight` | Item weight; `255` = immovable
`0x09` | `1` | `UBYTE` | `Quality` | Layer or light ID depending on tile type
`0x0A` | `2` | `UWORD` | `Unknown` | Purpose unknown
`0x0C` | `1` | `UBYTE` | `Unknown1` | Purpose unknown
`0x0D` | `1` | `UBYTE` | `Quantity` | Class or armor value
`0x0E` | `2` | `UWORD` | `AnimID` | Body ID; add 50,000 or 60,000 for gump index
`0x10` | `1` | `UBYTE` | `Unknown2` | Purpose unknown
`0x11` | `1` | `UBYTE` | `Hue` | Colored light value
`0x12` | `2` | `UWORD` | `Unknown3` | Purpose unknown
`0x14` | `1` | `UBYTE` | `Height` | Z-height or container capacity
`0x15` | `20` | `CHAR[20]` | `Name` | Tile name (null-padded)

---

## Tile flags

| Flag name | Hex value | Meaning |
| --- | --- | --- |
| `Background` | `0x00000001` | Non-interactive background tile |
| `Weapon` | `0x00000002` | Weapon item |
| `Transparent` | `0x00000004` | See-through |
| `Translucent` | `0x00000008` | Semi-transparent |
| `Wall` | `0x00000010` | Collision wall |
| `Damaging` | `0x00000020` | Causes damage on contact |
| `Impassable` | `0x00000040` | Blocks movement |
| `Wet` | `0x00000080` | Water tile |
| `Unknown` | `0x00000100` | Unknown |
| `Surface` | `0x00000200` | Character can stand on this tile |
| `Bridge` | `0x00000400` | Bridge structure (half-height surface) |
| `Stackable` | `0x00000800` | Item can be stacked |
| `Window` | `0x00001000` | Window tile |
| `NoShoot` | `0x00002000` | Blocks projectiles |
| `PrefixA` | `0x00004000` | Use article "a" |
| `PrefixAn` | `0x00008000` | Use article "an" |
| `Internal` | `0x00010000` | Hair, beard — internal use |
| `Foliage` | `0x00020000` | Plant/leaf tile |
| `PartialHue` | `0x00040000` | Partial hue applied |
| `Unknown1` | `0x00080000` | Unknown |
| `Map` | `0x00100000` | Map item |
| `Container` | `0x00200000` | Container item |
| `Wearable` | `0x00400000` | Can be equipped |
| `LightSource` | `0x00800000` | Emits light |
| `Animated` | `0x01000000` | Has animation frames |
| `NoDiagonal` | `0x02000000` | Prevents diagonal movement |
| `Unknown2` | `0x04000000` | Unknown |
| `Armor` | `0x08000000` | Armor piece |
| `Roof` | `0x10000000` | Roof structure |
| `Door` | `0x20000000` | Door item |
| `StairBack` | `0x40000000` | Back staircase component |
| `StairRight` | `0x80000000` | Right staircase component |
