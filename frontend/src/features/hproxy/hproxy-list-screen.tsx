import { Link, useSearchParams } from 'react-router-dom'
import { useHProxyList } from '../../apiclient/api-hproxy'
import { useEffect, useState } from 'react'
import {
    BookDetailsLink,
    HProxyBookLink,
    HProxyListLink,
} from '../../core/routing'
import { useSystemHandle } from '../../apiclient/api-system-handle'
import {
    ColorizedTextWidget,
    ContainerWidget,
    ErrorTextWidget,
} from '../../widgets/design-system'
import {
    BookImagePreviewWidget,
    ImageSize,
    PreviewSizeWidget,
} from '../../widgets/book'

export function HProxyListScreen() {
    const [searchParams, setSearchParams] = useSearchParams()
    const [hProxyListResponse, doHProxyList] = useHProxyList()
    const [systemHandleResponse, doSystemHandle] = useSystemHandle()
    const [imageSize, setImageSize] = useState<ImageSize>('medium')
    const [bookOnRow, setBookOnRow] = useState(4)

    const [currentURL, setCurrentURL] = useState('')
    const [isReadOnlyMode, setIsReadOnlyMode] = useState(false)

    useEffect(() => {
        const u = searchParams.get('url')
        if (!u) {
            return
        }

        setCurrentURL(u)
        doHProxyList({ url: u })
    }, [doHProxyList, searchParams])

    const downloadedFlags = {
        is_deleted: false,
        is_rebuild: false,
        is_verified: true,
        parsed_name: false,
        parsed_page: false,
    }

    return (
        <ContainerWidget
            direction='column'
            gap='medium'
        >
            <ErrorTextWidget value={hProxyListResponse} />

            <ContainerWidget
                appContainer
                direction='row'
                gap='medium'
                wrap
            >
                <input
                    className='app'
                    value={currentURL}
                    onChange={(e) => setCurrentURL(e.target.value)}
                />
                <button
                    className='app'
                    onClick={() => {
                        searchParams.set('url', currentURL)
                        setSearchParams(searchParams)
                    }}
                >
                    Перейти
                </button>

                <PreviewSizeWidget
                    value={imageSize}
                    onChange={setImageSize}
                />
                <input
                    className='app'
                    value={bookOnRow}
                    type='number'
                    onChange={(e) => setBookOnRow(e.target.valueAsNumber)}
                />

                <label>
                    <span>Расчет</span>
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
                            auto_verify: false,
                            is_multi: false,
                            read_only_mode: isReadOnlyMode,
                            urls:
                                hProxyListResponse.data?.books?.map(
                                    (b) => b.ext_url
                                ) ?? [],
                        })
                    }}
                >
                    Загрузить страницу
                </button>
                <button
                    className='app'
                    onClick={() => {
                        doSystemHandle({
                            auto_verify: false,
                            is_multi: true,
                            read_only_mode: isReadOnlyMode,
                            urls: [currentURL],
                        })
                    }}
                >
                    Загрузить все
                </button>

                <ContainerWidget
                    direction='row'
                    gap='small'
                    wrap
                >
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

            <ContainerWidget
                direction='columns'
                columns={bookOnRow}
                gap='medium'
                wrap
            >
                {hProxyListResponse.data?.books?.map((book, i) => (
                    <ContainerWidget
                        key={i}
                        appContainer
                        direction='column'
                        gap='medium'
                        style={{
                            alignItems: 'center',
                        }}
                    >
                        <BookImagePreviewWidget
                            preview_url={book.preview_url}
                            previewSize={imageSize}
                            flags={
                                (book.exists_ids?.length ?? 0 > 0)
                                    ? downloadedFlags
                                    : undefined
                            }
                        />
                        <ColorizedTextWidget bold>
                            {book.name}
                        </ColorizedTextWidget>
                        {book.exists_ids?.map((id) => (
                            <Link
                                className='app-button'
                                to={BookDetailsLink(id)}
                                key={id}
                            >
                                {id}
                            </Link>
                        ))}
                        <Link
                            to={HProxyBookLink(book.ext_url)}
                            className='app-button'
                        >
                            Открыть в HProxy
                        </Link>
                    </ContainerWidget>
                ))}
            </ContainerWidget>

            <ContainerWidget
                appContainer
                direction='row'
                gap='medium'
                wrap
            >
                {hProxyListResponse.data?.pagination?.map((page, i) => (
                    <Link
                        key={i}
                        to={HProxyListLink(page.ext_url)}
                        className='app-button'
                    >
                        {page.name}
                    </Link>
                ))}
            </ContainerWidget>
        </ContainerWidget>
    )
}
