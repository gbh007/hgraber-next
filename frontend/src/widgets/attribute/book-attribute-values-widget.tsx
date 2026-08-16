import { AttributeColor } from '../../apiclient/api-attribute'
import { BookAttributeValue } from '../../apiclient/model-book'
import { ContainerWidget } from '../design-system'
import { BookAttributeValueWidget } from './book-attribute'
import { MassloadLinkWidget } from './massload-link-widget'

export function BookAttributeValuesWidget(props: {
    code: string
    values: Array<string>
    valuesV2?: Array<BookAttributeValue>
    colors?: Array<AttributeColor>
}) {
    const colors = props.colors?.filter((color) => color.code == props.code)

    if (props.valuesV2) {
        return (
            <>
                {props.valuesV2.map((value) => (
                    <ContainerWidget direction='row'>
                        <BookAttributeValueWidget
                            code={props.code}
                            value={value.name}
                            color={colors?.find(
                                (color) => color.value == value.name
                            )}
                            key={value.name}
                        />
                        {value.massloads_by_name?.map((ml) => (
                            <MassloadLinkWidget
                                key={ml.id}
                                value={ml}
                            />
                        ))}
                    </ContainerWidget>
                ))}
            </>
        )
    }

    return (
        <>
            {props.values.map((value) => (
                <BookAttributeValueWidget
                    code={props.code}
                    value={value}
                    color={colors?.find((color) => color.value == value)}
                    key={value}
                />
            ))}
        </>
    )
}
