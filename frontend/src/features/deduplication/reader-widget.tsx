import { PropsWithChildren, useCallback, useState } from 'react'
import styles from './reader-widget.module.css'
import { BookSimplePage } from '../../apiclient/model-book'
import { PageBadgesWidget } from '../../widgets/book'
import { ContainerWidget } from '../../widgets/design-system'

export function ReaderWidget(
    props: PropsWithChildren & {
        bookID: string
        pageCount?: number
        pages?: Array<BookSimplePage>
        currentPage?: BookSimplePage
        currentIndex: number
        prevPage: () => void
        nextPage: () => void
        goPage: (i: number) => void
    }
) {
    const [goToIndex, setGoToIndex] = useState(0)

    if (!props.pages) {
        return null
    }

    const goGo = useCallback(
        (event: React.MouseEvent<HTMLImageElement, MouseEvent>) => {
            const pos = event.currentTarget.getBoundingClientRect()
            const dx = (event.pageX - pos.left) / (pos.right - pos.left)
            if (dx < 0.3) {
                props.prevPage()
            } else {
                props.nextPage()
            }
        },
        [props.prevPage, props.nextPage]
    )

    return (
        <ContainerWidget
            direction='column'
            gap='medium'
        >
            <ContainerWidget
                appContainer
                direction='column'
                gap='medium'
            >
                <ContainerWidget
                    direction='row'
                    gap='small'
                >
                    <input
                        className='app'
                        type='number'
                        value={goToIndex}
                        onChange={(e) => setGoToIndex(e.target.valueAsNumber)}
                    />
                    <button
                        className='app'
                        onClick={() => props.goPage(goToIndex)}
                    >
                        Перейти
                    </button>
                </ContainerWidget>
                <div className={styles.pageViewActions}>
                    <span>
                        <button
                            className='app'
                            onClick={props.prevPage}
                        >
                            <span
                                className={styles.pageViewActionsPageNavigate}
                            >
                                {'<'}
                            </span>
                        </button>
                        <button
                            className='app'
                            onClick={props.nextPage}
                        >
                            <span
                                className={styles.pageViewActionsPageNavigate}
                            >
                                {'>'}
                            </span>
                        </button>
                    </span>
                    <PageBadgesWidget
                        flags={props.currentPage}
                        badgeSize='medium'
                    />
                    <span>
                        {`Страница ${props.currentPage?.page_number}` +
                            ` из ${props.pageCount ?? props.pages.length ?? 0}` +
                            (props.pageCount &&
                            props.pageCount != props.pages.length
                                ? ` (${props.currentIndex + 1}/${props.pages.length})`
                                : '')}
                    </span>
                </div>
            </ContainerWidget>
            <div className={styles.pageView}>
                {props.currentPage?.preview_url ? (
                    <img
                        src={props.currentPage?.preview_url}
                        style={{
                            // TODO: подумать что с таким делать
                            maxWidth: '100%',
                            maxHeight: '100%',
                        }}
                        onClick={goGo}
                    />
                ) : null}
            </div>
            <ContainerWidget
                appContainer
                direction='column'
                gap='medium'
            >
                {props.children ? (
                    <div className={styles.pageViewActions}>
                        {props.children}
                    </div>
                ) : null}
            </ContainerWidget>
        </ContainerWidget>
    )
}
