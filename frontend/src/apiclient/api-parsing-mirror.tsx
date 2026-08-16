import {
    PostAction,
    useAPIPost,
    Response,
    useAPIGet,
    GetAction,
} from './client-hooks'

export interface ParsingMirrorCreateRequest {
    name: string
    description?: string
    prefixes: Array<string>
}

export function useParsingMirrorCreate(): [
    Response<void | null>,
    PostAction<ParsingMirrorCreateRequest>,
] {
    const [response, fetchData] = useAPIPost<ParsingMirrorCreateRequest, void>(
        '/api/parsing/mirror/create'
    )

    return [response, fetchData]
}

export interface ParsingMirrorUpdateRequest {
    id: string
    name: string
    description?: string
    prefixes: Array<string>
}

export function useParsingMirrorUpdate(): [
    Response<void | null>,
    PostAction<ParsingMirrorUpdateRequest>,
] {
    const [response, fetchData] = useAPIPost<ParsingMirrorUpdateRequest, void>(
        '/api/parsing/mirror/update'
    )

    return [response, fetchData]
}

export interface ParsingMirrorDeleteRequest {
    id: string
}

export function useParsingMirrorDelete(): [
    Response<void | null>,
    PostAction<ParsingMirrorDeleteRequest>,
] {
    const [response, fetchData] = useAPIPost<ParsingMirrorDeleteRequest, void>(
        '/api/parsing/mirror/delete'
    )

    return [response, fetchData]
}

export interface ParsingMirrorListResponse {
    mirrors?: Array<ParsingMirrorListResponseLabel>
}

export interface ParsingMirrorListResponseLabel {
    id: string
    name: string
    description?: string
    prefixes: Array<string>
}

export function useParsingMirrorList(): [
    Response<ParsingMirrorListResponse | null>,
    GetAction,
] {
    const [response, fetchData] = useAPIGet<ParsingMirrorListResponse>(
        '/api/parsing/mirror/list'
    )

    return [response, fetchData]
}

// TODO: разделить модели
export function useParsingMirrorGet(): [
    Response<ParsingMirrorListResponseLabel | null>,
    PostAction<ParsingMirrorDeleteRequest>,
] {
    const [response, fetchData] = useAPIPost<
        ParsingMirrorDeleteRequest,
        ParsingMirrorListResponseLabel
    >('/api/parsing/mirror/get')

    return [response, fetchData]
}
