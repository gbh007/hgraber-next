import { useContext, useEffect, useState } from 'react'
import { ThemeContext } from '../../core/context'
import { useUserLogin } from '../../apiclient/api-user-login'
import { useAppSettings } from '../../apiclient/settings'
import { ContainerWidget, ErrorTextWidget } from '../../widgets/design-system'

export function SettingsScreen() {
    const theme = useContext(ThemeContext)
    const [loginResponse, doLogin] = useUserLogin()

    const [appSettings, setSettings] = useAppSettings()

    const [token, setToken] = useState('')
    const [pageCount, setPageCount] = useState(appSettings.book_on_page)

    useEffect(() => {
        setPageCount(appSettings.book_on_page)
    }, [appSettings, setPageCount])

    useEffect(() => {
        localStorage.setItem('theme', theme.theme)
    }, [theme.theme])

    return (
        <ContainerWidget
            direction='column'
            gap='big'
        >
            <ContainerWidget
                appContainer
                direction='column'
                gap='medium'
            >
                <b>Данные приложения</b>
                <ContainerWidget
                    direction='2-column'
                    gap='small'
                >
                    <span>Количество на странице</span>
                    <input
                        className='app'
                        type='number'
                        id='on-page'
                        value={pageCount}
                        onChange={(e) => setPageCount(parseInt(e.target.value))}
                    />
                </ContainerWidget>
                <div>
                    <button
                        className='app'
                        onClick={() =>
                            setSettings({
                                ...appSettings,
                                book_on_page: pageCount,
                            })
                        }
                    >
                        Сохранить
                    </button>
                </div>
            </ContainerWidget>

            <ContainerWidget
                appContainer
                direction='row'
                gap='smaller'
            >
                <ErrorTextWidget value={loginResponse} />
                <input
                    className='app'
                    value={token}
                    onChange={(e) => setToken(e.target.value)}
                    placeholder='Токен'
                    id='token'
                    type='password'
                />
                <button
                    className='app'
                    onClick={() => {
                        doLogin({ token }).then(() => setToken(''))
                    }}
                >
                    Авторизоваться
                </button>
            </ContainerWidget>
            <ContainerWidget
                appContainer
                direction='column'
                gap='medium'
            >
                <b>Выбрать тему</b>
                <label>
                    <input
                        className='app'
                        type='radio'
                        checked={theme.theme == 'light'}
                        onChange={() => theme.setTheme('light')}
                    />
                    <span>Светлая</span>
                </label>
                <label>
                    <input
                        className='app'
                        type='radio'
                        checked={theme.theme == 'dark'}
                        onChange={() => theme.setTheme('dark')}
                    />
                    <span>Темная</span>
                </label>
            </ContainerWidget>
        </ContainerWidget>
    )
}
