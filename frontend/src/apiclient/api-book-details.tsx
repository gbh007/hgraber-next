import { useGenPost } from './gen-hooks'
import type { PostAction, Response } from './gen-hooks'
import { postApiBookDetails } from './gen/book/book'
import type {
    PostApiBookDetailsBody,
    PostApiBookDetails200,
} from './gen/hGraberNext.schemas'

export function useBookDetails(): [
    Response<PostApiBookDetails200 | null>,
    PostAction<PostApiBookDetailsBody>,
] {
    return useGenPost<PostApiBookDetailsBody, PostApiBookDetails200>(
        postApiBookDetails
    )
}
