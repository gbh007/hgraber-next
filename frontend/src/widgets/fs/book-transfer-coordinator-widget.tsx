import { useEffect, useState } from 'react'
import { useFSList, useFSTransferBook } from '../../apiclient/api-fs'
import {
    ColorizedTextWidget,
    ContainerWidget,
    DialogWidget,
    ErrorTextWidget,
} from '../design-system'

export function BookTransferCoordinatorWidget(props: {
    bookID: string
    pageNumber?: number
}) {
    const [show, setShow] = useState(false)
    const [fsID, setFSID] = useState('')
    const [onlyPreview, setOnlyPreview] = useState(false)

    const [fsListResponse, fetchFSList] = useFSList()
    const [fsTransferBookResponse, doFSTransferBook] = useFSTransferBook()

    useEffect(() => {
        setShow(false)
    }, [props.bookID, props.pageNumber])

    return (
        <>
            <button
                className='app'
                onClick={() => {
                    fetchFSList({})
                    setShow(true)
                }}
            >
                Перенести между ФС
            </button>
            <DialogWidget
                open={show}
                onClose={() => setShow(false)}
            >
                <ErrorTextWidget value={fsListResponse} />
                <ErrorTextWidget value={fsTransferBookResponse} />

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
                        {props.pageNumber ? (
                            <>
                                <span>Страница</span>
                                <span>{props.pageNumber}</span>
                            </>
                        ) : null}
                        <span>В</span>
                        <select
                            className='app'
                            value={fsID}
                            onChange={(e) => setFSID(e.target.value)}
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
                            checked={onlyPreview}
                            onChange={(e) => setOnlyPreview(e.target.checked)}
                        />
                        <span>Только превью</span>
                    </label>
                    <button
                        className='app'
                        onClick={() => {
                            if (
                                !confirm(
                                    'Перенести книгу между файловыми системами?'
                                )
                            ) {
                                return
                            }

                            doFSTransferBook({
                                book_id: props.bookID,
                                to: fsID,
                                only_preview_pages: onlyPreview,
                                page_number: props.pageNumber,
                            })
                        }}
                        disabled={fsTransferBookResponse.isLoading}
                    >
                        <ColorizedTextWidget color='danger-lite'>
                            Перенести файлы
                        </ColorizedTextWidget>
                    </button>
                </ContainerWidget>
            </DialogWidget>
        </>
    )
}
