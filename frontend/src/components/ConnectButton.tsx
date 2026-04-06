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
      className="flex items-center gap-2 px-4 py-2 text-sm font-medium rounded-xl bg-indigo-600 hover:bg-indigo-500 disabled:opacity-40 disabled:cursor-not-allowed transition-all duration-150 shadow-lg shadow-indigo-500/20 active:scale-95"
    >
      <span className="text-base leading-none">🦊</span>
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
      <div className="flex items-center gap-2">
        <div className="flex items-center gap-2 px-3 py-1.5 rounded-xl bg-white/5 border border-white/10 text-sm">
          <span className="w-2 h-2 rounded-full bg-emerald-400 shadow-sm shadow-emerald-400/50"></span>
          <span className="font-mono text-gray-300 text-xs">
            {address?.slice(0, 6)}…{address?.slice(-4)}
          </span>
        </div>
        <button
          onClick={() => disconnect()}
          className="px-3 py-1.5 text-xs font-medium rounded-xl border border-white/10 text-gray-400 hover:text-white hover:bg-white/5 transition-all duration-150"
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