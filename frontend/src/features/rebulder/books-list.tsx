import { BookShortInfo } from '../../apiclient/api-book-list'
import { BooksSimpleWidget } from '../../widgets/book'
import styles from './book-rebuilder.module.css'

export function BooksList(props: {
    value?: Array<BookShortInfo>
    selected?: string
    onChange: (v: BookShortInfo) => void
}) {
    return (
        <div className={styles.preview}>
            {props.value?.map((book) => (
                <BooksSimpleWidget
                    value={book.info}
                    key={book.info.id}
                >
                    <button
                        className='app'
                        onClick={() => props.onChange(book)}
                        disabled={book.info.id == props.selected}
                    >
                        Выбрать
                    </button>
                </BooksSimpleWidget>
            ))}
        </div>
    )
}
