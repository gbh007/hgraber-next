import { useEffect, useState } from 'react'
import {
    systemWorkerRunnerConfig,
    useSystemWorkerConfig,
} from '../../apiclient/api-system-worker-config'
import { useSystemInfoWorkers } from '../../apiclient/api-system-info'

export function RPCScreen() {
    return (
        <>
            <h1 style={{ color: 'var(--app-color-danger)' }}>
                Операции на данной странице не являются обратимыми.
                <br />
                Перед их выполнением рекомендуется произвести бекап как
                файлового хранилища так и БД.
                <br />
                Выполнение операций во время загрузки данных или парсинга или
                прочей обработки гарантированно приводит к повреждению данных.
            </h1>
            <h2 style={{ color: 'var(--app-color-danger)' }}>
                Операции на данной странице также являются медленными и
                тежеловесными, не рекомендуется их выполнять без крайней
                необходимости.
            </h2>

            <div
                style={{
                    display: 'flex',
                    flexDirection: 'column',
                    gap: '20px',
                }}
            >
                <RunnersWidget />
            </div>
        </>
    )
}

function RunnersWidget() {
    const [runners, setRunners] = useState<Array<systemWorkerRunnerConfig>>()
    const [runnerObj, fetchRunnerData] = useSystemInfoWorkers()

    useEffect(() => {
        fetchRunnerData()
    }, [fetchRunnerData])
    useEffect(() => {
        let newRunners: Array<systemWorkerRunnerConfig> = []

        ;(runnerObj.data?.workers || []).map((e) => {
            newRunners.push({
                count: e.runners,
                name: e.name,
            })
        })

        setRunners(newRunners)
    }, [runnerObj.data])

    const [{ isError, errorText, isLoading }, doRequest] =
        useSystemWorkerConfig()

    return (
        <div className='app-container'>
            <h3>Изменение количества раннеров</h3>
            {isError ? (
                <div className='app-error-container'>{errorText}</div>
            ) : null}
            {runnerObj.isError ? (
                <div className='app-error-container'>{runnerObj.errorText}</div>
            ) : null}
            <table>
                <thead>
                    <tr>
                        <td>Название</td>
                        <td>Раннеров</td>
                    </tr>
                </thead>
                <tbody>
                    {runners?.map((worker) => (
                        <tr key={worker.name}>
                            <td>{worker.name}</td>
                            <td>
                                <input
                                    type='number'
                                    value={worker.count}
                                    onChange={(e) => {
                                        setRunners(
                                            runners.map((runner) =>
                                                runner.name == worker.name
                                                    ? {
                                                          name: worker.name,
                                                          count: e.target
                                                              .valueAsNumber,
                                                      }
                                                    : runner
                                            )
                                        )
                                    }}
                                />
                            </td>
                        </tr>
                    )) || null}
                </tbody>
            </table>
            <button
                className='app'
                onClick={() => {
                    doRequest({ runners_count: runners! })
                }}
                disabled={isLoading || !runners}
            >
                Установить количество раннеров
            </button>
        </div>
    )
}
