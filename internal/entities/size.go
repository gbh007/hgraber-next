package entities

import "strconv"

func PrettySize(raw int64) string {
	if raw < 1 {
		return "0 б"
	}

	var div, mod int64

	const divider = 1024

	div = raw
	step := 0

	for div/divider > 0 {
		step++

		mod = div % divider
		div = div / divider
	}

	return strconv.FormatInt(div, 10) + "." + strconv.FormatInt(mod*10/1024, 10) + " " + SizeUnitFromStep(step)
}

func SizeUnitFromStep(step int) string {
	switch step {
	case 0:
		return "б"
	case 1:
		return "Кб"
	case 2:
		return "Мб"
	case 3:
		return "Гб"
	case 4:
		return "Тб"
	default:
		return "??"
	}
}
