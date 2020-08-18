package huego

import (
	"encoding/json"
	"fmt"
)

// Entertainment represents a bridge entertainment group https://developers.meethue.com/develop/hue-entertainment/philips-hue-entertainment-api/
type EntertainmentGroup struct {
	Name       string                   `json:"name,omitempty"`
	Lights     []string                 `json:"lights,omitempty"`
	Type       string                   `json:"type,omitempty"`
	GroupState *EntertainmentGroupState `json:"state,omitempty"`
	Recycle    bool                     `json:"recycle,omitempty"`
	Class      string                   `json:"class,omitempty"`
	State      *State                   `json:"action,omitempty"`
	Stream     *Stream                  `json:"stream,omitempty"`
	Locations  map[int]*Locations       `json:"locations,omitempty,string"`
	ID         int                      `json:"-"`
	bridge     *Bridge
}

// EntertainmentGroupState defines the state on a group.
// Can be used to control the state of all lights in a group rather than controlling them individually
type EntertainmentGroupState struct {
	AllOn bool `json:"all_on,omitempty"`
	AnyOn bool `json:"any_on,omitempty"`
}

type Stream struct {
	ProxyMode string `json:"proxymode,omitempty"`
	ProxyNode string `json:"proxynode,omitempty"`
	Active    bool   `json:"active,omitempty"`
	Owner     string `json:"owner,omitempty"`
}

type Locations struct {
	X float64
	Y float64
	Z float64
}

func (n *Locations) UnmarshalJSON(buf []byte) error {
	tmp := []interface{}{&n.X, &n.Y, &n.Z}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}
	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields in Locations: %d != %d", g, e)
	}
	return nil
}
