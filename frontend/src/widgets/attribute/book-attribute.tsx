import { Link } from 'react-router-dom'
import {
    AttributeColor,
    AttributeCountResponseAttribute,
} from '../../apiclient/api-attribute'
import { attributeCodes } from './codes'
import { BookListLinkAttribute } from '../../core/routing'

export function BookOneAttributeWidget(props: {
    code: string
    value: string
    colors?: Array<AttributeColor>
}) {
    const color = props.colors
        ?.filter((color) => color.code == props.code)
        ?.find((color) => color.value == props.value)

    return (
        <BookAttributeValueWidget
            code={props.code}
            value={props.value}
            color={color}
            key={props.value}
        />
    )
}

export function BookAttributeValueWidget(props: {
    value: string
    code: string
    color?: {
        text_color: string
        background_color: string
    }
}) {
    return (
        <span
            key={props.value}
            style={{
                borderRadius: '3px',
                padding: '3px',
                margin: '2px',
                color: props.color?.text_color ?? 'var(--app-attribute-color)',
                backgroundColor:
                    props.color?.background_color ??
                    'var(--app-attribute-background-color)',
                display: 'inline-block',
            }}
        >
            <Link
                style={{
                    textDecoration: 'none',
                    color: 'unset',
                }}
                to={BookListLinkAttribute(props.code, props.value)}
            >
                {props.value}
            </Link>
        </span>
    )
}

export function BookAttributeAutocompleteWidget(props: {
    attributeCount?: Array<AttributeCountResponseAttribute>
    isOrigin?: boolean
}) {
    if (!props.attributeCount) {
        return null
    }

    return (
        <>
            {attributeCodes.map((code) => (
                <datalist
                    key={code}
                    id={BookAttributeAutocompleteList(code, props.isOrigin)}
                >
                    {props.attributeCount
                        ?.filter((e) => e.code == code)
                        .map((attr) => (
                            <option
                                value={attr.value}
                                key={attr.value}
                            />
                        ))}
                </datalist>
            ))}
        </>
    )
}

export function BookAttributeAutocompleteList(
    code: string,
    isOrigin?: boolean
) {
    return (isOrigin ? 'origin-' : '') + 'attribute-autocomplete-' + code
}
