import { Fragment } from 'react'
import { Link, useLocation, useNavigate } from 'react-router-dom'
import { buildBreadcrumbs } from '../../core/breadcrumbs'
import { ContainerWidget } from './container-widget'

export function BreadcrumbWidget() {
    const location = useLocation()
    const navigate = useNavigate()

    const crumbs = buildBreadcrumbs(location.pathname)

    if (crumbs.length < 2) {
        return null
    }

    return (
        <ContainerWidget
            appContainer
            direction='row'
            gap='medium'
            wrap
        >
            <button
                className='app'
                onClick={() => navigate(-1)}
                title='Вернуться назад'
            >
                ← Назад
            </button>

            {crumbs.map((crumb, index) => (
                <Fragment key={crumb.path}>
                    {index > 0 && ' / '}
                    {index == crumbs.length - 1 ? (
                        <span className='app-button app-disabled'>
                            {crumb.label}
                        </span>
                    ) : (
                        <Link
                            className='app-button'
                            to={crumb.path}
                        >
                            {crumb.label}
                        </Link>
                    )}
                </Fragment>
            ))}
        </ContainerWidget>
    )
}
