import { useGenPost } from './gen-hooks'
import type { PostAction, Response } from './client-hooks'
import { postApiSystemWorkerConfig } from './gen/system/system'
import type { PostApiSystemWorkerConfigBody } from './gen/hGraberNext.schemas'

export function useSystemWorkerConfig(): [
    Response<void | null>,
    PostAction<PostApiSystemWorkerConfigBody>,
] {
    return useGenPost<PostApiSystemWorkerConfigBody, void>(
        postApiSystemWorkerConfig
    )
}
