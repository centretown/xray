package flags

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
	Quick              bool
}

func (cmdl *FlagSet) Dump() {
	buf, err := yaml.Marshal(cmdl)
	if err == nil {
		log.Println("\nFlags\n", string(buf))
	}
}

var Flags FlagSet

func Setup(options ...string) {

	var (
		majorUsage   = "major version number"
		minorUsage   = "minor version number"
		keyUsage     = "uuid key"
		installUsage = "install build to folder"
		quickUsage   = "quick build and run in temporary memory database"

		flagMajor   = "major"
		flagMinor   = "minor"
		flagKey     = "key"
		flagInstall = "install"
		flagQuick   = "quick"

		flagM = "m"
		flagN = "n"
		flagK = "k"
		flagI = "i"
		flagQ = "q"
	)

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

		case flagInstall:
			flag.StringVar(&Flags.Install, flagInstall, Flags.Install, installUsage)
			flag.StringVar(&Flags.Install, flagI, Flags.Install, sameAs(flagInstall))
		case flagI:
			flag.StringVar(&Flags.Install, flagInstall, Flags.Install, installUsage)
			flag.StringVar(&Flags.Install, flagI, Flags.Install, sameAs(flagInstall))

		case flagQuick:
			flag.BoolVar(&Flags.Quick, flagQuick, Flags.Quick, quickUsage)
			flag.BoolVar(&Flags.Quick, flagQ, Flags.Quick, sameAs(flagQuick))
		case flagQ:
			flag.BoolVar(&Flags.Quick, flagQuick, Flags.Quick, quickUsage)
			flag.BoolVar(&Flags.Quick, flagQ, Flags.Quick, sameAs(flagQuick))
		}
	}

}
