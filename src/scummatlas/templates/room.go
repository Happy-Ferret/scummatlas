package templates

import (
	"fmt"
	"html/template"
	_ "io/ioutil"
	"os"
	"scummatlas"
)

type RoomData struct {
	Index      int
	Title      string
	Background string
	Boxes      [][4]scummatlas.Point
	scummatlas.Room
}

const roomTpl = `
<html>
    <head>
        <title>Room {{.Index}}</title>
    </head>
    <body>
		<h1>{{.Title}}</h1>
		<h2>Background</h2>
		<img width="100%" src="{{.Background}}"/>
		<h2>Walking boxes</h2>
		<svg width="{{.Width}}" width="{{.Height}}">

		{{range .Boxes}}
		  <polygon points="
		  {{range .}}{{.X}},{{.Y}} {{end}}
		  " style="fill:#ccc;stroke:black;stroke-width:1" />
		{{end}}

		    Sorry, your browser does not support inline SVG.
			</svg>

		<h2>Scripts</h2>
		<h2>Objects</h2>
		<h2>Palette</h2>
    </body>
</html>`

func WriteRoom(room scummatlas.Room, index int, outputdir string) {

	t := template.Must(template.New("index").Parse(roomTpl))

	bgPath := fmt.Sprintf("./room%02d_bg.png", index)
	htmlPath := fmt.Sprintf("%v/room%02d.html", outputdir, index)
	file, err := os.Create(htmlPath)
	if err != nil {
		panic("Can't create room file")
	}

	var boxes [][4]scummatlas.Point

	for _, v := range room.Boxes {
		fmt.Println(v.Corners())
		boxes = append(boxes, v.Corners())
	}

	data := RoomData{
		index,
		"A room",
		bgPath,
		boxes,
		room,
	}
	t.Execute(file, data)
}
