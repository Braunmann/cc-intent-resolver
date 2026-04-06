import { useConnect, useConnectors, useConnection, useDisconnect } from 'wagmi'
import { useState, useEffect } from 'react'
import type { Connector } from 'wagmi'

function WalletOption({ connector, onClick }: { connector: Connector; onClick: () => void }) {
  const [ready, setReady] = useState(false)

  useEffect(() => {
    ;(async () => {
      const provider = await connector.getProvider()
      setReady(!!provider)
    })()
  }, [connector])

  return (
    <button
      disabled={!ready}
      onClick={onClick}
      className="px-4 py-2 text-sm bg-blue-600 rounded-lg hover:bg-blue-700 disabled:opacity-50"
    >
      {connector.name}
    </button>
  )
}

export function ConnectButton() {
  const { isConnected, address } = useConnection()
  const { connect } = useConnect()
  const { disconnect } = useDisconnect()
  const connectors = useConnectors()

  if (isConnected) {
    return (
      <div className="flex items-center gap-3">
        <span className="text-sm text-gray-400">
          {address?.slice(0, 6)}...{address?.slice(-4)}
        </span>
        <button
          onClick={() => disconnect()}
          className="px-4 py-2 text-sm border border-gray-600 rounded-lg hover:bg-gray-800"
        >
          Disconnect
        </button>
      </div>
    )
  }

  return (
    <div className="flex gap-2">
      {connectors.map((connector) => (
        <WalletOption
          key={connector.uid}
          connector={connector}
          onClick={() => connect({ connector })}
        />
      ))}
    </div>
  )
}