import {
    AttributeColor,
    AttributeCountResponseAttribute,
} from '../../apiclient/api-attribute'
import { BookRawAttribute } from '../../apiclient/model-book'
import {
    attributeCodes,
    BookAttributeAutocompleteList,
    BookAttributeAutocompleteWidget,
} from '../attribute'
import {
    ContainerWidget,
    DeleteButtonWidget,
    StringArrayPickerWidget,
} from '../design-system'

export function BookAttributeInfoEditorWidget(props: {
    value: Array<BookRawAttribute>
    onChange: (v: Array<BookRawAttribute>) => void
    attributeCount?: Array<AttributeCountResponseAttribute>
    colors?: Array<AttributeColor>
}) {
    return (
        <ContainerWidget
            appContainer
            direction='column'
            gap='medium'
        >
            <ContainerWidget
                direction='row'
                gap='small'
            >
                <b>Аттрибуты</b>
                <button
                    className='app'
                    onClick={() => {
                        props.onChange([
                            ...props.value,
                            {
                                code: 'tag', // FIXME: убрать этот хардкод
                                values: [],
                            },
                        ])
                    }}
                >
                    создать
                </button>
            </ContainerWidget>
            {props.value.map((label, i) => (
                <AttributeEditor
                    key={i}
                    value={label}
                    onChange={(e) =>
                        props.onChange(
                            props.value.map((v, index) => (i == index ? e : v))
                        )
                    }
                    onDelete={() =>
                        props.onChange(
                            props.value.filter((_, index) => i != index)
                        )
                    }
                    colors={props.colors}
                />
            ))}
            <BookAttributeAutocompleteWidget
                attributeCount={props.attributeCount}
                isOrigin={true}
            />
        </ContainerWidget>
    )
}

function AttributeEditor(props: {
    value: BookRawAttribute
    onChange: (v: BookRawAttribute) => void
    onDelete: () => void
    colors?: Array<AttributeColor>
}) {
    const colors = props.colors?.filter(
        (color) => color.code == props.value.code
    )

    return (
        <ContainerWidget
            direction='row'
            gap='small'
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
            <DeleteButtonWidget onClick={props.onDelete} />
            <StringArrayPickerWidget
                value={props.value.values}
                onChange={(e) => props.onChange({ ...props.value, values: e })}
                autoCompleteID={BookAttributeAutocompleteList(
                    props.value.code,
                    true
                )}
                colors={colors}
            />
        </ContainerWidget>
    )
}
