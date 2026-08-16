export function AgentStatusWidget(props: { value?: string }) {
    const color =
        props.value == 'ok'
            ? 'var(--app-agent-status-color-ok)'
            : props.value == 'error'
              ? 'var(--app-agent-status-color-error)'
              : props.value == 'warning'
                ? 'var(--app-agent-status-color-warning)'
                : props.value == 'offline'
                  ? 'var(--app-agent-status-color-offline)'
                  : 'var(--app-agent-status-color-unknown)'

    return (
        <div>
            <div
                style={{
                    backgroundColor: color,
                    padding: '10px',
                    borderRadius: '10px',
                }}
            />
        </div>
    )
}
