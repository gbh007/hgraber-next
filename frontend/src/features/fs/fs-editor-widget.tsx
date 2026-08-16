import { AgentListResponse } from '../../apiclient/api-agent'
import { FileSystemInfo } from '../../apiclient/api-fs'
import { ContainerWidget } from '../../widgets/design-system'

export function FSEditorWidget(props: {
    value: FileSystemInfo
    onChange: (v: FileSystemInfo) => void
    agents?: Array<AgentListResponse>
}) {
    return (
        <ContainerWidget
            direction='2-column'
            gap='medium'
        >
            <span>Название</span>
            <input
                className='app'
                placeholder='Название'
                type='text'
                autoComplete='off'
                value={props.value.name}
                onChange={(e) => {
                    props.onChange({ ...props.value, name: e.target.value })
                }}
            />

            <span>Описание</span>
            <input
                className='app'
                placeholder='Описание'
                type='text'
                autoComplete='off'
                value={props.value.description}
                onChange={(e) => {
                    props.onChange({
                        ...props.value,
                        description: e.target.value,
                    })
                }}
            />

            <span>Agent ID</span>
            <select
                className='app'
                value={props.value.agent_id}
                onChange={(e) => {
                    props.onChange({ ...props.value, agent_id: e.target.value })
                }}
            >
                <option value=''>Не выбран</option>
                {props.agents?.map((agent) => (
                    <option
                        value={agent.info.id}
                        key={agent.info.id}
                    >
                        {agent.info.name}
                    </option>
                ))}
            </select>

            <span>Путь в локальной системе</span>
            <input
                className='app'
                placeholder='Путь в локальной системе'
                type='text'
                autoComplete='off'
                value={props.value.path}
                onChange={(e) => {
                    props.onChange({ ...props.value, path: e.target.value })
                }}
            />

            <span>Приоритет загрузки</span>
            <input
                className='app'
                placeholder='Приоритет загрузки'
                type='number'
                autoComplete='off'
                value={props.value.download_priority}
                onChange={(e) => {
                    props.onChange({
                        ...props.value,
                        download_priority: e.target.valueAsNumber,
                    })
                }}
            />

            <span>Приоритет дедупликации</span>
            <input
                className='app'
                placeholder='Приоритет дедупликации'
                type='number'
                autoComplete='off'
                value={props.value.deduplicate_priority}
                onChange={(e) => {
                    props.onChange({
                        ...props.value,
                        deduplicate_priority: e.target.valueAsNumber,
                    })
                }}
            />

            <span>Флаги</span>
            <ContainerWidget
                direction='column'
                gap='small'
            >
                <label>
                    <input
                        className='app'
                        placeholder='Включен highway'
                        type='checkbox'
                        autoComplete='off'
                        checked={props.value.highway_enabled}
                        onChange={(e) => {
                            props.onChange({
                                ...props.value,
                                highway_enabled: e.target.checked,
                            })
                        }}
                    />
                    <span>Включен highway</span>
                </label>
            </ContainerWidget>

            <span>Внешний адрес highway</span>
            <input
                className='app'
                placeholder='Внешний адрес highway'
                type='text'
                autoComplete='off'
                value={props.value.highway_addr}
                onChange={(e) => {
                    props.onChange({
                        ...props.value,
                        highway_addr: e.target.value,
                    })
                }}
            />
        </ContainerWidget>
    )
}
