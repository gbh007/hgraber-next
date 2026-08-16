import { PropsWithChildren } from 'react'
import { ContainerWidget } from './container-widget'

export function HumanTimeWidget(props: { value: string }) {
    return <span>{new Date(props.value).toLocaleString()}</span>
}

export function prettyPercent(raw: number): number {
    return Math.round(raw * 1000) / 10
}

export function ColorizedTextWidget(
    props: PropsWithChildren<{
        color?: 'danger' | 'danger-lite' | 'good'
        bold?: boolean
    }>
) {
    const color =
        props.color == 'danger'
            ? 'var(--app-color-danger)'
            : props.color == 'danger-lite'
              ? 'var(--app-color-danger-lite)'
              : props.color == 'good'
                ? 'var(--app-color-good)'
                : undefined

    if (props.bold) {
        return <b style={{ color: color }}>{props.children}</b>
    }

    return <span style={{ color: color }}>{props.children}</span>
}

export function PrettySizeWidget(
    props: PropsWithChildren<{
        value?: number
    }>
) {
    if (props.value == undefined || props.value == null) {
        return null
    }

    return (
        <>
            {props.children}
            <span>{prettySize(props.value)}</span>
        </>
    )
}

export function PrettyDualSizeWidget(
    props: PropsWithChildren<{
        first?: number
        second?: number
    }>
) {
    const emptyFirst = props.first == undefined || props.first == null
    const emptySecond = props.second == undefined || props.second == null

    if (emptyFirst && emptySecond) {
        return null
    }

    return (
        <ContainerWidget
            direction='row'
            gap='small'
        >
            {props.children}
            <PrettySizeWidget value={props.first} />
            {!emptyFirst && !emptySecond && props.first != props.second ? (
                <span>({prettySize(props.second!)})</span>
            ) : null}
        </ContainerWidget>
    )
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

    let modFirstDigit = Math.floor((mod * 10) / divider)

    return (
        (hasMinus ? '-' : '') +
        value.toString() +
        (modFirstDigit != 0 ? '.' + modFirstDigit.toString() : '') +
        ' ' +
        unit
    )
}
