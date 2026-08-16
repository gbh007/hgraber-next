import { useState } from 'react'
import { MassloadInfoExternalLink } from '../../apiclient/api-massload'
import {
    ColorizedTextWidget,
    ContainerWidget,
    HumanTimeWidget,
} from '../../widgets/design-system'

interface linkData {
    url: string
    auto_check?: boolean
}

export function MassloadExternalLinkEditorWidget(props: {
    value?: Array<MassloadInfoExternalLink>
    onCreate: (v: linkData) => void
    onUpdate: (v: linkData) => void
    onDelete: (url: string) => void
}) {
    const [value, setValue] = useState('')
    const [autoCheck, setAutoCheck] = useState(false)
    const [edit, setEdit] = useState(false)

    return (
        <ContainerWidget
            appContainer
            gap='medium'
            direction='column'
        >
            <ContainerWidget
                direction='row'
                gap='small'
                style={{ alignItems: 'center' }}
            >
                <input
                    className='app'
                    value={value}
                    onChange={(e) => setValue(e.target.value)}
                />
                <label>
                    <input
                        className='app'
                        checked={autoCheck}
                        onChange={(e) => {
                            setAutoCheck(e.target.checked)
                        }}
                        type='checkbox'
                    />
                    <span>Авто проверка</span>
                </label>
                <button
                    className='app'
                    onClick={() => {
                        if (edit) {
                            props.onUpdate({
                                url: value,
                                auto_check: autoCheck,
                            })
                        } else {
                            props.onCreate({
                                url: value,
                                auto_check: autoCheck,
                            })
                        }
                    }}
                >
                    <ColorizedTextWidget color='good'>
                        {edit ? 'Сохранить' : 'Создать'}
                    </ColorizedTextWidget>
                </button>
                <button
                    className='app'
                    onClick={() => {
                        setValue('')
                        setAutoCheck(false)
                        setEdit(false)
                    }}
                >
                    <ColorizedTextWidget color='danger-lite'>
                        Очистить
                    </ColorizedTextWidget>
                </button>
            </ContainerWidget>

            <table style={{ borderSpacing: '20px' }}>
                <thead>
                    <tr>
                        <td>Значение</td>
                        <td>Авто проверка</td>
                        <td>Создана</td>
                        <td>Действия</td>
                    </tr>
                </thead>
                <tbody>
                    {props.value?.map((link, i) => (
                        <tr key={i}>
                            <td>{link.url}</td>
                            <td>
                                {link.auto_check ? (
                                    <ColorizedTextWidget
                                        color='good'
                                        bold
                                    >
                                        Да
                                    </ColorizedTextWidget>
                                ) : (
                                    <ColorizedTextWidget color='danger-lite'>
                                        Нет
                                    </ColorizedTextWidget>
                                )}
                            </td>
                            <td>
                                <HumanTimeWidget value={link.created_at} />
                            </td>
                            <td>
                                <ContainerWidget
                                    direction='column'
                                    gap='small'
                                >
                                    <button
                                        className='app'
                                        onClick={() => {
                                            props.onDelete(link.url)
                                        }}
                                    >
                                        <ColorizedTextWidget color='danger-lite'>
                                            Удалить
                                        </ColorizedTextWidget>
                                    </button>
                                    <button
                                        className='app'
                                        onClick={() => {
                                            setValue(link.url)
                                            setAutoCheck(link.auto_check)
                                            setEdit(true)
                                        }}
                                    >
                                        Изменить
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
