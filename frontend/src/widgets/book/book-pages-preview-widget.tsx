import { useEffect, useState } from 'react'
import { BookSimplePage } from '../../apiclient/model-book'
import { ImageSize } from './image-size'
import {
    PageImagePreviewWidget,
    PreviewSizeWidget,
} from './page-image-preview-widget'
import { Link } from 'react-router-dom'
import { BookReaderLink } from '../../core/routing'
import { ContainerWidget } from '../design-system'

export function BookPagesPreviewWidget(props: {
    bookID: string
    pages?: Array<BookSimplePage>
    pageLimit?: number
}) {
    const [pageLimit, setPageLimit] = useState(20)

    useEffect(() => {
        setPageLimit(props.pageLimit ?? 20)
    }, [setPageLimit, props.pageLimit, props.bookID])

    const [imageSize, setImageSize] = useState<ImageSize>('medium')

    const scrollToTop = () => {
        document.getElementById('BookPagesPreviewWidgetTop')?.scrollIntoView({
            behavior: 'smooth',
        })
    }

    if (!props.pages?.length) {
        return null
    }

    const notAllPages = pageLimit != -1 && pageLimit < props.pages.length

    return (
        <ContainerWidget
            direction='column'
            gap='medium'
            id='BookPagesPreviewWidgetTop'
        >
            <ContainerWidget
                appContainer
                direction='row'
                gap='medium'
            >
                {notAllPages ? (
                    <button
                        className='app'
                        onClick={() => setPageLimit(-1)}
                    >
                        Показать все страницы
                    </button>
                ) : null}
                <PreviewSizeWidget
                    value={imageSize}
                    onChange={setImageSize}
                />
            </ContainerWidget>
            <ContainerWidget
                direction='row'
                gap='medium'
                wrap
            >
                {props.pages
                    ?.filter((page) => page.preview_url)
                    .filter((_, i) => pageLimit == -1 || i < pageLimit)
                    .map((page) => (
                        <ContainerWidget
                            appContainer
                            direction='column'
                            style={{ flexGrow: 1, alignItems: 'center' }}
                            key={page.page_number}
                        >
                            <Link
                                to={BookReaderLink(
                                    props.bookID,
                                    page.page_number
                                )}
                            >
                                <PageImagePreviewWidget
                                    previewSize={imageSize}
                                    flags={page}
                                    preview_url={page.preview_url}
                                />
                            </Link>
                        </ContainerWidget>
                    ))}
            </ContainerWidget>
            {!notAllPages ? (
                <button
                    className='app'
                    onClick={scrollToTop}
                >
                    Наверх
                </button>
            ) : null}
        </ContainerWidget>
    )
}
