import { useEffect, useState } from 'react'
import { Link, useParams } from 'react-router-dom'
import { BookRaw } from '../../apiclient/model-book'
import { useBookRaw, useBookUpdate } from '../../apiclient/api-book'
import { useLabelPresetList } from '../../apiclient/api-labels'
import {
    useAttributeColorList,
    useAttributeOriginCount,
} from '../../apiclient/api-attribute'
import { BookDetailsLink } from '../../core/routing'
import { BookEditorWidget } from './book-editor-widget'
import { ContainerWidget, ErrorTextWidget } from '../../widgets/design-system'

export function BookEditorScreen() {
    const params = useParams()
    const bookID = params.bookID!

    const [book, setBook] = useState<BookRaw>({
        create_at: new Date().toJSON(),
        id: '',
        name: '',
        page_count: 0,
    })

    const [bookRawResponse, fetchBookRaw] = useBookRaw()
    const [bookUpdateResponse, doBookUpdate] = useBookUpdate()
    const [labelPresetsResponse, fetchLabelPresets] = useLabelPresetList()

    const [attributeOriginCountResponse, getAttributeOriginCount] =
        useAttributeOriginCount()

    const [attributeColorListResponse, fetchAttributeColorList] =
        useAttributeColorList()

    useEffect(() => {
        fetchBookRaw({ id: bookID })
    }, [fetchBookRaw, bookID])

    useEffect(() => {
        fetchLabelPresets()
    }, [fetchLabelPresets])
    useEffect(() => {
        getAttributeOriginCount()
    }, [getAttributeOriginCount])

    useEffect(() => {
        if (bookRawResponse.data) {
            setBook(bookRawResponse.data)
        }
    }, [bookRawResponse.data])

    useEffect(() => {
        fetchAttributeColorList()
    }, [fetchAttributeColorList])
    return (
        <ContainerWidget
            direction='column'
            gap='medium'
        >
            <ErrorTextWidget value={bookRawResponse} />
            <ErrorTextWidget value={bookUpdateResponse} />
            <ErrorTextWidget value={labelPresetsResponse} />
            <ErrorTextWidget value={attributeOriginCountResponse} />
            <ContainerWidget
                direction='row'
                gap='medium'
            >
                <Link
                    className='app-button'
                    to={BookDetailsLink(bookID)}
                >
                    На страницу книги
                </Link>
                <button
                    className='app'
                    onClick={() => {
                        doBookUpdate(book).then(() =>
                            fetchBookRaw({ id: bookID })
                        )
                    }}
                >
                    Сохранить
                </button>
            </ContainerWidget>
            <BookEditorWidget
                value={book}
                onChange={(e) => setBook(e)}
                labelsAutoComplete={labelPresetsResponse.data?.presets}
                attributeCount={attributeOriginCountResponse.data?.attributes}
                attributeColors={attributeColorListResponse.data?.colors}
            />
        </ContainerWidget>
    )
}
