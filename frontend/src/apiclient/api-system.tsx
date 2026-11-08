import { useGenGet, useGenPost } from './gen-hooks'
import type { GetAction, PostAction, Response } from './gen-hooks'
import { postApiUserLogin } from './gen/user/user'
import {
    getApiSystemInfoSize,
    getApiSystemInfoWorkers,
    postApiSystemTaskCreate,
    getApiSystemTaskResults,
    postApiSystemWorkerConfig,
} from './gen/system/system'
import { postApiParsingHandle } from './gen/parsing/parsing'
import type {
    PostApiUserLoginBody,
    GetApiSystemInfoSize200,
    GetApiSystemInfoWorkers200,
    PostApiSystemTaskCreateBody,
    GetApiSystemTaskResults200,
    PostApiSystemWorkerConfigBody,
    PostApiParsingHandleBody,
    PostApiParsingHandle200,
} from './gen/hGraberNext.schemas'

export function useUserLogin(): [
    Response<void | null>,
    PostAction<PostApiUserLoginBody>,
] {
    return useGenPost<PostApiUserLoginBody, void>(postApiUserLogin)
}

export function useSystemWorkerConfig(): [
    Response<void | null>,
    PostAction<PostApiSystemWorkerConfigBody>,
] {
    return useGenPost<PostApiSystemWorkerConfigBody, void>(
        postApiSystemWorkerConfig
    )
}

export function useSystemInfoSize(): [
    Response<GetApiSystemInfoSize200 | null>,
    GetAction,
] {
    return useGenGet<GetApiSystemInfoSize200>(getApiSystemInfoSize)
}

export function useSystemInfoWorkers(): [
    Response<GetApiSystemInfoWorkers200 | null>,
    GetAction,
] {
    return useGenGet<GetApiSystemInfoWorkers200>(getApiSystemInfoWorkers)
}

export function useSystemHandle(): [
    Response<PostApiParsingHandle200 | null>,
    PostAction<PostApiParsingHandleBody>,
] {
    return useGenPost<PostApiParsingHandleBody, PostApiParsingHandle200>(
        postApiParsingHandle
    )
}

export function useTaskCreate(): [
    Response<void | null>,
    PostAction<PostApiSystemTaskCreateBody>,
] {
    return useGenPost<PostApiSystemTaskCreateBody, void>(postApiSystemTaskCreate)
}

export function useTaskResults(): [
    Response<GetApiSystemTaskResults200 | null>,
    GetAction,
] {
    return useGenGet<GetApiSystemTaskResults200>(getApiSystemTaskResults)
}
