import { useCallback, useEffect, useState } from 'react'
import { useAgentList } from '../../apiclient/api-agent'
import {
    FSTransferRequest,
    useFSDelete,
    useFSList,
    useFSRemoveMismatch,
    useFSTransfer,
    useFSValidate,
} from '../../apiclient/api-fs'
import { Link } from 'react-router-dom'
import { FSEditLink } from '../../core/routing'
import { FSInfoWidget } from './fs-info-widget'
import {
    ColorizedTextWidget,
    ContainerWidget,
    ErrorTextWidget,
} from '../../widgets/design-system'

export function FSListScreen() {
    const [fsListResponse, fetchFSList] = useFSList()
    const [fsDeleteResponse, doFSDelete] = useFSDelete()
    const [fsValidateResponse, doFSValidate] = useFSValidate()
    const [fsRemoveMismatchResponse, doFSRemoveMismatch] = useFSRemoveMismatch()
    const [fsTransferResponse, doFSTransfer] = useFSTransfer()
    const [agentListResponse, fetchAgentList] = useAgentList()
    const [includeAvailableSize, setIncludeAvailableSize] = useState(true)
    const [includeDbFileSize, setIncludeDbFileSize] = useState(true)

    const fetchFS = useCallback(() => {
        fetchFSList({
            include_db_file_size: includeDbFileSize,
            include_available_size: includeAvailableSize,
        })
    }, [fetchFSList, includeDbFileSize, includeAvailableSize])

    useEffect(() => {
        fetchFS()
    }, [fetchFS])
    useEffect(() => {
        fetchAgentList({ include_status: true })
    }, [fetchAgentList])

    const [transferRequest, setTransferRequest] = useState<FSTransferRequest>({
        from: '',
        to: '',
    })

    return (
        <ContainerWidget
            direction='column'
            gap='big'
        >
            <ErrorTextWidget value={fsListResponse} />
            <ErrorTextWidget value={fsDeleteResponse} />
            <ErrorTextWidget value={fsValidateResponse} />
            <ErrorTextWidget value={fsRemoveMismatchResponse} />
            <ErrorTextWidget value={fsTransferResponse} />
            <ErrorTextWidget value={agentListResponse} />

            <ContainerWidget
                appContainer
                direction='row'
                gap='medium'
                wrap
            >
                <Link
                    className='app-button'
                    to={FSEditLink()}
                >
                    Новая
                </Link>
                <label>
                    <input
                        className='app'
                        type='checkbox'
                        checked={includeAvailableSize}
                        onChange={(e) =>
                            setIncludeAvailableSize(e.target.checked)
                        }
                    />
                    <span>Показывать доступное место</span>
                </label>
                <label>
                    <input
                        className='app'
                        type='checkbox'
                        checked={includeDbFileSize}
                        onChange={(e) => setIncludeDbFileSize(e.target.checked)}
                    />
                    <span>Показывать занятое место в БД</span>
                </label>
                <button
                    className='app'
                    onClick={() => {
                        fetchFS()
                        fetchAgentList({ include_status: true })
                    }}
                >
                    Обновить данные
                </button>
            </ContainerWidget>

            <ContainerWidget
                appContainer
                direction='column'
                gap='medium'
            >
                <b>Перенести файлы</b>
                <ContainerWidget
                    direction='2-column'
                    gap='medium'
                >
                    <span>Из</span>
                    <select
                        className='app'
                        value={transferRequest.from}
                        onChange={(e) =>
                            setTransferRequest({
                                ...transferRequest,
                                from: e.target.value,
                            })
                        }
                    >
                        <option value=''>Не выбрана</option>
                        {fsListResponse.data?.file_systems?.map((fs) => (
                            <option
                                key={fs.info.id}
                                value={fs.info.id}
                            >
                                {fs.info.name}
                            </option>
                        ))}
                    </select>
                    <span>В</span>
                    <select
                        className='app'
                        value={transferRequest.to}
                        onChange={(e) =>
                            setTransferRequest({
                                ...transferRequest,
                                to: e.target.value,
                            })
                        }
                    >
                        <option value=''>Не выбрана</option>
                        {fsListResponse.data?.file_systems?.map((fs) => (
                            <option
                                key={fs.info.id}
                                value={fs.info.id}
                            >
                                {fs.info.name}
                            </option>
                        ))}
                    </select>
                </ContainerWidget>
                <label>
                    <input
                        type='checkbox'
                        checked={transferRequest.only_preview_pages ?? false}
                        onChange={(e) =>
                            setTransferRequest({
                                ...transferRequest,
                                only_preview_pages: e.target.checked,
                            })
                        }
                    />
                    <span>Только превью</span>
                </label>
                <button
                    className='app'
                    onClick={() => {
                        if (
                            !confirm(
                                'Перенести файлы между файловыми системами?'
                            )
                        ) {
                            return
                        }

                        doFSTransfer(transferRequest)
                    }}
                    disabled={fsTransferResponse.isLoading}
                >
                    <ColorizedTextWidget color='danger-lite'>
                        Перенести файлы
                    </ColorizedTextWidget>
                </button>
            </ContainerWidget>

            {fsListResponse.data?.file_systems?.map((fs) => (
                <FSInfoWidget
                    key={fs.info.id}
                    value={fs}
                    onDelete={() => {
                        if (
                            !confirm(
                                'Удалить файловую систему и все ее файлы (ЭТО НЕОБРАТИМО)?'
                            )
                        ) {
                            return
                        }

                        doFSDelete({ id: fs.info.id }).then(() => {
                            fetchFS()
                        })
                    }}
                    validationLoading={fsValidateResponse.isLoading}
                    onValidate={() => {
                        doFSValidate({ id: fs.info.id })
                    }}
                    onRemoveMismatch={() => {
                        if (
                            !confirm(
                                'Удалить рассинхронизированные файлы (ЭТО НЕОБРАТИМО)?'
                            )
                        ) {
                            return
                        }

                        doFSRemoveMismatch({ id: fs.info.id })
                    }}
                    agents={agentListResponse.data ?? undefined}
                />
            ))}
        </ContainerWidget>
    )
}
