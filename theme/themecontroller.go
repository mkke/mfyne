package theme

import (
	"cmp"

	"fyne.io/fyne/v2"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

const PreferencesKeyTheme = "theme"

type ThemeController struct {
	app          fyne.App
	themes       map[string]*LabeledTheme
	themePref    string
	menuItemMaps []ThemeMenuItemMap
}

func NewThemeController(app fyne.App, defaultTheme string) *ThemeController {
	tc := &ThemeController{
		app:       app,
		themes:    make(map[string]*LabeledTheme),
		themePref: app.Preferences().StringWithFallback(PreferencesKeyTheme, defaultTheme),
	}

	return tc
}

func (tc *ThemeController) RegisterTheme(themes ...*LabeledTheme) *ThemeController {
	for _, t := range themes {
		tc.themes[t.PreferencesValue] = t
		if t.PreferencesValue == tc.themePref {
			tc.SetTheme(t)
		}
	}
	return tc
}

func (tc *ThemeController) Themes() []*LabeledTheme {
	sortedThemes := maps.Values(tc.themes)
	slices.SortFunc(sortedThemes, func(a, b *LabeledTheme) int {
		return cmp.Compare(a.LocalizedLabel(), b.LocalizedLabel())
	})
	return sortedThemes
}

func (tc *ThemeController) NewThemeMenu(label string) *fyne.Menu {
	var (
		themeMenuItems   []*fyne.MenuItem
		themeMenuItemMap = make(ThemeMenuItemMap)
	)
	for _, t := range tc.Themes() {
		themeMenuItem := fyne.NewMenuItem(t.LocalizedLabel(), func() {
			tc.SetTheme(t)
		})
		if t.PreferencesValue == tc.themePref {
			themeMenuItem.Checked = true
		}
		themeMenuItems = append(themeMenuItems, themeMenuItem)
		themeMenuItemMap[t] = themeMenuItem
	}
	tc.menuItemMaps = append(tc.menuItemMaps, themeMenuItemMap)
	return fyne.NewMenu(label, themeMenuItems...)
}

func (tc *ThemeController) NewThemeMenuItem(label string) *fyne.MenuItem {
	themeMenuItem := fyne.NewMenuItem(label, func() {})
	themeMenuItem.ChildMenu = tc.NewThemeMenu("")
	return themeMenuItem
}

func (tc *ThemeController) SetTheme(theme *LabeledTheme) {
	tc.app.Settings().SetTheme(theme)
	tc.app.Preferences().SetString(PreferencesKeyTheme, theme.PreferencesValue)
	for _, menuItemMap := range tc.menuItemMaps {
		menuItemMap.SetChecked(theme)
	}
}

type ThemeMenuItemMap map[*LabeledTheme]*fyne.MenuItem

func (tmim ThemeMenuItemMap) SetChecked(theme *LabeledTheme) {
	for menuItemTheme, menuItem := range tmim {
		menuItem.Checked = menuItemTheme == theme
	}
}
