import { Link, useParams, useSearchParams } from 'react-router-dom'
import {
    LabelDeleteRequest,
    LabelSetRequest,
    useLabelDelete,
    useLabelGet,
    useLabelPresetList,
    useLabelSet,
} from '../../apiclient/api-labels'
import { useEffect } from 'react'
import { BookDetailsLink } from '../../core/routing'
import { BookLabelEditorWidget } from './book-label-editor-widget'
import { ContainerWidget, ErrorTextWidget } from '../../widgets/design-system'

export function BookLabelsEditorScreen() {
    const params = useParams()
    const bookID = params.bookID!

    const [searchParams, setSearchParams] = useSearchParams()

    const pageNumber = parseInt(searchParams.get('pageNumber') ?? '0')

    const [labelsResponse, fetchLabels] = useLabelGet()
    const [labelPresetsResponse, fetchLabelPresets] = useLabelPresetList()
    const [labelSetResponse, doSetLabel] = useLabelSet()
    const [labelDeleteResponse, doDeleteLabel] = useLabelDelete()

    useEffect(() => {
        fetchLabelPresets()
    }, [fetchLabelPresets])

    useEffect(() => {
        fetchLabels({ book_id: bookID })
    }, [fetchLabels, bookID])

    const labels = labelsResponse.data?.labels?.filter(
        (label) => !pageNumber || label.page_number == pageNumber
    )

    return (
        <ContainerWidget
            direction='column'
            gap='big'
        >
            <ErrorTextWidget value={labelsResponse} />
            <ErrorTextWidget value={labelPresetsResponse} />
            <ErrorTextWidget value={labelSetResponse} />
            <ErrorTextWidget value={labelDeleteResponse} />
            <ContainerWidget
                appContainer
                direction='row'
                gap='medium'
            >
                <Link
                    className='app-button'
                    to={BookDetailsLink(bookID)}
                >
                    На страницу книги
                </Link>
                {pageNumber ? (
                    <>
                        <span>Выбрана страница {pageNumber}</span>
                        <button
                            className='app'
                            onClick={() => {
                                searchParams.delete('pageNumber')
                                setSearchParams(searchParams)
                            }}
                        >
                            смотреть все метки
                        </button>
                    </>
                ) : null}
            </ContainerWidget>
            <BookLabelEditorWidget
                bookID={bookID}
                onCreate={(v: LabelSetRequest) => {
                    doSetLabel(v).then(() => {
                        fetchLabels({ book_id: bookID })
                    })
                }}
                onDelete={(v: LabelDeleteRequest) => {
                    doDeleteLabel(v).then(() => {
                        fetchLabels({ book_id: bookID })
                    })
                }}
                value={labels}
                labelsAutoComplete={labelPresetsResponse.data?.presets}
            />
        </ContainerWidget>
    )
}
