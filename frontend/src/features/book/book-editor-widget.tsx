import type {
    AttributeColor,
    GetApiAttributeCount200AttributesItem,
    GetApiLabelPresetList200PresetsItem,
} from '../../apiclient/gen/hGraberNext.schemas'
import { BookRaw } from '../../apiclient/gen/hGraberNext.schemas'
import {
    BookAttributeInfoEditorWidget,
    BookLabelInfoEditorWidget,
    BookMainInfoEditorWidget,
} from '../../widgets/book'
import { ContainerWidget } from '../../widgets/design-system'

export function BookEditorWidget(props: {
    value: BookRaw
    onChange: (v: BookRaw) => void
    labelsAutoComplete?: Array<GetApiLabelPresetList200PresetsItem>
    attributeCount?: Array<GetApiAttributeCount200AttributesItem>
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
