import { matchPath } from 'react-router-dom'

export type Breadcrumb = {
    /** Route pattern of the screen, for example `/agent/edit/:id`. */
    path: string
    /** Human readable name of the screen. */
    label: string
    /** Route pattern of the parent screen. */
    parent?: string
}

const BREADCRUMBS: Breadcrumb[] = [
    {
        path: '/menu',
        label: 'Меню',
    },

    {
        path: '/agent/list',
        label: 'Агенты',
        parent: '/menu',
    },
    {
        path: '/agent/edit/:id',
        label: 'Редактор агента',
        parent: '/agent/list',
    },

    {
        path: '/rpc',
        label: 'RPC',
        parent: '/menu',
    },

    {
        path: '/tasks',
        label: 'Задачи',
        parent: '/menu',
    },

    {
        path: '/label/presets',
        label: 'Пресеты меток',
        parent: '/menu',
    },
    {
        path: '/label/preset/edit/:name',
        label: 'Редактор пресета',
        parent: '/label/presets',
    },
    {
        path: '/label/preset/edit',
        label: 'Новый пресет',
        parent: '/label/presets',
    },

    {
        path: '/attribute/color/list',
        label: 'Цвета аттрибутов',
        parent: '/menu',
    },
    {
        path: '/attribute/color/edit/:code/:value',
        label: 'Редактор цвета',
        parent: '/attribute/color/list',
    },
    {
        path: '/attribute/color/edit',
        label: 'Новый цвет',
        parent: '/attribute/color/list',
    },

    {
        path: '/attribute/remap/list',
        label: 'Ремапинг аттрибутов',
        parent: '/menu',
    },

    {
        path: '/fs/list',
        label: 'Файловые системы',
        parent: '/menu',
    },
    {
        path: '/fs/edit/:id',
        label: 'Редактор файловой системы',
        parent: '/fs/list',
    },

    {
        path: '/parsing/mirrors',
        label: 'Зеркала для парсинга',
        parent: '/menu',
    },
    {
        path: '/parsing/mirror/edit/:id',
        label: 'Редактор зеркала',
        parent: '/parsing/mirrors',
    },
    {
        path: '/parsing/mirror/edit',
        label: 'Новое зеркало',
        parent: '/parsing/mirrors',
    },

    {
        path: '/massload/list',
        label: 'Массовые загрузки',
        parent: '/menu',
    },
    {
        path: '/massload/edit/:id',
        label: 'Редактор загрузки',
        parent: '/massload/list',
    },
    {
        path: '/massload/view/:id',
        label: 'Просмотр загрузки',
        parent: '/massload/list',
    },

    {
        path: '/massload/flag/list',
        label: 'Флаги массовых загрузок',
        parent: '/menu',
    },
    {
        path: '/massload/flag/edit/:code',
        label: 'Редактор флага',
        parent: '/massload/flag/list',
    },
    {
        path: '/massload/flag/edit',
        label: 'Новый флаг',
        parent: '/massload/flag/list',
    },

    {
        path: '/settings',
        label: 'Настройки',
        parent: '/menu',
    },
]

function findBreadcrumb(pathname: string): Breadcrumb | null {
    for (const breadcrumb of BREADCRUMBS) {
        if (matchPath(breadcrumb.path, pathname)) {
            return breadcrumb
        }
    }

    return null
}

/** Build crumbs from root to target for provided pathname. */
export function buildBreadcrumbs(pathname: string): Breadcrumb[] {
    const crumbs: Breadcrumb[] = []

    let current: Breadcrumb | null = findBreadcrumb(pathname)

    while (current) {
        crumbs.unshift(current)

        current = current.parent ? findBreadcrumb(current.parent) : null
    }

    return crumbs
}
