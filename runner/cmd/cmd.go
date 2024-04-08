package cmd

import "flag"

var (
	MajorKey, MinorKey int64
	Key                string
	Path               string
	majorUsage         = "major key number"
	minorUsage         = "minor key number"
	keyUsage           = "uuid key"
	pathUsage          = "data base path name"
)
var (
	flagMajor = "major"
	flagJ     = "mj"
	flagMinor = "minor"
	flagN     = "mn"
	flagKey   = "key"
	flagK     = "k"
	flagP     = "p"
	flagPath  = "file path"
)

func Setup() {
	sameAs := func(s string) string {
		return "same as -" + s
	}

	flag.Int64Var(&MajorKey, flagMajor, MajorKey, majorUsage)
	flag.Int64Var(&MajorKey, flagJ, MajorKey, sameAs(flagMajor))
	flag.Int64Var(&MinorKey, flagMinor, MinorKey, minorUsage)
	flag.Int64Var(&MinorKey, flagN, MinorKey, sameAs(flagMinor))
	flag.StringVar(&Key, flagKey, Key, keyUsage)
	flag.StringVar(&Key, flagK, Key, sameAs(flagKey))
	flag.StringVar(&Path, flagPath, Path, pathUsage)
	flag.StringVar(&Path, flagP, Path, sameAs(flagPath))

}
