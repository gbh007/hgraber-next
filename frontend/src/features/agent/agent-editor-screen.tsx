import { useNavigate, useParams } from 'react-router-dom'
import {
    Agent,
    useAgentGet,
    useAgentNew,
    useAgentUpdate,
} from '../../apiclient/api-agent'
import { useCallback, useEffect, useState } from 'react'
import { AgentListLink } from '../../core/routing'
import { AgentEditorWidget } from './agent-editor-widget'
import { ContainerWidget, ErrorTextWidget } from '../../widgets/design-system'

export function AgentEditorScreen() {
    const params = useParams()
    const agentID = decodeURIComponent(params.id ?? 'new')

    const navigate = useNavigate()

    const [agentGetResponse, fetchAgent] = useAgentGet()
    const [agentNewResponse, doAgentNew] = useAgentNew()
    const [agentUpdateResponse, doAgentUpdate] = useAgentUpdate()

    const [data, setData] = useState<Agent>({
        id: '',
        addr: '',
        can_export: false,
        can_parse: false,
        can_parse_multi: false,
        has_fs: false,
        has_hproxy: false,
        name: '',
        priority: 0,
        token: '',
        created_at: new Date().toJSON(),
    })

    useEffect(() => {
        if (agentGetResponse.data) {
            setData(agentGetResponse.data!)
        }
    }, [agentGetResponse.data])

    const isExists = agentID != 'new'

    useEffect(() => {
        if (isExists) {
            fetchAgent({ id: agentID })
        }
    }, [fetchAgent, agentID, isExists])

    const useSave = useCallback(() => {
        if (isExists) {
            doAgentUpdate(data)
        } else {
            doAgentNew(data).then(() => navigate(AgentListLink()))
        }
    }, [doAgentUpdate, doAgentNew, isExists, data])

    return (
        <div>
            <ErrorTextWidget value={agentGetResponse} />
            <ErrorTextWidget value={agentNewResponse} />
            <ErrorTextWidget value={agentUpdateResponse} />

            <ContainerWidget
                appContainer
                direction='column'
                gap='medium'
            >
                <b>Редактор агента</b>

                <AgentEditorWidget
                    value={data}
                    onChange={setData}
                />

                <button
                    className='app'
                    onClick={useSave}
                >
                    Сохранить
                </button>
            </ContainerWidget>
        </div>
    )
}
