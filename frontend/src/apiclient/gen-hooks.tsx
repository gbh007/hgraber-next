import axios from 'axios'
import { useCallback, useState } from 'react'
import type { Response, PostAction, GetAction } from './client-hooks'
import type { AxiosResponse } from 'axios'

// Обёртки над сгенерированным в ./gen клиентом (orval)
// с интерфейсом Response из ./client-hooks
export function useGenPost<RequestType, ResponseType>(
    generatePost: (data: RequestType) => Promise<AxiosResponse<ResponseType>>
): [Response<ResponseType | null>, PostAction<RequestType>] {
    const [data, setData] = useState<ResponseType | null>(null)
    const [isLoading, setIsLoading] = useState(false)
    const [isError, setIsError] = useState(false)
    const [isUnauthorize, setIsUnauthorize] = useState(false)
    const [errorText, setErrorText] = useState('')

    const fetchData = useCallback(
        async (requestData: RequestType) => {
            setIsError(false)
            setIsUnauthorize(false)
            setIsLoading(true)
            setErrorText('')

            try {
                const result = await generatePost(requestData)

                setData(result.data)
            } catch (error) {
                let text = 'unknown error'

                if (error instanceof axios.AxiosError) {
                    text = error.response?.data?.details || 'unknown error'

                    if (error.response?.status == 401) {
                        setIsUnauthorize(true)
                    }

                    if (error.response?.status == 403) {
                        setIsUnauthorize(true)
                    }
                }

                setErrorText(text)
                setIsError(true)
                setIsLoading(false)

                throw text
            }

            setIsLoading(false)
        },
        [generatePost]
    )

    return [
        {
            data,
            isLoading,
            isError,
            errorText,
            isUnauthorize,
            cleanData: () => {
                setData(null)
            },
        },
        fetchData,
    ]
}

export function useGenGet<ResponseType>(
    generateGet: () => Promise<AxiosResponse<ResponseType>>
): [Response<ResponseType | null>, GetAction] {
    const [data, setData] = useState<ResponseType | null>(null)
    const [isLoading, setIsLoading] = useState(false)
    const [isError, setIsError] = useState(false)
    const [isUnauthorize, setIsUnauthorize] = useState(false)
    const [errorText, setErrorText] = useState('')

    const fetchData = useCallback(async () => {
        setIsError(false)
        setIsUnauthorize(false)
        setIsLoading(true)
        setErrorText('')

        try {
            const result = await generateGet()

            setData(result.data)
        } catch (error) {
            let text = 'unknown error'

            if (error instanceof axios.AxiosError) {
                text = error.response?.data?.details || 'unknown error'

                if (error.response?.status == 401) {
                    setIsUnauthorize(true)
                }

                if (error.response?.status == 403) {
                    setIsUnauthorize(true)
                }
            }

            setErrorText(text)
            setIsError(true)
            setIsLoading(false)

            throw text
        }

        setIsLoading(false)
    }, [generateGet])

    return [
        {
            data,
            isLoading,
            isError,
            errorText,
            isUnauthorize,
            cleanData: () => {
                setData(null)
            },
        },
        fetchData,
    ]
}
