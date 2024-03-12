# How To Write An `if` Statement in Go
### Notes on branchless programming

Issue #6011: can be found here.
[https://github.com/golang/go/issues/6011](https://github.com/golang/go/issues/6011)


> robpike commented on Aug 2, 2013
> Comment 1:
> 
> The compiler could provide speed without new syntax just by generating better code.
> Labels changed: added priority-later, performance, removed priority-triage.
> 
> Status changed to Accepted.


After reading this, I decided to see if it had been implemented. It has and proven to be useful. The following code represents the most efficient form.

```
func Bool2int(b bool) int {
	var i int
	if b {
		i = 1
	} else {
		i = 0
	}
	return i
}
```

I then tried the generic form and inspected the results.

```
func B2N[T int | uint | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | float32 | float64](b bool) T {
	var i int
	if b {
		i = 1
	} else {
		i = 0
	}
	return T(i)
}
```

## b2 build script
Builds to the code and create an object dump the functions B2N and Inspect.
```
#! /bin/bash
go build -gcflags='-l -v' b2.go
go tool objdump -S -s B2N b2 >b2n.txt
go tool objdump -S -s Inspect b2 >inspect.txt
```
```
go tool objdump -S -s B2N b2 >b2n.txt
```
The code [bool2.go](../b2i/bool2.go) and the output [b2n.txt](b2n.txt) which is empty.
```
go tool objdump -S -s Inspect b2 >inspect.txt
```
The code [inspect.go](../b2i/inspect.go) and the output [inspect.txt](inspect.txt).

## The Important Parts

No conditional branches!

```
bi = B2N[int](i > 0)
  0x4821f9		48833daf3f0a0000	CMPQ xray/b2i.i(SB), $0x0	
  0x482201		0f9fc1			    SETG CL				
	return T(i)
```

```
bi16 = B2N[int16](i16 > 0)
  0x48222c		66833d303f0a0000	CMPW xray/b2i.i16(SB), $0x0	
  0x482234		400f9fc6		    SETG SI				
	return T(i)
  0x482238		400fb6f6		    MOVZX SI, SI		
  0x48223c		6689742442		    MOVW SI, 0x42(SP)	

bi32 = B2N[int32](i32 > 0)
0x482241		833d203f0a0000	    CMPL xray/b2i.i32(SB), $0x0	
  0x482248		410f9fc0		    SETG R8				
	return T(i)
  0x48224c		450fb6c0		    MOVZX R8, R8		
  0x482250		4489442448		    MOVL R8, 0x48(SP)	

bi64 = B2N[int64](i64 > 0)
  0x482255		48833d5b3f0a0000    CMPQ xray/b2i.i64(SB), $0x0	
  0x48225d		410f9fc1		    SETG R9				
	return T(i)
  0x482261		450fb6c9		    MOVZX R9, R9		
  0x482265		4c894c2468		    MOVQ R9, 0x68(SP)	
```
