import { useEffect, useState } from 'react'
import {
    AttributeColor,
    AttributeCountResponseAttribute,
    AttributeOriginCountResponseAttribute,
    AttributeRemapCreateRequest,
    AttributeRemapDeleteRequest,
    AttributeRemapUpdateRequest,
} from '../../apiclient/api-attribute'
import {
    ColorizedTextWidget,
    ContainerWidget,
} from '../../widgets/design-system'
import {
    attributeCodes,
    BookAttributeAutocompleteList,
    BookOneAttributeWidget,
} from '../../widgets/attribute'

export function AttributeRemapEditorWidget(props: {
    value: {
        code: string
        value: string
        to_code?: string
        to_value?: string
        is_delete?: boolean
    }
    isNew: boolean
    colors?: Array<AttributeColor>
    onCreate: (v: AttributeRemapCreateRequest) => void
    onUpdate: (v: AttributeRemapUpdateRequest) => void
    onDelete: (v: AttributeRemapDeleteRequest) => void
    attributeCount?: Array<AttributeCountResponseAttribute>
    originAttributeCount?: Array<AttributeOriginCountResponseAttribute>
}) {
    const [remap, setRemap] = useState<{
        code: string
        value: string
        to_code?: string
        to_value?: string
        is_delete?: boolean
    }>({ code: 'tag', value: '' })

    useEffect(() => {
        setRemap(props.value)
    }, [props.value])

    const count = props.attributeCount?.find(
        (attr) => attr.code == remap.to_code && attr.value == remap.to_value
    )?.count
    const originCount = props.originAttributeCount?.find(
        (attr) => attr.code == remap.code && attr.value == remap.value
    )?.count

    return (
        <ContainerWidget
            direction='row'
            gap='medium'
            style={{ alignItems: 'center' }}
        >
            <BookOneAttributeWidget
                code={remap.code}
                value={remap.value}
                colors={props.colors}
            />
            {originCount != undefined ? <span>({originCount})</span> : null}
            <select
                className='app'
                value={remap.code}
                onChange={(e) => {
                    setRemap({ ...remap, code: e.target.value })
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
            <input
                className='app'
                value={remap.value}
                onChange={(e) => setRemap({ ...remap, value: e.target.value })}
                list={BookAttributeAutocompleteList(remap.code, true)}
            />
            <label style={{ display: 'flex' }}>
                <input
                    className='app'
                    type='checkbox'
                    checked={remap?.is_delete ?? false}
                    onChange={(e) =>
                        setRemap({ ...remap, is_delete: e.target.checked })
                    }
                />
                <span>удалить</span>
            </label>
            {remap.is_delete ? null : (
                <>
                    <select
                        className='app'
                        value={remap.to_code}
                        onChange={(e) => {
                            setRemap({ ...remap, to_code: e.target.value })
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
                    <input
                        className='app'
                        value={remap.to_value}
                        onChange={(e) =>
                            setRemap({ ...remap, to_value: e.target.value })
                        }
                        list={BookAttributeAutocompleteList(
                            remap.to_code ?? ''
                        )}
                    />
                    {remap.to_code && remap.to_value ? (
                        <>
                            <BookOneAttributeWidget
                                code={remap.to_code}
                                value={remap.to_value}
                                colors={props.colors}
                            />
                            {count != undefined ? <span>({count})</span> : null}
                        </>
                    ) : null}
                </>
            )}
            {props.isNew ? (
                <button
                    className='app'
                    onClick={() => {
                        props.onCreate(remap)
                    }}
                >
                    <ColorizedTextWidget color='good'>
                        Создать
                    </ColorizedTextWidget>
                </button>
            ) : (
                <>
                    <button
                        className='app'
                        onClick={() => {
                            props.onUpdate(remap)
                        }}
                    >
                        <ColorizedTextWidget color='good'>
                            Сохранить
                        </ColorizedTextWidget>
                    </button>
                    <button
                        className='app'
                        onClick={() => {
                            props.onDelete(remap)
                        }}
                    >
                        <ColorizedTextWidget color='danger'>
                            Удалить
                        </ColorizedTextWidget>
                    </button>
                </>
            )}
        </ContainerWidget>
    )
}
