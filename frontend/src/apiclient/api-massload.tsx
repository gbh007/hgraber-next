import { useGenGet, useGenPost } from './gen-hooks'
import type { GetAction, PostAction, Response } from './gen-hooks'
import {
    postApiMassloadInfoCreate,
    postApiMassloadInfoUpdate,
    postApiMassloadInfoDelete,
    postApiMassloadInfoGet,
    postApiMassloadInfoList,
    postApiMassloadCalculate,
    postApiMassloadFlagCreate,
    postApiMassloadFlagUpdate,
    postApiMassloadFlagDelete,
    postApiMassloadFlagGet,
    getApiMassloadFlagList,
    postApiMassloadInfoExternalLinkCreate,
    postApiMassloadInfoExternalLinkUpdate,
    postApiMassloadInfoExternalLinkDelete,
    postApiMassloadInfoAttributeCreate,
    postApiMassloadInfoAttributeDelete,
} from './gen/massload/massload'
import type {
    MassloadInfo,
    MassloadFlag,
    PostApiMassloadInfoCreateBody,
    PostApiMassloadInfoCreate200,
    PostApiMassloadInfoUpdateBody,
    PostApiMassloadInfoDeleteBody,
    PostApiMassloadInfoGetBody,
    PostApiMassloadInfoListBody,
    PostApiMassloadInfoList200,
    PostApiMassloadCalculateBody,
    PostApiMassloadFlagCreateBody,
    PostApiMassloadFlagUpdateBody,
    PostApiMassloadFlagDeleteBody,
    PostApiMassloadFlagGetBody,
    GetApiMassloadFlagList200,
    PostApiMassloadInfoExternalLinkCreateBody,
    PostApiMassloadInfoExternalLinkUpdateBody,
    PostApiMassloadInfoExternalLinkDeleteBody,
    PostApiMassloadInfoAttributeCreateBody,
    PostApiMassloadInfoAttributeDeleteBody,
} from './gen/hGraberNext.schemas'

export function useMassloadInfoCreate(): [
    Response<PostApiMassloadInfoCreate200 | null>,
    PostAction<PostApiMassloadInfoCreateBody>,
] {
    return useGenPost<
        PostApiMassloadInfoCreateBody,
        PostApiMassloadInfoCreate200
    >(postApiMassloadInfoCreate)
}

export function useMassloadInfoUpdate(): [
    Response<void | null>,
    PostAction<PostApiMassloadInfoUpdateBody>,
] {
    return useGenPost<PostApiMassloadInfoUpdateBody, void>(
        postApiMassloadInfoUpdate
    )
}

export function useMassloadInfoDelete(): [
    Response<void | null>,
    PostAction<PostApiMassloadInfoDeleteBody>,
] {
    return useGenPost<PostApiMassloadInfoDeleteBody, void>(
        postApiMassloadInfoDelete
    )
}

export function useMassloadInfoGet(): [
    Response<MassloadInfo | null>,
    PostAction<PostApiMassloadInfoGetBody>,
] {
    return useGenPost<PostApiMassloadInfoGetBody, MassloadInfo>(
        postApiMassloadInfoGet
    )
}

export function useMassloadInfoList(): [
    Response<PostApiMassloadInfoList200 | null>,
    PostAction<PostApiMassloadInfoListBody>,
] {
    return useGenPost<PostApiMassloadInfoListBody, PostApiMassloadInfoList200>(
        postApiMassloadInfoList
    )
}

export function useMassloadCalculate(): [
    Response<void | null>,
    PostAction<PostApiMassloadCalculateBody>,
] {
    return useGenPost<PostApiMassloadCalculateBody, void>(
        postApiMassloadCalculate
    )
}

export function useMassloadFlagCreate(): [
    Response<void | null>,
    PostAction<PostApiMassloadFlagCreateBody>,
] {
    return useGenPost<PostApiMassloadFlagCreateBody, void>(
        postApiMassloadFlagCreate
    )
}

export function useMassloadFlagUpdate(): [
    Response<void | null>,
    PostAction<PostApiMassloadFlagUpdateBody>,
] {
    return useGenPost<PostApiMassloadFlagUpdateBody, void>(
        postApiMassloadFlagUpdate
    )
}

export function useMassloadFlagDelete(): [
    Response<void | null>,
    PostAction<PostApiMassloadFlagDeleteBody>,
] {
    return useGenPost<PostApiMassloadFlagDeleteBody, void>(
        postApiMassloadFlagDelete
    )
}

export function useMassloadFlagGet(): [
    Response<MassloadFlag | null>,
    PostAction<PostApiMassloadFlagGetBody>,
] {
    return useGenPost<PostApiMassloadFlagGetBody, MassloadFlag>(
        postApiMassloadFlagGet
    )
}

export function useMassloadFlagList(): [
    Response<GetApiMassloadFlagList200 | null>,
    GetAction,
] {
    return useGenGet<GetApiMassloadFlagList200>(getApiMassloadFlagList)
}

export function useMassloadExternalLinkCreate(): [
    Response<void | null>,
    PostAction<PostApiMassloadInfoExternalLinkCreateBody>,
] {
    return useGenPost<PostApiMassloadInfoExternalLinkCreateBody, void>(
        postApiMassloadInfoExternalLinkCreate
    )
}

export function useMassloadExternalLinkUpdate(): [
    Response<void | null>,
    PostAction<PostApiMassloadInfoExternalLinkUpdateBody>,
] {
    return useGenPost<PostApiMassloadInfoExternalLinkUpdateBody, void>(
        postApiMassloadInfoExternalLinkUpdate
    )
}

export function useMassloadExternalLinkDelete(): [
    Response<void | null>,
    PostAction<PostApiMassloadInfoExternalLinkDeleteBody>,
] {
    return useGenPost<PostApiMassloadInfoExternalLinkDeleteBody, void>(
        postApiMassloadInfoExternalLinkDelete
    )
}

export function useMassloadAttributeCreate(): [
    Response<void | null>,
    PostAction<PostApiMassloadInfoAttributeCreateBody>,
] {
    return useGenPost<PostApiMassloadInfoAttributeCreateBody, void>(
        postApiMassloadInfoAttributeCreate
    )
}

export function useMassloadAttributeDelete(): [
    Response<void | null>,
    PostAction<PostApiMassloadInfoAttributeDeleteBody>,
] {
    return useGenPost<PostApiMassloadInfoAttributeDeleteBody, void>(
        postApiMassloadInfoAttributeDelete
    )
}
