package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/amimof/huego"
)

const (
	ColorSpaceRGB = iota
	ColorSpaceXY
)
const (
	StreamType = iota
)

type Stream struct {
	ProtocolName string
	Version      [2]int
	SequenceID   int
	ColorSpace   int
	Lights       []Light
}

func toHex(i int) string {
	//return fmt.Sprintf("\\x%02x", i)
	return strconv.FormatInt(int64(i), 16)
}

func (s Stream) Serialize() []byte {
	b := []byte(s.ProtocolName)
	b = append(b, byte(s.Version[0]))
	b = append(b, byte(s.Version[1]))
	b = append(b, byte(s.SequenceID))
	b = append(b, byte(0))
	b = append(b, byte(0))
	b = append(b, byte(s.ColorSpace))
	b = append(b, byte(0))

	for _, l := range s.Lights {
		b = append(b, l.Serialize()...)
	}
	return b
}

type Light struct {
	Type int
	ID   int
	RGB  [3]int
}

func (l Light) Serialize() []byte {
	b := []byte{}
	b = append(b, byte(l.Type))
	b = append(b, byte(0))
	b = append(b, byte(l.ID))
	b = append(b, byte(l.RGB[0]))
	b = append(b, byte(0))
	b = append(b, byte(l.RGB[1]))
	b = append(b, byte(0))
	b = append(b, byte(l.RGB[2]))
	b = append(b, byte(0))

	return b
}

func main() {

	bridge, _ := huego.Discover()
	bridge = bridge.Login("h0vOCJQAkvarW9rEyN-omCBd4vGZ2HeDDo3Nv1GY", "B01FE23CE8CC3031DFF28D42B7448E4C") // Link button needs to be pressed
	//fmt.Println(user, clientkey)
	//if err != nil {
	//	fmt.Printf("Error creating user: %s", err.Error())
	//}
	//bridge = bridge.Login(user)
	eGroup, err := bridge.GetEntertainmentGroup(7)

	s, err := eGroup.StartStream()
	if err != nil {
		fmt.Println(err)
	}
	defer s.StopStream()
	l := make(map[int][]float32)

	l[11] = []float32{.215, .711, .5}
	l[17] = []float32{.7, .299, .5}
	s.Set(l)
	time.Sleep(30 * time.Second)
}
