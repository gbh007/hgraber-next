import { PropsWithChildren } from 'react'
import { ContainerWidget } from './container-widget'
import { prettySize } from './pretty-utils'

export function HumanTimeWidget(props: { value: string }) {
    return <span>{new Date(props.value).toLocaleString()}</span>
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
