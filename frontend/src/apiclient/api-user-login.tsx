import { useGenPost } from './gen-hooks'
import type { PostAction, Response } from './client-hooks'
import { postApiUserLogin } from './gen/user/user'
import type { PostApiUserLoginBody } from './gen/hGraberNext.schemas'

export function useUserLogin(): [
    Response<void | null>,
    PostAction<PostApiUserLoginBody>,
] {
    return useGenPost<PostApiUserLoginBody, void>(postApiUserLogin)
}
