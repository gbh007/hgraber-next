import { useGenPost } from './gen-hooks'
import type { PostAction, Response } from './gen-hooks'
import { postApiHproxyList, postApiHproxyBook } from './gen/hproxy/hproxy'
import type {
    PostApiHproxyListBody,
    PostApiHproxyList200,
    PostApiHproxyBookBody,
    PostApiHproxyBook200,
} from './gen/hGraberNext.schemas'

export function useHProxyList(): [
    Response<PostApiHproxyList200 | null>,
    PostAction<PostApiHproxyListBody>,
] {
    return useGenPost<PostApiHproxyListBody, PostApiHproxyList200>(
        postApiHproxyList
    )
}

export function useHProxyBook(): [
    Response<PostApiHproxyBook200 | null>,
    PostAction<PostApiHproxyBookBody>,
] {
    return useGenPost<PostApiHproxyBookBody, PostApiHproxyBook200>(
        postApiHproxyBook
    )
}
