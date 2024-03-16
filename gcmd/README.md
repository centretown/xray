# GCMD

## [gcmd.go](gcmd.go)
Command line flag parser and runner for testpad.

## [cmds.go](cmds.go)

Commands or tests to carry out.
```
func LastButtonPressed(cmd *GCmd)

func IsButtonUp(cmd *GCmd) 

func IsButtonDown(cmd *GCmd)

func IsButtonReleased(cmd *GCmd)

func IsButtonPressed(cmd *GCmd)

func GetAxisValues(cmd *GCmd)

func GetAxisMovement(cmd *GCmd)

func DumpPad(cmd *GCmd)

func TestKeys(cmd *GCmd)

func TestAxes(cmd *GCmd) 
```
