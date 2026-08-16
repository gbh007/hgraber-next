import { useState } from 'react'
import { ContainerWidget } from './container-widget'
import { DeleteButtonWidget } from './buttons'

export function ManyStringSelectWidget(props: {
    value: Array<string>
    onChange: (v: Array<string>) => void
    autoCompleteID?: string
}) {
    const [value, setValue] = useState('')
    return (
        <ContainerWidget
            direction='column'
            gap='smaller'
        >
            <div>
                <input
                    className='app'
                    value={value}
                    onChange={(e) => setValue(e.target.value)}
                    list={props.autoCompleteID}
                />
                <button
                    className='app'
                    onClick={() => {
                        props.onChange([...props.value, value])
                        setValue('')
                    }}
                    disabled={value == ''}
                >
                    добавить
                </button>
            </div>
            {props.value.map((v, i) => (
                <ContainerWidget
                    key={i}
                    direction='row'
                    gap='smaller'
                >
                    <span>{v}</span>
                    <DeleteButtonWidget
                        onClick={() => {
                            props.onChange(
                                props.value.filter((_, index) => index != i)
                            )
                        }}
                    />
                </ContainerWidget>
            ))}
        </ContainerWidget>
    )
}

export function StringArrayPickerWidget(props: {
    name?: string
    value: Array<string>
    onChange: (v: Array<string>) => void
    autoCompleteID?: string
    colors?: Array<{
        value: string
        text_color: string
        background_color: string
    }>
}) {
    return (
        <ContainerWidget
            direction='column'
            gap='smaller'
        >
            <div>
                {props.name ? <span>{props.name}: </span> : null}
                <button
                    className='app'
                    onClick={() => {
                        props.onChange([...props.value, ''])
                    }}
                >
                    Добавить
                </button>
            </div>
            {props.value.map((value, i) => (
                <ContainerWidget
                    key={i}
                    direction='row'
                    gap='smaller'
                >
                    <input
                        className='app'
                        type='text'
                        value={value}
                        onChange={(e) => {
                            props.onChange(
                                props.value.map((value, ind) =>
                                    ind == i ? e.target.value : value
                                )
                            )
                        }}
                        list={props.autoCompleteID}
                        style={{
                            backgroundColor:
                                props.colors?.find((v) => v.value == value)
                                    ?.background_color || undefined,
                            color:
                                props.colors?.find((v) => v.value == value)
                                    ?.text_color || undefined,
                        }}
                    />
                    <DeleteButtonWidget
                        onClick={() => {
                            props.onChange(
                                props.value.filter((_, ind) => ind != i)
                            )
                        }}
                    />
                </ContainerWidget>
            ))}
        </ContainerWidget>
    )
}
