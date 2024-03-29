package try

import "fmt"

var (
	i    int     = -1
	i8   int8    = -8
	i16  int16   = -16
	i32  int32   = -32
	i64  int64   = -64
	ui   uint    = 1
	ui8  uint8   = 8
	ui16 uint16  = 16
	ui32 uint32  = 32
	ui64 uint64  = 64
	f32  float32 = 32.0
	f64  float64 = 64.0
)

func Inspect() {
	var (
		bi    int     = -1
		bi8   int8    = -8
		bi16  int16   = -16
		bi32  int32   = -32
		bi64  int64   = -64
		bui   uint    = 1
		bui8  uint8   = 8
		bui16 uint16  = 16
		bui32 uint32  = 32
		bui64 uint64  = 64
		bf32  float32 = 32.0
		bf64  float64 = 64.0
	)

	// INSPECT THIS BLOCK
	bi = As[int](i > 0)
	bi8 = As[int8](i8 > 0)
	bi16 = As[int16](i16 > 0)
	bi32 = As[int32](i32 > 0)
	bi64 = As[int64](i64 > 0)
	bui = As[uint](ui > 0)
	bui8 = As[uint8](ui8 > 0)
	bui16 = As[uint16](ui16 > 0)
	bui32 = As[uint32](ui32 > 0)
	bui64 = As[uint64](ui64 > 0)
	bf32 = As[float32](f32 > 0)
	bf64 = As[float64](f64 > 0)
	//

	fmt.Printf(""+
		"f32=%f "+
		"f64=%f "+
		"i=%d "+
		"i8=%d "+
		"i16=%d "+
		"i32=%d "+
		"i64=%d "+
		"ui=%d "+
		"ui8=%d "+
		"ui16=%d "+
		"ui32=%d "+
		"ui64=%d\n",
		f32,
		f64,
		i,
		i8,
		i16,
		i32,
		i64,
		ui,
		ui8,
		ui16,
		ui32,
		ui64)

	fmt.Printf(""+
		"bf32=%f "+
		"bf64=%f "+
		"bi=%d "+
		"bi8=%d "+
		"bi16=%d "+
		"bi32=%d "+
		"bi64=%d "+
		"bui=%d "+
		"bui8=%d "+
		"bui16=%d "+
		"bui32=%d "+
		"bui64=%d\n",
		bf32,
		bf64,
		bi,
		bi8,
		bi16,
		bi32,
		bi64,
		bui,
		bui8,
		bui16,
		bui32,
		bui64)
}
