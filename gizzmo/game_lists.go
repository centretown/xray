package gizzmo

import (
	"fmt"
	"slices"
)

// BuildLists for game operation
func (gs *Game) BuildLists() {
	gs.BuildDrawerList()
	gs.BuildTextureList()
	gs.Content.depthList = make([]DeepDrawer, len(gs.Content.drawerList))

	var (
		drawerList = gs.Content.drawerList
		depthList  = gs.Content.depthList
		depth      int32
	)

	// initialize depths
	for i := range drawerList {
		deepDr, hasDepth := drawerList[i].(HasDepth)
		if hasDepth {
			depth = deepDr.GetDepth()
		} else {
			depth = Deepest
		}
		depthList[i] = DeepDrawer{
			Drawer: drawerList[i],
			Depth:  depth}
	}

	fmt.Println("BUILDLISTS depthlist", len(gs.Content.depthList))

	gs.SortDepthList()
}

// SortDepthList all drawers plus depth sorted by depth (ascending)
func (gs *Game) SortDepthList() []DeepDrawer {
	for _, dl := range gs.Content.depthList {
		deepDr, hasDepth := dl.Drawer.(HasDepth)
		if hasDepth {
			dl.Depth = deepDr.GetDepth()
		} else {
			dl.Depth = Deepest
		}
	}
	slices.SortStableFunc(gs.Content.depthList, CompareDepths)
	return gs.Content.depthList
}

// BuildDrawerList of all drawers
func (gs *Game) BuildDrawerList() {
	list := make([]Drawer, 0)
	for _, mv := range gs.Content.movers {
		list = append(list, mv)
	}
	list = append(list, gs.Content.drawers...)
	gs.Content.drawerList = list
}

// BuildTextureList of all textures from all drawers
func (gs *Game) BuildTextureList() {
	list := gs.Content.textureList

	for _, mv := range gs.Content.movers {
		if t, ok := mv.GetDrawer().(*Texture); ok {
			list = append(list, t)
		}
	}

	for _, dr := range gs.Content.drawers {
		if t, ok := dr.(*Texture); ok {
			list = append(list, t)
		}
	}
	gs.Content.textureList = list
}
