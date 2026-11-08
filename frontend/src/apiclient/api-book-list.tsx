import { useGenPost } from './gen-hooks'
import type { PostAction, Response } from './gen-hooks'
import { postApiBookList } from './gen/book/book'
import type { BookFilter, PostApiBookList200 } from './gen/hGraberNext.schemas'

export function useBookList(): [
    Response<PostApiBookList200 | null>,
    PostAction<BookFilter>,
] {
    return useGenPost<BookFilter, PostApiBookList200>(postApiBookList)
}
