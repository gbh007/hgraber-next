import { useGenPost } from './gen-hooks'
import type { PostAction, Response } from './gen-hooks'
import {
    postApiFsGet,
    postApiFsCreate,
    postApiFsUpdate,
    postApiFsDelete,
    postApiFsList,
    postApiFsValidate,
    postApiFsRemoveMismatch,
    postApiFsTransfer,
    postApiFsTransferBook,
} from './gen/fs/fs'
import type {
    FileSystemInfo,
    PostApiFsGetBody,
    PostApiFsCreateBody,
    PostApiFsCreate200,
    PostApiFsUpdateBody,
    PostApiFsDeleteBody,
    PostApiFsListBody,
    PostApiFsList200,
    PostApiFsValidateBody,
    PostApiFsRemoveMismatchBody,
    PostApiFsTransferBody,
    PostApiFsTransferBookBody,
} from './gen/hGraberNext.schemas'

export function useFSGet(): [
    Response<FileSystemInfo | null>,
    PostAction<PostApiFsGetBody>,
] {
    return useGenPost<PostApiFsGetBody, FileSystemInfo>(postApiFsGet)
}

export function useFSCreate(): [
    Response<PostApiFsCreate200 | null>,
    PostAction<PostApiFsCreateBody>,
] {
    return useGenPost<PostApiFsCreateBody, PostApiFsCreate200>(postApiFsCreate)
}

export function useFSUpdate(): [
    Response<void | null>,
    PostAction<PostApiFsUpdateBody>,
] {
    return useGenPost<PostApiFsUpdateBody, void>(postApiFsUpdate)
}

export function useFSDelete(): [
    Response<void | null>,
    PostAction<PostApiFsDeleteBody>,
] {
    return useGenPost<PostApiFsDeleteBody, void>(postApiFsDelete)
}

export function useFSList(): [
    Response<PostApiFsList200 | null>,
    PostAction<PostApiFsListBody>,
] {
    return useGenPost<PostApiFsListBody, PostApiFsList200>(postApiFsList)
}

export function useFSValidate(): [
    Response<void | null>,
    PostAction<PostApiFsValidateBody>,
] {
    return useGenPost<PostApiFsValidateBody, void>(postApiFsValidate)
}

export function useFSRemoveMismatch(): [
    Response<void | null>,
    PostAction<PostApiFsRemoveMismatchBody>,
] {
    return useGenPost<PostApiFsRemoveMismatchBody, void>(
        postApiFsRemoveMismatch
    )
}

export function useFSTransfer(): [
    Response<void | null>,
    PostAction<PostApiFsTransferBody>,
] {
    return useGenPost<PostApiFsTransferBody, void>(postApiFsTransfer)
}

export function useFSTransferBook(): [
    Response<void | null>,
    PostAction<PostApiFsTransferBookBody>,
] {
    return useGenPost<PostApiFsTransferBookBody, void>(postApiFsTransferBook)
}
