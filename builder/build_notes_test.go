package builder

import (
	"fmt"
	"testing"

	"github.com/centretown/xray/entries"
	"github.com/centretown/xray/gizzmo"
	"github.com/centretown/xray/notes"
	"github.com/centretown/xray/numbers"
)

const (
	testEnglish = iota
	testFrench
)

func TestNotes(t *testing.T) {
	gs := &gizzmo.Game{}
	BuildNotes(gs)

	content := &gs.Content

	langEntry := content.CaptureNotes.Notes[0].(*notes.Chooser[*notes.Language])
	testLanguageChooser(t, content.CaptureNotes, &content.Languages, langEntry)

	fontEntry := content.CaptureNotes.Notes[1].(*notes.Ranger[float64])
	testRanger(t, content.CaptureNotes, &content.Languages, fontEntry)

	monitor := content.CaptureNotes.Notes[2].(*entries.MonitorEntry)
	testMonitor(t, content.CaptureNotes, &content.Languages, monitor)

	screen := content.CaptureNotes.Notes[3].(*entries.ScreenEntry)
	testScreen(t, content.CaptureNotes, &content.Languages, screen)

	draw := func(i int, label, value string) {
		t.Log(label, value)
	}

	nts := content.CaptureNotes
	languages := &content.Languages

	t.Log("")
	t.Log("DRAW ALL en_US map")
	nts.Fetch(languages.Items["en_US"])
	nts.DrawAll(draw)

	t.Log("")
	t.Log("DRAW ALL fr map")
	nts.Fetch(languages.Items["fr"])
	nts.DrawAll(draw)

	t.Log("")
	t.Log("DRAW ALL en_US list")
	nts.Fetch(languages.List[testEnglish])
	nts.DrawAll(draw)

	t.Log("")
	t.Log("DRAW ALL fr list")
	nts.Fetch(languages.List[testFrench])
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
	mon *entries.MonitorEntry) {

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
	scr *entries.ScreenEntry) {
	custom := scr.Custom
	custom.Width = 1920
	custom.Height = 1080
	nts.Fetch(languages.List[testEnglish])
}
