import { PostAction, useAPIPost, Response } from './client-hooks'
import { BookFilter } from './model-book-filter'

interface AgentListRequest {
    can_parse?: boolean
    can_export?: boolean
    can_parse_multi?: boolean
    has_fs?: boolean
    has_hproxy?: boolean
    include_status?: boolean
}

export interface Agent {
    id: string
    name: string
    token: string
    addr: string
    can_parse: boolean
    can_parse_multi: boolean
    can_export: boolean
    has_fs: boolean
    has_hproxy: boolean
    priority: number
    created_at: string
}

export interface AgentListResponse {
    status?: AgentListResponseStatus
    info: Agent
}

interface AgentListResponseStatus {
    start_at?: string
    check_status_error?: string
    status: string
    problems?: Array<AgentListResponseStatusProblems>
}

export interface AgentListResponseStatusProblems {
    type: string
    details: string
}

export function useAgentList(): [
    Response<Array<AgentListResponse> | null>,
    PostAction<AgentListRequest>,
] {
    const [response, fetchData] = useAPIPost<
        AgentListRequest,
        Array<AgentListResponse>
    >('/api/agent/list')

    return [response, fetchData]
}

interface AgentDeleteRequest {
    id: string
}

export function useAgentDelete(): [
    Response<void | null>,
    PostAction<AgentDeleteRequest>,
] {
    const [response, fetchData] = useAPIPost<AgentDeleteRequest, void>(
        '/api/agent/delete'
    )

    return [response, fetchData]
}

export interface AgentNewRequest {
    name: string
    addr: string
    token: string
    can_parse: boolean
    can_parse_multi: boolean
    can_export: boolean
    has_fs: boolean
    has_hproxy: boolean
    priority: number
}

export function useAgentNew(): [
    Response<void | null>,
    PostAction<AgentNewRequest>,
] {
    const [response, fetchData] = useAPIPost<AgentNewRequest, void>(
        '/api/agent/new'
    )

    return [response, fetchData]
}

interface AgentTaskExportRequest {
    book_filter: BookFilter
    exporter: string
    delete_after: boolean
}

export function useAgentTaskExport(): [
    Response<void | null>,
    PostAction<AgentTaskExportRequest>,
] {
    const [response, fetchData] = useAPIPost<AgentTaskExportRequest, void>(
        '/api/agent/task/export'
    )

    return [response, fetchData]
}

export interface AgentUpdateRequest {
    id: string
    name: string
    token: string
    addr: string
    can_parse: boolean
    can_parse_multi: boolean
    can_export: boolean
    has_fs: boolean
    has_hproxy: boolean
    priority: number
}

export function useAgentUpdate(): [
    Response<void | null>,
    PostAction<AgentUpdateRequest>,
] {
    const [response, fetchData] = useAPIPost<AgentUpdateRequest, void>(
        '/api/agent/update'
    )

    return [response, fetchData]
}

interface AgentGetRequest {
    id: string
}

export function useAgentGet(): [
    Response<Agent | null>,
    PostAction<AgentGetRequest>,
] {
    const [response, fetchData] = useAPIPost<AgentGetRequest, Agent>(
        '/api/agent/get'
    )

    return [response, fetchData]
}
