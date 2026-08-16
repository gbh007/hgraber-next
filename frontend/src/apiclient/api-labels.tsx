import {
    PostAction,
    useAPIPost,
    Response,
    useAPIGet,
    GetAction,
} from './client-hooks'

export interface LabelSetRequest {
    book_id: string
    page_number?: number
    name: string
    value: string
}

export function useLabelSet(): [
    Response<void | null>,
    PostAction<LabelSetRequest>,
] {
    const [response, fetchData] = useAPIPost<LabelSetRequest, void>(
        '/api/label/set'
    )

    return [response, fetchData]
}

export interface LabelDeleteRequest {
    book_id: string
    page_number?: number
    name: string
}

export function useLabelDelete(): [
    Response<void | null>,
    PostAction<LabelDeleteRequest>,
] {
    const [response, fetchData] = useAPIPost<LabelDeleteRequest, void>(
        '/api/label/delete'
    )

    return [response, fetchData]
}

export interface LabelGetRequest {
    book_id: string
}

export interface LabelGetResponse {
    labels?: Array<LabelGetResponseLabel>
}

export interface LabelGetResponseLabel {
    book_id: string
    page_number: number
    name: string
    value: string
    created_at: string
}

export function useLabelGet(): [
    Response<LabelGetResponse | null>,
    PostAction<LabelGetRequest>,
] {
    const [response, fetchData] = useAPIPost<LabelGetRequest, LabelGetResponse>(
        '/api/label/get'
    )

    return [response, fetchData]
}

export interface LabelPresetCreateRequest {
    name: string
    description?: string
    values: Array<string>
}

export function useLabelPresetCreate(): [
    Response<void | null>,
    PostAction<LabelPresetCreateRequest>,
] {
    const [response, fetchData] = useAPIPost<LabelPresetCreateRequest, void>(
        '/api/label/preset/create'
    )

    return [response, fetchData]
}

export interface LabelPresetUpdateRequest {
    name: string
    description?: string
    values: Array<string>
}

export function useLabelPresetUpdate(): [
    Response<void | null>,
    PostAction<LabelPresetUpdateRequest>,
] {
    const [response, fetchData] = useAPIPost<LabelPresetUpdateRequest, void>(
        '/api/label/preset/update'
    )

    return [response, fetchData]
}

export interface LabelPresetDeleteRequest {
    name: string
}

export function useLabelPresetDelete(): [
    Response<void | null>,
    PostAction<LabelPresetDeleteRequest>,
] {
    const [response, fetchData] = useAPIPost<LabelPresetDeleteRequest, void>(
        '/api/label/preset/delete'
    )

    return [response, fetchData]
}

export interface LabelPresetListResponse {
    presets?: Array<LabelPresetListResponseLabel>
}

export interface LabelPresetListResponseLabel {
    name: string
    description?: string
    values: Array<string>
    created_at: string
    updated_at?: string
}

export function useLabelPresetList(): [
    Response<LabelPresetListResponse | null>,
    GetAction,
] {
    const [response, fetchData] = useAPIGet<LabelPresetListResponse>(
        '/api/label/preset/list'
    )

    return [response, fetchData]
}

// TODO: разделить модели
export function useLabelPresetGet(): [
    Response<LabelPresetListResponseLabel | null>,
    PostAction<LabelPresetDeleteRequest>,
] {
    const [response, fetchData] = useAPIPost<
        LabelPresetDeleteRequest,
        LabelPresetListResponseLabel
    >('/api/label/preset/get')

    return [response, fetchData]
}
