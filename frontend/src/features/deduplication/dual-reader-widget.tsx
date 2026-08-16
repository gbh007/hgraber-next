import { useCallback, useEffect, useState } from 'react'
import { BookSimplePage } from '../../apiclient/model-book'
import { ReaderWidget } from './reader-widget'
import { BookReadActionButtonWidget } from '../../widgets/book'
import {
    ColorizedTextWidget,
    ContainerWidget,
} from '../../widgets/design-system'

export function DualReaderWidget(props: {
    aBookID: string
    aPageCount?: number
    aPages?: Array<BookSimplePage>
    bBookID: string
    bPageCount?: number
    bPages?: Array<BookSimplePage>
    onCreateDeadHash: (bookID: string, page: BookSimplePage) => void
    onDeleteDeadHash: (bookID: string, page: BookSimplePage) => void
    onDeletePage: (bookID: string, page: BookSimplePage) => void
    onDeleteAllPages: (bookID: string, page: BookSimplePage) => void
    onDeleteAllPagesWithDeadHash: (bookID: string, page: BookSimplePage) => void
}) {
    const [currentAIndex, setCurrentAIndex] = useState(0)
    const [currentAPage, setCurrentAPage] = useState<BookSimplePage>()

    const [currentBIndex, setCurrentBIndex] = useState(0)
    const [currentBPage, setCurrentBPage] = useState<BookSimplePage>()

    useEffect(() => {
        setCurrentAIndex(0)
    }, [props.aBookID])

    useEffect(() => {
        setCurrentAPage(props.aPages![currentAIndex])
    }, [currentAIndex, props.aPages])

    useEffect(() => {
        setCurrentBIndex(0)
    }, [props.bBookID])

    useEffect(() => {
        setCurrentBPage(props.bPages![currentBIndex])
    }, [currentBIndex, props.bPages])

    const prevAPage = useCallback(() => {
        if (currentAIndex == 0) return
        setCurrentAIndex(currentAIndex - 1)
    }, [props.aPages, currentAIndex])

    const nextAPage = useCallback(() => {
        if (currentAIndex == props.aPages!.length - 1) return
        setCurrentAIndex(currentAIndex + 1)
    }, [props.aPages, currentAIndex])

    const goAPage = useCallback(
        (i: number) => {
            if (i < 1 || i >= props.aPages!.length - 1) return
            setCurrentAIndex(i - 1)
        },
        [props.aPages, currentAIndex]
    )

    const prevBPage = useCallback(() => {
        if (currentBIndex == 0) return
        setCurrentBIndex(currentBIndex - 1)
    }, [props.bPages, currentBIndex])

    const nextBPage = useCallback(() => {
        if (currentBIndex == props.bPages!.length - 1) return
        setCurrentBIndex(currentBIndex + 1)
    }, [props.bPages, currentBIndex])

    const goBPage = useCallback(
        (i: number) => {
            if (i < 1 || i >= props.bPages!.length - 1) return
            setCurrentBIndex(i - 1)
        },
        [props.bPages, currentBIndex]
    )

    const prevBothPage = useCallback(() => {
        if (currentBIndex == 0) return
        if (currentAIndex == 0) return
        setCurrentBIndex(currentBIndex - 1)
        setCurrentAIndex(currentAIndex - 1)
    }, [props.bPages, currentBIndex, props.aPages, currentAIndex])

    const nextBothPage = useCallback(() => {
        if (currentAIndex == props.aPages!.length - 1) return
        if (currentBIndex == props.bPages!.length - 1) return
        setCurrentBIndex(currentBIndex + 1)
        setCurrentAIndex(currentAIndex + 1)
    }, [props.bPages, currentBIndex, props.aPages, currentAIndex])

    return (
        <ContainerWidget
            direction='column'
            gap='small'
        >
            <ContainerWidget
                appContainer
                direction='row'
                gap='small'
            >
                <button
                    className='app'
                    onClick={prevBothPage}
                >
                    Пролистать назад
                </button>
                <button
                    className='app'
                    onClick={nextBothPage}
                >
                    Пролистать вперед
                </button>
            </ContainerWidget>
            <ContainerWidget
                direction='row'
                gap='small'
            >
                <ReaderWidget
                    bookID={props.aBookID}
                    currentIndex={currentAIndex}
                    nextPage={nextAPage}
                    prevPage={prevAPage}
                    currentPage={currentAPage}
                    pageCount={props.aPageCount}
                    pages={props.aPages}
                    goPage={goAPage}
                >
                    {currentAPage ? (
                        <>
                            <BookReadActionButtonWidget
                                bookID={props.aBookID}
                                pageNumber={currentAPage.page_number}
                                currentPage={currentAPage}
                                onCreateDeadHash={() =>
                                    props.onCreateDeadHash(
                                        props.aBookID,
                                        currentAPage
                                    )
                                }
                                onDeletePage={() =>
                                    props.onDeletePage(
                                        props.aBookID,
                                        currentAPage
                                    )
                                }
                                onDeleteAllPages={() =>
                                    props.onDeleteAllPages(
                                        props.aBookID,
                                        currentAPage
                                    )
                                }
                                onDeleteDeadHash={() =>
                                    props.onDeleteDeadHash(
                                        props.aBookID,
                                        currentAPage
                                    )
                                }
                            />
                            <button
                                className='app'
                                onClick={() =>
                                    props.onDeleteAllPagesWithDeadHash(
                                        props.aBookID,
                                        currentAPage
                                    )
                                }
                            >
                                <ColorizedTextWidget
                                    bold
                                    color='danger'
                                >
                                    Удалить такие страницы
                                </ColorizedTextWidget>
                            </button>
                            <span>
                                {/* это специальная заглушка для равномерного размещения компонентов */}
                            </span>
                        </>
                    ) : null}
                </ReaderWidget>
                <ReaderWidget
                    bookID={props.bBookID}
                    currentIndex={currentBIndex}
                    nextPage={nextBPage}
                    prevPage={prevBPage}
                    currentPage={currentBPage}
                    pageCount={props.bPageCount}
                    pages={props.bPages}
                    goPage={goBPage}
                >
                    {currentBPage ? (
                        <>
                            <BookReadActionButtonWidget
                                bookID={props.bBookID}
                                pageNumber={currentBPage.page_number}
                                currentPage={currentBPage}
                                onCreateDeadHash={() =>
                                    props.onCreateDeadHash(
                                        props.bBookID,
                                        currentBPage
                                    )
                                }
                                onDeletePage={() =>
                                    props.onDeletePage(
                                        props.bBookID,
                                        currentBPage
                                    )
                                }
                                onDeleteAllPages={() =>
                                    props.onDeleteAllPages(
                                        props.bBookID,
                                        currentBPage
                                    )
                                }
                                onDeleteDeadHash={() =>
                                    props.onDeleteDeadHash(
                                        props.bBookID,
                                        currentBPage
                                    )
                                }
                            />
                            <button
                                className='app'
                                onClick={() =>
                                    props.onDeleteAllPagesWithDeadHash(
                                        props.bBookID,
                                        currentBPage
                                    )
                                }
                            >
                                <ColorizedTextWidget
                                    bold
                                    color='danger'
                                >
                                    Удалить такие страницы
                                </ColorizedTextWidget>
                            </button>
                            <span>
                                {/* это специальная заглушка для равномерного размещения компонентов */}
                            </span>
                        </>
                    ) : null}
                </ReaderWidget>
            </ContainerWidget>
        </ContainerWidget>
    )
}
