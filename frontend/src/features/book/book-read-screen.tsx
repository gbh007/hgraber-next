import { Link, useNavigate, useParams } from 'react-router-dom'
import styles from './book-read-screen.module.css'
import { useCallback, useEffect, useState } from 'react'
import { BookSimplePage } from '../../apiclient/model-book'
import { useBookDetails } from '../../apiclient/api-book-details'
import { useSetDeadHash } from '../../apiclient/api-deduplicate'
import { useDeleteBookPage } from '../../apiclient/api-book-delete'
import { BookDetailsLink, BookReaderLink } from '../../core/routing'
import {
    ColorizedTextWidget,
    ErrorTextWidget,
} from '../../widgets/design-system'
import {
    BookReadActionButtonWidget,
    PageBadgesWidget,
} from '../../widgets/book'

export function BookReadScreen() {
    const params = useParams()
    const bookID = params.bookID!
    const pageNumber = parseInt(params.page!)

    const [currentPage, setCurrentPage] = useState<BookSimplePage>()
    const [showPageOnlyWithPreview, setShowPageOnlyWithPreview] =
        useState(false)
    const [showPageWithDeadHash, setShowPageWithDeadHash] = useState(true)

    const [bookDetailsResponse, getBookDetails] = useBookDetails()
    const [setDeadHashResponse, doSetDeadHash] = useSetDeadHash()
    const [deletePageResponse, doDeletePage] = useDeleteBookPage()

    const navigate = useNavigate()

    useEffect(() => {
        getBookDetails({ id: bookID })
    }, [getBookDetails, bookID])

    useEffect(() => {
        // TODO: заменить на вычисляемую функцию
        setCurrentPage(
            bookDetailsResponse.data?.pages
                ?.filter((page) => page.page_number == pageNumber)
                .pop()
        )
    }, [bookDetailsResponse.data, bookID, pageNumber])

    const goPage = useCallback(
        (page: number) => {
            navigate(BookReaderLink(bookID, page))
        },
        [bookID]
    )

    const pages = bookDetailsResponse.data?.pages
        ?.filter((page) => !showPageOnlyWithPreview || page.preview_url)
        .filter((page) => showPageWithDeadHash || !page.has_dead_hash)

    const prevPage = useCallback(() => {
        if (pageNumber == 1) return

        const pageNumberInArray = pages
            ?.map((e) => e.page_number)
            .filter((e) => e < pageNumber)
            .sort((a: number, b: number) => a - b)
            .reduce((_, cur) => cur, 0)

        if (pageNumberInArray) {
            goPage(pageNumberInArray)
        }
    }, [bookDetailsResponse.data, bookID, pageNumber, goPage, pages])

    const nextPage = useCallback(() => {
        if (pageNumber == bookDetailsResponse.data?.info.page_count) return

        const pageNumberInArray = pages
            ?.map((e) => e.page_number)
            .filter((e) => e > pageNumber)
            .sort((a: number, b: number) => b - a)
            .reduce((_, cur) => cur, 0)

        if (pageNumberInArray) {
            goPage(pageNumberInArray)
        }
    }, [bookDetailsResponse.data, bookID, pageNumber, goPage, pages])

    const goGo = useCallback(
        (event: any) => {
            const pos = document
                .getElementById('main-image')!
                .getBoundingClientRect()
            const dx = (event.pageX - pos.left) / (pos.right - pos.left)
            if (dx < 0.3) {
                prevPage()
            } else {
                nextPage()
            }
        },
        [prevPage, nextPage]
    )

    useEffect(() => {
        const eventHandler = (event: KeyboardEvent) => {
            if (event.keyCode === 37) prevPage()
            if (event.keyCode === 39) nextPage()
        }

        window.addEventListener('keydown', eventHandler)

        return () => {
            window.removeEventListener('keydown', eventHandler)
        }
    }, [prevPage, nextPage])

    return (
        <div>
            <ErrorTextWidget value={bookDetailsResponse} />
            <ErrorTextWidget value={setDeadHashResponse} />
            <ErrorTextWidget value={deletePageResponse} />
            <div className={styles.viewScreen}>
                <div className={'app-container ' + styles.actions}>
                    <Link
                        className='app-button'
                        to={BookDetailsLink(bookID)}
                    >
                        На страницу книги
                    </Link>
                    <PageBadgesWidget
                        flags={currentPage}
                        badgeSize='medium'
                    />
                    <span>
                        {`Страница ${pageNumber} из ${bookDetailsResponse.data?.info.page_count ?? 0}`}
                    </span>
                </div>
                <div className={styles.view}>
                    {currentPage?.preview_url ? (
                        <img
                            src={currentPage?.preview_url}
                            id='main-image'
                            style={{
                                // TODO: подумать что с таким делать
                                maxWidth: '100%',
                                maxHeight: '100%',
                            }}
                            onClick={goGo}
                        />
                    ) : null}
                </div>
                <div className={'app-container ' + styles.actions}>
                    <span>
                        <button
                            className='app'
                            onClick={prevPage}
                        >
                            <span className={styles.pageNavigate}>{'<'}</span>
                        </button>
                        <button
                            className='app'
                            onClick={nextPage}
                        >
                            <span className={styles.pageNavigate}>{'>'}</span>
                        </button>
                    </span>
                    {currentPage ? (
                        <button
                            className='app'
                            onClick={() => {
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
                                    page_number: currentPage.page_number,
                                    set_dead_hash: true,
                                }).then(() => {
                                    getBookDetails({ id: bookID })
                                })
                            }}
                        >
                            <ColorizedTextWidget
                                bold
                                color='danger'
                            >
                                Удалить такие страницы
                            </ColorizedTextWidget>
                        </button>
                    ) : null}
                    <BookReadActionButtonWidget
                        bookID={bookID}
                        pageNumber={pageNumber}
                        currentPage={currentPage}
                        onCreateDeadHash={() => {
                            if (!currentPage) {
                                return
                            }

                            if (
                                !confirm(
                                    'Создать мертвых хеш для таких страниц?'
                                )
                            ) {
                                return
                            }

                            doSetDeadHash({
                                book_id: bookID,
                                page_number: currentPage.page_number,
                                value: true,
                            }).then(() => {
                                getBookDetails({ id: bookID })
                            })
                        }}
                        onDeleteDeadHash={() => {
                            if (!currentPage) {
                                return
                            }

                            if (
                                !confirm(
                                    'Удалить мертвых хеш для таких страниц?'
                                )
                            ) {
                                return
                            }

                            doSetDeadHash({
                                book_id: bookID,
                                page_number: currentPage.page_number,
                                value: false,
                            }).then(() => {
                                getBookDetails({ id: bookID })
                            })
                        }}
                        onDeletePage={() => {
                            if (!currentPage) {
                                return
                            }

                            if (
                                !confirm(
                                    'Удалить эту страницу? (ЭТО НЕОБРАТИМО)'
                                )
                            ) {
                                return
                            }

                            doDeletePage({
                                type: 'one',
                                book_id: bookID,
                                page_number: currentPage.page_number,
                            }).then(() => {
                                getBookDetails({ id: bookID })
                            })
                        }}
                        onDeleteAllPages={() => {
                            if (!currentPage) {
                                return
                            }

                            if (
                                !confirm(
                                    'Удалить такие страницы? (ЭТО НЕОБРАТИМО)'
                                )
                            ) {
                                return
                            }

                            const setDeadHash = confirm(
                                'Установить для текущих страниц мертвый хеш?'
                            )

                            doDeletePage({
                                type: 'all_copy',
                                book_id: bookID,
                                page_number: currentPage.page_number,
                                set_dead_hash: setDeadHash,
                            }).then(() => {
                                getBookDetails({ id: bookID })
                            })
                        }}
                    >
                        <label>
                            <input
                                type='checkbox'
                                className='app'
                                checked={showPageOnlyWithPreview}
                                onChange={(e) =>
                                    setShowPageOnlyWithPreview(e.target.checked)
                                }
                            />
                            Только страницы с изображением
                        </label>
                        <label>
                            <input
                                type='checkbox'
                                className='app'
                                checked={showPageWithDeadHash}
                                onChange={(e) =>
                                    setShowPageWithDeadHash(e.target.checked)
                                }
                            />
                            Показывать с мертвым хешом
                        </label>
                    </BookReadActionButtonWidget>
                    <span>
                        <button
                            className='app'
                            onClick={prevPage}
                        >
                            <span className={styles.pageNavigate}>{'<'}</span>
                        </button>
                        <button
                            className='app'
                            onClick={nextPage}
                        >
                            <span className={styles.pageNavigate}>{'>'}</span>
                        </button>
                    </span>
                </div>
            </div>
        </div>
    )
}
