import { PostAction, useAPIPost, Response } from './client-hooks'
import {
    BookAttribute,
    BookDetailsSize,
    BookSimple,
    BookSimplePage,
} from './model-book'
import { FSDBFilesInfo } from './model-fs'

export interface BookDetailsRequest {
    id: string
}

export interface BookDetails {
    info: BookSimple
    page_loaded_percent: number
    attributes?: Array<BookAttribute>
    pages?: Array<BookSimplePage>
    size?: BookDetailsSize
    fs_disposition?: Array<BookDetailsFSDisposition>
}

export interface BookDetailsFSDisposition {
    id: string
    name: string
    files: FSDBFilesInfo
}

export function useBookDetails(): [
    Response<BookDetails | null>,
    PostAction<BookDetailsRequest>,
] {
    const [response, fetchData] = useAPIPost<BookDetailsRequest, BookDetails>(
        '/api/book/details'
    )

    return [response, fetchData]
}
