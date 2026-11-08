import { useGenGet, useGenPost } from './gen-hooks'
import type { GetAction, PostAction, Response } from './gen-hooks'
import {
    getApiAttributeCount,
    getApiAttributeOriginCount,
    getApiAttributeColorList,
    getApiAttributeRemapList,
    postApiAttributeColorDelete,
    postApiAttributeColorGet,
    postApiAttributeColorCreate,
    postApiAttributeColorUpdate,
    postApiAttributeRemapDelete,
    postApiAttributeRemapGet,
    postApiAttributeRemapCreate,
    postApiAttributeRemapUpdate,
} from './gen/attribute/attribute'
import type {
    GetApiAttributeCount200,
    GetApiAttributeOriginCount200,
    GetApiAttributeColorList200,
    GetApiAttributeRemapList200,
    AttributeColor,
    PostApiAttributeColorDeleteBody,
    PostApiAttributeColorGetBody,
    PostApiAttributeColorCreateBody,
    PostApiAttributeColorUpdateBody,
    AttributeRemap,
    PostApiAttributeRemapDeleteBody,
    PostApiAttributeRemapGetBody,
    PostApiAttributeRemapCreateBody,
    PostApiAttributeRemapUpdateBody,
} from './gen/hGraberNext.schemas'

export function useAttributeCount(): [
    Response<GetApiAttributeCount200 | null>,
    GetAction,
] {
    return useGenGet<GetApiAttributeCount200>(getApiAttributeCount)
}

export function useAttributeOriginCount(): [
    Response<GetApiAttributeOriginCount200 | null>,
    GetAction,
] {
    return useGenGet<GetApiAttributeOriginCount200>(getApiAttributeOriginCount)
}

export function useAttributeColorList(): [
    Response<GetApiAttributeColorList200 | null>,
    GetAction,
] {
    return useGenGet<GetApiAttributeColorList200>(getApiAttributeColorList)
}

export function useAttributeColorDelete(): [
    Response<void | null>,
    PostAction<PostApiAttributeColorDeleteBody>,
] {
    return useGenPost<PostApiAttributeColorDeleteBody, void>(
        postApiAttributeColorDelete
    )
}

export function useAttributeColorGet(): [
    Response<AttributeColor | null>,
    PostAction<PostApiAttributeColorGetBody>,
] {
    return useGenPost<PostApiAttributeColorGetBody, AttributeColor>(
        postApiAttributeColorGet
    )
}

export function useAttributeColorCreate(): [
    Response<void | null>,
    PostAction<PostApiAttributeColorCreateBody>,
] {
    return useGenPost<PostApiAttributeColorCreateBody, void>(
        postApiAttributeColorCreate
    )
}

export function useAttributeColorUpdate(): [
    Response<void | null>,
    PostAction<PostApiAttributeColorUpdateBody>,
] {
    return useGenPost<PostApiAttributeColorUpdateBody, void>(
        postApiAttributeColorUpdate
    )
}

export function useAttributeRemapList(): [
    Response<GetApiAttributeRemapList200 | null>,
    GetAction,
] {
    return useGenGet<GetApiAttributeRemapList200>(getApiAttributeRemapList)
}

export function useAttributeRemapDelete(): [
    Response<void | null>,
    PostAction<PostApiAttributeRemapDeleteBody>,
] {
    return useGenPost<PostApiAttributeRemapDeleteBody, void>(
        postApiAttributeRemapDelete
    )
}

export function useAttributeRemapGet(): [
    Response<AttributeRemap | null>,
    PostAction<PostApiAttributeRemapGetBody>,
] {
    return useGenPost<PostApiAttributeRemapGetBody, AttributeRemap>(
        postApiAttributeRemapGet
    )
}

export function useAttributeRemapCreate(): [
    Response<void | null>,
    PostAction<PostApiAttributeRemapCreateBody>,
] {
    return useGenPost<PostApiAttributeRemapCreateBody, void>(
        postApiAttributeRemapCreate
    )
}

export function useAttributeRemapUpdate(): [
    Response<void | null>,
    PostAction<PostApiAttributeRemapUpdateBody>,
] {
    return useGenPost<PostApiAttributeRemapUpdateBody, void>(
        postApiAttributeRemapUpdate
    )
}
