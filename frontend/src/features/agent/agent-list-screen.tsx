import { useEffect } from 'react'
import { useAgentDelete, useAgentList } from '../../apiclient/api-agent'
import { Link } from 'react-router-dom'
import { AgentEditLink } from '../../core/routing'
import { AgentInfoWidget } from './agent-info-widget'
import { ContainerWidget, ErrorTextWidget } from '../../widgets/design-system'

export function AgentListScreen() {
    const [agentListResponse, fetchAgentList] = useAgentList()
    const [agentDeleteResponse, doAgentDelete] = useAgentDelete()

    useEffect(() => {
        fetchAgentList({ include_status: true })
    }, [fetchAgentList])

    return (
        <ContainerWidget
            direction='column'
            gap='big'
        >
            <ErrorTextWidget value={agentListResponse} />
            <ErrorTextWidget value={agentDeleteResponse} />

            <div>
                <Link
                    className='app-button'
                    to={AgentEditLink()}
                >
                    Новый
                </Link>
            </div>

            {agentListResponse.data?.map((agent) => (
                <AgentInfoWidget
                    key={agent.info.id}
                    value={agent}
                    onDelete={() => {
                        doAgentDelete({ id: agent.info.id }).then(() => {
                            fetchAgentList({ include_status: true })
                        })
                    }}
                />
            ))}
        </ContainerWidget>
    )
}
