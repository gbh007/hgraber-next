import { useState } from 'react'
import { BookSimplePage } from '../../apiclient/model-book'
import styles from './book-rebuilder.module.css'
import { Link } from 'react-router-dom'
import { BookReaderLink } from '../../core/routing'
import { ContainerWidget } from '../../widgets/design-system'
import { ImageSize, PageImagePreviewWidget } from '../../widgets/book'

export function PageListPreview(props: {
    value: Array<number>
    onChange: (v: Array<number>) => void
    bookID: string
    pages?: Array<BookSimplePage>
    previewSize: ImageSize
    pageDragAndDrop?: (a: number, b: number) => void
    enablePageReOrder: boolean
    imageColumns?: number
    pageOrder: Array<number>
}) {
    const fontSize =
        props.previewSize == 'superbig'
            ? 60
            : props.previewSize == 'big'
              ? 30
              : props.previewSize == 'medium'
                ? 15
                : 10

    const [aIndex, setAIndex] = useState(-1)
    const [bIndex, setBIndex] = useState(-1)

    const scrollToTop = () => {
        document.getElementById('pageSelector')?.scrollIntoView({
            behavior: 'smooth',
        })
    }

    const imagePreviews = props.pages?.map((page, index) => (
        <ContainerWidget
            appContainer
            direction='column'
            key={page.page_number}
            style={{
                borderLeft:
                    bIndex == index
                        ? '10px dashed var(--app-color)'
                        : undefined,
                flexGrow: 1,
                alignItems: 'center',
            }}
        >
            {page.preview_url ? (
                <PageImagePreviewWidget
                    previewSize={props.previewSize}
                    flags={page}
                    preview_url={page.preview_url}
                    onClick={() => {
                        const checked = props.value.includes(page.page_number)
                        props.onChange(
                            !checked
                                ? [...props.value, page.page_number]
                                : props.value.filter(
                                      (v) => v != page.page_number
                                  )
                        )
                    }}
                >
                    {props.previewSize != 'small' &&
                    props.previewSize != 'medium' &&
                    props.value.includes(page.page_number) ? (
                        <div
                            style={{
                                position: 'absolute',
                                top: '50%',
                                width: 'calc(100% - 40px)',
                                textAlign: 'center',
                                fontSize: fontSize,
                                color: 'var(--app-color-good)',
                                fontWeight: 'bolder',
                                background: 'var(--app-background)',
                                padding: '20px',
                                borderRadius: '10px',
                            }}
                        >
                            Выбрана
                        </div>
                    ) : null}
                </PageImagePreviewWidget>
            ) : null}
            <span
                style={
                    props.enablePageReOrder
                        ? {
                              cursor: 'grab',
                              userSelect: 'none',
                          }
                        : undefined
                }
                draggable={props.enablePageReOrder ? 'true' : 'false'}
                onDragStart={() => {
                    setAIndex(index)
                }}
                onDragOver={() => {
                    setBIndex(index)
                }}
                onDragEnd={() => {
                    props.pageDragAndDrop?.(aIndex, bIndex)
                    setBIndex(-1)
                }}
            >
                Страница: {page.page_number}
                {props.enablePageReOrder &&
                props.pageOrder.findIndex((v) => v == page.page_number) + 1 !=
                    page.page_number
                    ? ` (${props.pageOrder.findIndex((v) => v == page.page_number) + 1})`
                    : ''}
            </span>
            <label>
                <input
                    className='app'
                    type='checkbox'
                    checked={props.value.includes(page.page_number)}
                    onChange={(e) =>
                        props.onChange(
                            e.target.checked
                                ? [...props.value, page.page_number]
                                : props.value.filter(
                                      (v) => v != page.page_number
                                  )
                        )
                    }
                />
                выбрать
            </label>
            <Link
                className='app-button'
                to={BookReaderLink(props.bookID, page.page_number)}
            >
                Открыть в читалке
            </Link>
        </ContainerWidget>
    ))

    if (props.imageColumns) {
        return (
            <>
                <ContainerWidget
                    direction='columns'
                    gap='medium'
                    columns={props.imageColumns}
                >
                    {imagePreviews}
                </ContainerWidget>
                <button
                    className='app'
                    onClick={scrollToTop}
                >
                    Наверх
                </button>
            </>
        )
    }

    return (
        <>
            <div className={styles.preview}>{imagePreviews}</div>
            <button
                className='app'
                onClick={scrollToTop}
            >
                Наверх
            </button>
        </>
    )
}
