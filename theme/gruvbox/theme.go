package gruvbox

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type Contrast int

const (
	Medium Contrast = iota
	Hard
	Soft
)

type Theme struct {
	Contrast Contrast
}

func (t Theme) Color(themeColorName fyne.ThemeColorName, themeVariant fyne.ThemeVariant) color.Color {
	palette := &DarkMedium
	switch themeVariant {
	case theme.VariantDark:
		switch t.Contrast {
		case Medium:
			palette = &DarkMedium
		case Hard:
			palette = &DarkHard
		case Soft:
			palette = &DarkSoft
		}
	case theme.VariantLight:
		switch t.Contrast {
		case Medium:
			palette = &LightMedium
		case Hard:
			palette = &LightHard
		case Soft:
			palette = &LightSoft
		}
	}

	switch themeColorName {
	case theme.ColorNameBackground:
		return palette.Bg0 // 0xff
	case theme.ColorNameButton:
		return palette.Bg2 // 0xf5
	case theme.ColorNameDisabled:
		return palette.Gray // 0xe3
	case theme.ColorNameDisabledButton:
		return palette.Bg2 // 0xf5
	case theme.ColorNameError:
		return palette.Red
	case theme.ColorNameFocus:
		return palette.Orange2
	case theme.ColorNameForeground:
		return palette.Fg0
	case theme.ColorNameHover:
		return palette.Fg4 // 0x0f
	case theme.ColorNameHeaderBackground:
		return palette.Bg1 // 0xf9
	case theme.ColorNameHyperlink:
		return palette.Blue
	case theme.ColorNameInputBackground:
		return palette.Bg3 // 0xf3
	case theme.ColorNameInputBorder:
		return palette.Gray // 0xe3
	case theme.ColorNameMenuBackground:
		return palette.Bg2 // 0xf5
	case theme.ColorNameOverlayBackground:
		return palette.Bg0 // 0xff
	case theme.ColorNamePlaceHolder:
		return palette.Gray // 0x88
	case theme.ColorNamePressed:
		return color.NRGBA{A: 0x19}
	case theme.ColorNamePrimary:
		return palette.Orange
	case theme.ColorNameScrollBar:
		return color.NRGBA{A: 0x99}
	case theme.ColorNameSelection:
		return palette.Orange2
	case theme.ColorNameSeparator:
		return palette.Gray // 0xe3
	case theme.ColorNameShadow:
		return color.NRGBA{A: 0x66}
	case theme.ColorNameSuccess:
		return palette.Green
	case theme.ColorNameWarning:
		return palette.Yellow
	default:
		return theme.DefaultTheme().Color(themeColorName, themeVariant)
	}
}

func (Theme) Font(textStyle fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(textStyle)
}

func (Theme) Icon(themeIconName fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(themeIconName)
}

func (Theme) Size(themeSizeName fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(themeSizeName)
}
