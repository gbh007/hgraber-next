import { AttributeColor } from '../../apiclient/api-attribute'
import { ContainerWidget } from '../design-system'
import {
    BookAttributeAutocompleteList,
    BookAttributeValueWidget,
} from './book-attribute'
import { attributeCodes } from './codes'

export function AttributeColorEditorWidget(props: {
    value: AttributeColor
    onChange: (v: AttributeColor) => void
    isNew?: boolean
}) {
    return (
        <ContainerWidget
            direction='2-column'
            gap='medium'
        >
            <span>Код</span>
            <select
                className='app'
                value={props.value.code}
                onChange={(e) => {
                    props.onChange({ ...props.value, code: e.target.value })
                }}
                disabled={!props.isNew}
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
            <span>Значение</span>
            <input
                className='app'
                value={props.value.value}
                onChange={(e) =>
                    props.onChange({ ...props.value, value: e.target.value })
                }
                disabled={!props.isNew}
                list={BookAttributeAutocompleteList(props.value.code)}
            />
            <span>Цвет текста</span>
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
            <span>Цвет фона</span>
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

            <span>Итог</span>
            <BookAttributeValueWidget
                value={props.value.value}
                color={props.value}
                code={props.value.code}
            />
        </ContainerWidget>
    )
}
