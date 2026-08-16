import { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'
import {
    ParsingMirrorUpdateRequest,
    useParsingMirrorCreate,
    useParsingMirrorGet,
    useParsingMirrorUpdate,
} from '../../apiclient/api-parsing-mirror'
import {
    ContainerWidget,
    ErrorTextWidget,
    StringArrayPickerWidget,
} from '../../widgets/design-system'

export function ParsingMirrorEditorScreen() {
    const params = useParams()
    const mirrorID = params.id ?? ''

    const [parsingMirrorCreateResponse, doCreateParsingMirror] =
        useParsingMirrorCreate()
    const [parsingMirrorUpdateResponse, doUpdateParsingMirror] =
        useParsingMirrorUpdate()
    const [parsingMirrorGetResponse, doGetParsingMirror] = useParsingMirrorGet()

    const [data, setData] = useState<ParsingMirrorUpdateRequest>({
        id: '',
        name: '',
        prefixes: [],
    })

    useEffect(() => {
        if (parsingMirrorGetResponse.data) {
            setData(parsingMirrorGetResponse.data!)
        }
    }, [parsingMirrorGetResponse.data])

    if (mirrorID) {
        useEffect(() => {
            doGetParsingMirror({ id: mirrorID })
        }, [doGetParsingMirror, mirrorID])
    }

    return (
        <ContainerWidget
            appContainer
            direction='column'
            gap='small'
        >
            <ErrorTextWidget value={parsingMirrorCreateResponse} />
            <ErrorTextWidget value={parsingMirrorUpdateResponse} />
            <ErrorTextWidget value={parsingMirrorGetResponse} />
            <input
                className='app'
                value={data.name}
                onChange={(e) => {
                    setData({ ...data, name: e.target.value })
                }}
            />
            <textarea
                className='app'
                value={data.description}
                onChange={(e) => {
                    setData({ ...data, description: e.target.value })
                }}
            />
            <StringArrayPickerWidget
                value={data.prefixes}
                onChange={(e) => {
                    setData({ ...data, prefixes: e })
                }}
            />
            <button
                className='app'
                onClick={() => {
                    if (mirrorID) {
                        doUpdateParsingMirror(data)
                    } else {
                        doCreateParsingMirror(data) // TODO: сделать роутинг на редактирование
                    }
                }}
            >
                Сохранить
            </button>
        </ContainerWidget>
    )
}
