import { CSSProperties, useEffect } from 'react'
import { Link, useParams } from 'react-router-dom'
import { useBookDetails } from '../../apiclient/api-book-details'
import { useBookDelete } from '../../apiclient/api-book-delete'
import { useBookRestore, useBookStatusSet } from '../../apiclient/api-book'
import {
    useDeduplicateBookByPageBody,
    useSetDeadHash,
} from '../../apiclient/api-deduplicate'
import { useAttributeColorList } from '../../apiclient/api-attribute'
import { BookDetailInfoWidget } from './book-detail-info-widget'
import {
    BookEditLink,
    BookLabelEditLink,
    BookReaderLink,
    BookRebuildLink,
    BookUniquePagesLink,
} from '../../core/routing'
import { BookTransferCoordinatorWidget } from '../../widgets/fs'
import {
    ColorizedTextWidget,
    ContainerWidget,
    ErrorTextWidget,
    prettySize,
} from '../../widgets/design-system'

export function BookDetailsScreen() {
    const mainButton: CSSProperties = { textAlign: 'center', flexGrow: 1 }

    const params = useParams()
    const bookID = params.bookID!

    const [bookDetailsResponse, getBookDetails] = useBookDetails()
    const [bookDeleteResponse, postBookDelete] = useBookDelete()
    const [bookSetStatusResponse, doBookSetStatus] = useBookStatusSet()
    const [bookDeduplicateResponse, doBookDeduplicate] =
        useDeduplicateBookByPageBody()
    const [setDeadHashResponse, doSetDeadHash] = useSetDeadHash()
    const [bookRestoreResponse, doBookRestore] = useBookRestore()

    const [attributeColorListResponse, fetchAttributeColorList] =
        useAttributeColorList()
    useEffect(() => {
        fetchAttributeColorList()
    }, [fetchAttributeColorList])

    useEffect(() => {
        getBookDetails({ id: bookID })
    }, [getBookDetails, bookID])

    useEffect(() => {
        bookDeduplicateResponse.cleanData()
    }, [bookID])

    const hasPages = (bookDetailsResponse.data?.pages?.length ?? 0) > 0
    const hasUniquePages =
        !bookDetailsResponse.data?.info.flags.is_deleted &&
        (!bookDetailsResponse.data?.size ||
            bookDetailsResponse.data.size.unique != 0)
    const hasSharedPages =
        !bookDetailsResponse.data?.info.flags.is_deleted &&
        (!bookDetailsResponse.data?.size ||
            bookDetailsResponse.data.size.shared != 0)
    const hasDeletedPages =
        bookDetailsResponse.data?.pages &&
        bookDetailsResponse.data?.pages.length !=
            bookDetailsResponse.data?.info.page_count

    return (
        <ContainerWidget
            direction='column'
            gap='big'
        >
            <ErrorTextWidget value={bookDetailsResponse} />
            <ErrorTextWidget value={bookDeleteResponse} />
            <ErrorTextWidget value={bookSetStatusResponse} />
            <ErrorTextWidget value={bookDeduplicateResponse} />
            <ErrorTextWidget value={setDeadHashResponse} />
            <ErrorTextWidget value={bookRestoreResponse} />
            <ErrorTextWidget value={attributeColorListResponse} />
            {bookDetailsResponse.data ? (
                <BookDetailInfoWidget
                    book={bookDetailsResponse.data}
                    deduplicateBookInfo={bookDeduplicateResponse.data?.result}
                    colors={attributeColorListResponse.data?.colors}
                >
                    <ContainerWidget
                        direction='column'
                        gap='big'
                    >
                        <ContainerWidget
                            direction='row'
                            gap='medium'
                            wrap
                        >
                            <button
                                className='app'
                                style={mainButton}
                                onClick={() => {
                                    window.open(
                                        '/api/book/archive/' + bookID,
                                        '_blank'
                                    )
                                }}
                            >
                                Скачать
                            </button>
                            {hasPages ? (
                                <Link
                                    className='app-button'
                                    style={mainButton}
                                    to={BookReaderLink(bookID)}
                                >
                                    Читать
                                </Link>
                            ) : null}
                            {bookDetailsResponse.data.info.flags
                                .is_deleted ? null : (
                                <button
                                    className='app'
                                    style={mainButton}
                                    disabled={bookDeleteResponse.isLoading}
                                    onClick={() => {
                                        if (
                                            !confirm(
                                                `Удалить книгу: ${bookDetailsResponse.data?.info.name}?`
                                            )
                                        ) {
                                            return
                                        }

                                        if (
                                            bookDetailsResponse.data?.size
                                                ?.unique &&
                                            !confirm(
                                                `У книги ${bookDetailsResponse.data?.info.name} есть ${prettySize(bookDetailsResponse.data?.size?.unique)} уникального контента, точно хотите ее удалить?`
                                            )
                                        ) {
                                            return
                                        }

                                        postBookDelete({
                                            book_id: bookID,
                                            type: 'soft',
                                        }).then(() => {
                                            getBookDetails({ id: bookID })
                                        })
                                    }}
                                >
                                    <ColorizedTextWidget color='danger'>
                                        Удалить
                                    </ColorizedTextWidget>
                                </button>
                            )}
                            {bookDetailsResponse.data.info.flags.is_verified ? (
                                <button
                                    className='app'
                                    style={mainButton}
                                    disabled={bookSetStatusResponse.isLoading}
                                    onClick={() => {
                                        if (
                                            !confirm(
                                                `Снять подтверждение с книги: ${bookDetailsResponse.data?.info.name}?`
                                            )
                                        ) {
                                            return
                                        }

                                        doBookSetStatus({
                                            id: bookID,
                                            status: 'verify',
                                            value: false,
                                        }).then(() => {
                                            getBookDetails({ id: bookID })
                                        })
                                    }}
                                >
                                    <ColorizedTextWidget color='danger-lite'>
                                        Снять статус подтвержденной
                                    </ColorizedTextWidget>
                                </button>
                            ) : (
                                <button
                                    className='app'
                                    style={mainButton}
                                    disabled={bookSetStatusResponse.isLoading}
                                    onClick={() => {
                                        if (
                                            !confirm(
                                                `Подтвердить книгу: ${bookDetailsResponse.data?.info.name}?`
                                            )
                                        ) {
                                            return
                                        }

                                        doBookSetStatus({
                                            id: bookID,
                                            status: 'verify',
                                            value: true,
                                        }).then(() => {
                                            getBookDetails({ id: bookID })
                                        })
                                    }}
                                >
                                    <ColorizedTextWidget color='good'>
                                        Подтвердить
                                    </ColorizedTextWidget>
                                </button>
                            )}
                            {hasSharedPages ? (
                                <button
                                    className='app'
                                    style={mainButton}
                                    disabled={bookDeduplicateResponse.isLoading}
                                    onClick={() => {
                                        doBookDeduplicate({ book_id: bookID })
                                    }}
                                >
                                    Показать дубли
                                </button>
                            ) : null}
                            {hasUniquePages ? (
                                <Link
                                    className='app-button'
                                    style={mainButton}
                                    to={BookUniquePagesLink(
                                        bookDetailsResponse.data.info.id
                                    )}
                                >
                                    Показать уникальные страницы
                                </Link>
                            ) : null}
                        </ContainerWidget>
                        <details className='app'>
                            <summary>дополнительные действия</summary>
                            <ContainerWidget
                                direction='row'
                                gap='medium'
                                wrap
                            >
                                <Link
                                    className='app-button'
                                    to={BookEditLink(
                                        bookDetailsResponse.data.info.id
                                    )}
                                >
                                    Редактировать
                                </Link>
                                <Link
                                    className='app-button'
                                    to={BookRebuildLink(
                                        bookDetailsResponse.data.info.id
                                    )}
                                >
                                    Пересобрать
                                </Link>
                                <Link
                                    className='app-button'
                                    to={BookLabelEditLink(
                                        bookDetailsResponse.data.info.id
                                    )}
                                >
                                    Редактировать метки
                                </Link>
                                {hasPages ? (
                                    <>
                                        <button
                                            className='app'
                                            disabled={
                                                setDeadHashResponse.isLoading
                                            }
                                            onClick={() => {
                                                if (
                                                    !confirm(
                                                        'Создать мертвых хеш для всех страниц этой книги?'
                                                    )
                                                ) {
                                                    return
                                                }

                                                doSetDeadHash({
                                                    book_id: bookID,
                                                    value: true,
                                                }).then(() => {
                                                    getBookDetails({
                                                        id: bookID,
                                                    })
                                                })
                                            }}
                                        >
                                            <ColorizedTextWidget color='danger-lite'>
                                                Создать мертвый хеш
                                            </ColorizedTextWidget>
                                        </button>
                                        <button
                                            className='app'
                                            disabled={
                                                setDeadHashResponse.isLoading
                                            }
                                            onClick={() => {
                                                if (
                                                    !confirm(
                                                        'Удалить мертвых хеш для всех страниц этой книги?'
                                                    )
                                                ) {
                                                    return
                                                }

                                                doSetDeadHash({
                                                    book_id: bookID,
                                                    value: false,
                                                }).then(() => {
                                                    getBookDetails({
                                                        id: bookID,
                                                    })
                                                })
                                            }}
                                        >
                                            <ColorizedTextWidget color='danger-lite'>
                                                Удалить мертвый хеш
                                            </ColorizedTextWidget>
                                        </button>
                                        <button
                                            className='app'
                                            disabled={
                                                bookDeleteResponse.isLoading
                                            }
                                            onClick={() => {
                                                if (
                                                    !confirm(
                                                        'Удалить все страницы из этой книги с мертвым хешом? (ЭТО НЕОБРАТИМО)'
                                                    )
                                                ) {
                                                    return
                                                }

                                                postBookDelete({
                                                    book_id: bookID,
                                                    type: 'dead_hashed_pages',
                                                }).then(() => {
                                                    getBookDetails({
                                                        id: bookID,
                                                    })
                                                })
                                            }}
                                        >
                                            <ColorizedTextWidget color='danger'>
                                                Удалить все страницы из книги с
                                                мертвым хешом
                                            </ColorizedTextWidget>
                                        </button>
                                        <button
                                            className='app'
                                            disabled={
                                                bookDeleteResponse.isLoading
                                            }
                                            onClick={() => {
                                                if (
                                                    !confirm(
                                                        'Удалить все страницы из этой книги и их копии? (ЭТО НЕОБРАТИМО)'
                                                    )
                                                ) {
                                                    return
                                                }

                                                const markAsDeletedEmptyBook =
                                                    confirm(
                                                        'Удалить книги которые станут пустыми?'
                                                    )

                                                postBookDelete({
                                                    book_id: bookID,
                                                    type: 'page_and_copy',
                                                    mark_as_deleted_empty_book:
                                                        markAsDeletedEmptyBook,
                                                }).then(() => {
                                                    getBookDetails({
                                                        id: bookID,
                                                    })
                                                })
                                            }}
                                        >
                                            <ColorizedTextWidget
                                                bold
                                                color='danger'
                                            >
                                                Удалить все страницы из книги и
                                                их копии
                                            </ColorizedTextWidget>
                                        </button>
                                        <BookTransferCoordinatorWidget
                                            bookID={bookID}
                                        />
                                    </>
                                ) : null}
                                {!hasPages ||
                                hasDeletedPages ||
                                bookDetailsResponse.data.info.flags
                                    .is_deleted ? (
                                    <button
                                        className='app'
                                        disabled={bookRestoreResponse.isLoading}
                                        onClick={() => {
                                            if (
                                                !confirm('Восстановить книгу?')
                                            ) {
                                                return
                                            }

                                            const onlyPages = confirm(
                                                'Восстановить только страницы?'
                                            )

                                            doBookRestore({
                                                book_id: bookID,
                                                only_pages: onlyPages,
                                            }).then(() => {
                                                getBookDetails({ id: bookID })
                                            })
                                        }}
                                    >
                                        <ColorizedTextWidget color='good'>
                                            Восстановить книгу
                                        </ColorizedTextWidget>
                                    </button>
                                ) : null}

                                {bookDetailsResponse.data.info.flags
                                    .is_rebuild ? (
                                    <button
                                        className='app'
                                        style={mainButton}
                                        disabled={
                                            bookSetStatusResponse.isLoading
                                        }
                                        onClick={() => {
                                            if (
                                                !confirm(
                                                    `Снять статус пересобранной с книги: ${bookDetailsResponse.data?.info.name}?`
                                                )
                                            ) {
                                                return
                                            }

                                            doBookSetStatus({
                                                id: bookID,
                                                status: 'rebuild',
                                                value: false,
                                            }).then(() => {
                                                getBookDetails({ id: bookID })
                                            })
                                        }}
                                    >
                                        <ColorizedTextWidget color='danger-lite'>
                                            Снять статус пересобранной
                                        </ColorizedTextWidget>
                                    </button>
                                ) : (
                                    <button
                                        className='app'
                                        style={mainButton}
                                        disabled={
                                            bookSetStatusResponse.isLoading
                                        }
                                        onClick={() => {
                                            if (
                                                !confirm(
                                                    `Установить статус пересобранной книге: ${bookDetailsResponse.data?.info.name}?`
                                                )
                                            ) {
                                                return
                                            }

                                            doBookSetStatus({
                                                id: bookID,
                                                status: 'rebuild',
                                                value: true,
                                            }).then(() => {
                                                getBookDetails({ id: bookID })
                                            })
                                        }}
                                    >
                                        <ColorizedTextWidget color='danger-lite'>
                                            Установить статус пересобранной
                                        </ColorizedTextWidget>
                                    </button>
                                )}
                            </ContainerWidget>
                        </details>
                    </ContainerWidget>
                </BookDetailInfoWidget>
            ) : null}
        </ContainerWidget>
    )
}
