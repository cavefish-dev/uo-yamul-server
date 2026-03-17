# (0x82) Login Error

Sent by the login server to indicate that the login request failed. Contains a reason code explaining the failure. Fixed size of 2 bytes.

Note: this packet is sometimes listed as "Connect to Game Server" in older references, but that is incorrect — 0x8C is the relay packet. This packet is the login failure response.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x82`
`0x01` | `BYTE` | `reason` | Failure reason: `0x00`=invalid credentials, `0x01`=account in use, `0x02`=account blocked, `0x03`=credentials invalid, `0xFE`=idle timeout, `0xFF`=communication error

## Example packet (hex)

```
-- Login error: invalid credentials
82 00

-- Login error: account in use
82 01

-- Login error: idle timeout
82 FE
```
