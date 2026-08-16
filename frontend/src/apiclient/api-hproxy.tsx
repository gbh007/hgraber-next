import { PostAction, Response, useAPIPost } from './client-hooks'

export interface HProxyListResponse {
    books?: Array<HProxyListResponseBook>
    pagination?: Array<HProxyListResponsePage>
}

export interface HProxyListResponseBook {
    ext_url: string
    name?: string
    preview_url?: string
    exists_ids?: Array<string>
}

export interface HProxyListResponsePage {
    ext_url: string
    name: string
}

export interface HProxyListRequest {
    url: string
}

export function useHProxyList(): [
    Response<HProxyListResponse | null>,
    PostAction<HProxyListRequest>,
] {
    const [response, fetchData] = useAPIPost<
        HProxyListRequest,
        HProxyListResponse
    >('/api/hproxy/list')

    return [response, fetchData]
}

export interface HProxyBookResponse {
    name: string
    ext_url: string
    preview_url?: string
    page_count: number
    exists_ids?: Array<string>
    pages: Array<HProxyBookResponsePage>
    attributes: Array<HProxyBookResponseAttribute>
}

export interface HProxyBookResponsePage {
    page_number: number
    preview_url: string
}

export interface HProxyBookResponseAttribute {
    code: string
    name: string
    values: Array<HProxyBookResponseAttributeValue>
}

export interface HProxyBookResponseAttributeValue {
    ext_name: string
    name?: string
    ext_url?: string
    massloads_by_name: Array<HProxyBookResponseAttributeValueMassload>
    massloads_by_ext_url: Array<HProxyBookResponseAttributeValueMassload>
}

export interface HProxyBookResponseAttributeValueMassload {
    id: number
    name: string
}

export interface HProxyBookRequest {
    url: string
    page_limit?: number
}

export function useHProxyBook(): [
    Response<HProxyBookResponse | null>,
    PostAction<HProxyBookRequest>,
] {
    const [response, fetchData] = useAPIPost<
        HProxyBookRequest,
        HProxyBookResponse
    >('/api/hproxy/book')

    return [response, fetchData]
}
