import { useCallback, useEffect, useState } from 'react'
import { BookSimplePage } from '../../apiclient/model-book'
import styles from './book-rebuilder.module.css'
import { ContainerWidget } from '../../widgets/design-system'
import { PageBadgesWidget } from '../../widgets/book'

export function PageSelectorReaderWidget(props: {
    value: Array<number>
    onChange: (v: Array<number>) => void
    bookID: string
    pageCount?: number
    pages?: Array<BookSimplePage>
}) {
    if (!props.pages) {
        return null
    }

    const [currentIndex, setCurrentIndex] = useState(0)
    const [currentPage, setCurrentPage] = useState<BookSimplePage>()

    useEffect(() => {
        setCurrentIndex(0)
    }, [props.bookID, props.pages.length])

    useEffect(() => {
        setCurrentPage(props.pages![currentIndex])
    }, [currentIndex])

    const prevPage = useCallback(() => {
        if (currentIndex == 0) return
        setCurrentIndex(currentIndex - 1)
    }, [props.pages, currentIndex])

    const nextPage = useCallback(() => {
        if (currentIndex == props.pages!.length - 1) return
        setCurrentIndex(currentIndex + 1)
    }, [props.pages, currentIndex])

    const goGo = useCallback(
        (event: React.MouseEvent<HTMLImageElement, MouseEvent>) => {
            const pos = event.currentTarget.getBoundingClientRect()
            const dx = (event.pageX - pos.left) / (pos.right - pos.left)
            if (dx < 0.3) {
                prevPage()
            } else {
                nextPage()
            }
        },
        [prevPage, nextPage]
    )

    return (
        <ContainerWidget
            direction='column'
            gap='medium'
        >
            <ContainerWidget appContainer>
                <div className={styles.pageViewActions}>
                    <span>
                        <button
                            className='app'
                            onClick={prevPage}
                        >
                            <span
                                className={styles.pageViewActionsPageNavigate}
                            >
                                {'<'}
                            </span>
                        </button>
                        <button
                            className='app'
                            onClick={nextPage}
                        >
                            <span
                                className={styles.pageViewActionsPageNavigate}
                            >
                                {'>'}
                            </span>
                        </button>
                    </span>
                    <PageBadgesWidget
                        flags={currentPage}
                        badgeSize='medium'
                    />
                    {currentPage ? (
                        <label>
                            <input
                                className='app'
                                type='checkbox'
                                checked={props.value.includes(
                                    currentPage!.page_number
                                )}
                                onChange={(e) =>
                                    props.onChange(
                                        e.target.checked
                                            ? [
                                                  ...props.value,
                                                  currentPage!.page_number,
                                              ]
                                            : props.value.filter(
                                                  (v) =>
                                                      v !=
                                                      currentPage!.page_number
                                              )
                                    )
                                }
                            />
                            выбрать
                        </label>
                    ) : null}
                    <span>
                        {`Страница ${currentPage?.page_number}` +
                            ` из ${props.pageCount ?? props.pages.length ?? 0}` +
                            (props.pageCount &&
                            props.pageCount != props.pages.length
                                ? ` (${currentIndex + 1}/${props.pages.length})`
                                : '')}
                    </span>
                </div>
            </ContainerWidget>
            <div className={styles.pageView}>
                {currentPage?.preview_url ? (
                    <img
                        src={currentPage?.preview_url}
                        style={{
                            // TODO: подумать что с таким делать
                            maxWidth: '100%',
                            maxHeight: '100%',
                        }}
                        onClick={goGo}
                    />
                ) : null}
                {props.value.includes(currentPage?.page_number ?? -1) ? (
                    <div
                        style={{
                            position: 'absolute',
                            top: '50%',
                            width: '100%',
                            textAlign: 'center',
                            fontSize: 60,
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
            </div>
        </ContainerWidget>
    )
}
