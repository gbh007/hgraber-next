import { useState } from 'react'
import { AttributeColor } from '../../apiclient/api-attribute'
import { MassloadInfoAttribute } from '../../apiclient/api-massload'
import {
    ColorizedTextWidget,
    ContainerWidget,
    HumanTimeWidget,
} from '../../widgets/design-system'
import {
    attributeCodes,
    BookAttributeAutocompleteList,
    BookOneAttributeWidget,
} from '../../widgets/attribute'

export function MassloadAttributeEditorWidget(props: {
    value?: Array<MassloadInfoAttribute>
    onCreate: (code: string, text: string) => void
    onDelete: (code: string, text: string) => void
    colors?: Array<AttributeColor>
}) {
    const [code, setCode] = useState('')
    const [value, setValue] = useState('')

    return (
        <ContainerWidget
            appContainer
            gap='medium'
            direction='column'
        >
            <ContainerWidget
                direction='row'
                gap='small'
            >
                <select
                    className='app'
                    value={code}
                    onChange={(e) => {
                        setCode(e.target.value)
                    }}
                >
                    <option
                        value=''
                        key=''
                    >
                        Не выбрано
                    </option>
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
                    value={value}
                    onChange={(e) => setValue(e.target.value)}
                    list={BookAttributeAutocompleteList(code)}
                />
                <button
                    className='app'
                    onClick={() => {
                        props.onCreate(code, value)
                    }}
                >
                    <ColorizedTextWidget color='good'>
                        Создать
                    </ColorizedTextWidget>
                </button>
            </ContainerWidget>

            <table style={{ borderSpacing: '20px' }}>
                <thead>
                    <tr>
                        <td>Аттрибут</td>
                        <td>Создан</td>
                        <td>Действия</td>
                    </tr>
                </thead>
                <tbody>
                    {props.value?.map((attr, i) => (
                        <tr key={i}>
                            <td>
                                <BookOneAttributeWidget
                                    value={attr.value}
                                    colors={props.colors}
                                    code={attr.code}
                                />
                            </td>
                            <td>
                                <HumanTimeWidget value={attr.created_at} />
                            </td>
                            <td>
                                <ContainerWidget
                                    direction='column'
                                    gap='small'
                                >
                                    <button
                                        className='app'
                                        onClick={() => {
                                            props.onDelete(
                                                attr.code,
                                                attr.value
                                            )
                                        }}
                                    >
                                        <ColorizedTextWidget color='danger-lite'>
                                            Удалить
                                        </ColorizedTextWidget>
                                    </button>
                                </ContainerWidget>
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </ContainerWidget>
    )
}
