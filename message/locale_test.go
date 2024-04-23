package message

import (
	"fmt"
	"testing"

	"unknwon.dev/i18n"
)

func TestLocale(t *testing.T) {
	s := i18n.NewStore()
	l1, err := s.AddLocale("en-US", "English", "locale/locale_en-US.ini")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(l1.Translate("messages::test1", 1, 2))
	// => "This patch has 1 changed file and deleted 2 files"

	fmt.Println(l1.Translate("messages::capture"),
		l1.Translate("messages::capturev", 1000, 750))
	fmt.Println(l1.Translate("messages::monitor"),
		l1.Translate("messages::monitorv", 0, 2560, 1440, 60))
	fmt.Println(l1.Translate("messages::view"),
		l1.Translate("messages::viewv", 1920, 1080))
	fmt.Println(l1.Translate("messages::framerate"),
		l1.Translate("messages::frameratev", 25.6))
	fmt.Println(l1.Translate("messages::duration"),
		l1.Translate("messages::durationv", 30))
}
