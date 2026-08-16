import { BookFilterLabel } from '../../apiclient/model-book-filter'
import {
    ContainerWidget,
    DeleteButtonWidget,
    ManyStringSelectWidget,
} from '../design-system'

export function BookFilterLabelsWidget(props: {
    value: Array<BookFilterLabel>
    onChange: (v: Array<BookFilterLabel>) => void
}) {
    return (
        <ContainerWidget
            direction='column'
            gap='small'
        >
            <div>
                <span>Метки </span>
                <button
                    className='app'
                    onClick={() => {
                        props.onChange([
                            ...props.value,
                            {
                                name: '',
                                type: 'in', // TODO: не прибивать гвоздями
                            },
                        ])
                    }}
                >
                    Добавить фильтр
                </button>
            </div>
            {props.value.map((v, i) => (
                <ContainerWidget
                    key={i}
                    direction='row'
                    gap='medium'
                >
                    <BookFilterLabelWidget
                        value={v}
                        onChange={(e) => {
                            props.onChange(
                                props.value.map((ov, index) =>
                                    index == i ? e : ov
                                )
                            )
                        }}
                    />
                    <DeleteButtonWidget
                        onClick={() => {
                            props.onChange(
                                props.value.filter((_, index) => index != i)
                            )
                        }}
                    ></DeleteButtonWidget>
                </ContainerWidget>
            ))}
        </ContainerWidget>
    )
}

function BookFilterLabelWidget(props: {
    value: BookFilterLabel
    onChange: (v: BookFilterLabel) => void
}) {
    return (
        <ContainerWidget
            direction='row'
            gap='medium'
            wrap
        >
            <input
                className='app'
                list='label-preset-names'
                value={props.value.name}
                onChange={(e) => {
                    props.onChange({ ...props.value, name: e.target.value })
                }}
                placeholder='название'
            />
            <select
                className='app'
                value={props.value.type}
                onChange={(e) => {
                    props.onChange({ ...props.value, type: e.target.value })
                }}
            >
                <option value='like'>LIKE</option>
                <option value='in'>IN</option>
                <option value='count_eq'>=</option>
                <option value='count_gt'>{'>'}</option>
                <option value='count_lt'>{'<'}</option>
            </select>
            {props.value.type == 'count_eq' ||
            props.value.type == 'count_gt' ||
            props.value.type == 'count_lt' ? (
                <input
                    className='app'
                    type='number'
                    value={props.value.count ?? 0}
                    onChange={(e) => {
                        props.onChange({
                            ...props.value,
                            count: e.target.valueAsNumber,
                        })
                    }}
                />
            ) : (
                <ManyStringSelectWidget
                    value={props.value.values ?? []}
                    onChange={(e) => {
                        props.onChange({ ...props.value, values: e })
                    }}
                    autoCompleteID={'label-preset-values-' + props.value.name}
                />
            )}
        </ContainerWidget>
    )
}
