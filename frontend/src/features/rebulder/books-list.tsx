import type { PostApiBookList200BooksItem } from '../../apiclient/gen/hGraberNext.schemas'
import { BooksSimpleWidget } from '../../widgets/book'
import styles from './book-rebuilder.module.css'

export function BooksList(props: {
    value?: Array<PostApiBookList200BooksItem>
    selected?: string
    onChange: (v: PostApiBookList200BooksItem) => void
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
