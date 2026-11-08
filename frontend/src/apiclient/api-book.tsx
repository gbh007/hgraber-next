import { useGenPost } from './gen-hooks'
import type { PostAction, Response } from './gen-hooks'
import {
    postApiBookUpdate,
    postApiBookRaw,
    postApiBookRebuild,
    postApiBookRestore,
    postApiBookStatusSet,
    postApiBookDetails,
    postApiBookList,
    postApiBookDelete,
    postApiBookPageDelete,
} from './gen/book/book'
import type {
    BookRaw,
    BookFilter,
    PostApiBookRawBody,
    PostApiBookRebuildBody,
    PostApiBookRebuild200,
    PostApiBookRestoreBody,
    PostApiBookStatusSetBody,
    PostApiBookDetailsBody,
    PostApiBookDetails200,
    PostApiBookList200,
    PostApiBookDeleteBody,
    PostApiBookPageDeleteBody,
} from './gen/hGraberNext.schemas'

export function useBookUpdate(): [Response<void | null>, PostAction<BookRaw>] {
    return useGenPost<BookRaw, void>(postApiBookUpdate)
}

export function useBookRaw(): [
    Response<BookRaw | null>,
    PostAction<PostApiBookRawBody>,
] {
    return useGenPost<PostApiBookRawBody, BookRaw>(postApiBookRaw)
}

export function useBookRebuild(): [
    Response<PostApiBookRebuild200 | null>,
    PostAction<PostApiBookRebuildBody>,
] {
    return useGenPost<PostApiBookRebuildBody, PostApiBookRebuild200>(
        postApiBookRebuild
    )
}

export function useBookRestore(): [
    Response<void | null>,
    PostAction<PostApiBookRestoreBody>,
] {
    return useGenPost<PostApiBookRestoreBody, void>(postApiBookRestore)
}

export function useBookStatusSet(): [
    Response<void | null>,
    PostAction<PostApiBookStatusSetBody>,
] {
    return useGenPost<PostApiBookStatusSetBody, void>(postApiBookStatusSet)
}

export function useBookDetails(): [
    Response<PostApiBookDetails200 | null>,
    PostAction<PostApiBookDetailsBody>,
] {
    return useGenPost<PostApiBookDetailsBody, PostApiBookDetails200>(
        postApiBookDetails
    )
}

export function useBookList(): [
    Response<PostApiBookList200 | null>,
    PostAction<BookFilter>,
] {
    return useGenPost<BookFilter, PostApiBookList200>(postApiBookList)
}

export function useBookDelete(): [
    Response<void | null>,
    PostAction<PostApiBookDeleteBody>,
] {
    return useGenPost<PostApiBookDeleteBody, void>(postApiBookDelete)
}

export function useDeleteBookPage(): [
    Response<void | null>,
    PostAction<PostApiBookPageDeleteBody>,
] {
    return useGenPost<PostApiBookPageDeleteBody, void>(postApiBookPageDelete)
}
