import { useEffect, useState } from 'react'
import { Link, useParams } from 'react-router-dom'
import {
    useDeduplicateCompare,
    useSetDeadHash,
} from '../../apiclient/api-deduplicate'
import { useDeleteBookPage } from '../../apiclient/api-book-delete'
import { useAttributeColorList } from '../../apiclient/api-attribute'
import { DualReaderWidget } from './dual-reader-widget'
import { BookSimple, BookSimplePage } from '../../apiclient/model-book'
import { BookDetailsLink } from '../../core/routing'
import {
    ContainerWidget,
    ErrorTextWidget,
    HumanTimeWidget,
    PrettyDualSizeWidget,
    prettyPercent,
} from '../../widgets/design-system'
import {
    BookAttributesWidget,
    BookImagePreviewWidget,
    BookPagesPreviewWidget,
} from '../../widgets/book'

export function CompareBookScreen() {
    const params = useParams()
    const originBookID = params.origin!
    const targetBookID = params.target!

    const [currentShow, setCurrentShow] = useState('origin')
    const [deadHashSelector, setDeadHashSelector] = useState('all')

    const [compareResult, doCompare] = useDeduplicateCompare()
    const [setDeadHashResponse, doSetDeadHash] = useSetDeadHash()
    const [deletePageResponse, doDeletePage] = useDeleteBookPage()

    const [attributeColorListResponse, fetchAttributeColorList] =
        useAttributeColorList()
    useEffect(() => {
        fetchAttributeColorList()
    }, [fetchAttributeColorList])

    useEffect(() => {
        doCompare({
            origin_book_id: originBookID,
            target_book_id: targetBookID,
        })
    }, [doCompare, originBookID, targetBookID])

    const originBookName = compareResult.data?.origin.name
    const targetBookName = compareResult.data?.target.name

    const originPages = compareResult.data?.origin_pages?.filter(
        (page) =>
            deadHashSelector == 'all' ||
            (deadHashSelector == 'without' && page.has_dead_hash === false) ||
            (deadHashSelector == 'only' && page.has_dead_hash === true)
    )

    const bothPages = compareResult.data?.both_pages?.filter(
        (page) =>
            deadHashSelector == 'all' ||
            (deadHashSelector == 'without' && page.has_dead_hash === false) ||
            (deadHashSelector == 'only' && page.has_dead_hash === true)
    )

    const targetPages = compareResult.data?.target_pages?.filter(
        (page) =>
            deadHashSelector == 'all' ||
            (deadHashSelector == 'without' && page.has_dead_hash === false) ||
            (deadHashSelector == 'only' && page.has_dead_hash === true)
    )

    return (
        <ContainerWidget
            direction='column'
            gap='bigger'
        >
            <ErrorTextWidget value={compareResult} />
            <ErrorTextWidget value={setDeadHashResponse} />
            <ErrorTextWidget value={deletePageResponse} />
            <ErrorTextWidget value={attributeColorListResponse} />

            <ContainerWidget
                appContainer
                direction='row'
                gap='medium'
                wrap
            >
                <ContainerWidget
                    direction='row'
                    gap='small'
                >
                    <BookImagePreviewWidget
                        flags={compareResult.data?.origin.flags}
                        previewSize='superbig'
                        preview_url={compareResult.data?.origin.preview_url}
                    />
                    <BookShortInfo
                        value={compareResult.data?.origin}
                        covered_target={
                            compareResult.data?.origin_covered_target
                        }
                        covered_target_without_dead_hashes={
                            compareResult.data
                                ?.origin_covered_target_without_dead_hashes
                        }
                        book_size={compareResult.data?.origin_book_size}
                        page_size={compareResult.data?.origin_page_avg_size}
                    />
                </ContainerWidget>
                <div style={{ flexGrow: 1, textAlign: 'center' }}>
                    <ContainerWidget
                        direction='column'
                        gap='small'
                    >
                        <select
                            className='app'
                            value={currentShow}
                            onChange={(e) => setCurrentShow(e.target.value)}
                        >
                            <option value='origin'>Страницы оригинала</option>
                            <option value='both'>Страницы общие</option>
                            <option value='target'>Страницы цели</option>
                            <option value='dual_read'>
                                Посмотреть отличия постранично
                            </option>
                        </select>
                        <span>Показывать страницы с мертвыми хешами</span>
                        <select
                            className='app'
                            value={deadHashSelector}
                            onChange={(e) =>
                                setDeadHashSelector(e.target.value)
                            }
                        >
                            <option value='all'>Все</option>
                            <option value='without'>Кроме</option>
                            <option value='only'>Только</option>
                        </select>
                    </ContainerWidget>
                </div>
                <ContainerWidget
                    direction='row'
                    gap='small'
                >
                    <BookShortInfo
                        value={compareResult.data?.target}
                        covered_target={
                            compareResult.data?.target_covered_origin
                        }
                        covered_target_without_dead_hashes={
                            compareResult.data
                                ?.target_covered_origin_without_dead_hashes
                        }
                        book_size={compareResult.data?.target_book_size}
                        page_size={compareResult.data?.target_page_avg_size}
                    />
                    <BookImagePreviewWidget
                        flags={compareResult.data?.target.flags}
                        previewSize='superbig'
                        preview_url={compareResult.data?.target.preview_url}
                    />
                </ContainerWidget>
            </ContainerWidget>

            <ContainerWidget
                appContainer
                direction='row'
                gap='medium'
                style={{ justifyContent: 'space-between' }}
            >
                <ContainerWidget
                    direction='row'
                    gap='small'
                >
                    <b>Аттрибуты оригинала</b>
                    <BookAttributesWidget
                        value={compareResult.data?.origin_attributes}
                        colors={attributeColorListResponse.data?.colors}
                    />
                </ContainerWidget>
                <ContainerWidget
                    direction='row'
                    gap='small'
                >
                    <b>Аттрибуты общие</b>
                    <BookAttributesWidget
                        value={compareResult.data?.both_attributes}
                        colors={attributeColorListResponse.data?.colors}
                    />
                </ContainerWidget>
                <ContainerWidget
                    direction='row'
                    gap='small'
                >
                    <b>Аттрибуты цели</b>
                    <BookAttributesWidget
                        value={compareResult.data?.target_attributes}
                        colors={attributeColorListResponse.data?.colors}
                    />
                </ContainerWidget>
            </ContainerWidget>

            {/* FIXME: Переделать это непотребство */}
            {currentShow == 'origin' ? (
                <BookPagesPreviewWidget
                    bookID={originBookID}
                    pages={originPages}
                />
            ) : currentShow == 'both' ? (
                <BookPagesPreviewWidget
                    bookID={originBookID}
                    pages={bothPages}
                />
            ) : currentShow == 'target' ? (
                <BookPagesPreviewWidget
                    bookID={targetBookID}
                    pages={targetPages}
                />
            ) : currentShow == 'dual_read' && compareResult.data ? (
                <DualReaderWidget
                    aBookID={compareResult.data.origin.id}
                    bBookID={compareResult.data.target.id}
                    onCreateDeadHash={(
                        bookID: string,
                        page: BookSimplePage
                    ) => {
                        const bookName =
                            originBookID == bookID
                                ? originBookName
                                : targetBookID == bookID
                                  ? targetBookName
                                  : undefined
                        if (!bookName) {
                            return
                        }

                        if (
                            !confirm(
                                `Создать мертвых хеш для ${bookName} (${page.page_number})?`
                            )
                        ) {
                            return
                        }

                        doSetDeadHash({
                            book_id: bookID,
                            page_number: page.page_number,
                            value: true,
                        }).then(() =>
                            doCompare({
                                origin_book_id: originBookID,
                                target_book_id: targetBookID,
                            })
                        )
                    }}
                    onDeleteDeadHash={(
                        bookID: string,
                        page: BookSimplePage
                    ) => {
                        const bookName =
                            originBookID == bookID
                                ? originBookName
                                : targetBookID == bookID
                                  ? targetBookName
                                  : undefined
                        if (!bookName) {
                            return
                        }

                        if (
                            !confirm(
                                `Удалить мертвых хеш для ${bookName} (${page.page_number})?`
                            )
                        ) {
                            return
                        }

                        doSetDeadHash({
                            book_id: bookID,
                            page_number: page.page_number,
                            value: false,
                        }).then(() =>
                            doCompare({
                                origin_book_id: originBookID,
                                target_book_id: targetBookID,
                            })
                        )
                    }}
                    onDeletePage={(bookID: string, page: BookSimplePage) => {
                        const bookName =
                            originBookID == bookID
                                ? originBookName
                                : targetBookID == bookID
                                  ? targetBookName
                                  : undefined
                        if (!bookName) {
                            return
                        }

                        if (
                            !confirm(
                                `Удалить эту страницу ${bookName} (${page.page_number})? (ЭТО НЕОБРАТИМО)`
                            )
                        ) {
                            return
                        }

                        doDeletePage({
                            type: 'one',
                            book_id: bookID,
                            page_number: page.page_number,
                        }).then(() =>
                            doCompare({
                                origin_book_id: originBookID,
                                target_book_id: targetBookID,
                            })
                        )
                    }}
                    onDeleteAllPages={(
                        bookID: string,
                        page: BookSimplePage
                    ) => {
                        const bookName =
                            originBookID == bookID
                                ? originBookName
                                : targetBookID == bookID
                                  ? targetBookName
                                  : undefined
                        if (!bookName) {
                            return
                        }

                        if (
                            !confirm(
                                `Удалить такие страницы ${bookName} (${page.page_number})? (ЭТО НЕОБРАТИМО)`
                            )
                        ) {
                            return
                        }

                        const setDeadHash = confirm(
                            'Установить для страниц мертвый хеш?'
                        )

                        doDeletePage({
                            type: 'all_copy',
                            book_id: bookID,
                            page_number: page.page_number,
                            set_dead_hash: setDeadHash,
                        }).then(() =>
                            doCompare({
                                origin_book_id: originBookID,
                                target_book_id: targetBookID,
                            })
                        )
                    }}
                    onDeleteAllPagesWithDeadHash={(
                        bookID: string,
                        page: BookSimplePage
                    ) => {
                        if (
                            !confirm(
                                'Удалить такие страницы и добавить их в мертвый хеш? (ЭТО НЕОБРАТИМО)'
                            )
                        ) {
                            return
                        }

                        doDeletePage({
                            type: 'all_copy',
                            book_id: bookID,
                            page_number: page.page_number,
                            set_dead_hash: true,
                        }).then(() =>
                            doCompare({
                                origin_book_id: originBookID,
                                target_book_id: targetBookID,
                            })
                        )
                    }}
                    aPageCount={compareResult.data.origin.page_count}
                    bPageCount={compareResult.data.target.page_count}
                    aPages={originPages}
                    bPages={targetPages}
                />
            ) : null}
        </ContainerWidget>
    )
}

function BookShortInfo(props: {
    value?: BookSimple
    covered_target?: number
    covered_target_without_dead_hashes?: number
    book_size?: number
    page_size?: number
}) {
    if (!props.value) {
        return
    }

    const originDomain = /https?:\/\/([^\/]+)\/.*/.exec(
        props.value.origin_url ?? ''
    )?.[1]

    return (
        <ContainerWidget
            direction='column'
            gap='small'
            style={{ maxWidth: '500px' }}
        >
            <b style={{ wordBreak: 'break-all' }}>{props.value.name}</b>
            <span>
                Создана: <HumanTimeWidget value={props.value.created_at} />
            </span>
            <span>Страниц: {props.value.page_count}</span>
            {props.value.origin_url ? (
                <a href={props.value.origin_url}>Ссылка на первоисточник</a>
            ) : null}
            {originDomain ? <span>Домен: {originDomain}</span> : null}
            {props.covered_target != undefined &&
            props.covered_target_without_dead_hashes != undefined ? (
                <span title='Сколько страниц этой книги есть в другой'>
                    Покрытие: {prettyPercent(props.covered_target)}% (
                    {prettyPercent(props.covered_target_without_dead_hashes)}%)
                </span>
            ) : null}
            <PrettyDualSizeWidget
                first={props.book_size}
                second={props.page_size}
            >
                <span>Размер:</span>
            </PrettyDualSizeWidget>
            <Link
                className='app-button'
                to={BookDetailsLink(props.value.id)}
            >
                Страница книги
            </Link>
        </ContainerWidget>
    )
}
