import { BookRebuildRequestFlags } from '../../apiclient/api-book'
import {
    ColorizedTextWidget,
    ContainerWidget,
} from '../../widgets/design-system'

export function BuilderFlagsWidget(props: {
    value?: BookRebuildRequestFlags
    onChange: (v: BookRebuildRequestFlags) => void
}) {
    return (
        <ContainerWidget
            direction='column'
            gap='small'
        >
            <label>
                <input
                    className='app'
                    type='checkbox'
                    checked={props.value?.set_origin_labels ?? false}
                    onChange={(e) =>
                        props.onChange({
                            ...props.value,
                            set_origin_labels: e.target.checked,
                        })
                    }
                />
                <ColorizedTextWidget color='good'>
                    Проставить каждой страницы метки об оригинальной книге если
                    ее нет
                </ColorizedTextWidget>
            </label>
            <label>
                <input
                    className='app'
                    type='checkbox'
                    checked={props.value?.auto_verify ?? false}
                    onChange={(e) =>
                        props.onChange({
                            ...props.value,
                            auto_verify: e.target.checked,
                        })
                    }
                />
                <ColorizedTextWidget color='good'>
                    Проставить билду статус подтвержденной книги
                </ColorizedTextWidget>
            </label>
            <label>
                <input
                    className='app'
                    type='checkbox'
                    checked={props.value?.only_unique ?? false}
                    onChange={(e) =>
                        props.onChange({
                            ...props.value,
                            only_unique: e.target.checked,
                        })
                    }
                />
                Оставить только уникальные страницы в результате (без дублей)
            </label>
            <label>
                <input
                    className='app'
                    type='checkbox'
                    checked={props.value?.exclude_dead_hash_pages ?? false}
                    onChange={(e) =>
                        props.onChange({
                            ...props.value,
                            exclude_dead_hash_pages: e.target.checked,
                        })
                    }
                />
                Исключить страницы с мертвыми хешами
            </label>
            <label>
                <input
                    className='app'
                    type='checkbox'
                    checked={props.value?.only_1_copy ?? false}
                    onChange={(e) =>
                        props.onChange({
                            ...props.value,
                            only_1_copy: e.target.checked,
                        })
                    }
                />
                Только уникальные страницы в системе (без копий) и без дублей
            </label>
            <label>
                <input
                    className='app'
                    type='checkbox'
                    checked={props.value?.page_re_order ?? false}
                    onChange={(e) =>
                        props.onChange({
                            ...props.value,
                            page_re_order: e.target.checked,
                        })
                    }
                />
                Применить новый порядок страниц
            </label>
            <label>
                <input
                    className='app'
                    type='checkbox'
                    checked={props.value?.extract_mode ?? false}
                    onChange={(e) =>
                        props.onChange({
                            ...props.value,
                            extract_mode: e.target.checked,
                        })
                    }
                />
                <ColorizedTextWidget color='danger-lite'>
                    Режим экстракции - вынос страниц в новую книгу с удалением
                    их только из исходной
                </ColorizedTextWidget>
            </label>
            <label>
                <input
                    className='app'
                    type='checkbox'
                    checked={
                        props.value?.mark_unused_pages_as_dead_hash ?? false
                    }
                    onChange={(e) =>
                        props.onChange({
                            ...props.value,
                            mark_unused_pages_as_dead_hash: e.target.checked,
                        })
                    }
                />
                <ColorizedTextWidget color='danger-lite'>
                    Отметить страницы что не вошли в ребилд как мертвый хеш
                </ColorizedTextWidget>
            </label>
            <label>
                <input
                    className='app'
                    type='checkbox'
                    checked={props.value?.mark_unused_pages_as_deleted ?? false}
                    onChange={(e) =>
                        props.onChange({
                            ...props.value,
                            mark_unused_pages_as_deleted: e.target.checked,
                        })
                    }
                />
                <ColorizedTextWidget color='danger'>
                    Удалить страницы что не вошли в ребилд и их копии в системе
                </ColorizedTextWidget>
            </label>
            <label>
                <input
                    className='app'
                    type='checkbox'
                    checked={
                        props.value
                            ?.mark_empty_book_as_deleted_after_remove_pages ??
                        false
                    }
                    onChange={(e) =>
                        props.onChange({
                            ...props.value,
                            mark_empty_book_as_deleted_after_remove_pages:
                                e.target.checked,
                        })
                    }
                />
                <ColorizedTextWidget color='danger-lite'>
                    Отметить удаленным книги что остались без страниц после их
                    удаления
                </ColorizedTextWidget>
            </label>
        </ContainerWidget>
    )
}
