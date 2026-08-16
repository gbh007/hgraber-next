import { PostAction, useAPIPost, Response } from './client-hooks'
import { FSDBFilesInfo } from './model-fs'

export interface FileSystemInfo {
    id: string
    name: string
    description?: string
    agent_id?: string
    path?: string
    download_priority: number
    deduplicate_priority: number
    highway_enabled: boolean
    highway_addr?: string
    created_at: string
}

export interface FSGetRequest {
    id: string
}

export function useFSGet(): [
    Response<FileSystemInfo | null>,
    PostAction<FSGetRequest>,
] {
    const [response, fetchData] = useAPIPost<FSGetRequest, FileSystemInfo>(
        '/api/fs/get'
    )

    return [response, fetchData]
}

export interface FSCreateRequest {
    name: string
    description?: string
    agent_id?: string
    path?: string
    download_priority: number
    deduplicate_priority: number
    highway_enabled: boolean
    highway_addr?: string
}

export interface FSCreateResponse {
    id: string
}

export function useFSCreate(): [
    Response<FSCreateResponse | null>,
    PostAction<FSCreateRequest>,
] {
    const [response, fetchData] = useAPIPost<FSCreateRequest, FSCreateResponse>(
        '/api/fs/create'
    )

    return [response, fetchData]
}

export interface FSUpdateRequest {
    id: string
    name: string
    description?: string
    agent_id?: string
    path?: string
    download_priority: number
    deduplicate_priority: number
    highway_enabled: boolean
    highway_addr?: string
}

export function useFSUpdate(): [
    Response<void | null>,
    PostAction<FSUpdateRequest>,
] {
    const [response, fetchData] = useAPIPost<FSUpdateRequest, void>(
        '/api/fs/update'
    )

    return [response, fetchData]
}

export interface FSDeleteRequest {
    id: string
}

export function useFSDelete(): [
    Response<void | null>,
    PostAction<FSDeleteRequest>,
] {
    const [response, fetchData] = useAPIPost<FSDeleteRequest, void>(
        '/api/fs/delete'
    )

    return [response, fetchData]
}

export interface FSListRequest {
    include_db_file_size?: boolean
    include_available_size?: boolean
}

export interface FSListResponse {
    file_systems?: Array<FSListResponseUnit>
}

export interface FSListResponseUnit {
    info: FileSystemInfo
    is_legacy: boolean
    db_files_info?: FSDBFilesInfo
    db_invalid_files_info?: FSDBFilesInfo
    db_detached_files_info?: FSDBFilesInfo
    available_size?: number
}

export function useFSList(): [
    Response<FSListResponse | null>,
    PostAction<FSListRequest>,
] {
    const [response, fetchData] = useAPIPost<FSListRequest, FSListResponse>(
        '/api/fs/list'
    )

    return [response, fetchData]
}

export interface FSValidateRequest {
    id: string
}

export function useFSValidate(): [
    Response<void | null>,
    PostAction<FSValidateRequest>,
] {
    const [response, fetchData] = useAPIPost<FSValidateRequest, void>(
        '/api/fs/validate'
    )

    return [response, fetchData]
}

export interface FSRemoveMismatchRequest {
    id: string
}

export function useFSRemoveMismatch(): [
    Response<void | null>,
    PostAction<FSRemoveMismatchRequest>,
] {
    const [response, fetchData] = useAPIPost<FSRemoveMismatchRequest, void>(
        '/api/fs/remove-mismatch'
    )

    return [response, fetchData]
}

export interface FSTransferRequest {
    from: string
    to: string
    only_preview_pages?: boolean
}

export function useFSTransfer(): [
    Response<void | null>,
    PostAction<FSTransferRequest>,
] {
    const [response, fetchData] = useAPIPost<FSTransferRequest, void>(
        '/api/fs/transfer'
    )

    return [response, fetchData]
}

export interface FSTransferBookRequest {
    book_id: string
    to: string
    page_number?: number
    only_preview_pages?: boolean
}

export function useFSTransferBook(): [
    Response<void | null>,
    PostAction<FSTransferBookRequest>,
] {
    const [response, fetchData] = useAPIPost<FSTransferBookRequest, void>(
        '/api/fs/transfer/book'
    )

    return [response, fetchData]
}
