package cmdl

import (
	"flag"
	"fmt"

	"gopkg.in/yaml.v3"
)

type CmdLineFlags struct {
	MajorKey, MinorKey int
	Key                string
	Path               string
	Resource           string
	Install            bool
	Test               bool
}

func (cmdl *CmdLineFlags) Dump() {
	buf, err := yaml.Marshal(cmdl)
	if err == nil {
		fmt.Println(string(buf))
	}
}

var Cmdl CmdLineFlags

var (
	majorUsage    = "major version number"
	minorUsage    = "minor version number"
	keyUsage      = "uuid key"
	pathUsage     = "build destination folder"
	installUsage  = "install build"
	resourceUsage = "build resource location if different from path"
	testUsage     = "run test version"

	flagMajor    = "major"
	flagM        = "m"
	flagMinor    = "minor"
	flagN        = "n"
	flagKey      = "key"
	flagK        = "k"
	flagPath     = "path"
	flagP        = "p"
	flagInstall  = "install"
	flagI        = "i"
	flagTest     = "test"
	flagT        = "t"
	flagResource = "resource"
	flagR        = "r"
)

func Setup(options ...string) {
	sameAs := func(s string) string {
		return "same as -" + s
	}

	for _, option := range options {
		switch option {

		case flagMajor:
			flag.IntVar(&Cmdl.MajorKey, flagMajor, Cmdl.MajorKey, majorUsage)
			flag.IntVar(&Cmdl.MajorKey, flagM, Cmdl.MajorKey, sameAs(flagMajor))
		case flagM:
			flag.IntVar(&Cmdl.MajorKey, flagMajor, Cmdl.MajorKey, majorUsage)
			flag.IntVar(&Cmdl.MajorKey, flagM, Cmdl.MajorKey, sameAs(flagMajor))

		case flagMinor:
			flag.IntVar(&Cmdl.MinorKey, flagMinor, Cmdl.MinorKey, minorUsage)
			flag.IntVar(&Cmdl.MinorKey, flagN, Cmdl.MinorKey, sameAs(flagMinor))
		case flagN:
			flag.IntVar(&Cmdl.MinorKey, flagMinor, Cmdl.MinorKey, minorUsage)
			flag.IntVar(&Cmdl.MinorKey, flagN, Cmdl.MinorKey, sameAs(flagMinor))

		case flagKey:
			flag.StringVar(&Cmdl.Key, flagKey, Cmdl.Key, keyUsage)
			flag.StringVar(&Cmdl.Key, flagK, Cmdl.Key, sameAs(flagKey))
		case flagK:
			flag.StringVar(&Cmdl.Key, flagKey, Cmdl.Key, keyUsage)
			flag.StringVar(&Cmdl.Key, flagK, Cmdl.Key, sameAs(flagKey))

		case flagP:
			flag.StringVar(&Cmdl.Path, flagPath, Cmdl.Path, pathUsage)
			flag.StringVar(&Cmdl.Path, flagP, Cmdl.Path, sameAs(flagPath))
		case flagPath:
			flag.StringVar(&Cmdl.Path, flagPath, Cmdl.Path, pathUsage)
			flag.StringVar(&Cmdl.Path, flagP, Cmdl.Path, sameAs(flagPath))

		case flagInstall:
			flag.BoolVar(&Cmdl.Install, flagInstall, Cmdl.Install, installUsage)
			flag.BoolVar(&Cmdl.Install, flagI, Cmdl.Install, sameAs(flagInstall))
		case flagI:
			flag.BoolVar(&Cmdl.Install, flagInstall, Cmdl.Install, installUsage)
			flag.BoolVar(&Cmdl.Install, flagI, Cmdl.Install, sameAs(flagInstall))

		case flagTest:
			flag.BoolVar(&Cmdl.Test, flagTest, Cmdl.Test, testUsage)
			flag.BoolVar(&Cmdl.Test, flagT, Cmdl.Test, sameAs(flagTest))
		case flagT:
			flag.BoolVar(&Cmdl.Test, flagTest, Cmdl.Test, testUsage)
			flag.BoolVar(&Cmdl.Test, flagT, Cmdl.Test, sameAs(flagTest))

		case flagResource:
			flag.StringVar(&Cmdl.Resource, flagResource, Cmdl.Resource, resourceUsage)
			flag.StringVar(&Cmdl.Resource, flagR, Cmdl.Resource, sameAs(flagResource))
		case flagR:
			flag.StringVar(&Cmdl.Resource, flagResource, Cmdl.Resource, resourceUsage)
			flag.StringVar(&Cmdl.Resource, flagR, Cmdl.Resource, sameAs(flagResource))
		}
	}

}
