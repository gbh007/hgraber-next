import { useEffect, useState } from 'react'
import { Link, useParams } from 'react-router-dom'
import {
    BookRebuildRequest,
    useBookRaw,
    useBookRebuild,
} from '../../apiclient/api-book'
import { BookFilter } from '../../apiclient/model-book-filter'
import { useBookList } from '../../apiclient/api-book-list'
import { useBookDetails } from '../../apiclient/api-book-details'
import { useLabelPresetList } from '../../apiclient/api-labels'
import {
    useAttributeColorList,
    useAttributeOriginCount,
} from '../../apiclient/api-attribute'
import { BookDetailsLink } from '../../core/routing'
import { BookRebuilderWidget } from './book-rebuilder-widget'
import { ContainerWidget, ErrorTextWidget } from '../../widgets/design-system'

export function BookRebuilderScreen() {
    const params = useParams()
    const bookID = params.bookID!

    const [bookRebuildData, setBookRebuildData] = useState<BookRebuildRequest>({
        old_book: {
            create_at: new Date().toJSON(),
            id: '',
            name: '',
            page_count: 0,
        },
        flags: {
            set_origin_labels: true,
            auto_verify: true,
        },
        selected_pages: [],
        page_order: [],
    })

    const [bookFilter, setBookFilter] = useState<BookFilter>({
        pagination: {
            count: 10,
            page: 1,
        },
        filter: {
            flags: {
                delete_status: 'except',
                download_status: 'only',
                verify_status: 'only',
                show_rebuilded: 'only',
            },
        },
        sort: {
            field: 'created_at',
            desc: true,
        },
    })

    const [bookRawResponse, fetchBookRaw] = useBookRaw()
    const [bookRebuildResponse, doBookRebuild] = useBookRebuild()
    const [booksResponse, getBooks] = useBookList()

    // TODO: сделать другой запрос, который возвращает только нужные данные страниц
    const [bookDetailsResponse, getBookDetails] = useBookDetails()
    const [labelPresetsResponse, fetchLabelPresets] = useLabelPresetList()
    const [attributeOriginCountResponse, getAttributeOriginCount] =
        useAttributeOriginCount()

    const [attributeColorListResponse, fetchAttributeColorList] =
        useAttributeColorList()

    useEffect(() => {
        fetchBookRaw({ id: bookID })
    }, [fetchBookRaw, bookID])

    useEffect(() => {
        getBookDetails({ id: bookID })
    }, [getBookDetails, bookID])

    useEffect(() => {
        getBooks(bookFilter)
    }, [getBooks])
    useEffect(() => {
        fetchLabelPresets()
    }, [fetchLabelPresets])
    useEffect(() => {
        getAttributeOriginCount()
    }, [getAttributeOriginCount])

    useEffect(() => {
        if (bookRawResponse.data) {
            setBookRebuildData({
                ...bookRebuildData,
                old_book: {
                    ...bookRawResponse.data,
                    name: 'Rebuild: ' + bookRawResponse.data.name,
                },
                selected_pages: [],
                page_order:
                    bookRawResponse.data.pages?.map((e) => e.page_number) ?? [],
            })
        }
    }, [bookRawResponse.data])

    useEffect(() => {
        fetchAttributeColorList()
    }, [fetchAttributeColorList])

    return (
        <ContainerWidget
            direction='column'
            gap='medium'
        >
            <ErrorTextWidget value={bookRawResponse} />
            <ErrorTextWidget value={bookDetailsResponse} />
            <ErrorTextWidget value={labelPresetsResponse} />
            <ErrorTextWidget value={attributeOriginCountResponse} />
            <ErrorTextWidget value={bookRebuildResponse} />
            <ErrorTextWidget value={booksResponse} />
            <ErrorTextWidget value={attributeColorListResponse} />
            <ContainerWidget
                direction='row'
                gap='small'
            >
                <Link
                    className='app-button'
                    to={BookDetailsLink(bookID)}
                >
                    На страницу исходной книги
                </Link>
                {bookRebuildResponse.data?.id ? (
                    <Link
                        className='app-button'
                        to={BookDetailsLink(bookRebuildResponse.data.id)}
                    >
                        На страницу собранной книги
                    </Link>
                ) : (
                    <button
                        className='app'
                        onClick={() => {
                            if (
                                bookRebuildData.flags?.extract_mode &&
                                !confirm(
                                    'Удалить все страницы из этой книги (экстракция) что вошли в ребилд? (ЭТО НЕОБРАТИМО)'
                                )
                            ) {
                                return
                            }

                            if (
                                bookRebuildData.flags
                                    ?.mark_unused_pages_as_dead_hash &&
                                !confirm(
                                    'Создать мертвых хеш для всех страницы из этой книги что не вошли в ребилд?'
                                )
                            ) {
                                return
                            }

                            if (
                                bookRebuildData.flags
                                    ?.mark_unused_pages_as_deleted &&
                                !confirm(
                                    'Удалить все страницы из этой книги и их копии что не вошли в ребилд? (ЭТО НЕОБРАТИМО)'
                                )
                            ) {
                                return
                            }

                            if (
                                bookRebuildData.flags
                                    ?.mark_empty_book_as_deleted_after_remove_pages &&
                                !confirm(
                                    'Удалить все книги что останутся без страниц после ребилда? (ЭТО НЕОБРАТИМО)'
                                )
                            ) {
                                return
                            }

                            doBookRebuild(bookRebuildData)
                        }}
                    >
                        Пересобрать
                    </button>
                )}
            </ContainerWidget>
            {!bookRebuildResponse.data?.id ? (
                <BookRebuilderWidget
                    value={bookRebuildData}
                    onChange={(e) => setBookRebuildData(e)}
                    labelsAutoComplete={labelPresetsResponse.data?.presets}
                    attributeCount={
                        attributeOriginCountResponse.data?.attributes
                    }
                    pages={bookDetailsResponse.data?.pages}
                    pageCount={bookDetailsResponse.data?.info.page_count}
                    targetBookFilter={bookFilter}
                    targetBookFilterChange={setBookFilter}
                    getTargetBooks={(e) => getBooks(e)}
                    targetBookResponse={booksResponse.data ?? undefined}
                    attributeColors={attributeColorListResponse.data?.colors}
                />
            ) : null}
        </ContainerWidget>
    )
}
