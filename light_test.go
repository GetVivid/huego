package huego

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLights(t *testing.T) {
	b := New(hostname, username, clientKey)
	lights, err := b.GetLights()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Found %d lights", len(lights))
	for _, l := range lights {
		t.Logf("ID: %d", l.ID)
		t.Logf("  Type: %s", l.Type)
		t.Logf("  Name: %s", l.Name)
		t.Logf("  ModelID: %s", l.ModelID)
		t.Logf("  ManufacturerName: %s", l.ManufacturerName)
		t.Logf("  UniqueID: %s", l.UniqueID)
		t.Logf("  SwVersion: %s", l.SwVersion)
		t.Logf("  SwConfigID: %s", l.SwConfigID)
		t.Logf("  ProductID: %s", l.ProductID)
	}
	contains := func(name string, ss []Light) bool {
		for _, s := range ss {
			if s.Name == name {
				return true
			}
		}
		return false
	}

	assert.True(t, contains("Huecolorlamp7", lights))
	assert.True(t, contains("Huelightstripplus1", lights))

	b.Host = badHostname
	_, err = b.GetLights()
	assert.NotNil(t, err)
}

func TestGetLight(t *testing.T) {
	b := New(hostname, username, clientKey)
	l, err := b.GetLight(1)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("ID: %d", l.ID)
		t.Logf("Type: %s", l.Type)
		t.Logf("Name: %s", l.Name)
		t.Logf("ModelID: %s", l.ModelID)
		t.Logf("ManufacturerName: %s", l.ManufacturerName)
		t.Logf("UniqueID: %s", l.UniqueID)
		t.Logf("SwVersion: %s", l.SwVersion)
		t.Logf("SwConfigID: %s", l.SwConfigID)
		t.Logf("ProductID: %s", l.ProductID)
		t.Logf("State:")
		t.Logf("  On: %t", l.State.On)
		t.Logf("  Bri: %d", l.State.Bri)
		t.Logf("  Hue: %d", l.State.Hue)
		t.Logf("  Sat: %d", l.State.Sat)
		t.Logf("  Xy: %b", l.State.Xy)
		t.Logf("  Ct: %d", l.State.Ct)
		t.Logf("  Alert: %s", l.State.Alert)
		t.Logf("  Effect: %s", l.State.Effect)
		t.Logf("  TransitionTime: %d", l.State.TransitionTime)
		t.Logf("  BriInc: %d", l.State.BriInc)
		t.Logf("  SatInc: %d", l.State.SatInc)
		t.Logf("  HueInc: %d", l.State.HueInc)
		t.Logf("  CtInc: %d", l.State.CtInc)
		t.Logf("  XyInc: %d", l.State.XyInc)
		t.Logf("  ColorMode: %s", l.State.ColorMode)
		t.Logf("  Reachable: %t", l.State.Reachable)
	}

	b.Host = badHostname
	_, err = b.GetLight(1)
	assert.NotNil(t, err)

}

func TestSetLight(t *testing.T) {
	b := New(hostname, username, clientKey)
	id := 1
	state := State{
		On:  true,
		Bri: 254,
	}
	resp, err := b.SetLightState(id, state)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("Light %d state updated", id)
		for k, v := range resp.Success {
			t.Logf("%v: %s", k, v)
		}
	}

	b.Host = badHostname
	_, err = b.SetLightState(id, state)
	assert.NotNil(t, err)
}

func TestFindLights(t *testing.T) {
	b := New(hostname, username, clientKey)
	resp, err := b.FindLights()
	if err != nil {
		t.Fatal(err)
	} else {
		for k, v := range resp.Success {
			t.Logf("%v: %s", k, v)
		}
	}

	b.Host = badHostname
	_, err = b.FindLights()
	assert.NotNil(t, err)
}

func TestUpdateLight(t *testing.T) {
	b := New(hostname, username, clientKey)
	id := 1
	light := Light{
		Name: "New Light",
	}
	resp, err := b.UpdateLight(id, light)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("Light %d updated", id)
		for k, v := range resp.Success {
			t.Logf("%v: %s", k, v)
		}
	}
	b.Host = badHostname
	_, err = b.UpdateLight(id, light)
	assert.NotNil(t, err)
}

func TestTurnOffLight(t *testing.T) {
	b := New(hostname, username, clientKey)
	id := 1
	light, err := b.GetLight(id)
	if err != nil {
		t.Fatal(err)
	}
	err = light.Off()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Turned off light with id %d", light.ID)

	b.Host = badHostname
	err = light.Off()
	assert.NotNil(t, err)
}

func TestTurnOnLight(t *testing.T) {
	b := New(hostname, username, clientKey)
	id := 1
	light, err := b.GetLight(id)
	if err != nil {
		t.Fatal(err)
	}
	err = light.On()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Turned on light with id %d", light.ID)

	b.Host = badHostname
	err = light.On()
	assert.NotNil(t, err)
}

func TestIfLightIsOn(t *testing.T) {
	b := New(hostname, username, clientKey)
	id := 1
	light, err := b.GetLight(id)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Is light %d on?: %t", light.ID, light.IsOn())
}

func TestRenameLight(t *testing.T) {
	b := New(hostname, username, clientKey)
	id := 1
	light, err := b.GetLight(id)
	if err != nil {
		t.Fatal(err)
	}
	name := "Color Lamp 1"
	err = light.Rename(name)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Renamed light to '%s'", light.Name)

	b.Host = badHostname
	err = light.Rename(name)
	assert.NotNil(t, err)
}

func TestSetLightBri(t *testing.T) {
	b := New(hostname, username, clientKey)
	id := 1
	light, err := b.GetLight(id)
	if err != nil {
		t.Fatal(err)
	}
	err = light.Bri(254)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Brightness of light %d set to %d", light.ID, light.State.Bri)

	b.Host = badHostname
	err = light.Bri(254)
	assert.NotNil(t, err)
}

func TestSetLightHue(t *testing.T) {
	b := New(hostname, username, clientKey)
	id := 1
	light, err := b.GetLight(id)
	if err != nil {
		t.Fatal(err)
	}
	err = light.Hue(65535)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Hue of light %d set to %d", light.ID, light.State.Hue)

	b.Host = badHostname
	err = light.Hue(65535)
	assert.NotNil(t, err)
}

func TestSetLightSat(t *testing.T) {
	b := New(hostname, username, clientKey)
	id := 1
	light, err := b.GetLight(id)
	if err != nil {
		t.Fatal(err)
	}
	err = light.Sat(254)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Sat of light %d set to %d", light.ID, light.State.Sat)

	b.Host = badHostname
	err = light.Sat(254)
	assert.NotNil(t, err)
}

func TestSetLightXy(t *testing.T) {
	b := New(hostname, username, clientKey)
	id := 1
	light, err := b.GetLight(id)
	if err != nil {
		t.Fatal(err)
	}
	xy := []float32{0.1, 0.5}
	err = light.Xy(xy)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Xy of light %d set to %+v", light.ID, light.State.Xy)

	b.Host = badHostname
	err = light.Xy(xy)
	assert.NotNil(t, err)
}

func TestSetLightCt(t *testing.T) {
	b := New(hostname, username, clientKey)
	id := 1
	light, err := b.GetLight(id)
	if err != nil {
		t.Fatal(err)
	}
	err = light.Ct(16)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Ct of light %d set to %d", light.ID, light.State.Ct)

	b.Host = badHostname
	err = light.Ct(16)
	assert.NotNil(t, err)
}

func TestSetLightTransitionTime(t *testing.T) {
	b := New(hostname, username, clientKey)
	id := 1
	light, err := b.GetLight(id)
	if err != nil {
		t.Fatal(err)
	}
	err = light.TransitionTime(10)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("TransitionTime of light %d set to %d", light.ID, light.State.TransitionTime)

	b.Host = badHostname
	err = light.TransitionTime(10)
	assert.NotNil(t, err)
}

func TestSetLightEffect(t *testing.T) {
	b := New(hostname, username, clientKey)
	id := 1
	light, err := b.GetLight(id)
	if err != nil {
		t.Fatal(err)
	}
	err = light.Effect("colorloop")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Effect of light %d set to %s", light.ID, light.State.Effect)

	b.Host = badHostname
	err = light.Effect("colorloop")
	assert.NotNil(t, err)
}

func TestSetLightAlert(t *testing.T) {
	b := New(hostname, username, clientKey)
	id := 1
	light, err := b.GetLight(id)
	if err != nil {
		t.Fatal(err)
	}
	err = light.Alert("lselect")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Alert of light %d set to %s", light.ID, light.State.Alert)

	b.Host = badHostname
	err = light.Alert("lselect")
	assert.NotNil(t, err)
}

func TestSetStateLight(t *testing.T) {
	b := New(hostname, username, clientKey)
	id := 1
	light, err := b.GetLight(id)
	if err != nil {
		t.Fatal(err)
	}
	state := State{
		On:  true,
		Bri: 254,
	}
	err = light.SetState(state)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("State set successfully on light %d", id)

	b.Host = badHostname
	err = light.SetState(state)
	assert.NotNil(t, err)
}
