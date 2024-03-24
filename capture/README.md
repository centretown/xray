# Capture

A few helpful functions to capture whats on the screen. 

### CapturePNG
```go
CapturePNG(img image.Image)
```

Nothing much to see here.

### CaptureGIF

Rough but on the way.

```go
CaptureGIF(stop <-chan int, scr <-chan image.Image,
	colorMap map[color.Color]uint8, pal color.Palette)
```
Example usage. See [2d](../2d/2d.go)
```go
if pads.IsPadButtonDown(i, rl.GamepadButtonMiddleLeft) {
    capturing = true
    frameCount = 360
    go capture.CaptureGIF(stopChan, scrChan, colorMap, pal)
    return
}
    
```

### WriteGIF

```go
WriteGIF(pics []image.Image, colorsMap map[color.Color]uint8,
    pal color.Palette)
```
