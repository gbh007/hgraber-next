import { BookListResponsePages } from '../../apiclient/api-book-list'
import styles from './paginator-widget.module.css'

// TODO: выглядит как виджет для фичи книги, перенести
export function PaginatorWidget(props: {
    value: Array<BookListResponsePages>
    onChange: (v: number) => void
}) {
    return (
        <div>
            {props.value.map((page, index) => (
                <span
                    key={index}
                    className={styles.page}
                    data-current={page.is_current ? 'true' : 'false'}
                    data-separator={page.is_separator ? 'true' : 'false'}
                    onClick={() => {
                        if (page.is_separator) return
                        props.onChange(page.value)
                    }}
                >
                    {page.is_separator ? '...' : page.value}
                </span>
            ))}
        </div>
    )
}
