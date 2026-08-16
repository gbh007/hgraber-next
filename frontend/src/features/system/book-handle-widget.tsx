import { useEffect, useState } from 'react'
import { useSystemHandle } from '../../apiclient/api-system-handle'
import { ParseDetailsWidget } from './parse-details-widget'
import { ContainerWidget, ErrorTextWidget } from '../../widgets/design-system'

export function BookHandleWidget() {
    const [bookList, setBookList] = useState('')
    const [isMultiParse, setIsMultiParse] = useState(false)
    const [isAutoVerify, setIsAutoVerify] = useState(false)
    const [isReadOnlyMode, setIsReadOnlyMode] = useState(false)

    const [systemHandleResponse, doSystemHandle] = useSystemHandle()

    useEffect(() => {
        if (!systemHandleResponse.isError)
            setBookList(
                (systemHandleResponse.data?.not_handled || []).join('\n')
            )
    }, [systemHandleResponse.data, systemHandleResponse.isError, setBookList])

    return (
        <ContainerWidget
            direction='column'
            gap='big'
        >
            <ContainerWidget
                appContainer
                direction='row'
                gap='medium'
                wrap
            >
                <ContainerWidget direction='column'>
                    <textarea
                        className='app'
                        rows={10}
                        cols={50}
                        value={bookList}
                        onChange={(e) => {
                            setBookList(e.target.value)
                        }}
                        placeholder='Загрузить новые книги'
                    ></textarea>
                    <label>
                        <span>Множественный парсинг</span>
                        <input
                            className='app'
                            onChange={(e) => {
                                setIsMultiParse(e.target.checked)
                            }}
                            type='checkbox'
                            checked={isMultiParse}
                            autoComplete='off'
                        />
                    </label>
                    <label>
                        <span>Авто-подтверждение</span>
                        <input
                            className='app'
                            onChange={(e) => {
                                setIsAutoVerify(e.target.checked)
                            }}
                            type='checkbox'
                            checked={isAutoVerify}
                            autoComplete='off'
                        />
                    </label>
                    <label>
                        <span>Режим моделирования без изменений данных</span>
                        <input
                            className='app'
                            onChange={(e) => {
                                setIsReadOnlyMode(e.target.checked)
                            }}
                            type='checkbox'
                            checked={isReadOnlyMode}
                            autoComplete='off'
                        />
                    </label>
                    <button
                        className='app'
                        onClick={() => {
                            doSystemHandle({
                                urls: bookList
                                    .split('\n')
                                    .map((s) => s.trim())
                                    .filter((e) => e.length > 0),
                                is_multi: isMultiParse,
                                auto_verify: isAutoVerify,
                                read_only_mode: isReadOnlyMode,
                            })
                        }}
                        disabled={systemHandleResponse.isLoading}
                    >
                        Загрузить
                    </button>
                </ContainerWidget>
                <ContainerWidget direction='column'>
                    <ErrorTextWidget value={systemHandleResponse} />
                    <div>
                        <b>Всего: </b>
                        {systemHandleResponse.data?.total_count || 0}
                    </div>
                    <div>
                        <b>Загружено: </b>
                        {systemHandleResponse.data?.loaded_count || 0}
                    </div>
                    <div>
                        <b>Дубликаты: </b>
                        {systemHandleResponse.data?.duplicate_count || 0}
                    </div>
                    <div>
                        <b>Ошибки: </b>
                        {systemHandleResponse.data?.error_count || 0}
                    </div>
                </ContainerWidget>
            </ContainerWidget>
            <ParseDetailsWidget value={systemHandleResponse.data?.details} />
        </ContainerWidget>
    )
}
