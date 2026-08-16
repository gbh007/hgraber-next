import { useEffect, useState } from 'react'
import { HProxyBookResponsePage } from '../../apiclient/api-hproxy'
import { ContainerWidget } from '../../widgets/design-system'
import {
    ImageSize,
    PageImagePreviewWidget,
    PreviewSizeWidget,
} from '../../widgets/book'

export function HProxyBookPagesPreviewWidget(props: {
    url: string
    pages?: Array<HProxyBookResponsePage>
    pageLimit?: number
}) {
    const [pageLimit, setPageLimit] = useState(20)
    const [imageSize, setImageSize] = useState<ImageSize>('medium')

    useEffect(() => {
        setPageLimit(props.pageLimit ?? 20)
    }, [setPageLimit, props.pageLimit, props.url])

    const scrollToTop = () => {
        document
            .getElementById('HProxyBookPagesPreviewWidgetTop')
            ?.scrollIntoView({ behavior: 'smooth' })
    }

    if (!props.pages?.length) {
        return null
    }

    const notAllPages = pageLimit != -1 && pageLimit < props.pages.length

    return (
        <ContainerWidget
            direction='column'
            gap='medium'
            id='HProxyBookPagesPreviewWidgetTop'
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
                            <PageImagePreviewWidget
                                previewSize={imageSize}
                                preview_url={page.preview_url}
                            />
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
