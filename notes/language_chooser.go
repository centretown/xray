package notes

type LanguageChooser struct {
	Languages *Languages
	Chooser   *Chooser[*Language]
}

func NewLanguageChooser(languages *Languages) *LanguageChooser {
	lch := &LanguageChooser{
		Languages: languages,
		Chooser:   NewChooser(LanguageLabel, StringValue, &languages.List),
	}
	var _ Note = lch
	return lch
}

func (lch *LanguageChooser) Item() *NoteItem {
	return lch.Chooser.Item()
}

func (lch *LanguageChooser) Do(command COMMAND, args ...any) {
	lch.Chooser.Do(command, args...)
	var (
		language = lch.Current()
		locale   = language.locale
		item     = lch.Item()
		output   = &lch.Chooser.Output
	)
	output.Label = locale.TranslateWithFallback(language.fallback, item.LabelKey)
	output.Value = locale.TranslateWithFallback(language.fallback, item.FormatKey,
		lch.Values()...)
}

func (lch *LanguageChooser) Values() []any {
	return []any{(*lch.Chooser.List)[lch.Chooser.Current].String()}
}

func (lch *LanguageChooser) Current() *Language {
	return lch.Languages.List[lch.Chooser.Current]
}
