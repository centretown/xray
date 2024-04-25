package flags

import (
	"flag"
	"log"

	"github.com/centretown/xray/notes"
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
	notes.Initialize()

	var (
		majorUsage   = notes.Current.Translate(notes.MajorUsage)
		minorUsage   = notes.Current.Translate(notes.MinorUsage)
		keyUsage     = notes.Current.Translate(notes.KeyUsage)
		installUsage = notes.Current.Translate(notes.InstallUsage)
		quickUsage   = notes.Current.Translate(notes.QuickUsage)

		flagMajor   = notes.Current.Translate(notes.MajorFlag)
		flagMinor   = notes.Current.Translate(notes.MinorFlag)
		flagKey     = notes.Current.Translate(notes.KeyFlag)
		flagInstall = notes.Current.Translate(notes.InstallFlag)
		flagQuick   = notes.Current.Translate(notes.QuickFlag)

		flagM = notes.Current.Translate(notes.MajorShort)
		flagN = notes.Current.Translate(notes.MinorShort)
		flagK = notes.Current.Translate(notes.KeyShort)
		flagI = notes.Current.Translate(notes.InstallShort)
		flagQ = notes.Current.Translate(notes.QuickShort)
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
