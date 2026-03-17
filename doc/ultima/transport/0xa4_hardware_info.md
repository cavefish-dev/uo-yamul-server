# (0xA4) Hardware Info

Sent by the client to report system hardware information to the server. Contains screen resolution, color depth, CPU details, memory, video card info, and DirectX version. Fixed size of 149 bytes. The exact field layout varies slightly across client versions.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0xA4`
`0x01` | `BYTE` | `instance` | Client instance identifier
`0x02` | `UINT` | `screenWidth` | Screen width in pixels
`0x06` | `UINT` | `screenHeight` | Screen height in pixels
`0x0A` | `UINT` | `screenDepth` | Color depth in bits
`0x0E` | `UINT` | `refreshRate` | Monitor refresh rate in Hz
`0x12` | `BYTE` | `unknown1` | Unknown
`0x13` | `UINT` | `cpuSpeed` | CPU speed in MHz
`0x17` | `UINT` | `memory` | System memory in MB
`0x1B` | `UINT` | `unknown2` | Unknown (architecture or platform flags)
`0x1F` | `UINT` | `videoMemory` | Video card memory in MB
`0x23` | `UINT` | `videoVendorID` | Video card vendor ID
`0x27` | `UINT` | `videoDeviceID` | Video card device ID
`0x2B` | `UINT` | `videoDriverMajor` | Video driver major version
`0x2F` | `UINT` | `videoDriverMinor` | Video driver minor version
`0x33` | `UINT` | `videoDriverRevision` | Video driver revision
`0x37` | `UINT` | `videoDriverBuild` | Video driver build number
`0x3B` | `BYTE` | `distribution` | Client distribution type
`0x3C` | `BYTE` | `clientCount` | Number of UO clients running
`0x3D` | `BYTE` | `osType` | Operating system type
`0x3E` | `BYTE` | `osMajor` | Operating system major version
`0x3F` | `BYTE` | `osMinor` | Operating system minor version
`0x40` | `CHAR[64]` | `cpuVendor` | CPU vendor string, null-terminated
`0x80` | `BYTE[5]` | `padding` | Remaining zero-padded bytes
