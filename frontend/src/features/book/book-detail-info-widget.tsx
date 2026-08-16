import { PropsWithChildren } from 'react'
import { BookDetails } from '../../apiclient/api-book-details'
import { DeduplicateBookByPageBodyResponseResult } from '../../apiclient/api-deduplicate'
import { AttributeColor } from '../../apiclient/api-attribute'
import { BookSizeWidget } from './book-size-widget'
import { BookDuplicates } from './book-duplicates'
import {
    ColorizedTextWidget,
    ContainerWidget,
    prettySize,
} from '../../widgets/design-system'
import {
    BookAttributesWidget,
    BookImagePreviewWidget,
    BookPagesPreviewWidget,
} from '../../widgets/book'
import { HProxyBookLink } from '../../core/routing'
import { Link } from 'react-router-dom'

export function BookDetailInfoWidget(
    props: PropsWithChildren & {
        book: BookDetails
        deduplicateBookInfo?: Array<DeduplicateBookByPageBodyResponseResult>
        colors?: Array<AttributeColor>
    }
) {
    const originDomain = /https?:\/\/([^\/]+)\/.*/.exec(
        props.book.info.origin_url ?? ''
    )?.[1]
    return (
        <ContainerWidget
            direction='column'
            gap='big'
        >
            <ContainerWidget
                appContainer
                direction='row'
            >
                <div>
                    <BookImagePreviewWidget
                        flags={props.book.info.flags}
                        previewSize='superbig'
                        preview_url={props.book.info.preview_url}
                    />
                </div>
                <ContainerWidget
                    direction='column'
                    gap='medium'
                    style={{ flexGrow: 1, padding: '10px' }}
                >
                    <h1 style={{ wordBreak: 'break-all', margin: 0 }}>
                        <ColorizedTextWidget
                            color={
                                props.book.info.flags.parsed_name
                                    ? undefined
                                    : 'danger'
                            }
                        >
                            {props.book.info.name}
                        </ColorizedTextWidget>
                    </h1>
                    <ContainerWidget
                        direction='row'
                        wrap
                        style={{ justifyContent: 'space-between' }}
                    >
                        <span> #{props.book.info.id} </span>
                        <ColorizedTextWidget
                            color={
                                props.book.info.flags.parsed_page
                                    ? undefined
                                    : 'danger'
                            }
                        >
                            Страниц: {props.book.info.page_count}
                            {props.book.pages &&
                            props.book.pages.length !=
                                props.book.info.page_count
                                ? ` (${props.book.pages.length})`
                                : null}
                        </ColorizedTextWidget>
                        <ColorizedTextWidget
                            color={
                                props.book.page_loaded_percent != 100.0
                                    ? 'danger'
                                    : undefined
                            }
                        >
                            Загружено: {props.book.page_loaded_percent}%
                        </ColorizedTextWidget>
                        <span>
                            {new Date(
                                props.book.info.created_at
                            ).toLocaleString()}
                        </span>
                    </ContainerWidget>
                    <ContainerWidget
                        direction='row'
                        gap='small'
                    >
                        {props.book.info.origin_url ? (
                            <a href={props.book.info.origin_url}>
                                Ссылка на первоисточник
                            </a>
                        ) : null}
                        {originDomain ? <span>({originDomain})</span> : null}
                        {props.book.info.origin_url ? (
                            <Link
                                className='app-button'
                                to={HProxyBookLink(props.book.info.origin_url)}
                            >
                                Открыть в HProxy
                            </Link>
                        ) : null}
                    </ContainerWidget>
                    <BookAttributesWidget
                        value={props.book.attributes}
                        colors={props.colors}
                    />
                    <BookSizeWidget value={props.book.size} />
                    {props.children}
                </ContainerWidget>
            </ContainerWidget>
            <ContainerWidget
                direction='row'
                gap='medium'
                wrap
            >
                {props.book.fs_disposition?.map((fs) => (
                    <ContainerWidget
                        key={fs.id}
                        appContainer
                        direction='column'
                        gap='small'
                    >
                        <b>{fs.name}</b>
                        <span>{fs.id}</span>
                        <span>
                            {prettySize(fs.files.size)} ({fs.files.count})
                        </span>
                    </ContainerWidget>
                ))}
            </ContainerWidget>
            <BookDuplicates
                deduplicateBookInfo={props.deduplicateBookInfo}
                originID={props.book.info.id}
            />
            <BookPagesPreviewWidget
                bookID={props.book.info.id}
                pages={props.book.pages}
            />
        </ContainerWidget>
    )
}
