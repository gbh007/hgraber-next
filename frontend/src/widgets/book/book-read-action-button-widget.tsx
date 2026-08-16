import { PropsWithChildren, useState } from 'react'
import { BookSimplePage } from '../../apiclient/model-book'
import { Link } from 'react-router-dom'
import { BookLabelEditLink, DeduplicatePageLink } from '../../core/routing'
import { BookTransferCoordinatorWidget } from '../fs'
import { ColorizedTextWidget, DialogWidget } from '../design-system'

export function BookReadActionButtonWidget(
    props: PropsWithChildren & {
        bookID: string
        pageNumber: number
        currentPage?: BookSimplePage
        onCreateDeadHash: () => void
        onDeleteDeadHash: () => void
        onDeletePage: () => void
        onDeleteAllPages: () => void
    }
) {
    const [show, setShow] = useState(false)

    return (
        <>
            <button
                className='app'
                onClick={() => {
                    setShow(true)
                }}
            >
                опции
            </button>
            <DialogWidget
                open={show}
                onClose={() => setShow(false)}
            >
                <Link
                    className='app-button'
                    to={BookLabelEditLink(props.bookID, props.pageNumber)}
                >
                    Редактировать метки
                </Link>
                {props.currentPage?.has_dead_hash === false ? (
                    <button
                        className='app'
                        onClick={() => {
                            setShow(false)
                            props.onCreateDeadHash()
                        }}
                    >
                        <ColorizedTextWidget color='danger-lite'>
                            Создать мертвый хеш
                        </ColorizedTextWidget>
                    </button>
                ) : null}
                {props.currentPage?.has_dead_hash === true ? (
                    <button
                        className='app'
                        onClick={() => {
                            setShow(false)
                            props.onDeleteDeadHash()
                        }}
                    >
                        <ColorizedTextWidget color='danger-lite'>
                            Удалить мертвый хеш
                        </ColorizedTextWidget>
                    </button>
                ) : null}
                {props.currentPage ? (
                    <>
                        <button
                            className='app'
                            onClick={() => {
                                setShow(false)
                                props.onDeletePage()
                            }}
                        >
                            <ColorizedTextWidget
                                color='danger-lite'
                                bold
                            >
                                Удалить эту страницу
                            </ColorizedTextWidget>
                        </button>
                        <button
                            className='app'
                            onClick={() => {
                                setShow(false)
                                props.onDeleteAllPages()
                            }}
                        >
                            <ColorizedTextWidget
                                color='danger'
                                bold
                            >
                                Удалить такие страницы
                            </ColorizedTextWidget>
                        </button>
                        <Link
                            className='app-button'
                            to={DeduplicatePageLink(
                                props.bookID,
                                props.pageNumber
                            )}
                        >
                            Книги с этой страницей
                        </Link>
                    </>
                ) : null}
                {/* TODO: возможно не самое подходящее место */}
                <BookTransferCoordinatorWidget
                    bookID={props.bookID}
                    pageNumber={props.pageNumber}
                />
                {props.children}
            </DialogWidget>
        </>
    )
}
