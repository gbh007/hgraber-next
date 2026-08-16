import { createContext } from 'react'

export type AppTheme = 'light' | 'dark'

export type ThemeContextType = {
    theme: AppTheme
    setTheme: (theme: AppTheme) => void
}

export const ThemeContext = createContext<ThemeContextType>({
    theme: 'light',
    setTheme: () => {},
})
