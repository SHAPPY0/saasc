package widgets

import (
	"fmt"
)

var menuItemsLimit = 5

type Menu struct {
	*Flex
	Grid1		*MapView
	Grid2		*MapView
	Grid3		*MapView
	Items		[]Item
	Default 	[]Item
}

type Item struct {
	Name 		string
	Description string
	Icon 		string
}

var (
	EnterMenu = Item{
		Name:		"enter",
		Icon:		"enter",
		Description: "Select Row",
	}
	UpArrowMenu = Item{
		Name:		"up_arrow",
		Icon:		"↑",
		Description: "Move UP",
	}
	DownArrowMenu = Item{
		Name:		"down_arrow",
		Icon:		"↓",
		Description: "Move Down",
	}
	EscMenu = Item{
		Name:		"esc",
		Icon:		"esc",
		Description: "Go Back",
	}
	ContextMenu = Item{
		Name:		"region_namespace",
		Icon:		"2",
		Description: "Context",
	}
	CreateJobNMenu = Item{
		Name:		"create_new_jb",
		Icon:		"3",
		Description: "Create New Job",
	}
	PlanMenu = Item{
		Name:		"Plans",
		Icon:		"1",
		Description: "Plans",
	}
)

var DefaultMenus = []Item{
	EscMenu,
	UpArrowMenu,
	DownArrowMenu,
	EnterMenu,
}

var DefaultGlobalMenus = []Item{
	PlanMenu,
	ContextMenu,
	CreateJobNMenu,
}

func NewMenu() *Menu {
	m := &Menu{
		Flex:		NewFlex(),
		Grid1:		NewMapView(),
		Grid2:		NewMapView(),
		Grid3:		NewMapView(),
		Items:		make([]Item, 0),
		Default:	DefaultMenus,
	}
	m.Flex.AddItemX(m.Grid1, 0, 1, false)
	m.Flex.AddItemX(m.Grid2, 0, 1, false)
	m.Flex.AddItemX(m.Grid3, 0, 1, false)
	return m
}

func MenuExist(menu *Menu, name string) bool {
	found := false
	for _, Item := range menu.Items {
		if Item.Name == name {
			found = true
		}
	}
	return found
}

func (m *Menu) Add(menu Item, refresh bool) *Menu {
	if MenuExist(m, menu.Name) {
		return m
	}
	m.Items = append(m.Items, menu)
	if refresh {
		m.Render()
	}
	return m
}

func (m *Menu) Render() {
	// m.Grid1.Clear()
	m.Grid2.Clear()
	m.Grid3.Clear()
	for i, menu := range m.Items {
		key := fmt.Sprintf("[%s]<%s>", "orange", menu.Icon)
		value := fmt.Sprintf("[%s]%s\n", "DimGray", menu.Description)
		if i < menuItemsLimit {
			m.Grid2.SetMapKeyValue(key, value)
		} else if (i + 1) > menuItemsLimit && (i + 1) < (menuItemsLimit * 2) {
			m.Grid3.SetMapKeyValue(key, value)
		} 
		// else if (i + 1) > (menuItemsLimit * 2) && (i + 1) < (menuItemsLimit * 3) {
		// 	m.Grid3.SetMapKeyValue(key, value)
		// }
	}
	if m.Grid2.Size > 0 {
		m.Grid2.DrawMapView()
	}
	if m.Grid3.Size > 0 {
		m.Grid3.DrawMapView()
	}
	// if m.Grid3.Size > 0 {
	// 	m.Grid3.DrawMapView()
	// }
}

func (m *Menu) RenderGlobalMenus() {
	m.Grid1.Clear()
	for _, menu := range DefaultGlobalMenus {
		key := fmt.Sprintf("[%s]<%s>", "orange", menu.Icon)
		value := fmt.Sprintf("[%s]%s\n", "DimGray", menu.Description)
		m.Grid1.SetMapKeyValue(key, value)
	}
	m.Grid1.DrawMapView()
}

func (m *Menu) RenderMenu(menus []Item, addDefaults bool) {
	var allMenus = make([]Item, 0)
	if addDefaults {
		allMenus = append(allMenus, DefaultMenus...)
	}
	allMenus = append(allMenus, menus...)
	for _, menu :=  range allMenus {
		m.Add(menu, false)
	}
	m.Render()
}

func (m *Menu) Remove(menu Item) {
	for I, Menu := range m.Items {
		if Menu.Name == menu.Name {
			m.Items = append(m.Items[:I], m.Items[I + 1:]...)
		}
	}
	m.Render()
}

func (m *Menu) RemoveMenus(menus []Item) {
	for _, menu := range menus {
		m.Remove(menu)
	}
}

func (m *Menu) Replace(item1 Item, item2 Item) {
	for i, menu := range m.Items {
		if menu.Name == item1.Name {
			m.Items[i] = item2
			m.Render()
			break
		}
	}
}