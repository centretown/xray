package builder

import (
	"fmt"
	"testing"

	"github.com/centretown/xray/notes"
	"github.com/centretown/xray/numbers"
)

const (
	testEnglish = iota
	testFrench
)

func TestNotes(t *testing.T) {
	var (
		languages   = NewLanguageList()
		fontSize    float64
		monitorItem notes.MonitorItem
		viewItem    notes.ViewItem
	)

	nts := notes.NewNotes()
	langChooser := notes.NewChooser(notes.LanguageLabel, notes.StringValue, &languages.List)
	nts.Add(langChooser)
	testLanguageChooser(t, nts, languages, langChooser)

	fontRanger := notes.NewRanger(notes.FontSizeLabel, notes.FloatValue, &fontSize, 8, 100, 10)
	nts.Add(fontRanger)
	testRanger(t, nts, languages, fontRanger)

	monitor := notes.NewMonitor(notes.MonitorLabel, notes.MonitorValue, &monitorItem)
	nts.Add(monitor)
	testMonitor(t, nts, languages, monitor)

	screen := notes.NewScreen(notes.ViewLabel, notes.ViewValue, &viewItem)
	nts.Add(screen)
	testScreen(t, nts, languages, screen)

	draw := func(i int, label, value string) {
		t.Log(label, value)
	}

	t.Log("DRAW ALL")
	nts.DrawAll(draw)
}

func testLanguageChooser[T fmt.Stringer](t *testing.T, nts *notes.Notes, languages *notes.Languages,
	cho *notes.Chooser[T]) {

	item := cho.Item()
	output := &item.Output

	testChooser := func(label string, want string) {
		nts.Fetch(languages.List[cho.Current])
		t.Log(output.Label, output.Value)
		if output.Label != label {
			t.Fatalf("want: %s got %s", label, output.Label)
		}
		if output.Value != want {
			t.Fatalf("want: %s got %s", want, output.Value)
		}
	}

	// only two languages better to have at least 3

	testChooser("Language", "English")
	cho.Do(notes.INCREMENT)
	testChooser("Langue", "Français")
	cho.Do(notes.DECREMENT)
	testChooser("Language", "English")
	cho.Do(notes.DECREMENT)
	testChooser("Langue", "Français")
}

func testRanger[T numbers.NumberType](t *testing.T, nts *notes.Notes, languages *notes.Languages,
	rngr *notes.Ranger[T]) {

	item := rngr.Item()
	output := &item.Output
	test := func(label string, want string, lang int) {
		nts.Fetch(languages.List[lang])

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

	rngr.Do(notes.INCREMENT)
	test("Font Size", "9", testEnglish)
	rngr.Do(notes.DECREMENT)
	test("Taille de Police", "8", testFrench)
	rngr.Do(notes.INCREMENT)
	rngr.Do(notes.INCREMENT)
	rngr.Do(notes.INCREMENT)
	rngr.Do(notes.INCREMENT)
	rngr.Do(notes.INCREMENT)
	test("Taille de Police", "13", testFrench)
	rngr.Do(notes.INCREMENT_MORE)
	test("Taille de Police", "23", testFrench)
	test("Font Size", "23", testEnglish)
	rngr.Do(notes.DECREMENT_MORE)
	test("Font Size", "13", testEnglish)
}

func testMonitor(t *testing.T, nts *notes.Notes, languages *notes.Languages,
	mon *notes.Monitor) {

	item := mon.Item()
	output := &item.Output

	test_monitor := func(label string, want string, lang int) {

		nts.Fetch(languages.List[lang])
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

	test_monitor("Moniteur", "0 0x0 0 Mhz", testFrench)
	custom := mon.Custom
	custom.Num = 1
	custom.Width = 2560
	custom.Height = 1440
	custom.RefreshRate = 60
	test_monitor("Monitor", "1 2560x1440 60 Mhz", testEnglish)
	test_monitor("Moniteur", "1 2560x1440 60 Mhz", testFrench)

}

func testScreen(t *testing.T, nts *notes.Notes, languages *notes.Languages,
	scr *notes.View) {
	custom := scr.Custom
	custom.Width = 1920
	custom.Height = 1080
	nts.Fetch(languages.List[testEnglish])
}
