import { useGenPost } from './gen-hooks'
import type { PostAction, Response } from './gen-hooks'
import {
    postApiBookUpdate,
    postApiBookRaw,
    postApiBookRebuild,
    postApiBookRestore,
    postApiBookStatusSet,
} from './gen/book/book'
import type {
    BookRaw,
    PostApiBookRawBody,
    PostApiBookRebuildBody,
    PostApiBookRebuild200,
    PostApiBookRestoreBody,
    PostApiBookStatusSetBody,
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
