//go:build integration

package integration

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"time"
)

type uoClient struct {
	conn net.Conn
}

func newUOClient(conn net.Conn) *uoClient {
	return &uoClient{conn: conn}
}

func (c *uoClient) SetDeadline(d time.Duration) {
	_ = c.conn.SetDeadline(time.Now().Add(d))
}

func writeFixed(buf []byte, offset int, s string, length int) {
	for i := 0; i < length; i++ {
		if i < len(s) {
			buf[offset+i] = s[i]
		} else {
			buf[offset+i] = 0
		}
	}
}

func writeUInt32BE(buf []byte, offset int, v uint32) {
	binary.BigEndian.PutUint32(buf[offset:], v)
}

func writeUInt16BE(buf []byte, offset int, v uint16) {
	binary.BigEndian.PutUint16(buf[offset:], v)
}

// SendLoginRequest sends the 0x80 login request (62 bytes).
// An empty password (with username <= 19 chars) triggers the clean/no-encryption path.
func (c *uoClient) SendLoginRequest(user, pass string) error {
	buf := make([]byte, 62)
	buf[0] = 0x80
	writeFixed(buf, 1, user, 30)
	writeFixed(buf, 31, pass, 30)
	buf[61] = 0x00
	_, err := c.conn.Write(buf)
	return err
}

// ReadServerList reads the 0xA8 server list response and returns the server count.
func (c *uoClient) ReadServerList() (int, error) {
	cmd := make([]byte, 1)
	if _, err := io.ReadFull(c.conn, cmd); err != nil {
		return 0, err
	}
	if cmd[0] != 0xA8 {
		return 0, fmt.Errorf("expected 0xA8, got 0x%02X", cmd[0])
	}
	lenBuf := make([]byte, 2)
	if _, err := io.ReadFull(c.conn, lenBuf); err != nil {
		return 0, err
	}
	length := int(binary.BigEndian.Uint16(lenBuf))
	if length < 3 {
		return 0, fmt.Errorf("malformed 0xA8 response: length %d too small", length)
	}
	// remaining: flags(1) + count(2) + server data
	rest := make([]byte, length-3)
	if _, err := io.ReadFull(c.conn, rest); err != nil {
		return 0, err
	}
	count := int(binary.BigEndian.Uint16(rest[1:3]))
	return count, nil
}

// SendShardSelected sends the 0xA0 server selection (3 bytes).
func (c *uoClient) SendShardSelected(idx int) error {
	buf := make([]byte, 3)
	buf[0] = 0xA0
	writeUInt16BE(buf, 1, uint16(idx))
	_, err := c.conn.Write(buf)
	return err
}

// ReadRedirectToShard reads the 0x8C redirect packet (11 bytes) and returns the encryption key.
func (c *uoClient) ReadRedirectToShard() (uint32, error) {
	// 0x8C(1) + IP(4) + port(2) + key(4) = 11 bytes
	buf := make([]byte, 11)
	if _, err := io.ReadFull(c.conn, buf); err != nil {
		return 0, err
	}
	if buf[0] != 0x8C {
		return 0, fmt.Errorf("expected 0x8C, got 0x%02X", buf[0])
	}
	key := binary.BigEndian.Uint32(buf[7:11])
	return key, nil
}

// SendGameServerLogin sends the 0x91 game server login (65 bytes).
func (c *uoClient) SendGameServerLogin(key uint32, user, pass string) error {
	buf := make([]byte, 65)
	buf[0] = 0x91
	writeUInt32BE(buf, 1, key)
	writeFixed(buf, 5, user, 30)
	writeFixed(buf, 35, pass, 30)
	_, err := c.conn.Write(buf)
	return err
}

// ReadClientFeatures reads the 0xB9 client features packet (5 bytes).
func (c *uoClient) ReadClientFeatures() error {
	buf := make([]byte, 5)
	if _, err := io.ReadFull(c.conn, buf); err != nil {
		return err
	}
	if buf[0] != 0xB9 {
		return fmt.Errorf("expected 0xB9, got 0x%02X", buf[0])
	}
	return nil
}

// ReadCharacterList reads the 0xA9 character list packet and returns the character count.
func (c *uoClient) ReadCharacterList() (int, error) {
	cmd := make([]byte, 1)
	if _, err := io.ReadFull(c.conn, cmd); err != nil {
		return 0, err
	}
	if cmd[0] != 0xA9 {
		return 0, fmt.Errorf("expected 0xA9, got 0x%02X", cmd[0])
	}
	lenBuf := make([]byte, 2)
	if _, err := io.ReadFull(c.conn, lenBuf); err != nil {
		return 0, err
	}
	length := int(binary.BigEndian.Uint16(lenBuf))
	countBuf := make([]byte, 1)
	if _, err := io.ReadFull(c.conn, countBuf); err != nil {
		return 0, err
	}
	count := int(countBuf[0])
	remaining := make([]byte, length-4)
	if _, err := io.ReadFull(c.conn, remaining); err != nil {
		return 0, err
	}
	return count, nil
}

// SendPreLogin sends the 0x5D pre-login packet (73 bytes).
func (c *uoClient) SendPreLogin(name, pass string, slot, key uint32) error {
	buf := make([]byte, 73)
	buf[0] = 0x5D
	writeUInt32BE(buf, 1, 0) // ignored uint32
	writeFixed(buf, 5, name, 30)
	writeFixed(buf, 35, pass, 30)
	writeUInt32BE(buf, 65, slot)
	writeUInt32BE(buf, 69, key)
	_, err := c.conn.Write(buf)
	return err
}

// ReadPlayerStartConfirmation reads the 0x1B player start confirmation (37 bytes).
func (c *uoClient) ReadPlayerStartConfirmation() error {
	buf := make([]byte, 37)
	if _, err := io.ReadFull(c.conn, buf); err != nil {
		return err
	}
	if buf[0] != 0x1B {
		return fmt.Errorf("expected 0x1B, got 0x%02X", buf[0])
	}
	return nil
}

// ReadLoginComplete reads the 0x55 login complete packet (1 byte).
func (c *uoClient) ReadLoginComplete() error {
	buf := make([]byte, 1)
	if _, err := io.ReadFull(c.conn, buf); err != nil {
		return err
	}
	if buf[0] != 0x55 {
		return fmt.Errorf("expected 0x55, got 0x%02X", buf[0])
	}
	return nil
}
