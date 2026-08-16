import { PostAction, useAPIPost, Response } from './client-hooks'

interface userLoginRequest {
    token: string
}

export function useUserLogin(): [
    Response<void | null>,
    PostAction<userLoginRequest>,
] {
    const [response, fetchData] = useAPIPost<userLoginRequest, void>(
        '/api/user/login'
    )

    return [response, fetchData]
}
