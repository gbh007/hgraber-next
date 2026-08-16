import { BookFilterAttribute } from '../../apiclient/model-book-filter'
import { attributeCodes, BookAttributeAutocompleteList } from '../attribute'
import {
    ContainerWidget,
    DeleteButtonWidget,
    ManyStringSelectWidget,
} from '../design-system'

export function BookFilterAttributesWidget(props: {
    value: Array<BookFilterAttribute>
    onChange: (v: Array<BookFilterAttribute>) => void
}) {
    return (
        <ContainerWidget
            direction='column'
            gap='small'
        >
            <div>
                <span>Аттрибуты </span>
                <button
                    className='app'
                    onClick={() => {
                        props.onChange([
                            ...props.value,
                            {
                                code: 'tag', // TODO: не прибивать гвоздями
                                type: 'like', // TODO: не прибивать гвоздями
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
                    <BookFilterAttributeWidget
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

function BookFilterAttributeWidget(props: {
    value: BookFilterAttribute
    onChange: (v: BookFilterAttribute) => void
}) {
    return (
        <ContainerWidget
            direction='row'
            gap='medium'
            wrap
        >
            <select
                className='app'
                value={props.value.code}
                onChange={(e) => {
                    props.onChange({ ...props.value, code: e.target.value })
                }}
            >
                {attributeCodes.map((code) => (
                    <option
                        value={code}
                        key={code}
                    >
                        {code}
                    </option>
                ))}
            </select>
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
                    autoCompleteID={BookAttributeAutocompleteList(
                        props.value.code
                    )}
                />
            )}
        </ContainerWidget>
    )
}
