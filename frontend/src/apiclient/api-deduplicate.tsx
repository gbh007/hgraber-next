import { useGenPost } from './gen-hooks'
import type { PostAction, Response } from './gen-hooks'
import {
    postApiDeduplicateBookByPageBody,
    postApiDeduplicateCompare,
    postApiDeduplicateUniquePages,
    postApiDeduplicateBooksByPage,
    postApiDeduplicateDeadHashSet,
} from './gen/deduplicate/deduplicate'
import type {
    PostApiDeduplicateBookByPageBodyBody,
    PostApiDeduplicateBookByPageBody200,
    PostApiDeduplicateCompareBody,
    PostApiDeduplicateCompare200,
    PostApiDeduplicateUniquePagesBody,
    PostApiDeduplicateUniquePages200,
    PostApiDeduplicateBooksByPageBody,
    PostApiDeduplicateBooksByPage200,
    PostApiDeduplicateDeadHashSetBody,
} from './gen/hGraberNext.schemas'

export function useDeduplicateBookByPageBody(): [
    Response<PostApiDeduplicateBookByPageBody200 | null>,
    PostAction<PostApiDeduplicateBookByPageBodyBody>,
] {
    return useGenPost<
        PostApiDeduplicateBookByPageBodyBody,
        PostApiDeduplicateBookByPageBody200
    >(postApiDeduplicateBookByPageBody)
}

export function useDeduplicateCompare(): [
    Response<PostApiDeduplicateCompare200 | null>,
    PostAction<PostApiDeduplicateCompareBody>,
] {
    return useGenPost<
        PostApiDeduplicateCompareBody,
        PostApiDeduplicateCompare200
    >(postApiDeduplicateCompare)
}

export function useUniqueBookPages(): [
    Response<PostApiDeduplicateUniquePages200 | null>,
    PostAction<PostApiDeduplicateUniquePagesBody>,
] {
    return useGenPost<
        PostApiDeduplicateUniquePagesBody,
        PostApiDeduplicateUniquePages200
    >(postApiDeduplicateUniquePages)
}

export function useDeduplicateBooksByPage(): [
    Response<PostApiDeduplicateBooksByPage200 | null>,
    PostAction<PostApiDeduplicateBooksByPageBody>,
] {
    return useGenPost<
        PostApiDeduplicateBooksByPageBody,
        PostApiDeduplicateBooksByPage200
    >(postApiDeduplicateBooksByPage)
}

export function useSetDeadHash(): [
    Response<void | null>,
    PostAction<PostApiDeduplicateDeadHashSetBody>,
] {
    return useGenPost<PostApiDeduplicateDeadHashSetBody, void>(
        postApiDeduplicateDeadHashSet
    )
}
