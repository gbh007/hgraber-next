import { PostAction, useAPIPost, Response } from './client-hooks'
import { BookRaw } from './model-book'

export function useBookUpdate(): [Response<void | null>, PostAction<BookRaw>] {
    const [response, fetchData] = useAPIPost<BookRaw, void>('/api/book/update')

    return [response, fetchData]
}

export interface BookRawRequest {
    id?: string
    url?: string
}

export function useBookRaw(): [
    Response<BookRaw | null>,
    PostAction<BookRawRequest>,
] {
    const [response, fetchData] = useAPIPost<BookRawRequest, BookRaw>(
        '/api/book/raw'
    )

    return [response, fetchData]
}

export interface BookRebuildRequest {
    old_book: BookRaw
    selected_pages: Array<number>
    merge_with_book?: string
    flags?: BookRebuildRequestFlags
    page_order: Array<number>
}

export interface BookRebuildRequestFlags {
    only_unique?: boolean
    exclude_dead_hash_pages?: boolean
    only_1_copy?: boolean
    set_origin_labels?: boolean
    extract_mode?: boolean
    mark_unused_pages_as_dead_hash?: boolean
    mark_unused_pages_as_deleted?: boolean
    mark_empty_book_as_deleted_after_remove_pages?: boolean
    auto_verify?: boolean
    page_re_order?: boolean
}

export interface BookRebuildResponse {
    id: string
}

export function useBookRebuild(): [
    Response<BookRebuildResponse | null>,
    PostAction<BookRebuildRequest>,
] {
    const [response, fetchData] = useAPIPost<
        BookRebuildRequest,
        BookRebuildResponse
    >('/api/book/rebuild')

    return [response, fetchData]
}

export interface BookRestoreRequest {
    book_id: string
    only_pages?: boolean
}

export function useBookRestore(): [
    Response<void | null>,
    PostAction<BookRestoreRequest>,
] {
    const [response, fetchData] = useAPIPost<BookRestoreRequest, void>(
        '/api/book/restore'
    )

    return [response, fetchData]
}

export interface BookStatusSetRequest {
    id: string
    status: 'verify' | 'rebuild'
    value: boolean
}

export function useBookStatusSet(): [
    Response<void | null>,
    PostAction<BookStatusSetRequest>,
] {
    const [response, fetchData] = useAPIPost<BookStatusSetRequest, void>(
        '/api/book/status/set'
    )

    return [response, fetchData]
}
