import { useGenPost } from './gen-hooks'
import type { PostAction, Response } from './client-hooks'
import { postApiParsingHandle } from './gen/parsing/parsing'
import type {
    PostApiParsingHandleBody,
    PostApiParsingHandle200,
} from './gen/hGraberNext.schemas'

export function useSystemHandle(): [
    Response<PostApiParsingHandle200 | null>,
    PostAction<PostApiParsingHandleBody>,
] {
    return useGenPost<PostApiParsingHandleBody, PostApiParsingHandle200>(
        postApiParsingHandle
    )
}
