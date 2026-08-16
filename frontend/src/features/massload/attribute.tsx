import { MassloadInfoListRequestAttribute } from '../../apiclient/api-massload'
import {
    attributeCodes,
    BookAttributeAutocompleteList,
} from '../../widgets/attribute'
import {
    ContainerWidget,
    DeleteButtonWidget,
    ManyStringSelectWidget,
} from '../../widgets/design-system'

export function FilterAttributesWidget(props: {
    value: Array<MassloadInfoListRequestAttribute>
    onChange: (v: Array<MassloadInfoListRequestAttribute>) => void
}) {
    return (
        <ContainerWidget
            direction='column'
            gap='small'
        >
            <div>
                <span>Аттрибуты</span>
                <button
                    className='app'
                    onClick={() => {
                        props.onChange([
                            ...props.value,
                            {
                                code: 'tag', // TODO: не прибивать гвоздями
                                type: 'like', // TODO: не прибивать гвоздями
                                values: [],
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
                    <FilterAttributeWidget
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
                    />
                </ContainerWidget>
            ))}
        </ContainerWidget>
    )
}

export function FilterAttributeWidget(props: {
    value: MassloadInfoListRequestAttribute
    onChange: (v: MassloadInfoListRequestAttribute) => void
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
            </select>
            <ManyStringSelectWidget
                value={props.value.values ?? []}
                onChange={(e) => {
                    props.onChange({ ...props.value, values: e })
                }}
                autoCompleteID={BookAttributeAutocompleteList(props.value.code)}
            />
        </ContainerWidget>
    )
}
