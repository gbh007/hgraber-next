import { RouterProvider, Link, createHashRouter } from 'react-router-dom'
import { PropsWithChildren, useState } from 'react'
import {
    AgentListLink,
    AttributeColorListLink,
    AttributeRemapListLink,
    BookListLink,
    FSListLink,
    HProxyListLink,
    LabelPresetsLink,
    MainScreenLink,
    MassloadFlagListLink,
    MassloadListLink,
    MenuLink,
    ParsingMirrorsLink,
    RPCLink,
    SelectToCompareLink,
    SettingsLink,
    TasksLink,
} from './core/routing'
import { AppTheme, ThemeContext } from './core/context'
import { HProxyBookScreen } from './features/hproxy/hproxy-book-screen'
import { HProxyListScreen } from './features/hproxy/hproxy-list-screen'
import { MassloadListScreen } from './features/massload/massload-list-screen'
import { MassloadEditorScreen } from './features/massload/massload-editor-screen'
import { MassloadViewScreen } from './features/massload/massload-view-screen'
import { FSListScreen } from './features/fs/fs-list-screen'
import { FSEditorScreen } from './features/fs/fs-editor-screen'
import { AgentListScreen } from './features/agent/agent-list-screen'
import { AgentEditorScreen } from './features/agent/agent-editor-screen'
import { AttributeColorListScreen } from './features/attribute/attribute-color-list-screen'
import { AttributeColorEditorScreen } from './features/attribute/attribute-color-editor-screen'
import { AttributeRemapListScreen } from './features/attribute/attribute-remap-list-screen'
import { LabelPresetsScreen } from './features/label/label-presets-screen'
import { LabelPresetEditorScreen } from './features/label/label-preset-editor-screen'
import { BookLabelsEditorScreen } from './features/label/book-labels-editor-screen'
import { ParsingMirrorsScreen } from './features/parsing/parsing-mirrors-screen'
import { ParsingMirrorEditorScreen } from './features/parsing/parsing-mirror-editor-screen'
import { ListScreen } from './features/book/list-screen'
import { BookDetailsScreen } from './features/book/book-details-screen'
import { BookReadScreen } from './features/book/book-read-screen'
import { BookEditorScreen } from './features/book/book-editor-screen'
import { BookRebuilderScreen } from './features/rebulder/book-rebuilder-screen'
import { CompareBookScreen } from './features/deduplication/compare-book-screen'
import { UniqueBookPagesScreen } from './features/deduplication/unique-book-pages-screen'
import { BooksByPageScreen } from './features/deduplication/books-by-page-screen'
import { SelectToCompareScreen } from './features/deduplication/select-to-compare-screen'
import { MainScreen } from './features/system/main-screen'
import { SettingsScreen } from './features/system/settings-screen'
import { RPCScreen } from './features/system/rpc-screen'
import { TaskScreen } from './features/system/task-screen'
import { MassloadFlagListScreen } from './features/massload/massload-flag-list-screen'
import { MassloadFlagEditorScreen } from './features/massload/massload-flag-editor-screen'

const router = createHashRouter([
    {
        path: '/',
        element: (
            <SimpleWrapper>
                <MainScreen />
            </SimpleWrapper>
        ),
    },
    {
        path: '/settings',
        element: (
            <SimpleWrapper>
                <SettingsScreen />
            </SimpleWrapper>
        ),
    },
    {
        path: '/menu',
        element: (
            <SimpleWrapper>
                <MenuWidget />
            </SimpleWrapper>
        ),
    },
    {
        path: '/rpc',
        element: (
            <SimpleWrapper>
                <RPCScreen />
            </SimpleWrapper>
        ),
    },
    {
        path: '/agent/list',
        element: (
            <SimpleWrapper>
                <AgentListScreen />
            </SimpleWrapper>
        ),
    },
    {
        path: '/agent/edit/:id',
        element: (
            <SimpleWrapper>
                <AgentEditorScreen />
            </SimpleWrapper>
        ),
    },
    {
        path: '/list',
        element: (
            <SimpleWrapper>
                <ListScreen />
            </SimpleWrapper>
        ),
    },
    {
        path: '/book/:bookID',
        element: (
            <SimpleWrapper>
                <BookDetailsScreen />
            </SimpleWrapper>
        ),
    },
    {
        path: '/book/:bookID/read/:page',
        element: (
            <SimpleWrapper>
                <BookReadScreen />
            </SimpleWrapper>
        ),
    },
    {
        path: '/label/presets',
        element: (
            <SimpleWrapper>
                <LabelPresetsScreen />
            </SimpleWrapper>
        ),
    },
    {
        path: '/label/preset/edit/:name',
        element: (
            <SimpleWrapper>
                <LabelPresetEditorScreen />
            </SimpleWrapper>
        ),
    },
    {
        path: '/label/preset/edit',
        element: (
            <SimpleWrapper>
                <LabelPresetEditorScreen />
            </SimpleWrapper>
        ),
    },
    {
        path: '/tasks',
        element: (
            <SimpleWrapper>
                <TaskScreen />
            </SimpleWrapper>
        ),
    },
    {
        path: '/book/:origin/compare/:target',
        element: (
            <SimpleWrapper>
                <CompareBookScreen />
            </SimpleWrapper>
        ),
    },
    {
        path: '/book/:bookID/unique-pages',
        element: (
            <SimpleWrapper>
                <UniqueBookPagesScreen />
            </SimpleWrapper>
        ),
    },
    {
        path: '/deduplicate/:bookID/:page',
        element: (
            <SimpleWrapper>
                <BooksByPageScreen />
            </SimpleWrapper>
        ),
    },
    {
        path: '/book/:bookID/edit',
        element: (
            <SimpleWrapper>
                <BookEditorScreen />
            </SimpleWrapper>
        ),
    },
    {
        path: '/book/:bookID/rebuild',
        element: (
            <SimpleWrapper>
                <BookRebuilderScreen />
            </SimpleWrapper>
        ),
    },
    {
        path: '/select-to-compare',
        element: (
            <SimpleWrapper>
                <SelectToCompareScreen />
            </SimpleWrapper>
        ),
    },
    {
        path: '/book/:bookID/labels',
        element: (
            <SimpleWrapper>
                <BookLabelsEditorScreen />
            </SimpleWrapper>
        ),
    },
    {
        path: '/attribute/color/list',
        element: (
            <SimpleWrapper>
                <AttributeColorListScreen />
            </SimpleWrapper>
        ),
    },
    {
        path: '/attribute/color/edit',
        element: (
            <SimpleWrapper>
                <AttributeColorEditorScreen />
            </SimpleWrapper>
        ),
    },
    {
        path: '/attribute/color/edit/:code/:value',
        element: (
            <SimpleWrapper>
                <AttributeColorEditorScreen />
            </SimpleWrapper>
        ),
    },
    {
        path: '/fs/list',
        element: (
            <SimpleWrapper>
                <FSListScreen />
            </SimpleWrapper>
        ),
    },
    {
        path: '/fs/edit/:id',
        element: (
            <SimpleWrapper>
                <FSEditorScreen />
            </SimpleWrapper>
        ),
    },
    {
        path: '/parsing/mirrors',
        element: (
            <SimpleWrapper>
                <ParsingMirrorsScreen />
            </SimpleWrapper>
        ),
    },
    {
        path: '/parsing/mirror/edit/:id',
        element: (
            <SimpleWrapper>
                <ParsingMirrorEditorScreen />
            </SimpleWrapper>
        ),
    },
    {
        path: '/parsing/mirror/edit',
        element: (
            <SimpleWrapper>
                <ParsingMirrorEditorScreen />
            </SimpleWrapper>
        ),
    },
    {
        path: '/attribute/remap/list',
        element: (
            <SimpleWrapper>
                <AttributeRemapListScreen />
            </SimpleWrapper>
        ),
    },
    {
        path: '/hproxy/list',
        element: (
            <SimpleWrapper>
                <HProxyListScreen />
            </SimpleWrapper>
        ),
    },
    {
        path: '/hproxy/book',
        element: (
            <SimpleWrapper>
                <HProxyBookScreen />
            </SimpleWrapper>
        ),
    },
    {
        path: '/massload/list',
        element: (
            <SimpleWrapper>
                <MassloadListScreen />
            </SimpleWrapper>
        ),
    },
    {
        path: '/massload/edit/:id',
        element: (
            <SimpleWrapper>
                <MassloadEditorScreen />
            </SimpleWrapper>
        ),
    },
    {
        path: '/massload/view/:id',
        element: (
            <SimpleWrapper>
                <MassloadViewScreen />
            </SimpleWrapper>
        ),
    },
    {
        path: '/massload/flag/list',
        element: (
            <SimpleWrapper>
                <MassloadFlagListScreen />
            </SimpleWrapper>
        ),
    },
    {
        path: '/massload/flag/edit',
        element: (
            <SimpleWrapper>
                <MassloadFlagEditorScreen />
            </SimpleWrapper>
        ),
    },
    {
        path: '/massload/flag/edit/:code',
        element: (
            <SimpleWrapper>
                <MassloadFlagEditorScreen />
            </SimpleWrapper>
        ),
    },
])

function App() {
    const [theme, setTheme] = useState<AppTheme>(
        localStorage.getItem('theme') == 'dark' ? 'dark' : 'light'
    )

    return (
        <div
            data-theme={theme}
            style={{
                flexGrow: 1,
                display: 'flex',
                flexDirection: 'column',
                backgroundColor: 'var(--app-background)',
                color: 'var(--app-color)',
            }}
        >
            <ThemeContext.Provider
                value={{
                    theme: theme,
                    setTheme: setTheme,
                }}
            >
                <RouterProvider router={router} />
            </ThemeContext.Provider>
        </div>
    )
}

function SimpleWrapper(props: PropsWithChildren) {
    return (
        <>
            <div className='app-header'>
                <Link to={MainScreenLink()}>Главная</Link>
                <Link to={BookListLink()}>Список книг</Link>
                <Link to={HProxyListLink()}>HProxy</Link>
                <Link to={MenuLink()}>Меню</Link>
            </div>

            <div className='app-body'>{props.children}</div>
        </>
    )
}

function MenuWidget() {
    return (
        <div>
            <ol>
                <li>
                    <Link to={MainScreenLink()}>Главная</Link>
                </li>
                <li>
                    <Link to={BookListLink()}>Список книг</Link>
                </li>
                <li>
                    <Link to={AgentListLink()}>Агенты</Link>
                </li>
                <li>
                    <Link to={RPCLink()}>RPC</Link>
                </li>
                <li>
                    <Link to={TasksLink()}>Задачи</Link>
                </li>
                <li>
                    <Link to={LabelPresetsLink()}>Пресеты меток</Link>
                </li>
                <li>
                    <Link to={SelectToCompareLink()}>
                        Выбрать книги для сравнения
                    </Link>
                </li>
                <li>
                    <Link to={AttributeColorListLink()}>Цвета аттрибутов</Link>
                </li>
                <li>
                    <Link to={AttributeRemapListLink()}>
                        Ремапинг аттрибутов
                    </Link>
                </li>
                <li>
                    <Link to={FSListLink()}>Файловые системы</Link>
                </li>
                <li>
                    <Link to={ParsingMirrorsLink()}>Зеркала для парсинга</Link>
                </li>
                <li>
                    <Link to={HProxyListLink()}>HProxy</Link>
                </li>
                <li>
                    <Link to={MassloadListLink()}>Массовые загрузки</Link>
                </li>
                <li>
                    <Link to={MassloadFlagListLink()}>
                        Флаги массовых загрузок
                    </Link>
                </li>
                <li>
                    <Link to={SettingsLink()}>Настройки</Link>
                </li>
            </ol>
        </div>
    )
}

export default App
