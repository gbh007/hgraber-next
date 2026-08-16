import { Link } from 'react-router-dom'
import {
    AgentListResponse,
    AgentListResponseStatusProblems,
} from '../../apiclient/api-agent'
import { AgentStatusWidget } from '../../widgets/agent'
import { AgentEditLink } from '../../core/routing'
import {
    ColorizedTextWidget,
    ContainerWidget,
    HumanTimeWidget,
} from '../../widgets/design-system'

export function AgentInfoWidget(props: {
    value: AgentListResponse
    onDelete: () => void
}) {
    return (
        <ContainerWidget
            appContainer
            direction='column'
            gap='medium'
        >
            <h3>
                <ContainerWidget
                    direction='row'
                    gap='medium'
                >
                    <AgentStatusWidget value={props.value.status?.status} />
                    <span>{props.value.info.name}</span>
                </ContainerWidget>
            </h3>
            <ContainerWidget
                direction='2-column'
                gap='medium'
            >
                {props.value.status?.start_at ? (
                    <>
                        <b>Запущен:</b>
                        <HumanTimeWidget value={props.value.status.start_at} />
                    </>
                ) : null}
                <b>ID:</b>
                <span>{props.value.info.id}</span>
                <b>Адрес:</b>
                <span>{props.value.info.addr}</span>
                <b>Может обрабатывать новые:</b>
                <span>{props.value.info.can_parse ? 'Да' : 'Нет'}</span>
                <b>Может обрабатывать новые массово:</b>
                <span>{props.value.info.can_parse_multi ? 'Да' : 'Нет'}</span>
                <b>Может экспортировать:</b>
                <span>{props.value.info.can_export ? 'Да' : 'Нет'}</span>
                <b>Есть ФС:</b>
                <span>{props.value.info.has_fs ? 'Да' : 'Нет'}</span>
                <b>Есть HProxy:</b>
                <span>{props.value.info.has_hproxy ? 'Да' : 'Нет'}</span>
                <b>Приоритет:</b>
                <span>{props.value.info.priority}</span>
                <b>Создан:</b>
                <HumanTimeWidget value={props.value.info.created_at} />
            </ContainerWidget>
            <AgentStatusInfoWidget value={props.value.status?.problems} />
            <ContainerWidget
                direction='row'
                gap='medium'
            >
                <Link
                    className='app-button'
                    to={AgentEditLink(props.value.info.id)}
                >
                    Редактировать
                </Link>
                <button
                    className='app'
                    onClick={() => props.onDelete()}
                >
                    <ColorizedTextWidget
                        bold
                        color='danger'
                    >
                        Удалить
                    </ColorizedTextWidget>
                </button>
            </ContainerWidget>
        </ContainerWidget>
    )
}

function AgentStatusInfoWidget(props: {
    value?: Array<AgentListResponseStatusProblems>
}) {
    return (
        <ContainerWidget
            direction='column'
            gap='small'
        >
            {props.value?.map((problem, i) => (
                <span key={i}>
                    {problem.type}: {problem.details}
                </span>
            ))}
        </ContainerWidget>
    )
}
