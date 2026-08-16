import { PropsWithChildren } from 'react'

export function ContainerWidget(
    props: PropsWithChildren<{
        direction?: 'row' | 'column' | '2-column' | 'columns'
        gap?: 'smaller' | 'small' | 'medium' | 'big' | 'bigger' | 'biggest'
        wrap?: boolean
        appContainer?: boolean
        style?: React.CSSProperties
        columns?: number
        className?: string | undefined
        id?: string
    }>
) {
    return (
        <div
            className={
                (props.className ? props.className + ' ' : '') +
                (props.appContainer ? 'app-container' : '')
            }
            style={{
                ...{
                    display:
                        props.direction == 'row' || props.direction == 'column'
                            ? 'flex'
                            : props.direction == '2-column'
                              ? 'grid'
                              : props.direction == 'columns' && props.columns
                                ? 'grid'
                                : undefined,
                    flexDirection:
                        props.direction == 'row'
                            ? 'row'
                            : props.direction == 'column'
                              ? 'column'
                              : undefined,
                    gridTemplateColumns:
                        props.direction == '2-column'
                            ? 'auto 1fr'
                            : props.direction == 'columns' && props.columns
                              ? '1fr '.repeat(props.columns)
                              : undefined,
                    flexWrap: props.wrap ? 'wrap' : undefined,
                    gap:
                        props.gap == 'smaller'
                            ? '3px'
                            : props.gap == 'small'
                              ? '5px'
                              : props.gap == 'medium'
                                ? '10px'
                                : props.gap == 'big'
                                  ? '15px'
                                  : props.gap == 'bigger'
                                    ? '30px'
                                    : props.gap == 'biggest'
                                      ? '50px'
                                      : undefined,
                },
                ...props.style,
            }}
            id={props.id}
        >
            {props.children}
        </div>
    )
}
