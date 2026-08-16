import { Link } from 'react-router-dom'
import { AgentListResponse } from '../../apiclient/api-agent'
import { FSListResponseUnit } from '../../apiclient/api-fs'
import { FSEditLink } from '../../core/routing'
import { AgentStatusWidget } from '../../widgets/agent'
import {
    ColorizedTextWidget,
    ContainerWidget,
    HumanTimeWidget,
    prettySize,
    PrettySizeWidget,
} from '../../widgets/design-system'

export function FSInfoWidget(props: {
    value: FSListResponseUnit
    onDelete: () => void
    onValidate: () => void
    validationLoading: boolean
    onRemoveMismatch: () => void
    agents?: AgentListResponse[]
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
                    <span>{props.value.info.name}</span>
                </ContainerWidget>
            </h3>
            <ContainerWidget
                direction='2-column'
                gap='medium'
            >
                <b>ID:</b>
                <span>{props.value.info.id}</span>

                {props.value.info.description ? (
                    <>
                        <b>Описание:</b>
                        <span>{props.value.info.description}</span>
                    </>
                ) : null}

                {props.value.info.agent_id ? (
                    <>
                        <ContainerWidget
                            direction='row'
                            gap='medium'
                        >
                            <b>Агент ID:</b>
                            {props.agents
                                ?.filter(
                                    (agent) =>
                                        agent.info.id ==
                                        props.value.info.agent_id
                                )
                                .map((agent) => (
                                    <AgentStatusWidget
                                        key={agent.info.id}
                                        value={agent.status?.status}
                                    />
                                ))}
                        </ContainerWidget>
                        <span>{props.value.info.agent_id}</span>
                    </>
                ) : null}

                {props.value.info.path ? (
                    <>
                        <b>Путь на локальной ФС:</b>
                        <span>{props.value.info.path}</span>
                    </>
                ) : null}

                <b>Приоритет загрузки:</b>
                <span>{props.value.info.download_priority}</span>
                <b>Приоритет дедупликации:</b>
                <span>{props.value.info.deduplicate_priority}</span>

                <b>Включен highway:</b>
                <span>{props.value.info.highway_enabled ? 'Да' : 'Нет'}</span>
                <b>Адрес highway:</b>
                <span>{props.value.info.highway_addr}</span>

                {props.value.db_files_info ? (
                    <>
                        <b>Файлов:</b>
                        <span>
                            {`${prettySize(props.value.db_files_info.size)} (${props.value.db_files_info.count} шт)`}
                        </span>
                    </>
                ) : null}

                {props.value.db_invalid_files_info?.count ? (
                    <>
                        <b>Невалидных файлов:</b>
                        <span>
                            {`${prettySize(props.value.db_invalid_files_info.size)} (${props.value.db_invalid_files_info.count} шт)`}
                        </span>
                    </>
                ) : null}

                {props.value.db_detached_files_info?.count ? (
                    <>
                        <b>Неиспользуемых файлов:</b>
                        <span>
                            {`${prettySize(props.value.db_detached_files_info.size)} (${props.value.db_detached_files_info.count} шт)`}
                        </span>
                    </>
                ) : null}
                <PrettySizeWidget value={props.value.available_size}>
                    <b>Свободное место:</b>
                </PrettySizeWidget>

                <b>Создан:</b>
                <HumanTimeWidget value={props.value.info.created_at} />
            </ContainerWidget>
            <ContainerWidget
                direction='row'
                gap='medium'
                wrap
            >
                {!props.value.is_legacy ? (
                    <>
                        <Link
                            className='app-button'
                            to={FSEditLink(props.value.info.id)}
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
                    </>
                ) : null}
                <button
                    className='app'
                    onClick={() => props.onValidate()}
                    disabled={props.validationLoading}
                >
                    Запустить валидацию файлов
                </button>
                <button
                    className='app'
                    onClick={() => props.onRemoveMismatch()}
                >
                    <ColorizedTextWidget
                        bold
                        color='danger'
                    >
                        Удалить рассинхронизированные файлы
                    </ColorizedTextWidget>
                </button>
            </ContainerWidget>
        </ContainerWidget>
    )
}
