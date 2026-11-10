// утилиты форматирования, вынесенные из pretty.tsx — там только компоненты,
// иначе Fast Refresh (и react-refresh/only-export-components) ругается
export function prettyPercent(raw: number): number {
    return Math.round(raw * 1000) / 10
}

export function prettySize(value: number, div?: number) {
    let hasMinus = false
    if (value < 0) {
        hasMinus = true
        value *= -1
    }

    const divider = div ?? 1024

    let step = 0
    let mod = 0

    while (value / divider > 0.99) {
        step++
        mod = value % divider
        value = Math.floor(value / divider)
    }

    let unit = '??'

    switch (step) {
        case 0:
            unit = 'б'
            break
        case 1:
            unit = 'Кб'
            break
        case 2:
            unit = 'Мб'
            break
        case 3:
            unit = 'Гб'
            break
        case 4:
            unit = 'Тб'
            break
        case 5:
            unit = 'Пб'
            break
    }

    const modFirstDigit = Math.floor((mod * 10) / divider)

    return (
        (hasMinus ? '-' : '') +
        value.toString() +
        (modFirstDigit != 0 ? '.' + modFirstDigit.toString() : '') +
        ' ' +
        unit
    )
}
