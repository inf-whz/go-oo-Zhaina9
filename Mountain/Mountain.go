package Mountain

import (
	"github.com/geobe/go4j/loc"
)

type Mountain struct {
	loc.Location
	name string
	altitude int
}

func (m Mountain) Name() string {
	return m.name
}

func (m Mountain) Altitude() int {
	return m.altitude
}

var mountain_List = []Mountain{
	{Location: loc.New(15, 15), name: "Alpen", altitude: 15781},
	{Location: loc.New(15, 15), name: "Himalaya", altitude: 29029},
}

var mountain_name = mountain_List[0]
