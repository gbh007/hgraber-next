import { useGenGet, useGenPost } from './gen-hooks'
import type { GetAction, PostAction, Response } from './gen-hooks'
import {
    postApiParsingMirrorCreate,
    postApiParsingMirrorUpdate,
    postApiParsingMirrorDelete,
    postApiParsingMirrorGet,
    getApiParsingMirrorList,
} from './gen/parsing/parsing'
import type {
    PostApiParsingMirrorCreateBody,
    PostApiParsingMirrorUpdateBody,
    PostApiParsingMirrorDeleteBody,
    PostApiParsingMirrorGetBody,
    PostApiParsingMirrorGet200,
    GetApiParsingMirrorList200,
} from './gen/hGraberNext.schemas'

export function useParsingMirrorCreate(): [
    Response<void | null>,
    PostAction<PostApiParsingMirrorCreateBody>,
] {
    return useGenPost<PostApiParsingMirrorCreateBody, void>(
        postApiParsingMirrorCreate
    )
}

export function useParsingMirrorUpdate(): [
    Response<void | null>,
    PostAction<PostApiParsingMirrorUpdateBody>,
] {
    return useGenPost<PostApiParsingMirrorUpdateBody, void>(
        postApiParsingMirrorUpdate
    )
}

export function useParsingMirrorDelete(): [
    Response<void | null>,
    PostAction<PostApiParsingMirrorDeleteBody>,
] {
    return useGenPost<PostApiParsingMirrorDeleteBody, void>(
        postApiParsingMirrorDelete
    )
}

export function useParsingMirrorList(): [
    Response<GetApiParsingMirrorList200 | null>,
    GetAction,
] {
    return useGenGet<GetApiParsingMirrorList200>(getApiParsingMirrorList)
}

export function useParsingMirrorGet(): [
    Response<PostApiParsingMirrorGet200 | null>,
    PostAction<PostApiParsingMirrorGetBody>,
] {
    return useGenPost<PostApiParsingMirrorGetBody, PostApiParsingMirrorGet200>(
        postApiParsingMirrorGet
    )
}
