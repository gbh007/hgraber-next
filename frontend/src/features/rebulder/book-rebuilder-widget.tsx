import { useState } from 'react'
import {
    AttributeColor,
    AttributeCountResponseAttribute,
} from '../../apiclient/api-attribute'
import { BookRebuildRequest } from '../../apiclient/api-book'
import { BookListResponse, BookShortInfo } from '../../apiclient/api-book-list'
import { LabelPresetListResponseLabel } from '../../apiclient/api-labels'
import { BookSimplePage } from '../../apiclient/model-book'
import { BookFilter } from '../../apiclient/model-book-filter'
import { BuilderFlagsWidget } from './builder-flags-widget'
import { BooksList } from './books-list'
import { BookPagesSelectWidget } from './book-pages-select-widget'
import { ContainerWidget } from '../../widgets/design-system'
import {
    BookAttributeInfoEditorWidget,
    BookFilterWidget,
    BookLabelInfoEditorWidget,
    BookMainInfoEditorWidget,
    BooksSimpleWidget,
    PaginatorWidget,
} from '../../widgets/book'

export function BookRebuilderWidget(props: {
    value: BookRebuildRequest
    onChange: (v: BookRebuildRequest) => void
    targetBookFilter: BookFilter
    targetBookFilterChange: (v: BookFilter) => void
    getTargetBooks: (data: BookFilter) => void // TODO: Сильно перегруженный виджет, подумать над возможностью его упрощения
    targetBookResponse?: BookListResponse
    pages?: Array<BookSimplePage>
    pageCount?: number
    labelsAutoComplete?: Array<LabelPresetListResponseLabel>
    attributeCount?: Array<AttributeCountResponseAttribute>
    attributeColors?: Array<AttributeColor>
}) {
    const [targetPreview, setTargetPreview] = useState<BookShortInfo>()

    return (
        <ContainerWidget
            direction='column'
            gap='big'
        >
            <ContainerWidget
                appContainer
                direction='column'
                gap='small'
            >
                <ContainerWidget
                    direction='row'
                    gap='small'
                >
                    <span>Выбрана целевой</span>
                    <input
                        className='app'
                        value={props.value.merge_with_book ?? ''}
                        onChange={(e) =>
                            props.onChange({
                                ...props.value,
                                merge_with_book: e.target.value || undefined,
                            })
                        }
                        placeholder='книга с которой надо соединить'
                    />
                    <button
                        className='app'
                        onClick={() =>
                            props.onChange({
                                ...props.value,
                                merge_with_book: undefined,
                            })
                        }
                    >
                        Сбросить
                    </button>
                </ContainerWidget>
                <BuilderFlagsWidget
                    value={props.value.flags}
                    onChange={(e) =>
                        props.onChange({ ...props.value, flags: e })
                    }
                />
            </ContainerWidget>

            {targetPreview &&
            targetPreview.info.id == props.value.merge_with_book ? (
                <ContainerWidget
                    appContainer
                    direction='column'
                    gap='small'
                >
                    <span>Данные будут внесены в</span>
                    <BooksSimpleWidget
                        value={targetPreview.info}
                        align='start'
                        previewSize='medium'
                    />
                </ContainerWidget>
            ) : (
                <details className='app'>
                    <summary>Выбрать целевую книгу</summary>
                    <ContainerWidget
                        direction='column'
                        gap='big'
                    >
                        <ContainerWidget
                            appContainer
                            direction='column'
                            gap='small'
                        >
                            <BookFilterWidget
                                value={props.targetBookFilter}
                                onChange={props.targetBookFilterChange}
                            />
                            <button
                                className='app'
                                onClick={() => {
                                    props.getTargetBooks({
                                        ...props.targetBookFilter,
                                        pagination: {
                                            ...props.targetBookFilter
                                                .pagination,
                                            page: 1,
                                        },
                                    })
                                }}
                            >
                                Применить
                            </button>
                            <PaginatorWidget
                                onChange={(v: number) => {
                                    props.targetBookFilterChange({
                                        ...props.targetBookFilter,
                                        pagination: {
                                            ...props.targetBookFilter
                                                .pagination,
                                            page: v,
                                        },
                                    })
                                    props.getTargetBooks({
                                        ...props.targetBookFilter,
                                        pagination: {
                                            ...props.targetBookFilter
                                                .pagination,
                                            page: v,
                                        },
                                    })
                                }}
                                value={props.targetBookResponse?.pages || []}
                            />
                        </ContainerWidget>
                        <BooksList
                            value={props.targetBookResponse?.books}
                            onChange={(e) => {
                                props.onChange({
                                    ...props.value,
                                    merge_with_book: e.info.id || undefined,
                                })
                                setTargetPreview(e)
                            }}
                            selected={props.value.merge_with_book}
                        />
                    </ContainerWidget>
                </details>
            )}
            {!props.value.merge_with_book ? (
                <BookMainInfoEditorWidget
                    value={props.value.old_book}
                    onChange={(e) =>
                        props.onChange({ ...props.value, old_book: e })
                    }
                />
            ) : null}
            <BookLabelInfoEditorWidget
                value={props.value.old_book.labels ?? []}
                onChange={(e) =>
                    props.onChange({
                        ...props.value,
                        old_book: { ...props.value.old_book, labels: e },
                    })
                }
                labelsAutoComplete={props.labelsAutoComplete}
            />
            <BookAttributeInfoEditorWidget
                value={props.value.old_book.attributes ?? []}
                onChange={(e) =>
                    props.onChange({
                        ...props.value,
                        old_book: { ...props.value.old_book, attributes: e },
                    })
                }
                attributeCount={props.attributeCount}
                colors={props.attributeColors}
            />
            <BookPagesSelectWidget
                value={props.value.selected_pages}
                onChange={(e) =>
                    props.onChange({ ...props.value, selected_pages: e })
                }
                bookID={props.value.old_book.id}
                pages={props.pages}
                pageCount={props.pageCount}
                pageOrder={props.value.page_order}
                onPageOrderChange={(e) =>
                    props.onChange({ ...props.value, page_order: e })
                }
                enablePageReOrder={props.value.flags?.page_re_order ?? false}
            />
        </ContainerWidget>
    )
}
