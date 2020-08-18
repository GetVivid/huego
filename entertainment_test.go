package huego

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEntertainmentGroups(t *testing.T) {
	b := New(hostname, username)
	groups, err := b.GetEntertainmentGroups()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Found %d groups", len(groups))
	for i, g := range groups {
		t.Logf("%d:", i)
		t.Logf("  Name: %s", g.Name)
		t.Logf("  Lights: %s", g.Lights)
		t.Logf("  Type: %s", g.Type)
		t.Logf("  GroupState:")
		t.Logf("    AllOn: %t", g.GroupState.AllOn)
		t.Logf("    AnyOn: %t", g.GroupState.AnyOn)
		t.Logf("  Recycle: %t", g.Recycle)
		t.Logf("  Class: %s", g.Class)
		t.Logf("  State:")
		t.Logf("    On: %t", g.State.On)
		t.Logf("    Bri: %d", g.State.Bri)
		t.Logf("    Hue: %d", g.State.Hue)
		t.Logf("    Sat: %d", g.State.Sat)
		t.Logf("    Xy: %b", g.State.Xy)
		t.Logf("    Ct: %d", g.State.Ct)
		t.Logf("    Alert: %s", g.State.Alert)
		t.Logf("    Effect: %s", g.State.Effect)
		t.Logf("    TransitionTime: %d", g.State.TransitionTime)
		t.Logf("    BriInc: %d", g.State.BriInc)
		t.Logf("    SatInc: %d", g.State.SatInc)
		t.Logf("    HueInc: %d", g.State.HueInc)
		t.Logf("    CtInc: %d", g.State.CtInc)
		t.Logf("    XyInc: %d", g.State.XyInc)
		t.Logf("    ColorMode: %s", g.State.ColorMode)
		t.Logf("    Reachable: %t", g.State.Reachable)
		t.Logf("  Stream:")
		t.Logf("    ProxyMode: %s", g.Stream.ProxyMode)
		t.Logf("    ProxyNode: %s", g.Stream.ProxyNode)
		t.Logf("    Active: %t", g.Stream.Active)
		t.Logf("    Owner: %s", g.Stream.Owner)
		t.Logf("  Location:")
		for lid, l := range g.Locations {
			t.Logf("    %d:", lid)
			t.Logf("      X: %.2f", l.X)
			t.Logf("      Y: %.2f", l.Y)
			t.Logf("      Z: %.2f", l.Z)

		}
		t.Logf("  ID: %d", g.ID)
	}

	contains := func(name string, ss []EntertainmentGroup) bool {
		for _, s := range ss {
			if s.Name == name {
				return true
			}
		}
		return false
	}

	assert.True(t, contains("Group 3", groups))
	b.Host = badHostname
	_, err = b.GetEntertainmentGroups()
	assert.NotNil(t, err)
}

func TestGetEntertainmentGroup(t *testing.T) {
	b := New(hostname, username)
	g, err := b.GetEntertainmentGroup(3)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Name: %s", g.Name)
	t.Logf("Lights: %s", g.Lights)
	t.Logf("Type: %s", g.Type)
	t.Logf("GroupState:")
	t.Logf("  AllOn: %t", g.GroupState.AllOn)
	t.Logf("  AnyOn: %t", g.GroupState.AnyOn)
	t.Logf("Recycle: %t", g.Recycle)
	t.Logf("Class: %s", g.Class)
	t.Logf("State:")
	t.Logf("  On: %t", g.State.On)
	t.Logf("  Bri: %d", g.State.Bri)
	t.Logf("  Hue: %d", g.State.Hue)
	t.Logf("  Sat: %d", g.State.Sat)
	t.Logf("  Xy: %b", g.State.Xy)
	t.Logf("  Ct: %d", g.State.Ct)
	t.Logf("  Alert: %s", g.State.Alert)
	t.Logf("  Effect: %s", g.State.Effect)
	t.Logf("  TransitionTime: %d", g.State.TransitionTime)
	t.Logf("  BriInc: %d", g.State.BriInc)
	t.Logf("  SatInc: %d", g.State.SatInc)
	t.Logf("  HueInc: %d", g.State.HueInc)
	t.Logf("  CtInc: %d", g.State.CtInc)
	t.Logf("  XyInc: %d", g.State.XyInc)
	t.Logf("  ColorMode: %s", g.State.ColorMode)
	t.Logf("  Reachable: %t", g.State.Reachable)
	t.Logf("  Stream:")
	t.Logf("    ProxyMode: %s", g.Stream.ProxyMode)
	t.Logf("    ProxyNode: %s", g.Stream.ProxyNode)
	t.Logf("    Active: %t", g.Stream.Active)
	t.Logf("    Owner: %s", g.Stream.Owner)
	t.Logf("  Location:")
	for lid, l := range g.Locations {
		t.Logf("    %d:", lid)
		t.Logf("      X: %.2f", l.X)
		t.Logf("      Y: %.2f", l.Y)
		t.Logf("      Z: %.2f", l.Z)

	}
	t.Logf("ID: %d", g.ID)

	b.Host = badHostname
	_, err = b.GetEntertainmentGroup(3)
	assert.NotNil(t, err)
}
