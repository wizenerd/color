package Color

type Hue uint

const (
	Indigo Hue = iota
	Blue
	LightBlue
	Cyan
	Teal
	Green
	LightGreen
	Lime
	Yellow
	Amber
	Orange
	Brown
	BlueGrey
	Grey
	DeepOrange
	Red
	Pink
	Purple
	DeepPurple
)

func (h Hue) String() string {
	switch h {
	case Indigo:
		return "indigo"
	case Blue:
		return "blue"
	case LightBlue:
		return "light-blue"
	case Cyan:
		return "cyan"
	case Teal:
		return "teal"
	case Green:
		return "green"
	case LightGreen:
		return "light-green"
	case Lime:
		return "lime"
	case Yellow:
		return "yellow"
	case Amber:
		return "amber"
	case Orange:
		return "orange"
	case Brown:
		return "brown"
	case BlueGrey:
		return "blue-grey"
	case Grey:
		return "grey"
	case DeepOrange:
		return "deep-orange"
	case Red:
		return "red"
	case Pink:
		return "pink"
	case Purple:
		return "purple"
	case DeepPurple:
		return "deep-purple"
	}
	return ""
}

type Shade uint

const (
	S50 Shade = iota
	S100
	S200
	S300
	S400
	S500
	S600
	S700
	S800
	S900
	A100
	A200
	A400
	A700
)

func (s Shade) String() string {
	switch s {
	case S50:
		return "50"
	case S100:
		return "100"
	case S200:
		return "200"
	case S300:
		return "300"
	case S400:
		return "400"
	case S500:
		return "500"
	case S600:
		return "600"
	case S700:
		return "700"
	case S800:
		return "800"
	case S900:
		return "900"
	case A100:
		return "A100"
	case A200:
		return "A200"
	case A400:
		return "A400"
	case A700:
		return "A700"

	}
	return ""
}

type Color string

func Class(h Hue, s Shade) Color {
	return Color(h.String() + "-" + s.String())
}

const (
	Black           Color = "black"
	Primary         Color = "primary"
	PrimaryDark     Color = "primary-dark"
	PrimaryContrast Color = "primary-contrast"
	Accent          Color = "accent"
	AccentContrast  Color = "accent-contrast"
)

func (c Color) Background() string {
	return "mdl-color--" + string(c)
}

func (c Color) Text() string {
	return "mdl-color-text--" + string(c)
}
