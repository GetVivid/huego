package huego

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net"
	"net/url"
	"strconv"

	"github.com/pion/dtls/v2"
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
	Locations  map[int]*Location        `json:"locations,omitempty,string"`
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

type Location struct {
	X float64
	Y float64
	Z float64
}

type EntertainmentStream struct {
	sequenceId int
	conn       *dtls.Conn
}

func (n *Location) UnmarshalJSON(buf []byte) error {
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

// Rename sets the name property of the group
func (g *EntertainmentGroup) Rename(new string) error {
	return g.RenameContext(context.Background(), new)
}

// RenameContext sets the name property of the group
func (g *EntertainmentGroup) RenameContext(ctx context.Context, new string) error {
	update := EntertainmentGroup{Name: new}
	_, err := g.bridge.UpdateEntertainmentGroupContext(ctx, g.ID, update)
	if err != nil {
		return err
	}
	g.Name = new
	return nil
}

func (g *EntertainmentGroup) StartStream() (*EntertainmentStream, error) {
	return g.StartStreamContext(context.Background())
}

func (g *EntertainmentGroup) StartStreamContext(ctx context.Context) (*EntertainmentStream, error) {
	_, err := g.bridge.StartEntertainmentGroupContext(ctx, g.ID)
	if err != nil {
		return nil, err
	}

	config := &dtls.Config{
		PSK: func(hint []byte) ([]byte, error) {
			return hex.DecodeString("B01FE23CE8CC3031DFF28D42B7448E4C")
		},
		PSKIdentityHint: []byte("h0vOCJQAkvarW9rEyN-omCBd4vGZ2HeDDo3Nv1GY"),
		CipherSuites:    []dtls.CipherSuiteID{dtls.TLS_PSK_WITH_AES_128_GCM_SHA256},
	}

	// Connect to a DTLS server
	u, err := url.Parse(g.bridge.Host)
	if err != nil {
		return nil, err
	}
	h, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", u.Host, 2100))
	if err != nil {
		return nil, err
	}
	conn, err := dtls.DialWithContext(ctx, "udp", h, config)
	if err != nil {
		return nil, err
	}

	stream := EntertainmentStream{
		sequenceId: 0,
		conn:       conn,
	}
	return &stream, nil
}

// Set takes a map of the light ID, x value, y value, and brightness
// EX: []float32{.7, .299, .5}
func (g *EntertainmentStream) Set(lights map[int][]float32) {
	var msg []byte
	proto := []byte("HueStream")
	header := []byte{1, 0, 0, 0, 0, 1, 0}

	msg = append(proto, header...)
	for id, color := range lights {

		x := floatToXY(color[0])
		y := floatToXY(color[1])
		b := floatToXY(color[2])

		fmt.Println(x, y)
		msg = append(msg, []byte{0, 0, byte(id), x[0], x[1], y[0], y[1], b[0], b[1]}...)
	}

	g.conn.Write(msg)
}

// floatToXY converts the floats into hex between 0x0000 (0.0000) and 0xffff(1.0000)
// per hue api XY+Brightness with 16 bit resolution per element so that means
// we need to break the hex into 2 seperate parts
func floatToXY(num float32) [2]byte {
	b := fmt.Sprintf("%016s", strconv.FormatInt(int64(num*65535), 2))
	first, _ := strconv.ParseInt(b[0:8], 2, 32)
	second, _ := strconv.ParseInt(b[8:16], 2, 32)
	return [2]byte{byte(first), byte(second)}
}

func (s *EntertainmentStream) StopStream() {
	s.StopStreamContext()
}

func (s *EntertainmentStream) StopStreamContext() {
	s.conn.Close()
}
