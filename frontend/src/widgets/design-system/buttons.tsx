import { DeletedBadge } from './assets'

export function DeleteButtonWidget(props: {
    onClick: () => void
    disabled?: boolean
}) {
    return (
        <button
            className='app'
            onClick={props.onClick}
            disabled={props.disabled}
        >
            <img
                style={{
                    maxWidth: '18px',
                    maxHeight: '18px',
                    display: 'flex',
                    justifyContent: 'center',
                    alignItems: 'center',
                }}
                src={DeletedBadge}
            />
        </button>
    )
}
