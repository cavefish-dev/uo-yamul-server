//go:build integration

package integration

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	services "yamul-gateway/backend/services"
)

func TestFullLoginFlow(t *testing.T) {
	h := newGatewayHarness(t)
	h.backend.Login.AddCredential("testuser", "")
	c := h.client
	c.SetDeadline(10 * time.Second)

	// Login phase
	require.NoError(t, c.SendLoginRequest("testuser", ""))
	count, err := c.ReadServerList()
	require.NoError(t, err)
	assert.Equal(t, 1, count)

	// Server select → redirect
	require.NoError(t, c.SendShardSelected(1))
	encKey, err := c.ReadRedirectToShard()
	require.NoError(t, err)

	// Game server login → char list
	require.NoError(t, c.SendGameServerLogin(encKey, "testuser", ""))
	require.NoError(t, c.ReadClientFeatures())
	charCount, err := c.ReadCharacterList()
	require.NoError(t, err)
	assert.GreaterOrEqual(t, charCount, 1)

	// Character selection
	require.NoError(t, c.SendPreLogin("testchar", "", 0, encKey))

	// Mock backend receives character selection message
	pkg := <-h.backend.Game.Received
	assert.Equal(t, services.MsgType_TypeCharacterSelection, pkg.Type)

	// Backend sends game start
	h.backend.Game.ToSend <- &services.StreamPackage{
		Type: services.MsgType_TypePlayerStartConfirmation,
		Body: &services.Message{Msg: &services.Message_PlayerStartConfirmation{
			PlayerStartConfirmation: &services.MsgPlayerStartConfirmation{
				Id:          &services.ObjectId{Value: 1},
				Coordinates: &services.Coordinate{XLoc: 100, YLoc: 200},
				Direction:   services.ObjectDirection_north,
				GraphicId:   400,
			},
		}},
	}
	require.NoError(t, c.ReadPlayerStartConfirmation())

	// Backend sends login complete
	h.backend.Game.ToSend <- &services.StreamPackage{
		Type: services.MsgType_TypeLoginComplete,
		Body: &services.Message{},
	}
	require.NoError(t, c.ReadLoginComplete())
}
