import { Link, useNavigate, useParams } from 'react-router-dom'
import {
    FileSystemInfo,
    useFSCreate,
    useFSGet,
    useFSUpdate,
} from '../../apiclient/api-fs'
import { useCallback, useEffect, useState } from 'react'
import { FSListLink } from '../../core/routing'
import { useAgentList } from '../../apiclient/api-agent'
import { FSEditorWidget } from './fs-editor-widget'
import { ContainerWidget, ErrorTextWidget } from '../../widgets/design-system'

export function FSEditorScreen() {
    const params = useParams()
    const fsID = decodeURIComponent(params.id ?? 'new')

    const navigate = useNavigate()

    const [fsGetResponse, fetchFS] = useFSGet()
    const [fsCreateResponse, doFSCreate] = useFSCreate()
    const [fsUpdateResponse, doFSUpdate] = useFSUpdate()

    const [data, setData] = useState<FileSystemInfo>({
        id: 'ffffffff-ffff-ffff-ffff-ffffffffffff',
        name: '',
        download_priority: 0,
        deduplicate_priority: 0,
        highway_enabled: false,
        created_at: new Date().toJSON(),
    })

    useEffect(() => {
        if (fsGetResponse.data) {
            setData(fsGetResponse.data!)
        }
    }, [fsGetResponse.data])

    const isExists = fsID != 'new'

    useEffect(() => {
        if (isExists) {
            fetchFS({ id: fsID })
        }
    }, [fetchFS, fsID, isExists])

    const useSave = useCallback(() => {
        if (isExists) {
            doFSUpdate(data)
        } else {
            doFSCreate(data).then(() => navigate(FSListLink()))
        }
    }, [doFSUpdate, doFSCreate, isExists, data])

    const [agentsResponse, getAgents] = useAgentList()
    useEffect(() => {
        getAgents({ has_fs: true })
    }, [getAgents])

    return (
        <ContainerWidget
            direction='column'
            gap='big'
        >
            <ErrorTextWidget value={fsGetResponse} />
            <ErrorTextWidget value={fsCreateResponse} />
            <ErrorTextWidget value={fsUpdateResponse} />
            <ErrorTextWidget value={agentsResponse} />

            <div>
                <Link
                    className='app-button'
                    to={FSListLink()}
                >
                    Список файловых систем
                </Link>
            </div>

            <ContainerWidget
                appContainer
                direction='column'
                gap='medium'
            >
                <b>Редактор FS</b>

                <FSEditorWidget
                    value={data}
                    onChange={setData}
                    agents={agentsResponse.data ?? undefined}
                />

                <button
                    className='app'
                    onClick={useSave}
                >
                    Сохранить
                </button>
            </ContainerWidget>
        </ContainerWidget>
    )
}
