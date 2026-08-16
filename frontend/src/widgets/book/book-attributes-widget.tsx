import { AttributeColor } from '../../apiclient/api-attribute'
import { BookAttribute } from '../../apiclient/model-book'
import { BookAttributeValuesWidget } from '../attribute'
import { ContainerWidget } from '../design-system'

export function BookAttributesWidget(props: {
    value?: Array<BookAttribute>
    colors?: Array<AttributeColor>
}) {
    if (!props.value) {
        return null
    }

    return (
        <ContainerWidget
            direction='column'
            gap='small'
        >
            {props.value?.map((attr) => (
                <BookAttributeWidget
                    key={attr.code}
                    value={attr}
                    colors={props.colors}
                />
            ))}
        </ContainerWidget>
    )
}

function BookAttributeWidget(props: {
    value: BookAttribute
    colors?: Array<AttributeColor>
}) {
    return (
        <ContainerWidget
            direction='row'
            gap='small'
            wrap
            style={{ alignItems: 'center' }}
        >
            <span>{props.value.name}:</span>
            <BookAttributeValuesWidget
                code={props.value.code}
                values={props.value.values}
                colors={props.colors}
                valuesV2={props.value.values_v2}
            />
        </ContainerWidget>
    )
}
