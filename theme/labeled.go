package theme

import (
	"fyne.io/fyne/v2"
)

type LabeledTheme struct {
	fyne.Theme
	PreferencesValue string
	LocalizedLabel   func(...string) string
	OverrideFonts    func(textStyle fyne.TextStyle) fyne.Resource
	OverrideIcons    func(themeIconName fyne.ThemeIconName) fyne.Resource
	OverrideSize     func(themeSizeName fyne.ThemeSizeName) float32
}

var _ fyne.Theme = (*LabeledTheme)(nil)

func (tw *LabeledTheme) Font(textStyle fyne.TextStyle) fyne.Resource {
	if tw.OverrideFonts != nil {
		return tw.OverrideFonts(textStyle)
	} else {
		return tw.Theme.Font(textStyle)
	}
}

func (tw *LabeledTheme) Icon(themeIconName fyne.ThemeIconName) fyne.Resource {
	if tw.OverrideIcons != nil {
		return tw.OverrideIcons(themeIconName)
	} else {
		return tw.Theme.Icon(themeIconName)
	}
}

func (tw *LabeledTheme) Size(themeSizeName fyne.ThemeSizeName) float32 {
	if tw.OverrideSize != nil {
		return tw.OverrideSize(themeSizeName)
	} else {
		return tw.Theme.Size(themeSizeName)
	}
}
