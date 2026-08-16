import { PostAction, useAPIPost, Response } from './client-hooks'

interface systemWorkerConfigRequest {
    runners_count: Array<systemWorkerRunnerConfig>
}

export interface systemWorkerRunnerConfig {
    name: string
    count: number
}

export function useSystemWorkerConfig(): [
    Response<void | null>,
    PostAction<systemWorkerConfigRequest>,
] {
    const [response, fetchData] = useAPIPost<systemWorkerConfigRequest, void>(
        '/api/system/worker/config'
    )

    return [response, fetchData]
}
