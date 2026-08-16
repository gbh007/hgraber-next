import { Link } from 'react-router-dom'
import { DeduplicateBookByPageBodyResponseResult } from '../../apiclient/api-deduplicate'
import { BooksSimpleWidget } from '../../widgets/book'
import { BookCompareLink } from '../../core/routing'
import {
    ColorizedTextWidget,
    ContainerWidget,
    PrettyDualSizeWidget,
    prettyPercent,
    prettySize,
} from '../../widgets/design-system'

export function BookDuplicates(props: {
    originID: string
    deduplicateBookInfo?: Array<DeduplicateBookByPageBodyResponseResult>
}) {
    if (!props.deduplicateBookInfo) {
        return null
    }

    return (
        <ContainerWidget
            direction='row'
            gap='medium'
            wrap
        >
            {props.deduplicateBookInfo?.map((book) => (
                <BooksSimpleWidget
                    value={book.book}
                    align='center'
                    key={book.book.id}
                    actualPageCount={
                        book.book.page_count != book.target_page_count
                            ? book.target_page_count
                            : undefined
                    }
                >
                    <ContainerWidget
                        direction='column'
                        style={{ alignItems: 'center' }}
                    >
                        <ContainerWidget
                            direction='row'
                            gap='small'
                        >
                            <span title='Сколько страниц этой книги есть в открытой'>
                                Покрытие книги:
                            </span>
                            <ColorizedTextWidget
                                bold
                                color={
                                    book.origin_covered_target == 1
                                        ? 'good'
                                        : undefined
                                }
                            >
                                {prettyPercent(book.origin_covered_target)}%
                            </ColorizedTextWidget>
                            {book.origin_covered_target !=
                            book.origin_covered_target_without_dead_hashes ? (
                                <span>
                                    (
                                    {prettyPercent(
                                        book.origin_covered_target_without_dead_hashes
                                    )}
                                    %)
                                </span>
                            ) : null}
                        </ContainerWidget>
                        <ContainerWidget
                            direction='row'
                            gap='small'
                        >
                            <span title='Сколько страниц открытой книги есть в этой'>
                                Покрытие оригинала:
                            </span>
                            <ColorizedTextWidget
                                bold
                                color={
                                    book.target_covered_origin == 1
                                        ? 'good'
                                        : undefined
                                }
                            >
                                {prettyPercent(book.target_covered_origin)}%
                            </ColorizedTextWidget>
                            {book.target_covered_origin !=
                            book.target_covered_origin_without_dead_hashes ? (
                                <span>
                                    (
                                    {prettyPercent(
                                        book.target_covered_origin_without_dead_hashes
                                    )}
                                    %)
                                </span>
                            ) : null}
                        </ContainerWidget>
                    </ContainerWidget>
                    <ContainerWidget
                        direction='column'
                        style={{ alignItems: 'center' }}
                    >
                        <span>
                            {`Размер: ${prettySize(book.target_size)} / ${prettySize(book.target_avg_page_size)}`}
                        </span>
                        <PrettyDualSizeWidget
                            first={book.shared_size}
                            second={book.shared_size_without_dead_hashes}
                        >
                            <span>Общий размер:</span>
                        </PrettyDualSizeWidget>
                        <ContainerWidget
                            direction='row'
                            gap='small'
                        >
                            <span>
                                Общие страницы: {book.shared_page_count}
                            </span>
                            {book.shared_page_count !=
                            book.shared_page_count_without_dead_hashes ? (
                                <span>
                                    (
                                    {book.shared_page_count_without_dead_hashes}
                                    )
                                </span>
                            ) : null}
                        </ContainerWidget>
                        <Link
                            className='app-button'
                            to={BookCompareLink(props.originID, book.book.id)}
                        >
                            Сравнить
                        </Link>
                    </ContainerWidget>
                </BooksSimpleWidget>
            ))}
        </ContainerWidget>
    )
}
