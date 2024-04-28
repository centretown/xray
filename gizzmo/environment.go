package gizzmo

type Monitor struct {
	Num         int
	Width       int
	Height      int
	RefreshRate int
}
type Screen struct {
	Width  int64
	Height int64
}

type Environment struct {
	Monitor          Monitor
	Screen           Screen
	CurrentFrameRate int64
}
