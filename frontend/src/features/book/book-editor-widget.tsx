import {
    AttributeColor,
    AttributeCountResponseAttribute,
} from '../../apiclient/api-attribute'
import { LabelPresetListResponseLabel } from '../../apiclient/api-labels'
import { BookRaw } from '../../apiclient/model-book'
import {
    BookAttributeInfoEditorWidget,
    BookLabelInfoEditorWidget,
    BookMainInfoEditorWidget,
} from '../../widgets/book'
import { ContainerWidget } from '../../widgets/design-system'

export function BookEditorWidget(props: {
    value: BookRaw
    onChange: (v: BookRaw) => void
    labelsAutoComplete?: Array<LabelPresetListResponseLabel>
    attributeCount?: Array<AttributeCountResponseAttribute>
    attributeColors?: Array<AttributeColor>
}) {
    return (
        <ContainerWidget
            direction='column'
            gap='medium'
        >
            <BookMainInfoEditorWidget
                value={props.value}
                onChange={props.onChange}
            />
            <BookLabelInfoEditorWidget
                value={props.value.labels ?? []}
                onChange={(e) => props.onChange({ ...props.value, labels: e })}
                labelsAutoComplete={props.labelsAutoComplete}
            />
            <BookAttributeInfoEditorWidget
                value={props.value.attributes ?? []}
                onChange={(e) =>
                    props.onChange({ ...props.value, attributes: e })
                }
                attributeCount={props.attributeCount}
                colors={props.attributeColors}
            />
        </ContainerWidget>
    )
}
