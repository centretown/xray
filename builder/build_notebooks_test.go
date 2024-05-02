package builder

import (
	"testing"

	"github.com/centretown/xray/entries"
	"github.com/centretown/xray/gizzmo"
	"github.com/centretown/xray/notes"
	"github.com/centretown/xray/numbers"
)

const (
	testEnglishUS = iota
	testFrench
	testFrenchCA
	testEnglishCA
)

func TestNotes(t *testing.T) {
	gs := &gizzmo.Game{}
	BuildNotebooks(gs)

	options := gs.Options()

	chooser := options.Notes[0].(*notes.LanguageChooser)
	testLanguageChooser(t, options, chooser)

	fontEntry := options.Notes[1].(*notes.Ranger[float64])
	testRanger(t, options, fontEntry)

	monitor := options.Notes[2].(*entries.MonitorEntry)
	testMonitor(t, options, monitor)

	screen := options.Notes[3].(*entries.ScreenEntry)
	testScreen(t, options, screen)

	ntbk := options

	draw := func(i int, label, value string) {
		ntbk.Fetch()
		t.Log(label, value)
	}

	t.Log("")
	t.Log("DRAW ALL en_US map")
	chooser.Do(notes.SET, 0)
	ntbk.DrawAll(draw)

	t.Log("")
	t.Log("DRAW ALL fr map")
	chooser.Do(notes.SET, 1)
	ntbk.DrawAll(draw)

	t.Log("")
	t.Log("DRAW ALL fr_CA map")
	chooser.Do(notes.SET, 2)
	ntbk.DrawAll(draw)

	t.Log("")
	t.Log("DRAW ALL en_CA map")
	chooser.Do(notes.SET, 3)
	ntbk.DrawAll(draw)
}

func testLanguageChooser(t *testing.T, ntbk *notes.Notebook, chooser *notes.LanguageChooser) {

	item := chooser.Item()
	output := &item.Output

	testChooser := func(label string, want string) {
		ntbk.Fetch()
		t.Log(output.Label, output.Value)
		if output.Label != label {
			t.Fatalf("want: %s got %s", label, output.Label)
		}
		if output.Value != want {
			t.Fatalf("want: %s got %s", want, output.Value)
		}
	}

	chooser.Do(notes.SET, 1)
	testChooser("Langue", "Français")

	chooser.Do(notes.SET, 0)
	testChooser("Language", "English (US)")

	chooser.Do(notes.SET, 2)
	testChooser("Langue", "Français (CA)")

	chooser.Do(notes.SET, 3)
	testChooser("Language", "English (CA)")

}

func testRanger[T numbers.NumberType](t *testing.T, ntbk *notes.Notebook, rngr *notes.Ranger[T]) {

	item := rngr.Item()
	output := &item.Output
	test := func(label string, want string) {
		ntbk.Fetch()

		t.Log(output.Label, output.Value)
		if label != "" {
			if output.Label != label {
				t.Fatalf("want: %s got %s", label, output.Label)
			}
			if output.Value != want {
				t.Fatalf("want: %s got %s", want, output.Value)
			}
		}
	}

	ntbk.Language.Do(notes.SET, testEnglishUS)
	rngr.Do(notes.SET, 8.0)
	test("Font Size", "8")
	rngr.Do(notes.SET, 25.0)
	test("Font Size", "25")
	rngr.Do(notes.SET, 101.0)
	test("Font Size", "8")

	rngr.Do(notes.SET, 8.0)
	rngr.Do(notes.INCREMENT)
	test("Font Size", "9")

	ntbk.Language.Do(notes.SET, testFrench)
	rngr.Do(notes.DECREMENT)
	test("Taille de Police", "8")
	rngr.Do(notes.INCREMENT)
	rngr.Do(notes.INCREMENT)
	rngr.Do(notes.INCREMENT)
	rngr.Do(notes.INCREMENT)
	rngr.Do(notes.INCREMENT)

	ntbk.Language.Do(notes.SET, testFrenchCA)
	test("Taille de Police", "13")
	rngr.Do(notes.INCREMENT_MORE)
	test("Taille de Police", "23")

	ntbk.Language.Do(notes.SET, testEnglishCA)
	test("Font Size", "23")
	rngr.Do(notes.DECREMENT_MORE)
	test("Font Size", "13")
}

func testMonitor(t *testing.T, ntbk *notes.Notebook, mon *entries.MonitorEntry) {

	item := mon.Item()
	output := &item.Output

	test_monitor := func(label string, want string) {

		ntbk.Fetch()
		t.Log(output.Label, output.Value)
		if label != "" {
			if output.Label != label {
				t.Fatalf("want: '%s' got '%s'", label, output.Label)
			}
			if output.Value != want {
				t.Fatalf("want: '%s' got '%s'", want, output.Value)
			}
		}
	}

	ntbk.Language.Do(notes.SET, testFrench)
	test_monitor("Moniteur", "0 0x0 0 Mhz")
	custom := mon.Custom
	custom.Num = 1
	custom.Width = 2560
	custom.Height = 1440
	custom.RefreshRate = 60
	ntbk.Language.Do(notes.SET, testEnglishUS)
	test_monitor("Monitor", "1 2560x1440 60 Mhz")
	ntbk.Language.Do(notes.SET, testFrenchCA)
	test_monitor("Moniteur", "1 2560x1440 60 Mhz")

}

func testScreen(t *testing.T, ntbk *notes.Notebook, scr *entries.ScreenEntry) {
	custom := scr.Custom
	custom.Width = 1920
	custom.Height = 1080
	ntbk.Fetch()
}
