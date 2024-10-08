package widgets

import (
	"fmt"
)

type Container struct {
	*Flex
	Title 		string
	Rows 		int
	Col			int
	Grids		[]*MapView
}

func NewContainer(title string, rows, col int) *Container {
	c := Container{
		Flex:		NewFlex(),
		Title:		title,
		Rows:		rows,
		Col:		col,
		Grids:		make([]*MapView, 0),
	}
	c.SetTitle(fmt.Sprintf(" %s ", title))
	for i := 0; i < col; i++ {
		c.Grids = append(c.Grids, NewMapView())
		c.AddItemX(c.Grids[i], 0, 1, false)
	}
	c.Border(true)
	return &c
}

func (c *Container) SetKeyValue(key, value string) {
	for i := 0; i < c.Col; i++ {
		added := false
		if (c.Rows == 0 || c.Grids[i].GetRowCount() < c.Rows) {
			c.Grids[i].SetMapKeyValue(key, value)
			added = true
		}
		if added {
			c.Grids[i].DrawMapView()
			break
		}
	}
}

func (c *Container) Clear() {
	for i := 0; i < len(c.Grids); i++ {
		c.Grids[i].Clear()
	}
}