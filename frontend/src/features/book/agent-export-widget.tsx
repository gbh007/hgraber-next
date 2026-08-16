import { useEffect, useState } from 'react'
import { useAgentList, useAgentTaskExport } from '../../apiclient/api-agent'
import { BookFilter } from '../../apiclient/model-book-filter'
import { ContainerWidget, ErrorTextWidget } from '../../widgets/design-system'

export function AgentExportWidget(props: { filter: BookFilter }) {
    const [agentsResponse, getAgents] = useAgentList()
    const [agentID, setAgentID] = useState('')
    const [deleteAfterExport, setDeleteAfterExport] = useState(false)
    const [exportResponse, makeExport] = useAgentTaskExport()

    useEffect(() => {
        getAgents({ can_export: true })
    }, [getAgents])

    return (
        <details className='app'>
            <summary>Параметры экспорта</summary>
            <ContainerWidget
                direction='row'
                gap='medium'
            >
                <ErrorTextWidget value={agentsResponse} />

                <select
                    className='app'
                    value={agentID}
                    onChange={(e) => {
                        setAgentID(e.target.value)
                    }}
                >
                    <option value=''>Не выбран</option>
                    {agentsResponse.data?.map((agent) => (
                        <option
                            value={agent.info.id}
                            key={agent.info.id}
                        >
                            {agent.info.name}
                        </option>
                    ))}
                </select>
                <label>
                    <span>Удалить после экспорта</span>
                    <input
                        className='app'
                        checked={deleteAfterExport}
                        placeholder='Удалить после экспорта'
                        type='checkbox'
                        autoComplete='off'
                        onChange={(e) => {
                            setDeleteAfterExport(e.target.checked)
                        }}
                    />
                </label>
                <ErrorTextWidget value={exportResponse} />
                <button
                    className='app'
                    disabled={exportResponse.isLoading || !agentID}
                    onClick={() => {
                        if (
                            deleteAfterExport &&
                            !confirm(
                                'Книги будут удалены из системы после экспорта, продолжить?'
                            )
                        ) {
                            return
                        }

                        makeExport({
                            book_filter: {
                                ...props.filter,
                                pagination: undefined,
                            }, // Принудительно срезаем параметры пагинации.
                            delete_after: deleteAfterExport,
                            exporter: agentID,
                        })
                    }}
                >
                    Выгрузить
                </button>
            </ContainerWidget>
        </details>
    )
}
