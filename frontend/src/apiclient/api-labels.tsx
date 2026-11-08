import { useGenGet, useGenPost } from './gen-hooks'
import type { GetAction, PostAction, Response } from './gen-hooks'
import {
    postApiLabelSet,
    postApiLabelDelete,
    postApiLabelGet,
    postApiLabelPresetCreate,
    postApiLabelPresetUpdate,
    postApiLabelPresetDelete,
    postApiLabelPresetGet,
    getApiLabelPresetList,
} from './gen/label/label'
import type {
    PostApiLabelSetBody,
    PostApiLabelDeleteBody,
    PostApiLabelGetBody,
    PostApiLabelGet200,
    PostApiLabelPresetCreateBody,
    PostApiLabelPresetUpdateBody,
    PostApiLabelPresetDeleteBody,
    PostApiLabelPresetGetBody,
    PostApiLabelPresetGet200,
    GetApiLabelPresetList200,
} from './gen/hGraberNext.schemas'

export function useLabelSet(): [
    Response<void | null>,
    PostAction<PostApiLabelSetBody>,
] {
    return useGenPost<PostApiLabelSetBody, void>(postApiLabelSet)
}

export function useLabelDelete(): [
    Response<void | null>,
    PostAction<PostApiLabelDeleteBody>,
] {
    return useGenPost<PostApiLabelDeleteBody, void>(postApiLabelDelete)
}

export function useLabelGet(): [
    Response<PostApiLabelGet200 | null>,
    PostAction<PostApiLabelGetBody>,
] {
    return useGenPost<PostApiLabelGetBody, PostApiLabelGet200>(postApiLabelGet)
}

export function useLabelPresetCreate(): [
    Response<void | null>,
    PostAction<PostApiLabelPresetCreateBody>,
] {
    return useGenPost<PostApiLabelPresetCreateBody, void>(
        postApiLabelPresetCreate
    )
}

export function useLabelPresetUpdate(): [
    Response<void | null>,
    PostAction<PostApiLabelPresetUpdateBody>,
] {
    return useGenPost<PostApiLabelPresetUpdateBody, void>(
        postApiLabelPresetUpdate
    )
}

export function useLabelPresetDelete(): [
    Response<void | null>,
    PostAction<PostApiLabelPresetDeleteBody>,
] {
    return useGenPost<PostApiLabelPresetDeleteBody, void>(
        postApiLabelPresetDelete
    )
}

export function useLabelPresetList(): [
    Response<GetApiLabelPresetList200 | null>,
    GetAction,
] {
    return useGenGet<GetApiLabelPresetList200>(getApiLabelPresetList)
}

export function useLabelPresetGet(): [
    Response<PostApiLabelPresetGet200 | null>,
    PostAction<PostApiLabelPresetGetBody>,
] {
    return useGenPost<PostApiLabelPresetGetBody, PostApiLabelPresetGet200>(
        postApiLabelPresetGet
    )
}
