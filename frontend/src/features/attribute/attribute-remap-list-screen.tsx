import { useEffect, useState } from 'react'
import {
    AttributeRemapCreateRequest,
    AttributeRemapDeleteRequest,
    AttributeRemapUpdateRequest,
    useAttributeColorList,
    useAttributeCount,
    useAttributeOriginCount,
    useAttributeRemapCreate,
    useAttributeRemapDelete,
    useAttributeRemapList,
    useAttributeRemapUpdate,
} from '../../apiclient/api-attribute'
import { AttributeRemapEditorWidget } from './attribute-remap-editor-widget'
import { ContainerWidget, ErrorTextWidget } from '../../widgets/design-system'
import {
    attributeCodes,
    BookAttributeAutocompleteWidget,
} from '../../widgets/attribute'

export function AttributeRemapListScreen() {
    const [valueFilter, setValueFilter] = useState('')
    const [codeFilter, setCodeFilter] = useState('')
    const [showLimit, setShowLimit] = useState(true)
    const [maxShowOnLimit, setMaxShowOnLimit] = useState(100)

    const [attributeColorListResponse, fetchAttributeColorList] =
        useAttributeColorList()
    const [attributeCountResponse, getAttributeCount] = useAttributeCount()
    const [attributeOriginCountResponse, getAttributeOriginCount] =
        useAttributeOriginCount()
    const [attributeRemapListResponse, fetchAttributeRemapList] =
        useAttributeRemapList()

    const [attributeRemapDeleteResponse, doAttributeRemapDelete] =
        useAttributeRemapDelete()
    const [attributeRemapCreateResponse, doAttributeRemapCreate] =
        useAttributeRemapCreate()
    const [attributeRemapUpdateResponse, doAttributeRemapUpdate] =
        useAttributeRemapUpdate()

    useEffect(() => {
        fetchAttributeColorList()
    }, [fetchAttributeColorList])
    useEffect(() => {
        getAttributeCount()
    }, [getAttributeCount])
    useEffect(() => {
        getAttributeOriginCount()
    }, [getAttributeOriginCount])
    useEffect(() => {
        fetchAttributeRemapList()
    }, [fetchAttributeRemapList])

    const remappedAttributes = attributeRemapListResponse.data?.remaps
        ?.filter(
            (v) =>
                !v.is_delete &&
                (valueFilter == '' ||
                    v.value.toLowerCase().includes(valueFilter.toLowerCase()) ||
                    v.to_value
                        ?.toLowerCase()
                        .includes(valueFilter.toLowerCase())) &&
                (codeFilter == '' ||
                    v.code == codeFilter ||
                    v.to_code == codeFilter)
        )
        .filter((_, i) => !showLimit || i < maxShowOnLimit)
    const droppedAttributes = attributeRemapListResponse.data?.remaps
        ?.filter(
            (v) =>
                v.is_delete &&
                (valueFilter == '' ||
                    v.value.toLowerCase().includes(valueFilter.toLowerCase()) ||
                    v.to_value
                        ?.toLowerCase()
                        .includes(valueFilter.toLowerCase())) &&
                (codeFilter == '' ||
                    v.code == codeFilter ||
                    v.to_code == codeFilter)
        )
        .filter((_, i) => !showLimit || i < maxShowOnLimit)
    const newAttributes = attributeOriginCountResponse.data?.attributes
        ?.filter(
            (v) =>
                (valueFilter == '' ||
                    v.value
                        .toLowerCase()
                        .includes(valueFilter.toLowerCase())) &&
                (codeFilter == '' || v.code == codeFilter) &&
                !attributeRemapListResponse.data?.remaps?.find(
                    (attr) => attr.code == v.code && attr.value == v.value
                )
        )
        .filter((_, i) => !showLimit || i < maxShowOnLimit)

    return (
        <ContainerWidget
            direction='column'
            gap='big'
        >
            <ContainerWidget
                appContainer
                direction='column'
                gap='medium'
            >
                <ErrorTextWidget value={attributeColorListResponse} />
                <ErrorTextWidget value={attributeCountResponse} />
                <ErrorTextWidget value={attributeOriginCountResponse} />
                <ErrorTextWidget value={attributeRemapListResponse} />
                <ErrorTextWidget value={attributeRemapDeleteResponse} />
                <ErrorTextWidget value={attributeRemapCreateResponse} />
                <ErrorTextWidget value={attributeRemapUpdateResponse} />

                <ContainerWidget
                    direction='row'
                    gap='medium'
                    wrap
                >
                    <input
                        placeholder='Фильтр по названию'
                        className='app'
                        value={valueFilter}
                        onChange={(e) => setValueFilter(e.target.value)}
                    />

                    <select
                        className='app'
                        value={codeFilter}
                        onChange={(e) => {
                            setCodeFilter(e.target.value)
                        }}
                    >
                        <option value=''>Не выбрано</option>
                        {attributeCodes.map((code) => (
                            <option
                                value={code}
                                key={code}
                            >
                                {code}
                            </option>
                        ))}
                    </select>

                    <label style={{ display: 'flex' }}>
                        <input
                            className='app'
                            type='checkbox'
                            checked={showLimit}
                            onChange={(e) => setShowLimit(e.target.checked)}
                        />
                        <span>Ограничить вывод {maxShowOnLimit} первых</span>
                    </label>
                    <input
                        className='app'
                        type='number'
                        value={maxShowOnLimit}
                        onChange={(e) =>
                            setMaxShowOnLimit(e.target.valueAsNumber)
                        }
                    />
                </ContainerWidget>
            </ContainerWidget>

            {!remappedAttributes ? null : (
                <ContainerWidget
                    appContainer
                    direction='column'
                    gap='medium'
                >
                    <details className='app'>
                        <summary>Ремапинг</summary>
                        <ContainerWidget
                            direction='column'
                            gap='medium'
                        >
                            {remappedAttributes.map((attr) => (
                                <AttributeRemapEditorWidget
                                    isNew={false}
                                    value={attr}
                                    colors={
                                        attributeColorListResponse.data?.colors
                                    }
                                    key={attr.code + attr.value}
                                    onCreate={(
                                        v: AttributeRemapCreateRequest
                                    ) => {
                                        doAttributeRemapCreate(v).then(() =>
                                            fetchAttributeRemapList()
                                        )
                                    }}
                                    onUpdate={(
                                        v: AttributeRemapUpdateRequest
                                    ) => {
                                        doAttributeRemapUpdate(v).then(() =>
                                            fetchAttributeRemapList()
                                        )
                                    }}
                                    onDelete={(
                                        v: AttributeRemapDeleteRequest
                                    ) => {
                                        doAttributeRemapDelete(v).then(() =>
                                            fetchAttributeRemapList()
                                        )
                                    }}
                                    attributeCount={
                                        attributeCountResponse.data?.attributes
                                    }
                                    originAttributeCount={
                                        attributeOriginCountResponse.data
                                            ?.attributes
                                    }
                                />
                            ))}
                        </ContainerWidget>
                    </details>
                </ContainerWidget>
            )}

            {!droppedAttributes ? null : (
                <ContainerWidget
                    appContainer
                    direction='column'
                    gap='medium'
                >
                    <details className='app'>
                        <summary>Удаление атрибутов</summary>
                        <ContainerWidget
                            direction='column'
                            gap='medium'
                        >
                            {droppedAttributes.map((attr) => (
                                <AttributeRemapEditorWidget
                                    isNew={false}
                                    value={attr}
                                    colors={
                                        attributeColorListResponse.data?.colors
                                    }
                                    key={attr.code + attr.value}
                                    onCreate={(
                                        v: AttributeRemapCreateRequest
                                    ) => {
                                        doAttributeRemapCreate(v).then(() =>
                                            fetchAttributeRemapList()
                                        )
                                    }}
                                    onUpdate={(
                                        v: AttributeRemapUpdateRequest
                                    ) => {
                                        doAttributeRemapUpdate(v).then(() =>
                                            fetchAttributeRemapList()
                                        )
                                    }}
                                    onDelete={(
                                        v: AttributeRemapDeleteRequest
                                    ) => {
                                        doAttributeRemapDelete(v).then(() =>
                                            fetchAttributeRemapList()
                                        )
                                    }}
                                    attributeCount={
                                        attributeCountResponse.data?.attributes
                                    }
                                    originAttributeCount={
                                        attributeOriginCountResponse.data
                                            ?.attributes
                                    }
                                />
                            ))}
                        </ContainerWidget>
                    </details>
                </ContainerWidget>
            )}

            {!newAttributes ? null : (
                <ContainerWidget
                    appContainer
                    direction='column'
                    gap='medium'
                >
                    <details className='app'>
                        <summary>Новые аттрибуты</summary>
                        <ContainerWidget
                            direction='column'
                            gap='medium'
                        >
                            {newAttributes.map((attr) => (
                                <AttributeRemapEditorWidget
                                    isNew={true}
                                    value={attr}
                                    colors={
                                        attributeColorListResponse.data?.colors
                                    }
                                    key={attr.code + attr.value}
                                    onCreate={(
                                        v: AttributeRemapCreateRequest
                                    ) => {
                                        doAttributeRemapCreate(v).then(() =>
                                            fetchAttributeRemapList()
                                        )
                                    }}
                                    onUpdate={(
                                        v: AttributeRemapUpdateRequest
                                    ) => {
                                        doAttributeRemapUpdate(v).then(() =>
                                            fetchAttributeRemapList()
                                        )
                                    }}
                                    onDelete={(
                                        v: AttributeRemapDeleteRequest
                                    ) => {
                                        doAttributeRemapDelete(v).then(() =>
                                            fetchAttributeRemapList()
                                        )
                                    }}
                                    attributeCount={
                                        attributeCountResponse.data?.attributes
                                    }
                                    originAttributeCount={
                                        attributeOriginCountResponse.data
                                            ?.attributes
                                    }
                                />
                            ))}
                        </ContainerWidget>
                    </details>
                </ContainerWidget>
            )}

            <BookAttributeAutocompleteWidget
                attributeCount={attributeCountResponse.data?.attributes}
            />
            <BookAttributeAutocompleteWidget
                attributeCount={attributeOriginCountResponse.data?.attributes}
                isOrigin={true}
            />
        </ContainerWidget>
    )
}
