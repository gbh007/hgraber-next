import { MassloadFlag } from '../../apiclient/api-massload'
import { ContainerWidget } from '../../widgets/design-system'

export function MassloadFlagViewWidget(props: {
    flags?: Array<string>
    flagInfos?: Array<MassloadFlag>
}) {
    if (!props.flags) {
        return null
    }

    return (
        <ContainerWidget
            direction='row'
            gap='small'
            wrap
        >
            {props.flags.map((flag, i) => (
                <MassloadOneFlagViewWidget
                    key={i}
                    value={{
                        code: flag,
                        name: props.flagInfos?.find((v) => v.code == flag)
                            ?.name,
                        text_color: props.flagInfos?.find((v) => v.code == flag)
                            ?.text_color,
                        background_color: props.flagInfos?.find(
                            (v) => v.code == flag
                        )?.background_color,
                    }}
                />
            ))}
        </ContainerWidget>
    )
}

export function MassloadOneFlagViewWidget(props: {
    value: {
        code: string
        name?: string
        text_color?: string
        background_color?: string
    }
}) {
    const color = 'var(--app-color)'
    const backgroundColor = 'var(--app-background)'

    return (
        <span
            style={{
                borderRadius: '3px',
                padding: '3px',
                color: props.value.text_color || color,
                background: props.value.background_color || backgroundColor,
            }}
        >
            {props.value.name || props.value.code}
        </span>
    )
}

export function MassloadFlagPickerWidget(props: {
    value: Array<string>
    onChange: (v: Array<string>) => void
    flagInfos: Array<MassloadFlag>
}) {
    const color = 'var(--app-color)'
    const backgroundColor = 'var(--app-background)'

    return (
        <ContainerWidget
            direction='column'
            gap='small'
            wrap
        >
            {props.flagInfos.map((flag) => (
                <label
                    key={flag.code}
                    style={{
                        borderRadius: '3px',
                        padding: '3px',
                        color: flag.text_color || color,
                        background: flag.background_color || backgroundColor,
                    }}
                >
                    <input
                        className='app'
                        type='checkbox'
                        checked={props.value.includes(flag.code)}
                        onChange={(e) =>
                            props.onChange(
                                e.target.checked
                                    ? [...props.value, flag.code]
                                    : props.value.filter((v) => v != flag.code)
                            )
                        }
                    />
                    <span>{flag.name || flag.code}</span>
                </label>
            ))}
        </ContainerWidget>
    )
}
