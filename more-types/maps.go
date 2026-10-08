package main

type Vertexs struct {
	Lat  float64
	Long float64
}

var m = map[string]Vertexs{
	"Bell Labs": {
		40.68433, -74.39967,
	},
	"Google": {
		37.42202, -122.08408,
	},
}
