import { useGenGet, useGenPost } from './gen-hooks'
import type { GetAction, PostAction, Response } from './gen-hooks'
import {
    getApiSystemTaskResults,
    postApiSystemTaskCreate,
} from './gen/system/system'
import type {
    GetApiSystemTaskResults200,
    PostApiSystemTaskCreateBody,
} from './gen/hGraberNext.schemas'

export function useTaskCreate(): [
    Response<void | null>,
    PostAction<PostApiSystemTaskCreateBody>,
] {
    return useGenPost<PostApiSystemTaskCreateBody, void>(
        postApiSystemTaskCreate
    )
}

export function useTaskResults(): [
    Response<GetApiSystemTaskResults200 | null>,
    GetAction,
] {
    return useGenGet<GetApiSystemTaskResults200>(getApiSystemTaskResults)
}
