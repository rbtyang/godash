package dashnet_test

import (
	"testing"

	"github.com/rbtyang/godash/dashnet"
	"github.com/stretchr/testify/assert"
)

/*
TestIp2binary is a ...
*/
func TestIp2binary(t *testing.T) {
	{
		want := "11000000101010000011100000000100"
		recv := dashnet.Ip2binary("192.168.56.4")
		assert.Equal(t, want, recv)
	}
}

/*
TestMatchIP is a ...
*/
func TestMatchIP(t *testing.T) {
	{
		want := true
		recv := dashnet.MatchIp("192.168.56.4", "192.168.56.4")
		assert.Equal(t, want, recv)
	}
	{
		want := false
		recv := dashnet.MatchIp("192.168.56.4", "192.168.56.5")
		assert.Equal(t, want, recv)
	}

	{
		want := false
		recv := dashnet.MatchIp("192.168.56.4", "192.168.56.64/26")
		assert.Equal(t, want, recv)
	}
	{
		want := false
		recv := dashnet.MatchIp("192.168.56.63", "192.168.56.64/26")
		assert.Equal(t, want, recv)
	}
	{
		want := true
		recv := dashnet.MatchIp("192.168.56.64", "192.168.56.64/26")
		assert.Equal(t, want, recv)
	}
	{
		want := true
		recv := dashnet.MatchIp("192.168.56.65", "192.168.56.64/26")
		assert.Equal(t, want, recv)
	}
	{
		want := true
		recv := dashnet.MatchIp("192.168.56.100", "192.168.56.64/26")
		assert.Equal(t, want, recv)
	}
	{
		want := true
		recv := dashnet.MatchIp("192.168.56.126", "192.168.56.64/26")
		assert.Equal(t, want, recv)
	}
	{
		want := true
		recv := dashnet.MatchIp("192.168.56.127", "192.168.56.64/26")
		assert.Equal(t, want, recv)
	}
	{
		want := false
		recv := dashnet.MatchIp("192.168.56.128", "192.168.56.64/26")
		assert.Equal(t, want, recv)
	}
}
