import { useCallback, useEffect, useState } from 'react'
import { Link, useNavigate, useParams } from 'react-router-dom'
import { MassloadFlagEditLink, MassloadFlagListLink } from '../../core/routing'
import { ContainerWidget, ErrorTextWidget } from '../../widgets/design-system'
import {
    MassloadFlag,
    useMassloadFlagCreate,
    useMassloadFlagGet,
    useMassloadFlagUpdate,
} from '../../apiclient/api-massload'
import { MassloadFlagEditorWidget } from './massload-flag-editor-widget'

export function MassloadFlagEditorScreen() {
    const params = useParams()
    const attributeCode = decodeURIComponent(params.code ?? '')

    const navigate = useNavigate()

    const [data, setData] = useState<MassloadFlag>({
        code: attributeCode,
        name: '',
        order_weight: 0,
        created_at: new Date().toJSON(),
    })

    const [massloadFlagGetResponse, fetchMassloadFlagGet] = useMassloadFlagGet()
    const [massloadFlagCreateResponse, doMassloadFlagCreate] =
        useMassloadFlagCreate()
    const [massloadFlagUpdateResponse, doMassloadFlagUpdate] =
        useMassloadFlagUpdate()

    useEffect(() => {
        if (massloadFlagGetResponse.data) {
            setData(massloadFlagGetResponse.data!)
        }
    }, [massloadFlagGetResponse.data])

    const isExists = attributeCode.length > 0

    useEffect(() => {
        if (isExists) {
            fetchMassloadFlagGet({
                code: attributeCode,
            })
        }
    }, [fetchMassloadFlagGet, attributeCode, isExists])

    const useSave = useCallback(() => {
        if (isExists) {
            doMassloadFlagUpdate(data)
        } else {
            doMassloadFlagCreate(data).then(() => {
                navigate(MassloadFlagEditLink(data.code))
            })
        }
    }, [doMassloadFlagUpdate, doMassloadFlagCreate, navigate, isExists, data])

    return (
        <ContainerWidget
            appContainer
            direction='column'
            gap='big'
        >
            <ErrorTextWidget value={massloadFlagGetResponse} />
            <ErrorTextWidget value={massloadFlagCreateResponse} />
            <ErrorTextWidget value={massloadFlagUpdateResponse} />

            <Link
                className='app-button'
                to={MassloadFlagListLink()}
            >
                Список
            </Link>

            <MassloadFlagEditorWidget
                onChange={setData}
                value={data}
                isNew={!isExists}
            />
            <button
                className='app'
                onClick={useSave}
            >
                Сохранить
            </button>
        </ContainerWidget>
    )
}
