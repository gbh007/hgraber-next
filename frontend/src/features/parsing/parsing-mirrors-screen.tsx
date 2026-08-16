import { useEffect } from 'react'
import {
    useParsingMirrorDelete,
    useParsingMirrorList,
} from '../../apiclient/api-parsing-mirror'
import { Link } from 'react-router-dom'
import { ParsingMirrorEditLink } from '../../core/routing'
import { ContainerWidget, ErrorTextWidget } from '../../widgets/design-system'

export function ParsingMirrorsScreen() {
    const [parsingMirrorsResponse, fetchParsingMirrors] = useParsingMirrorList()
    useEffect(() => {
        fetchParsingMirrors()
    }, [fetchParsingMirrors])

    const [parsingMirrorDeleteResponse, doDeleteParsingMirror] =
        useParsingMirrorDelete()

    return (
        <ContainerWidget
            appContainer
            direction='column'
        >
            <ErrorTextWidget value={parsingMirrorsResponse} />
            <ErrorTextWidget value={parsingMirrorDeleteResponse} />
            <table>
                <thead>
                    <tr>
                        <td>
                            <ContainerWidget
                                direction='row'
                                gap='small'
                            >
                                <span>Название</span>
                                <Link
                                    className='app-button'
                                    to={ParsingMirrorEditLink()}
                                >
                                    новое
                                </Link>
                            </ContainerWidget>
                        </td>
                        <td>Описание</td>
                        <td>Значения</td>
                        <td>Действия</td>
                    </tr>
                </thead>
                <tbody>
                    {parsingMirrorsResponse.data?.mirrors?.map(
                        (labelPreset) => (
                            <tr key={labelPreset.id}>
                                <td>{labelPreset.name}</td>
                                <td>{labelPreset.description}</td>
                                <td>
                                    <ContainerWidget direction='column'>
                                        {labelPreset.prefixes.map((v) => (
                                            <span key={v}>{v}</span>
                                        ))}
                                    </ContainerWidget>
                                </td>
                                <td>
                                    <ContainerWidget
                                        direction='column'
                                        gap='smaller'
                                    >
                                        <Link
                                            className='app-button'
                                            to={ParsingMirrorEditLink(
                                                labelPreset.id
                                            )}
                                        >
                                            Редактировать
                                        </Link>
                                        <button
                                            className='app'
                                            onClick={() => {
                                                doDeleteParsingMirror({
                                                    id: labelPreset.id,
                                                }).then(() =>
                                                    fetchParsingMirrors()
                                                )
                                            }}
                                            disabled={
                                                parsingMirrorDeleteResponse.isLoading
                                            }
                                        >
                                            Удалить
                                        </button>
                                    </ContainerWidget>
                                </td>
                            </tr>
                        )
                    )}
                </tbody>
            </table>
        </ContainerWidget>
    )
}
