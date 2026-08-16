import { useCallback, useEffect, useState } from 'react'
import { useNavigate, useParams } from 'react-router-dom'
import {
    AttributeColor,
    useAttributeColorCreate,
    useAttributeColorGet,
    useAttributeColorUpdate,
    useAttributeCount,
} from '../../apiclient/api-attribute'
import { AttributeColorEditLink } from '../../core/routing'
import { ContainerWidget, ErrorTextWidget } from '../../widgets/design-system'
import {
    AttributeColorEditorWidget,
    BookAttributeAutocompleteWidget,
} from '../../widgets/attribute'

export function AttributeColorEditorScreen() {
    const params = useParams()
    const attributeCode = decodeURIComponent(params.code ?? '')
    const attributeValue = decodeURIComponent(params.value ?? '')

    const navigate = useNavigate()

    const [data, setData] = useState<AttributeColor>({
        code: attributeCode || 'tag',
        value: attributeValue || '',
        text_color: '#000000', // TODO: подумать можно ли это вытащить из цветов темы, и должны ли быть отдельные цвета на темы
        background_color: '#dfdfdf', // TODO: подумать можно ли это вытащить из цветов темы, и должны ли быть отдельные цвета на темы
        created_at: new Date().toJSON(),
    })

    const [attributeColorGetResponse, fetchAttributeColorGet] =
        useAttributeColorGet()
    const [attributeColorCreateResponse, doAttributeColorCreate] =
        useAttributeColorCreate()
    const [attributeColorUpdateResponse, doAttributeColorUpdate] =
        useAttributeColorUpdate()
    const [attributeCountResponse, getAttributeCount] = useAttributeCount()

    useEffect(() => {
        getAttributeCount()
    }, [getAttributeCount])

    useEffect(() => {
        if (attributeColorGetResponse.data) {
            setData(attributeColorGetResponse.data!)
        }
    }, [attributeColorGetResponse.data])

    const isExists = attributeCode && attributeValue

    useEffect(() => {
        if (isExists) {
            fetchAttributeColorGet({
                code: attributeCode,
                value: attributeValue,
            })
        }
    }, [fetchAttributeColorGet, attributeCode, attributeValue, isExists])

    const useSave = useCallback(() => {
        if (isExists) {
            doAttributeColorUpdate(data)
        } else {
            doAttributeColorCreate(data).then(() => {
                navigate(AttributeColorEditLink(data.code, data.value))
            })
        }
    }, [
        doAttributeColorUpdate,
        doAttributeColorCreate,
        navigate,
        isExists,
        data,
    ])

    return (
        <ContainerWidget
            appContainer
            direction='column'
            gap='big'
        >
            <ErrorTextWidget value={attributeColorGetResponse} />
            <ErrorTextWidget value={attributeColorCreateResponse} />
            <ErrorTextWidget value={attributeColorUpdateResponse} />
            <ErrorTextWidget value={attributeCountResponse} />
            <AttributeColorEditorWidget
                onChange={setData}
                value={data}
                isNew={!isExists}
            />
            <button
                className='app'
                onClick={useSave}
            >
                Сохранить
            </button>
            <BookAttributeAutocompleteWidget
                attributeCount={attributeCountResponse.data?.attributes}
            />
        </ContainerWidget>
    )
}
