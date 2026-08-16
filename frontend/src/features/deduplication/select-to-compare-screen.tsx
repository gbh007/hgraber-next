import { Link, useSearchParams } from 'react-router-dom'
import { BookShortInfo, useBookList } from '../../apiclient/api-book-list'
import { useAppSettings } from '../../apiclient/settings'
import styles from './select-to-compare-screen.module.css'
import { BookFilter } from '../../apiclient/model-book-filter'
import { useEffect, useState } from 'react'
import { useAttributeCount } from '../../apiclient/api-attribute'
import { useLabelPresetList } from '../../apiclient/api-labels'
import { BookCompareLink, BookListLink } from '../../core/routing'
import {
    ColorizedTextWidget,
    ContainerWidget,
    ErrorTextWidget,
} from '../../widgets/design-system'
import {
    BookFilterWidget,
    BooksSimpleWidget,
    PaginatorWidget,
} from '../../widgets/book'

export function SelectToCompareScreen() {
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

    const [attributeCountResponse, getAttributeCount] = useAttributeCount()
    useEffect(() => {
        getAttributeCount()
    }, [getAttributeCount])

    const [labelPresetsResponse, fetchLabelPresets] = useLabelPresetList()
    useEffect(() => {
        fetchLabelPresets()
    }, [fetchLabelPresets])

    const [originPreview, setOriginPreview] = useState<BookShortInfo>()
    const [targetPreview, setTargetPreview] = useState<BookShortInfo>()

    return (
        <ContainerWidget
            direction='column'
            gap='big'
        >
            <ErrorTextWidget value={booksResponse} />
            <ErrorTextWidget value={attributeCountResponse} />
            <ErrorTextWidget value={labelPresetsResponse} />
            <ContainerWidget
                appContainer
                direction='column'
                gap='big'
                style={{
                    position: 'sticky',
                    top: '0px',
                    zIndex: '1',
                }}
            >
                <ContainerWidget
                    direction='row'
                    gap='big'
                >
                    {originPreview ? (
                        <ContainerWidget
                            direction='column'
                            gap='small'
                        >
                            <h3>Оригинальная книга</h3>
                            <BooksSimpleWidget
                                value={originPreview.info}
                                align='start'
                            />
                            <button
                                className='app'
                                onClick={() => setOriginPreview(undefined)}
                            >
                                Сбросить
                            </button>
                        </ContainerWidget>
                    ) : null}
                    {targetPreview ? (
                        <ContainerWidget
                            direction='column'
                            gap='small'
                        >
                            <h3>Целевая книга</h3>
                            <BooksSimpleWidget
                                value={targetPreview.info}
                                align='start'
                            />
                            <button
                                className='app'
                                onClick={() => setTargetPreview(undefined)}
                            >
                                Сбросить
                            </button>
                        </ContainerWidget>
                    ) : null}
                </ContainerWidget>
                <ContainerWidget
                    direction='row'
                    gap='medium'
                >
                    <button
                        className='app'
                        onClick={() => {
                            setOriginPreview(undefined)
                            setTargetPreview(undefined)
                        }}
                    >
                        Сбросить все
                    </button>
                    {originPreview && targetPreview ? (
                        <Link
                            className='app-button'
                            to={BookCompareLink(
                                originPreview.info.id,
                                targetPreview.info.id
                            )}
                        >
                            Сравнить
                        </Link>
                    ) : null}
                </ContainerWidget>
            </ContainerWidget>
            <ContainerWidget
                appContainer
                direction='column'
                gap='medium'
            >
                <BookFilterWidget
                    value={bookFilter}
                    onChange={setBookFilter}
                    attributeCount={attributeCountResponse.data?.attributes}
                    labelsAutoComplete={labelPresetsResponse.data?.presets}
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
                        to={BookListLink(bookFilter)}
                    >
                        Перейти в список книг
                    </Link>
                </ContainerWidget>
                <span>Всего: {booksResponse.data?.count}</span>
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
                                pagination: {
                                    ...bookFilter.pagination,
                                    page: v,
                                },
                            })
                        )
                        setSearchParams(searchParams)
                    }}
                    value={booksResponse.data?.pages || []}
                />
            </ContainerWidget>
            <BooksList
                value={booksResponse.data?.books}
                onChangeOrigin={setOriginPreview}
                onChangeTarget={setTargetPreview}
                selectedOrigin={originPreview?.info.id}
                selectedTarget={targetPreview?.info.id}
            />
        </ContainerWidget>
    )
}

function BooksList(props: {
    value?: Array<BookShortInfo>
    selectedOrigin?: string
    selectedTarget?: string
    onChangeTarget: (v: BookShortInfo) => void
    onChangeOrigin: (v: BookShortInfo) => void
}) {
    return (
        <div className={styles.preview}>
            {props.value?.map((book) => (
                <BooksSimpleWidget
                    value={book.info}
                    key={book.info.id}
                >
                    <button
                        className='app'
                        onClick={() => props.onChangeOrigin(book)}
                        disabled={book.info.id == props.selectedOrigin}
                    >
                        Выбрать как оригинальную
                    </button>
                    <button
                        className='app'
                        onClick={() => props.onChangeTarget(book)}
                        disabled={book.info.id == props.selectedTarget}
                    >
                        Выбрать как целевую
                    </button>
                </BooksSimpleWidget>
            ))}
        </div>
    )
}
