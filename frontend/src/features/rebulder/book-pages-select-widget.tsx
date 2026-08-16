import { useState } from 'react'
import { BookSimplePage } from '../../apiclient/model-book'
import { PageSelectorReaderWidget } from './page-selector-reader-widget'
import { PageListPreview } from './page-list-preview'
import { ContainerWidget } from '../../widgets/design-system'
import { ImageSize, PreviewSizeWidget } from '../../widgets/book'

export function BookPagesSelectWidget(props: {
    value: Array<number>
    onChange: (v: Array<number>) => void
    bookID: string
    pages?: Array<BookSimplePage>
    pageCount?: number
    pageOrder: Array<number>
    onPageOrderChange: (v: Array<number>) => void
    enablePageReOrder: boolean
}) {
    const [viewMode, setViewMode] = useState('reader')
    const [showOnlySelected, setShowOnlySelected] = useState(false)
    const [showExceptSelected, setShowExceptSelected] = useState(false)
    const [showDeadHash, setShowDeadHash] = useState(true)
    const [imageOnRow, setImageOnRow] = useState(4)

    const [imageSize, setImageSize] = useState<ImageSize>('medium')

    const pages = props.pages
        ?.filter(
            (page) =>
                !showOnlySelected || props.value.includes(page.page_number)
        )
        .filter(
            (page) =>
                !showExceptSelected || !props.value.includes(page.page_number)
        )
        .filter((page) => showDeadHash || !page.has_dead_hash)
        .sort(
            (a: BookSimplePage, b: BookSimplePage) =>
                props.pageOrder.findIndex((v) => v == a.page_number) -
                props.pageOrder.findIndex((v) => v == b.page_number)
        )

    return (
        <ContainerWidget
            direction='column'
            gap='medium'
            id='pageSelector'
        >
            <ContainerWidget
                appContainer
                direction='column'
                gap='medium'
            >
                <ContainerWidget
                    direction='row'
                    gap='medium'
                >
                    <button
                        className='app'
                        onClick={() =>
                            props.onChange(
                                props.pages?.map((page) => page.page_number) ??
                                    []
                            )
                        }
                    >
                        Выбрать все
                    </button>
                    <button
                        className='app'
                        onClick={() => props.onChange([])}
                    >
                        Снять все
                    </button>
                    {props.enablePageReOrder ? (
                        <button
                            className='app'
                            onClick={() => {
                                if (!props.value) {
                                    return
                                }

                                const firstIndex = props.pageOrder.findIndex(
                                    (v) => props.value.includes(v)
                                )
                                if (firstIndex < 0) {
                                    return
                                }

                                let newPageOrder = props.pageOrder.filter(
                                    (_, i) => i < firstIndex
                                )
                                newPageOrder.push(
                                    ...props.pageOrder.filter((v) =>
                                        props.value.includes(v)
                                    )
                                )
                                newPageOrder.push(
                                    ...props.pageOrder.filter(
                                        (v, i) =>
                                            i > firstIndex &&
                                            !props.value.includes(v)
                                    )
                                )

                                props.onPageOrderChange?.(newPageOrder)
                            }}
                        >
                            Скомпоновать выделенные по порядку
                        </button>
                    ) : null}
                    <span>
                        Выбрано {props.value.length} из {props.pages?.length}
                    </span>
                </ContainerWidget>
                <ContainerWidget
                    direction='row'
                    gap='medium'
                >
                    <select
                        className='app'
                        value={viewMode}
                        onChange={(e) => setViewMode(e.target.value)}
                    >
                        <option value='reader'>
                            Режим просмотра: постранично
                        </option>
                        <option value='list'>
                            Режим просмотра: все страницы
                        </option>
                        <option value='columns'>
                            Режим просмотра: колонки
                        </option>
                    </select>
                    {viewMode == 'list' || viewMode == 'columns' ? (
                        <PreviewSizeWidget
                            value={imageSize}
                            onChange={setImageSize}
                        />
                    ) : null}
                    {viewMode == 'columns' ? (
                        <input
                            className='app'
                            value={imageOnRow}
                            type='number'
                            onChange={(e) =>
                                setImageOnRow(e.target.valueAsNumber)
                            }
                        />
                    ) : null}
                </ContainerWidget>
                <ContainerWidget
                    direction='column'
                    gap='medium'
                >
                    <label>
                        <input
                            type='checkbox'
                            className='app'
                            checked={showOnlySelected}
                            onChange={(e) =>
                                setShowOnlySelected(e.target.checked)
                            }
                        />
                        Только выбранные страницы
                    </label>
                    <label>
                        <input
                            type='checkbox'
                            className='app'
                            checked={showExceptSelected}
                            onChange={(e) =>
                                setShowExceptSelected(e.target.checked)
                            }
                        />
                        Только НЕ выбранные страницы
                    </label>
                    <label>
                        <input
                            type='checkbox'
                            className='app'
                            checked={showDeadHash}
                            onChange={(e) => setShowDeadHash(e.target.checked)}
                        />
                        Показывать с мертвым хешом
                    </label>
                </ContainerWidget>
            </ContainerWidget>
            {!pages?.length ? null : viewMode == 'reader' ? (
                <PageSelectorReaderWidget
                    bookID={props.bookID}
                    onChange={props.onChange}
                    value={props.value}
                    pages={pages}
                    pageCount={props.pageCount}
                />
            ) : viewMode == 'list' || viewMode == 'columns' ? (
                <PageListPreview
                    bookID={props.bookID}
                    onChange={props.onChange}
                    value={props.value}
                    pages={pages}
                    pageOrder={props.pageOrder}
                    previewSize={imageSize}
                    imageColumns={
                        viewMode == 'columns' ? imageOnRow : undefined
                    }
                    pageDragAndDrop={(a, b) => {
                        const valueA = pages[a]
                        const valueB = pages[b]

                        const targetIndex = props.pageOrder.findIndex(
                            (v) => v == valueA.page_number
                        )
                        const beforeIndex = props.pageOrder.findIndex(
                            (v) => v == valueB.page_number
                        )

                        props.onPageOrderChange?.(
                            movePageBefore(
                                targetIndex,
                                beforeIndex,
                                props.pageOrder
                            )
                        )
                    }}
                    enablePageReOrder={props.enablePageReOrder}
                />
            ) : null}
        </ContainerWidget>
    )
}

function movePageBefore(
    targetIndex: number,
    beforeIndex: number,
    pages?: Array<number>
): Array<number> {
    if (!pages) {
        return []
    }

    if (targetIndex < 0 || targetIndex > pages.length) {
        return pages
    }

    if (beforeIndex < 0 || beforeIndex > pages.length) {
        return pages
    }

    return [
        ...pages.filter((_, i) => i != targetIndex && i < beforeIndex),
        ...pages.filter((_, i) => i == targetIndex),
        ...pages.filter((_, i) => i != targetIndex && i >= beforeIndex),
    ]
}
