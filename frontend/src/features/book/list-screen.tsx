import { Link, useSearchParams } from 'react-router-dom'
import { useAppSettings } from '../../apiclient/settings'
import styles from './list-screen.module.css'
import { BookFilter } from '../../apiclient/model-book-filter'
import { useEffect, useState } from 'react'
import {
    useAttributeColorList,
    useAttributeCount,
} from '../../apiclient/api-attribute'
import { useLabelPresetList } from '../../apiclient/api-labels'
import { useBookList } from '../../apiclient/api-book-list'
import { SelectToCompareLink } from '../../core/routing'
import { AgentExportWidget } from './agent-export-widget'
import { BookShortInfoWidget } from './book-short-info-widget'
import {
    ColorizedTextWidget,
    ContainerWidget,
    ErrorTextWidget,
} from '../../widgets/design-system'
import { BookFilterWidget, PaginatorWidget } from '../../widgets/book'

export function ListScreen() {
    const [settings, _] = useAppSettings()
    const [searchParams, setSearchParams] = useSearchParams()

    const defaultFilterValue: BookFilter = {
        pagination: { count: settings.book_on_page, page: 1 },
        filter: {
            flags: {
                delete_status: 'except',
                download_status: 'only',
                verify_status: 'only',
            },
        },
        sort: { field: 'created_at', desc: true },
    }

    const [bookFilter, setBookFilter] = useState<BookFilter>(defaultFilterValue)

    const [attributeCountResponse, getAttributeCount] = useAttributeCount()
    useEffect(() => {
        getAttributeCount()
    }, [getAttributeCount])

    const [labelPresetsResponse, fetchLabelPresets] = useLabelPresetList()
    useEffect(() => {
        fetchLabelPresets()
    }, [fetchLabelPresets])

    const [attributeColorListResponse, fetchAttributeColorList] =
        useAttributeColorList()
    useEffect(() => {
        fetchAttributeColorList()
    }, [fetchAttributeColorList])

    const [booksResponse, getBooks] = useBookList()

    useEffect(() => {
        try {
            const filter = JSON.parse(searchParams.get('filter')!)

            if (!filter) {
                setBookFilter(defaultFilterValue)
                getBooks(defaultFilterValue)
            } else {
                setBookFilter({ ...bookFilter, ...filter })
                getBooks({ ...bookFilter, ...filter })
            }
        } catch (err) {
            console.log(err)
        }
    }, [setBookFilter, searchParams])

    return (
        <ContainerWidget
            direction='column'
            gap='medium'
        >
            <ErrorTextWidget value={booksResponse} />
            <ErrorTextWidget value={attributeCountResponse} />
            <ErrorTextWidget value={labelPresetsResponse} />
            <ErrorTextWidget value={attributeColorListResponse} />
            <ContainerWidget
                appContainer
                direction='column'
                gap='medium'
            >
                <details className='app'>
                    <summary>
                        Фильтр, всего {booksResponse.data?.count || 0}
                    </summary>
                    <ContainerWidget
                        direction='column'
                        gap='medium'
                    >
                        <BookFilterWidget
                            value={bookFilter}
                            onChange={setBookFilter}
                            attributeCount={
                                attributeCountResponse.data?.attributes
                            }
                            labelsAutoComplete={
                                labelPresetsResponse.data?.presets
                            }
                        />
                        <ContainerWidget
                            direction='row'
                            gap='medium'
                        >
                            <button
                                className='app'
                                onClick={() => {
                                    setBookFilter({
                                        ...bookFilter,
                                        pagination: {
                                            ...bookFilter.pagination,
                                            page: 1,
                                        },
                                    })
                                    searchParams.set(
                                        'filter',
                                        JSON.stringify({
                                            ...bookFilter,
                                            pagination: {
                                                ...bookFilter.pagination,
                                                page: 1,
                                            },
                                        })
                                    )
                                    setSearchParams(searchParams)
                                }}
                            >
                                Применить фильтр
                            </button>
                            <button
                                className='app'
                                onClick={() => {
                                    setBookFilter(defaultFilterValue)
                                    searchParams.delete('filter')
                                    setSearchParams(searchParams)
                                }}
                            >
                                <ColorizedTextWidget color='danger'>
                                    Очистить фильтр
                                </ColorizedTextWidget>
                            </button>
                            <button
                                className='app'
                                onClick={() => {
                                    getBooks(bookFilter)
                                }}
                            >
                                Обновить данные
                            </button>
                            <Link
                                className='app-button'
                                to={SelectToCompareLink(bookFilter)}
                            >
                                Перейти в выбор для сравнения
                            </Link>
                        </ContainerWidget>
                        <AgentExportWidget filter={bookFilter} />
                    </ContainerWidget>
                </details>
            </ContainerWidget>
            <PaginatorWidget
                onChange={(v: number) => {
                    setBookFilter({
                        ...bookFilter,
                        pagination: { ...bookFilter.pagination, page: v },
                    })
                    searchParams.set(
                        'filter',
                        JSON.stringify({
                            ...bookFilter,
                            pagination: { ...bookFilter.pagination, page: v },
                        })
                    )
                    setSearchParams(searchParams)
                }}
                value={booksResponse.data?.pages || []}
            />
            <div className={styles.bookList}>
                {booksResponse.data?.books?.map((book) => (
                    <BookShortInfoWidget
                        key={book.info.id}
                        value={book}
                        colors={attributeColorListResponse.data?.colors}
                    />
                ))}
            </div>
            <PaginatorWidget
                onChange={(v: number) => {
                    setBookFilter({
                        ...bookFilter,
                        pagination: { ...bookFilter.pagination, page: v },
                    })
                    searchParams.set(
                        'filter',
                        JSON.stringify({
                            ...bookFilter,
                            pagination: { ...bookFilter.pagination, page: v },
                        })
                    )
                    setSearchParams(searchParams)
                }}
                value={booksResponse.data?.pages || []}
            />
        </ContainerWidget>
    )
}
