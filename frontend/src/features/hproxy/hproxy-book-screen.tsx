import { useState, useEffect } from 'react'
import { useSearchParams, Link } from 'react-router-dom'
import { useHProxyBook } from '../../apiclient/api-hproxy'
import { BookDetailsLink, HProxyListLink } from '../../core/routing'
import { useAttributeColorList } from '../../apiclient/api-attribute'
import { useSystemHandle } from '../../apiclient/api-system-handle'
import { DeletedBadge } from '../../widgets/design-system/index'
import { HProxyBookPagesPreviewWidget } from './book-pages-preview-widget'
import { ContainerWidget, ErrorTextWidget } from '../../widgets/design-system'
import {
    BookOneAttributeWidget,
    MassloadLinkWidget,
} from '../../widgets/attribute'
import { BadgeWidget, BookImagePreviewWidget } from '../../widgets/book'

const defaultPageLimit = 10

export function HProxyBookScreen() {
    const [searchParams, setSearchParams] = useSearchParams()
    const [hProxyBookResponse, doHProxyBook] = useHProxyBook()
    const [systemHandleResponse, doSystemHandle] = useSystemHandle()

    const [currentURL, setCurrentURL] = useState('')
    const [pageLimit, setPageLimit] = useState(defaultPageLimit)
    const [isReadOnlyMode, setIsReadOnlyMode] = useState(false)

    const [attributeColorListResponse, fetchAttributeColorList] =
        useAttributeColorList()
    useEffect(() => {
        fetchAttributeColorList()
    }, [fetchAttributeColorList])

    useEffect(() => {
        const u = searchParams.get('url')
        if (!u) {
            return
        }

        setCurrentURL(u)
        doHProxyBook({ url: u, page_limit: pageLimit })
    }, [doHProxyBook, searchParams])

    const downloadedFlags = {
        is_deleted: false,
        is_rebuild: false,
        is_verified: true,
        parsed_name: false,
        parsed_page: false,
    }

    const isDownloaded = hProxyBookResponse?.data?.exists_ids?.length ?? 0 > 0

    return (
        <ContainerWidget
            direction='column'
            gap='medium'
        >
            <ErrorTextWidget value={hProxyBookResponse} />
            <ErrorTextWidget value={attributeColorListResponse} />

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
                <input
                    className='app'
                    value={pageLimit}
                    type='number'
                    onChange={(e) => setPageLimit(e.target.valueAsNumber)}
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
                <button
                    className='app'
                    onClick={() => {
                        const current = pageLimit + defaultPageLimit
                        setPageLimit(current)
                        doHProxyBook({ url: currentURL, page_limit: current })
                    }}
                    disabled={hProxyBookResponse.isLoading}
                >
                    Показать больше страниц
                </button>

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
                            urls: [currentURL],
                        })
                    }}
                >
                    Загрузить
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
                appContainer
                direction='row'
                gap='medium'
            >
                <div>
                    <BookImagePreviewWidget
                        previewSize='superbig'
                        preview_url={hProxyBookResponse.data?.preview_url}
                        flags={isDownloaded ? downloadedFlags : undefined}
                    />
                </div>
                <ContainerWidget
                    direction='column'
                    gap='medium'
                >
                    <h1 style={{ wordBreak: 'break-all', margin: 0 }}>
                        {hProxyBookResponse.data?.name}
                    </h1>
                    <span>Страниц: {hProxyBookResponse.data?.page_count}</span>
                    {hProxyBookResponse.data?.attributes.map((attr, i) => (
                        <ContainerWidget
                            key={i}
                            direction='row'
                            gap='medium'
                            wrap
                            style={{
                                alignItems: 'center',
                            }}
                        >
                            <span>{attr.name}:</span>
                            {attr.values.map((v, i) => (
                                <ContainerWidget
                                    key={i}
                                    direction='row'
                                    style={{
                                        // border: "1px solid var(--app-color)",
                                        backgroundColor: 'var(--app-secondary)',
                                        borderRadius: '5px',
                                        alignItems: 'center',
                                        padding: '3px',
                                    }}
                                >
                                    {v.ext_url ? (
                                        <Link
                                            to={HProxyListLink(v.ext_url)}
                                            className='app-button'
                                        >
                                            {v.ext_name}
                                        </Link>
                                    ) : (
                                        <span>{v.ext_name}</span>
                                    )}
                                    {v.massloads_by_ext_url.map((ml) => (
                                        <MassloadLinkWidget
                                            key={ml.id}
                                            value={ml}
                                        />
                                    ))}
                                    {v.name ? (
                                        <BookOneAttributeWidget
                                            code={attr.code}
                                            value={v.name}
                                            colors={
                                                attributeColorListResponse.data
                                                    ?.colors
                                            }
                                        />
                                    ) : (
                                        <BadgeWidget
                                            previewSize='small'
                                            src={DeletedBadge}
                                        />
                                    )}
                                    {v.massloads_by_name.map((ml) => (
                                        <MassloadLinkWidget
                                            key={ml.id}
                                            value={ml}
                                        />
                                    ))}
                                </ContainerWidget>
                            ))}
                        </ContainerWidget>
                    ))}
                    {isDownloaded ? (
                        <ContainerWidget
                            direction='row'
                            gap='medium'
                            wrap
                        >
                            <span>Скачано:</span>
                            {hProxyBookResponse.data?.exists_ids?.map((id) => (
                                <Link
                                    className='app-button'
                                    to={BookDetailsLink(id)}
                                    key={id}
                                >
                                    {id}
                                </Link>
                            ))}
                        </ContainerWidget>
                    ) : null}
                </ContainerWidget>
            </ContainerWidget>

            <HProxyBookPagesPreviewWidget
                url={hProxyBookResponse.data?.ext_url ?? ''}
                pages={hProxyBookResponse.data?.pages}
                pageLimit={10}
            />
        </ContainerWidget>
    )
}
