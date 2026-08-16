import { GetAction, useAPIGet, Response } from './client-hooks'

export interface SystemInfoSizeResponse {
    count: number
    downloaded_count: number
    verified_count: number
    rebuilded_count: number
    not_load_count: number
    deleted_count: number
    page_count: number
    not_load_page_count: number
    page_without_body_count: number
    deleted_page_count: number
    file_count: number
    unhashed_file_count: number
    invalid_file_count: number
    detached_file_count: number
    dead_hash_count: number
    pages_size: number
    files_size: number
}

export interface SystemInfoWorkersResponse {
    workers?: Array<SystemInfoWorkersResponseWorker>
}

export interface SystemInfoWorkersResponseWorker {
    name: string
    in_queue: number
    in_work: number
    runners: number
}

export function useSystemInfoSize(): [
    Response<SystemInfoSizeResponse | null>,
    GetAction,
] {
    const [response, fetchData] = useAPIGet<SystemInfoSizeResponse>(
        '/api/system/info/size'
    )

    return [response, fetchData]
}

export function useSystemInfoWorkers(): [
    Response<SystemInfoWorkersResponse | null>,
    GetAction,
] {
    const [response, fetchData] = useAPIGet<SystemInfoWorkersResponse>(
        '/api/system/info/workers'
    )

    return [response, fetchData]
}
