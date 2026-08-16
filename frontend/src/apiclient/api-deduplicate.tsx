import { PostAction, useAPIPost, Response } from './client-hooks'
import { BookAttribute, BookSimple, BookSimplePage } from './model-book'

export interface DeduplicateBookByPageBodyRequest {
    book_id: string
}

export interface DeduplicateBookByPageBodyResponse {
    result?: Array<DeduplicateBookByPageBodyResponseResult>
}

export interface DeduplicateBookByPageBodyResponseResult {
    book: BookSimple
    origin_covered_target: number
    target_covered_origin: number
    origin_covered_target_without_dead_hashes: number
    target_covered_origin_without_dead_hashes: number

    target_size: number
    shared_size: number
    shared_size_without_dead_hashes: number
    shared_page_count: number
    shared_page_count_without_dead_hashes: number

    target_avg_page_size: number
    target_page_count: number
}

export function useDeduplicateBookByPageBody(): [
    Response<DeduplicateBookByPageBodyResponse | null>,
    PostAction<DeduplicateBookByPageBodyRequest>,
] {
    const [response, fetchData] = useAPIPost<
        DeduplicateBookByPageBodyRequest,
        DeduplicateBookByPageBodyResponse
    >('/api/deduplicate/book-by-page-body')

    return [response, fetchData]
}

export interface DeduplicateCompareRequest {
    origin_book_id: string
    target_book_id: string
}

export interface DeduplicateCompareResponse {
    origin: BookSimple
    target: BookSimple

    origin_covered_target: number
    target_covered_origin: number
    origin_covered_target_without_dead_hashes: number
    target_covered_origin_without_dead_hashes: number

    origin_pages?: Array<BookSimplePage>
    both_pages?: Array<BookSimplePage>
    target_pages?: Array<BookSimplePage>

    origin_attributes?: Array<BookAttribute>
    both_attributes?: Array<BookAttribute>
    target_attributes?: Array<BookAttribute>

    origin_book_size?: number
    target_book_size?: number
    origin_page_avg_size?: number
    target_page_avg_size?: number
}

export function useDeduplicateCompare(): [
    Response<DeduplicateCompareResponse | null>,
    PostAction<DeduplicateCompareRequest>,
] {
    const [response, fetchData] = useAPIPost<
        DeduplicateCompareRequest,
        DeduplicateCompareResponse
    >('/api/deduplicate/compare')

    return [response, fetchData]
}

export interface UniqueBookPagesRequest {
    book_id: string
}

export interface UniqueBookPagesResponse {
    pages?: Array<BookSimplePage>
}

export function useUniqueBookPages(): [
    Response<UniqueBookPagesResponse | null>,
    PostAction<UniqueBookPagesRequest>,
] {
    const [response, fetchData] = useAPIPost<
        UniqueBookPagesRequest,
        UniqueBookPagesResponse
    >('/api/deduplicate/unique-pages')

    return [response, fetchData]
}

export interface DeduplicateBooksByPageRequest {
    book_id: string
    page_number: number
}

export interface DeduplicateBooksByPageResponse {
    books?: Array<BookSimple>
}

export function useDeduplicateBooksByPage(): [
    Response<DeduplicateBooksByPageResponse | null>,
    PostAction<DeduplicateBooksByPageRequest>,
] {
    const [response, fetchData] = useAPIPost<
        DeduplicateBooksByPageRequest,
        DeduplicateBooksByPageResponse
    >('/api/deduplicate/books-by-page')

    return [response, fetchData]
}

export interface SetDeadHashRequest {
    book_id: string
    page_number?: number
    value: boolean
}

export function useSetDeadHash(): [
    Response<void | null>,
    PostAction<SetDeadHashRequest>,
] {
    const [response, fetchData] = useAPIPost<SetDeadHashRequest, void>(
        '/api/deduplicate/dead-hash/set'
    )

    return [response, fetchData]
}
