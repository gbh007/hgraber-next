import { useGenPost } from './gen-hooks'
import type { PostAction, Response } from './gen-hooks'
import { postApiBookDelete, postApiBookPageDelete } from './gen/book/book'
import type {
    PostApiBookDeleteBody,
    PostApiBookPageDeleteBody,
} from './gen/hGraberNext.schemas'

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
