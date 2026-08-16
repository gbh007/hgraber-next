import { PostAction, useAPIPost, Response } from './client-hooks'
import { BookFilter } from './model-book-filter'
import { BookSimple } from './model-book'

export interface BookListResponse {
    books?: Array<BookShortInfo>
    count: number
    pages: Array<BookListResponsePages>
}

export interface BookShortInfo {
    info: BookSimple
    color_attributes?: Array<BookShortInfoColorAttributes>
}

export interface BookShortInfoColorAttributes {
    code: string
    value: string
    text_color: string
    background_color: string
}

export interface BookListResponsePages {
    value: number
    is_current: boolean
    is_separator: boolean
}

export function useBookList(): [
    Response<BookListResponse | null>,
    PostAction<BookFilter>,
] {
    const [response, fetchData] = useAPIPost<BookFilter, BookListResponse>(
        '/api/book/list'
    )

    return [response, fetchData]
}
