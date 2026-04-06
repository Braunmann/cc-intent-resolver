import { useEffect, useState } from 'react'
import { config } from '../config/env'

type Intent = {
  ID: string
  Maker: string
  Status: number
  InputAmount: number
  Deadline: number
}

const STATUS_LABELS: Record<number, string> = {
  0: 'Created',
  1: 'Fulfilled', 
  2: 'Settled',
  3: 'Cancelled',
}

export function IntentList() {
  const [intents, setIntents] = useState<Intent[]>([])

  useEffect(() => {
    fetch(`${config.apiUrl}/v1/intents`)
      .then(r => r.json())
      .then(setIntents)
  }, [])

  return (
    <div>
      <h2 className="text-lg font-medium mb-4">Intents</h2>
      <div className="flex flex-col gap-3">
        {intents.map(intent => (
          <div key={intent.ID} className="border border-gray-800 rounded-lg p-4">
            <div className="flex justify-between items-center">
              <span className="text-sm text-gray-400 font-mono">
                {intent.ID.slice(0, 10)}...
              </span>
              <span className="text-xs px-2 py-1 rounded-full bg-gray-800">
                {STATUS_LABELS[intent.Status]}
              </span>
            </div>
            <div className="mt-2 text-sm text-gray-300">
              Amount: {Number(intent.InputAmount) / 1e15} mETH
            </div>
          </div>
        ))}
        {intents.length === 0 && (
          <p className="text-gray-500 text-sm">No intents yet.</p>
        )}
      </div>
    </div>
  )
}