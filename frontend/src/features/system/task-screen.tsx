import { useEffect, useState } from 'react'
import { useTaskCreate, useTaskResults } from '../../apiclient/api-task'
import {
    ContainerWidget,
    ErrorTextWidget,
    prettyPercent,
} from '../../widgets/design-system'

export function TaskScreen() {
    const [taskCode, setTaskCode] = useState('')

    const [createTaskResponse, doCreateTask] = useTaskCreate()
    const [taskResultsResponse, fetchTaskResults] = useTaskResults()

    useEffect(() => {
        fetchTaskResults()
    }, [fetchTaskResults])

    return (
        <ContainerWidget
            direction='column'
            gap='medium'
        >
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
            <ContainerWidget
                appContainer
                direction='row'
                gap='smaller'
            >
                <ErrorTextWidget value={createTaskResponse} />
                <select
                    className='app'
                    value={taskCode}
                    onChange={(e) => setTaskCode(e.target.value)}
                >
                    <option value=''>Не выбрано</option>
                    {taskResultsResponse.data?.tasks.map((task) => (
                        <option
                            value={task.code}
                            key={task.code}
                            title={task.description}
                        >
                            {task.name}
                        </option>
                    ))}
                </select>
                <button
                    className='app'
                    onClick={() => {
                        doCreateTask({ code: taskCode })
                    }}
                >
                    Создать задачу
                </button>
                <button
                    className='app'
                    onClick={() => {
                        fetchTaskResults()
                    }}
                >
                    Обновить
                </button>
            </ContainerWidget>
            <ErrorTextWidget value={taskResultsResponse} />
            {taskResultsResponse.data?.results?.map((taskResult, i) => (
                <ContainerWidget
                    appContainer
                    key={i}
                    direction='column'
                    gap='small'
                >
                    <h4>{taskResult.name}</h4>
                    <span>{taskResult.duration_formatted}</span>
                    {taskResult.error ? <span>{taskResult.error}</span> : null}
                    {taskResult.result ? <pre>{taskResult.result}</pre> : null}
                    {taskResult.stages?.map((stage, i) => (
                        <ContainerWidget
                            direction='column'
                            gap='smaller'
                            key={i}
                        >
                            <h5>{stage.name}</h5>
                            <span>{stage.duration_formatted}</span>
                            <span>
                                {stage.progress}/{stage.total} (
                                {prettyPercent(stage.progress / stage.total)}%)
                            </span>
                            {stage.error ? <span>{stage.error}</span> : null}
                            {stage.result ? <pre>{stage.result}</pre> : null}
                        </ContainerWidget>
                    ))}
                </ContainerWidget>
            ))}
        </ContainerWidget>
    )
}
