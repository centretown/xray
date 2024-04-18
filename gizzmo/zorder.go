package gizzmo

const Deepest int32 = 10000

type DeepDrawer struct {
	Depth  int32
	Drawer Drawer
}

func CompareDepths(zo, other DeepDrawer) int {
	if zo.Depth == other.Depth {
		return 0
	}
	if zo.Depth < other.Depth {
		return -1
	}
	return 1
}
