import { MassloadFlag } from '../../apiclient/api-massload'
import {
    ContainerWidget,
    DeleteButtonWidget,
} from '../../widgets/design-system'
import { MassloadOneFlagViewWidget } from './flag'

export function MassloadFlagEditorWidget(props: {
    value: MassloadFlag
    onChange: (v: MassloadFlag) => void
    isNew?: boolean
}) {
    return (
        <ContainerWidget
            direction='2-column'
            gap='medium'
        >
            <span>Код</span>
            <input
                className='app'
                value={props.value.code}
                onChange={(e) =>
                    props.onChange({ ...props.value, code: e.target.value })
                }
                disabled={!props.isNew}
            />
            <span>Название</span>
            <input
                className='app'
                value={props.value.name}
                onChange={(e) =>
                    props.onChange({ ...props.value, name: e.target.value })
                }
            />
            <span>Описание</span>
            <textarea
                className='app'
                rows={10}
                cols={50}
                value={props.value.description}
                onChange={(e) =>
                    props.onChange({
                        ...props.value,
                        description: e.target.value,
                    })
                }
            />
            <span>Порядок сортировки (больше = раньше)</span>
            <input
                className='app'
                placeholder='Приоритет'
                type='number'
                autoComplete='off'
                value={props.value.order_weight}
                onChange={(e) => {
                    props.onChange({
                        ...props.value,
                        order_weight: e.target.valueAsNumber,
                    })
                }}
            />
            <span>Цвет текста</span>
            <ContainerWidget direction='row'>
                <input
                    className='app'
                    type='color'
                    value={props.value.text_color}
                    onChange={(e) =>
                        props.onChange({
                            ...props.value,
                            text_color: e.target.value,
                        })
                    }
                />
                {props.value.text_color ? (
                    <DeleteButtonWidget
                        onClick={() =>
                            props.onChange({
                                ...props.value,
                                text_color: undefined,
                            })
                        }
                    />
                ) : null}
            </ContainerWidget>
            <span>Цвет фона</span>
            <ContainerWidget direction='row'>
                <input
                    className='app'
                    type='color'
                    value={props.value.background_color}
                    onChange={(e) =>
                        props.onChange({
                            ...props.value,
                            background_color: e.target.value,
                        })
                    }
                />
                {props.value.background_color ? (
                    <DeleteButtonWidget
                        onClick={() =>
                            props.onChange({
                                ...props.value,
                                background_color: undefined,
                            })
                        }
                    />
                ) : null}
            </ContainerWidget>
            <span>Итог</span>
            <MassloadOneFlagViewWidget value={props.value} />
        </ContainerWidget>
    )
}
