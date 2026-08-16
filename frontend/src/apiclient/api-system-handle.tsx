import { PostAction, useAPIPost, Response } from './client-hooks'

interface systemHandleRequest {
    urls: Array<string>
    is_multi: boolean
    auto_verify: boolean
    read_only_mode: boolean
}

interface systemHandleResponse {
    total_count: number
    loaded_count: number
    duplicate_count: number
    error_count: number
    not_handled?: Array<string>
    details: Array<systemHandleResponseDetails>
}

export interface systemHandleResponseDetails {
    url: string
    is_duplicate: boolean
    duplicate_ids?: Array<string>
    is_handled: boolean
    id?: string
    error_reason?: string
}

export function useSystemHandle(): [
    Response<systemHandleResponse | null>,
    PostAction<systemHandleRequest>,
] {
    const [response, fetchData] = useAPIPost<
        systemHandleRequest,
        systemHandleResponse
    >('/api/parsing/handle')

    return [response, fetchData]
}
