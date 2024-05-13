package gruvbox

import (
	"image/color"
	"strconv"
)

// see https://github.com/morhetz/gruvbox

func parseHexByte(h string) uint8 {
	if v, err := strconv.ParseUint(string(h), 16, 32); err != nil {
		panic(err)
	} else {
		return uint8(v)
	}
}

func HTMLColor(c string) color.Color {
	return color.NRGBA{
		R: parseHexByte(c[0:2]),
		G: parseHexByte(c[2:4]),
		B: parseHexByte(c[4:6]),
		A: 0xff,
	}
}

func WithAlpha(c string, alpha uint8) color.Color {
	return color.NRGBA{
		R: parseHexByte(c[0:2]),
		G: parseHexByte(c[2:4]),
		B: parseHexByte(c[4:6]),
		A: alpha,
	}
}

type Palette struct {
	Bg0     color.Color
	Bg1     color.Color
	Bg2     color.Color
	Bg3     color.Color
	Bg4     color.Color
	Gray    color.Color
	Fg0     color.Color
	Fg1     color.Color
	Fg2     color.Color
	Fg3     color.Color
	Fg4     color.Color
	Red     color.Color
	Red2    color.Color
	Green   color.Color
	Green2  color.Color
	Yellow  color.Color
	Yellow2 color.Color
	Blue    color.Color
	Blue2   color.Color
	Purple  color.Color
	Purple2 color.Color
	Aqua    color.Color
	Aqua2   color.Color
	Orange  color.Color
	Orange2 color.Color
}

var (
	DarkMedium = Palette{
		Bg0:     HTMLColor("282828"),
		Bg1:     HTMLColor("3C3836"),
		Bg2:     HTMLColor("504945"),
		Bg3:     HTMLColor("665C54"),
		Bg4:     HTMLColor("7C6F64"),
		Gray:    HTMLColor("928374"),
		Fg0:     HTMLColor("FBF1C7"),
		Fg1:     HTMLColor("EBDBB2"),
		Fg2:     HTMLColor("D5C4A1"),
		Fg3:     HTMLColor("BDAE93"),
		Fg4:     HTMLColor("A89984"),
		Red:     HTMLColor("CC241D"),
		Red2:    HTMLColor("FB4934"),
		Green:   HTMLColor("98971A"),
		Green2:  HTMLColor("B8BB26"),
		Yellow:  HTMLColor("D79921"),
		Yellow2: HTMLColor("FABD2F"),
		Blue:    HTMLColor("458588"),
		Blue2:   HTMLColor("83A598"),
		Purple:  HTMLColor("B16286"),
		Purple2: HTMLColor("D3869B"),
		Aqua:    HTMLColor("689D6A"),
		Aqua2:   HTMLColor("8EC07C"),
		Orange:  HTMLColor("D65D0E"),
		Orange2: HTMLColor("FE8019"),
	}
	DarkHard Palette
	DarkSoft Palette

	LightMedium = Palette{
		Bg0:     HTMLColor("FBF1C7"),
		Bg1:     HTMLColor("EBDBB2"),
		Bg2:     HTMLColor("D5C4A1"),
		Bg3:     HTMLColor("BDAE93"),
		Bg4:     HTMLColor("A89984"),
		Gray:    HTMLColor("928374"),
		Fg0:     HTMLColor("282828"),
		Fg1:     HTMLColor("3C3836"),
		Fg2:     HTMLColor("504945"),
		Fg3:     HTMLColor("665C54"),
		Fg4:     HTMLColor("7C6F64"),
		Red:     HTMLColor("CC241D"),
		Red2:    HTMLColor("9D0006"),
		Green:   HTMLColor("98971A"),
		Green2:  HTMLColor("79740E"),
		Yellow:  HTMLColor("D79921"),
		Yellow2: HTMLColor("B57614"),
		Blue:    HTMLColor("458588"),
		Blue2:   HTMLColor("076678"),
		Purple:  HTMLColor("B16286"),
		Purple2: HTMLColor("8F3F71"),
		Aqua:    HTMLColor("689D6A"),
		Aqua2:   HTMLColor("427B58"),
		Orange:  HTMLColor("D65D0E"),
		Orange2: HTMLColor("AF3A03"),
	}
	LightHard Palette
	LightSoft Palette
)

func init() {
	DarkHard = DarkMedium
	DarkHard.Bg0 = HTMLColor("1D2021")
	DarkSoft = DarkMedium
	DarkSoft.Bg0 = HTMLColor("32302F")

	LightHard = LightMedium
	LightHard.Bg0 = HTMLColor("F9F5D7")
	LightSoft = LightMedium
	LightSoft.Bg0 = HTMLColor("F2E5BC")
}
