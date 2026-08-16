import {
    GetAction,
    PostAction,
    Response,
    useAPIGet,
    useAPIPost,
} from './client-hooks'

export interface AttributeCountResponse {
    attributes?: Array<AttributeCountResponseAttribute>
}

export interface AttributeCountResponseAttribute {
    code: string
    value: string
    count: number
}

export function useAttributeCount(): [
    Response<AttributeCountResponse | null>,
    GetAction,
] {
    const [response, fetchData] = useAPIGet<AttributeCountResponse>(
        '/api/attribute/count'
    )

    return [response, fetchData]
}

export interface AttributeOriginCountResponse {
    attributes?: Array<AttributeOriginCountResponseAttribute>
}

export interface AttributeOriginCountResponseAttribute {
    code: string
    value: string
    count: number
}

export function useAttributeOriginCount(): [
    Response<AttributeOriginCountResponse | null>,
    GetAction,
] {
    const [response, fetchData] = useAPIGet<AttributeOriginCountResponse>(
        '/api/attribute/origin/count'
    )

    return [response, fetchData]
}

export interface AttributeColor {
    code: string
    value: string
    text_color: string
    background_color: string
    created_at: string
}

export interface AttributeColorListResponse {
    colors?: Array<AttributeColor>
}

export function useAttributeColorList(): [
    Response<AttributeColorListResponse | null>,
    GetAction,
] {
    const [response, fetchData] = useAPIGet<AttributeColorListResponse>(
        '/api/attribute/color/list'
    )

    return [response, fetchData]
}

export interface AttributeColorDeleteRequest {
    code: string
    value: string
}

export function useAttributeColorDelete(): [
    Response<void | null>,
    PostAction<AttributeColorDeleteRequest>,
] {
    const [response, fetchData] = useAPIPost<AttributeColorDeleteRequest, void>(
        '/api/attribute/color/delete'
    )

    return [response, fetchData]
}

export interface AttributeColorGetRequest {
    code: string
    value: string
}

export function useAttributeColorGet(): [
    Response<AttributeColor | null>,
    PostAction<AttributeColorGetRequest>,
] {
    const [response, fetchData] = useAPIPost<
        AttributeColorGetRequest,
        AttributeColor
    >('/api/attribute/color/get')

    return [response, fetchData]
}

export interface AttributeColorCreateRequest {
    code: string
    value: string
    text_color: string
    background_color: string
}

export function useAttributeColorCreate(): [
    Response<void | null>,
    PostAction<AttributeColorCreateRequest>,
] {
    const [response, fetchData] = useAPIPost<AttributeColorCreateRequest, void>(
        '/api/attribute/color/create'
    )

    return [response, fetchData]
}

export interface AttributeColorUpdateRequest {
    code: string
    value: string
    text_color: string
    background_color: string
}

export function useAttributeColorUpdate(): [
    Response<void | null>,
    PostAction<AttributeColorUpdateRequest>,
] {
    const [response, fetchData] = useAPIPost<AttributeColorUpdateRequest, void>(
        '/api/attribute/color/update'
    )

    return [response, fetchData]
}

export interface AttributeRemap {
    code: string
    value: string
    to_code?: string
    to_value?: string
    is_delete?: boolean
    created_at: string
    updated_at?: string
}

export interface AttributeRemapListResponse {
    remaps?: Array<AttributeRemap>
}

export function useAttributeRemapList(): [
    Response<AttributeRemapListResponse | null>,
    GetAction,
] {
    const [response, fetchData] = useAPIGet<AttributeRemapListResponse>(
        '/api/attribute/remap/list'
    )

    return [response, fetchData]
}

export interface AttributeRemapDeleteRequest {
    code: string
    value: string
}

export function useAttributeRemapDelete(): [
    Response<void | null>,
    PostAction<AttributeRemapDeleteRequest>,
] {
    const [response, fetchData] = useAPIPost<AttributeRemapDeleteRequest, void>(
        '/api/attribute/remap/delete'
    )

    return [response, fetchData]
}

export interface AttributeRemapGetRequest {
    code: string
    value: string
}

export function useAttributeRemapGet(): [
    Response<AttributeRemap | null>,
    PostAction<AttributeRemapGetRequest>,
] {
    const [response, fetchData] = useAPIPost<
        AttributeRemapGetRequest,
        AttributeRemap
    >('/api/attribute/remap/get')

    return [response, fetchData]
}

export interface AttributeRemapCreateRequest {
    code: string
    value: string
    to_code?: string
    to_value?: string
    is_delete?: boolean
}

export function useAttributeRemapCreate(): [
    Response<void | null>,
    PostAction<AttributeRemapCreateRequest>,
] {
    const [response, fetchData] = useAPIPost<AttributeRemapCreateRequest, void>(
        '/api/attribute/remap/create'
    )

    return [response, fetchData]
}

export interface AttributeRemapUpdateRequest {
    code: string
    value: string
    to_code?: string
    to_value?: string
    is_delete?: boolean
}

export function useAttributeRemapUpdate(): [
    Response<void | null>,
    PostAction<AttributeRemapUpdateRequest>,
] {
    const [response, fetchData] = useAPIPost<AttributeRemapUpdateRequest, void>(
        '/api/attribute/remap/update'
    )

    return [response, fetchData]
}
