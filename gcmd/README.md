# GCMD

## [gcmd.go](gcmd.go)
Command line flag parser and runner for testpad.
```
go run ./testpad -h
Usage of /tmp/go-build3931522921/b001/exe/testpad:
  -a value
        same as axis
  -axis value
        one or more axes to test
        eg: -c down -c up -c last -a 0 -a 1 -a 3 [runs:  down axis 0, up axis 1, last axis 3]
  -b value
        same as button
  -button value
        one or more buttons to test
        eg: -c down -c up -c last -b 2 -b 5 [runs:  down button 2, up button 5, last button 5]
  -c value
        same as command
  -command value
        one or more commands eg: -c down -c up -c last 
                last - indicate last button pressed
                  up - indicate if selected button is up
                down - indicate if selected button is down
               press - indicate if selected button has been pressed
             release - indicate if selected button has been released
                move - indicate selected axis movement
                keys - indicate any key pressed
                axes - indicate any axis changes
                dump - dump maps and value corrections
  -d value
        same as duration
  -duration value
        one or more durations in seconds
        eg: -c down -c up -d 5 -d 6 [runs:  up 5s, down 6s]
  -j value
        same as joystick
  -joystick value
        one or more joysticks to test
        eg: -c down -c up -j 0 -j 1 [runs:  up joystick 0, down joystick 1]
  -k value
        same as keys
  -keys value
        test all keys
        eg: -c keys [runs: keys]
```
## [cmds.go](cmds.go)

Commands or tests to carry out.
### last
```
func LastButtonPressed(cmd *GCmd)
```
### up
```
func IsButtonUp(cmd *GCmd) 
```
### down
```
func IsButtonDown(cmd *GCmd)
```
### release
```
func IsButtonReleased(cmd *GCmd)
```
### press
```
func IsButtonPressed(cmd *GCmd)
```
### move
```
func GetAxisValues(cmd *GCmd)

func GetAxisMovement(cmd *GCmd)
```
### dump
```
func DumpPad(cmd *GCmd)
```
### keys
```
func TestKeys(cmd *GCmd)
```
### axes
```
func TestAxes(cmd *GCmd) 
```
