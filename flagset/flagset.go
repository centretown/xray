package flagset

import (
	"flag"
	"log"

	"gopkg.in/yaml.v3"
)

type FlagSet struct {
	MajorKey, MinorKey int
	Key                string
	Path               string
	Resource           string
	Install            string
	Test               bool
}

func (cmdl *FlagSet) Dump() {
	buf, err := yaml.Marshal(cmdl)
	if err == nil {
		log.Println("\nFlags\n", string(buf))
	}
}

var Flags FlagSet

var (
	majorUsage    = "major version number"
	minorUsage    = "minor version number"
	keyUsage      = "uuid key"
	pathUsage     = "build destination folder"
	installUsage  = "install build to folder"
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
			flag.IntVar(&Flags.MajorKey, flagMajor, Flags.MajorKey, majorUsage)
			flag.IntVar(&Flags.MajorKey, flagM, Flags.MajorKey, sameAs(flagMajor))
		case flagM:
			flag.IntVar(&Flags.MajorKey, flagMajor, Flags.MajorKey, majorUsage)
			flag.IntVar(&Flags.MajorKey, flagM, Flags.MajorKey, sameAs(flagMajor))

		case flagMinor:
			flag.IntVar(&Flags.MinorKey, flagMinor, Flags.MinorKey, minorUsage)
			flag.IntVar(&Flags.MinorKey, flagN, Flags.MinorKey, sameAs(flagMinor))
		case flagN:
			flag.IntVar(&Flags.MinorKey, flagMinor, Flags.MinorKey, minorUsage)
			flag.IntVar(&Flags.MinorKey, flagN, Flags.MinorKey, sameAs(flagMinor))

		case flagKey:
			flag.StringVar(&Flags.Key, flagKey, Flags.Key, keyUsage)
			flag.StringVar(&Flags.Key, flagK, Flags.Key, sameAs(flagKey))
		case flagK:
			flag.StringVar(&Flags.Key, flagKey, Flags.Key, keyUsage)
			flag.StringVar(&Flags.Key, flagK, Flags.Key, sameAs(flagKey))

		case flagP:
			flag.StringVar(&Flags.Path, flagPath, Flags.Path, pathUsage)
			flag.StringVar(&Flags.Path, flagP, Flags.Path, sameAs(flagPath))
		case flagPath:
			flag.StringVar(&Flags.Path, flagPath, Flags.Path, pathUsage)
			flag.StringVar(&Flags.Path, flagP, Flags.Path, sameAs(flagPath))

		case flagInstall:
			flag.StringVar(&Flags.Install, flagInstall, Flags.Install, installUsage)
			flag.StringVar(&Flags.Install, flagI, Flags.Install, sameAs(flagInstall))
		case flagI:
			flag.StringVar(&Flags.Install, flagInstall, Flags.Install, installUsage)
			flag.StringVar(&Flags.Install, flagI, Flags.Install, sameAs(flagInstall))

		case flagTest:
			flag.BoolVar(&Flags.Test, flagTest, Flags.Test, testUsage)
			flag.BoolVar(&Flags.Test, flagT, Flags.Test, sameAs(flagTest))
		case flagT:
			flag.BoolVar(&Flags.Test, flagTest, Flags.Test, testUsage)
			flag.BoolVar(&Flags.Test, flagT, Flags.Test, sameAs(flagTest))

		case flagResource:
			flag.StringVar(&Flags.Resource, flagResource, Flags.Resource, resourceUsage)
			flag.StringVar(&Flags.Resource, flagR, Flags.Resource, sameAs(flagResource))
		case flagR:
			flag.StringVar(&Flags.Resource, flagResource, Flags.Resource, resourceUsage)
			flag.StringVar(&Flags.Resource, flagR, Flags.Resource, sameAs(flagResource))
		}
	}

}
