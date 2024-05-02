package notes

type LanguageChooser struct {
	languages *VocabularyItem
	chooser   *Chooser[*Language]
}

func NewLanguageChooser(vocabulary *VocabularyItem) *LanguageChooser {
	lch := &LanguageChooser{
		languages: vocabulary,
		chooser:   NewChooser(LanguageLabel, StringValue, &vocabulary.Languages),
	}
	var _ Note = lch
	return lch
}

func (lch *LanguageChooser) Item() *NoteItem {
	return lch.chooser.Item()
}

func (lch *LanguageChooser) Do(command COMMAND, args ...any) {
	lch.chooser.Do(command, args...)
	var (
		language = lch.Current()
		locale   = language.locale
		item     = lch.Item()
		output   = &lch.chooser.Output
	)
	output.Label = locale.TranslateWithFallback(language.fallback, item.LabelKey)
	output.Value = locale.TranslateWithFallback(language.fallback, item.FormatKey,
		lch.Values()...)
}

func (lch *LanguageChooser) Values() []any {
	return []any{(*lch.chooser.List)[lch.chooser.Current].String()}
}

func (lch *LanguageChooser) Current() *Language {
	return lch.languages.Languages[lch.chooser.Current]
}
