import { useEffect, useState } from 'react'
import { config } from '../config/env'

type Intent = {
  ID: string
  Maker: string
  Status: number
  InputAmount: number
  Deadline: number
}

const STATUS_CONFIG: Record<number, { label: string; dot: string; badge: string }> = {
  0: { label: 'Created',   dot: 'bg-sky-400',     badge: 'bg-sky-400/10 text-sky-400 border-sky-400/20' },
  1: { label: 'Fulfilled', dot: 'bg-emerald-400',  badge: 'bg-emerald-400/10 text-emerald-400 border-emerald-400/20' },
  2: { label: 'Settled',   dot: 'bg-violet-400',   badge: 'bg-violet-400/10 text-violet-400 border-violet-400/20' },
  3: { label: 'Cancelled', dot: 'bg-rose-400',     badge: 'bg-rose-400/10 text-rose-400 border-rose-400/20' },
}

export function IntentList() {
  const [intents, setIntents] = useState<Intent[]>([])

  useEffect(() => {
    fetch(`${config.apiUrl}/v1/intents`)
      .then(r => r.json())
      .then(setIntents)
  }, [])

  return (
    <div className="rounded-2xl border border-white/8 bg-white/3 p-6 shadow-xl">
      <div className="flex items-center justify-between mb-5">
        <h2 className="text-sm font-semibold text-white">Intent History</h2>
        <span className="text-xs text-gray-600 tabular-nums">{intents.length} total</span>
      </div>

      <div className="flex flex-col gap-2">
        {intents.map(intent => {
          const status = STATUS_CONFIG[intent.Status] ?? STATUS_CONFIG[0]
          const deadline = new Date(intent.Deadline * 1000)
          const expired = Date.now() > intent.Deadline * 1000

          return (
            <div
              key={intent.ID}
              className="flex items-center justify-between rounded-xl border border-white/6 bg-white/2 px-4 py-3 hover:bg-white/5 transition-colors duration-150"
            >
              <div className="flex items-center gap-3 min-w-0">
                <span className={`shrink-0 w-2 h-2 rounded-full ${status.dot}`}></span>
                <span className="font-mono text-xs text-gray-400 truncate">
                  {intent.ID.slice(0, 18)}…
                </span>
              </div>
              <div className="flex items-center gap-4 shrink-0 ml-4">
                <span className="text-xs text-gray-400 tabular-nums hidden sm:block">
                  {(Number(intent.InputAmount) / 1e15).toFixed(3)} mETH
                </span>
                <span className={`text-xs font-medium ${expired && intent.Status === 0 ? 'text-rose-400' : 'text-gray-500'} hidden sm:block tabular-nums`}>
                  {expired ? 'expired' : deadline.toLocaleDateString()}
                </span>
                <span className={`text-xs font-medium px-2 py-0.5 rounded-full border ${status.badge}`}>
                  {status.label}
                </span>
              </div>
            </div>
          )
        })}

        {intents.length === 0 && (
          <div className="flex flex-col items-center justify-center py-10 gap-2">
            <span className="text-2xl opacity-20">⚡</span>
            <p className="text-sm text-gray-600">No intents yet</p>
          </div>
        )}
      </div>
    </div>
  )
}