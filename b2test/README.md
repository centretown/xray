## Branchless Programming in Go

Pipelining and branch prediction are mainstays of cpu design. In order to take advantage of this a Go program needs:
- compiler supported fast bool to numeric type conversions
- compiler supported function inlining

The Go compiler supports both of these requirements.
#### Issue #6011: Or How To Write An "if" Statement in Go

> by sjbogdan: on Aug 1, 2013

> Need a compiler level support for fast bool to numeric types ( int, 
> byte, float ) conversion.

> robpike commented on Aug 2, 2013

> Comment 1:
> 
> The compiler could provide speed without new syntax just by generating better code.
> Labels changed: added priority-later, performance, removed priority-triage.
> 
> Status changed to Accepted.

Issue #6011: can be found here.
[https://github.com/golang/go/issues/6011](https://github.com/golang/go/issues/6011)


After reading this, I decided to examine the results and see if it could be useful. The following code represents the most efficient form.

```
func ToInt(b bool) int {
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
func To[T int | uint | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | float32 | float64](b bool) T {
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
Builds and creates object dumps for the functions To and Inspect.
```
#! /bin/bash
go build -gcflags='-l -v' b2.go
go tool objdump -S -s To b2 >To.txt
go tool objdump -S -s Inspect b2 >inspect.txt
```
```
go tool objdump -S -s To b2 >To.txt
```
The code [bool2.go](../b2i/bool2.go) and the output [To.txt](To.txt) which is empty.
```
go tool objdump -S -s Inspect b2 >inspect.txt
```
The code [inspect.go](../b2i/inspect.go) and the output [inspect.txt](inspect.txt).

## The Important Parts

No conditional branches!

#### signed integers

```
bi = To[int](i > 0)
  0x4821f9		48833daf3f0a0000	CMPQ xray/b2i.i(SB), $0x0	
  0x482201		0f9fc1			    SETG CL				
	return T(i)
  0x482204		0fb6c9				MOVZX CL, CX		
  0x482207		48894c2450			MOVQ CX, 0x50(SP)	
```

```
bi16 = To[int16](i16 > 0)
  0x48222c		66833d303f0a0000	CMPW xray/b2i.i16(SB), $0x0	
  0x482234		400f9fc6		    SETG SI				
	return T(i)
  0x482238		400fb6f6		    MOVZX SI, SI		
  0x48223c		6689742442		    MOVW SI, 0x42(SP)	

bi32 = To[int32](i32 > 0)
0x482241		833d203f0a0000	    CMPL xray/b2i.i32(SB), $0x0	
  0x482248		410f9fc0		    SETG R8				
	return T(i)
  0x48224c		450fb6c0		    MOVZX R8, R8		
  0x482250		4489442448		    MOVL R8, 0x48(SP)	

bi64 = To[int64](i64 > 0)
  0x482255		48833d5b3f0a0000    CMPQ xray/b2i.i64(SB), $0x0	
  0x48225d		410f9fc1		    SETG R9				
	return T(i)
  0x482261		450fb6c9		    MOVZX R9, R9		
  0x482265		4c894c2468		    MOVQ R9, 0x68(SP)	
```

#### unsigned integers

```
bui = To[uint](ui > 0)
  0x48226a		48833d4e3f0a0000	CMPQ xray/b2i.ui(SB), $0x0	
  0x482272		410f97c2		SETA R10			
	return T(i)
  0x482276		450fb6d2		MOVZX R10, R10		
  0x48227a		4c89542460		MOVQ R10, 0x60(SP)	
```
```
bui8 = To[uint8](ui8 > 0)
  0x48227f		803ddb3e0a0000		CMPB xray/b2i.ui8(SB), $0x0	
  0x482286		410f97c3		SETA R11			
		bui8,
  0x48228a		450fb6db		MOVZX R11, R11		
  0x48228e		4e8d1cdb		LEAQ 0(BX)(R11*8), R11	
  0x482292		4c899c24f8010000	MOVQ R11, 0x1f8(SP)	
bui16 = To[uint16](ui16 > 0)
  0x48229a		66833dc43e0a0000	CMPW xray/b2i.ui16(SB), $0x0	
  0x4822a2		410f97c4		SETA R12			
	return T(i)
  0x4822a6		450fb6e4		MOVZX R12, R12		
  0x4822aa		664489642440		MOVW R12, 0x40(SP)	
bui32 = To[uint32](ui32 > 0)
  0x4822b0		833db53e0a0000		CMPL xray/b2i.ui32(SB), $0x0	
  0x4822b7		410f97c5		SETA R13			
	return T(i)
  0x4822bb		450fb6ed		MOVZX R13, R13		
  0x4822bf		44896c2444		MOVL R13, 0x44(SP)	
bui64 = To[uint64](ui64 > 0)
  0x4822c4		48833dfc3e0a0000	CMPQ xray/b2i.ui64(SB), $0x0	
  0x4822cc		410f97c7		SETA R15			
  0x4822d0		44887c243f		MOVB R15, 0x3f(SP)		
```

#### floats

```
bf32 = To[float32](f32 > 0)
  0x4821dd		f30f10058b3f0a00	MOVSS xray/b2i.f32(SB), X0	
  0x4821e5		f30f1144244c		MOVSS X0, 0x4c(SP)		
bf64 = To[float64](f64 > 0)
  0x4821eb		f20f100ddd3f0a00	MOVSD_XMM xray/b2i.f64(SB), X1	
  0x4821f3		f20f114c2470		MOVSD_XMM X1, 0x70(SP)		
```

## Inlining

[https://dave.cheney.net/2014/06/07/five-things-that-make-go-fast](https://dave.cheney.net/2014/06/07/five-things-that-make-go-fast)

"Go’s optimisations are always enabled by default. You can see the compiler’s escape analysis and inlining decisions with the -gcflags=-m switch."

