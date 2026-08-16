import { useCallback, useEffect, useState } from 'react'
import {
    useAttributeColorList,
    useAttributeCount,
} from '../../apiclient/api-attribute'
import {
    MassloadInfo,
    useMassloadAttributeCreate,
    useMassloadAttributeDelete,
    useMassloadExternalLinkCreate,
    useMassloadExternalLinkDelete,
    useMassloadExternalLinkUpdate,
    useMassloadFlagList,
    useMassloadInfoCreate,
    useMassloadInfoDelete,
    useMassloadInfoGet,
    useMassloadInfoUpdate,
} from '../../apiclient/api-massload'
import { Link, useNavigate, useParams } from 'react-router-dom'
import {
    MassloadEditorLink,
    MassloadListLink,
    MassloadViewLink,
} from '../../core/routing'
import { MassloadInfoEditorWidget } from './massload-info-editor-widget'
import { MassloadAttributeEditorWidget } from './massload-attribute-editor-widget'
import { MassloadExternalLinkEditorWidget } from './massload-external-link-editor-widget'
import { ContainerWidget, ErrorTextWidget } from '../../widgets/design-system'
import { BookAttributeAutocompleteWidget } from '../../widgets/attribute'

export function MassloadEditorScreen() {
    const [attributeColorListResponse, fetchAttributeColorList] =
        useAttributeColorList()
    const [attributeCountResponse, getAttributeCount] = useAttributeCount()
    const [massloadFlagListResponse, fetchMassloadFlagList] =
        useMassloadFlagList()

    useEffect(() => {
        fetchAttributeColorList()
    }, [fetchAttributeColorList])
    useEffect(() => {
        getAttributeCount()
    }, [getAttributeCount])
    useEffect(() => {
        fetchMassloadFlagList()
    }, [fetchMassloadFlagList])

    const params = useParams()
    const massloadID = params.id == 'new' ? 0 : parseInt(params.id ?? '0')
    const isExists = massloadID > 0

    const navigate = useNavigate()

    const [data, setData] = useState<MassloadInfo>({
        id: 0,
        name: '',
        created_at: new Date().toJSON(),
    })

    const [massLoadCreateResponse, doMassloadCreate] = useMassloadInfoCreate()
    const [massLoadUpdateResponse, doMassloadUpdate] = useMassloadInfoUpdate()
    const [massLoadGetResponse, fetchMassloadGet] = useMassloadInfoGet()
    const [massLoadDeleteResponse, doMassloadDelete] = useMassloadInfoDelete()

    const [massLoadAttributeCreateResponse, doMassloadAttributeCreate] =
        useMassloadAttributeCreate()
    const [massLoadAttributeDeleteResponse, doMassloadAttributeDelete] =
        useMassloadAttributeDelete()

    const [massLoadExternalLinkCreateResponse, doMassloadExternalLinkCreate] =
        useMassloadExternalLinkCreate()
    const [massLoadExternalLinkUpdateResponse, doMassloadExternalLinkUpdate] =
        useMassloadExternalLinkUpdate()
    const [massLoadExternalLinkDeleteResponse, doMassloadExternalLinkDelete] =
        useMassloadExternalLinkDelete()

    useEffect(() => {
        if (massLoadGetResponse.data) {
            setData(massLoadGetResponse.data!)
        }
    }, [massLoadGetResponse.data])

    useEffect(() => {
        if (massLoadCreateResponse.data?.id ?? 0 > 0) {
            navigate(MassloadEditorLink(massLoadCreateResponse.data?.id))
        }
    }, [massLoadCreateResponse.data?.id])

    useEffect(() => {
        if (isExists) {
            fetchMassloadGet({ id: massloadID })
        }
    }, [fetchMassloadGet, massloadID, isExists])

    const useSave = useCallback(() => {
        if (isExists) {
            doMassloadUpdate(data).then(() => {
                fetchMassloadGet({ id: massloadID })
            })
        } else {
            doMassloadCreate(data).then(() => {
                // FIXME: работает криво, нужно в промис докидывать данные ответа.
                // navigate(MassloadEditorLink(massLoadCreateResponse.data?.id))
            })
        }
    }, [doMassloadUpdate, doMassloadCreate, navigate, isExists, data])

    return (
        <ContainerWidget
            direction='column'
            gap='big'
        >
            <ErrorTextWidget value={attributeColorListResponse} />
            <ErrorTextWidget value={attributeCountResponse} />
            <ErrorTextWidget value={massLoadCreateResponse} />
            <ErrorTextWidget value={massLoadUpdateResponse} />
            <ErrorTextWidget value={massLoadGetResponse} />
            <ErrorTextWidget value={massLoadDeleteResponse} />
            <ErrorTextWidget value={massLoadAttributeCreateResponse} />
            <ErrorTextWidget value={massLoadAttributeDeleteResponse} />
            <ErrorTextWidget value={massLoadExternalLinkCreateResponse} />
            <ErrorTextWidget value={massLoadExternalLinkUpdateResponse} />
            <ErrorTextWidget value={massLoadExternalLinkDeleteResponse} />
            <ErrorTextWidget value={massloadFlagListResponse} />

            <ContainerWidget
                direction='row'
                gap='medium'
                wrap
            >
                <Link
                    className='app-button'
                    to={MassloadListLink()}
                >
                    Список
                </Link>
                {isExists ? (
                    <Link
                        className='app-button'
                        to={MassloadViewLink(massloadID)}
                    >
                        Посмотреть
                    </Link>
                ) : null}
            </ContainerWidget>

            <MassloadInfoEditorWidget
                onChange={setData}
                onDelete={() => {
                    if (
                        !confirm(
                            `Удалить массовую загрузку ${data.name} (${data.id})?`
                        )
                    ) {
                        return
                    }

                    doMassloadDelete({ id: data.id }).then(() =>
                        navigate(MassloadListLink())
                    )
                }}
                value={data}
                onSave={useSave}
                flagInfos={massloadFlagListResponse.data?.flags}
            />
            {isExists ? (
                <>
                    <MassloadAttributeEditorWidget
                        value={data.attributes}
                        onCreate={(code: string, value: string) => {
                            doMassloadAttributeCreate({
                                code: code,
                                massload_id: data.id,
                                value: value,
                            }).then(() => {
                                fetchMassloadGet({ id: massloadID })
                            })
                        }}
                        onDelete={(code: string, value: string) => {
                            doMassloadAttributeDelete({
                                code: code,
                                massload_id: data.id,
                                value: value,
                            }).then(() => {
                                fetchMassloadGet({ id: massloadID })
                            })
                        }}
                        colors={attributeColorListResponse.data?.colors}
                    />
                    <MassloadExternalLinkEditorWidget
                        value={data.external_links}
                        onCreate={(v) => {
                            doMassloadExternalLinkCreate({
                                url: v.url,
                                auto_check: v.auto_check,
                                massload_id: data.id,
                            }).then(() => {
                                fetchMassloadGet({ id: massloadID })
                            })
                        }}
                        onUpdate={(v) => {
                            doMassloadExternalLinkUpdate({
                                url: v.url,
                                auto_check: v.auto_check,
                                massload_id: data.id,
                            }).then(() => {
                                fetchMassloadGet({ id: massloadID })
                            })
                        }}
                        onDelete={(url: string) => {
                            doMassloadExternalLinkDelete({
                                url: url,
                                massload_id: data.id,
                            }).then(() => {
                                fetchMassloadGet({ id: massloadID })
                            })
                        }}
                    />
                </>
            ) : null}

            <BookAttributeAutocompleteWidget
                attributeCount={attributeCountResponse.data?.attributes}
            />
        </ContainerWidget>
    )
}
