import { ConnectButton } from './components/ConnectButton'
import { IntentList } from './components/IntentList'
import { CreateIntent } from './components/CreateIntent'
import { useState, useCallback } from 'react'

function App() {
  const [refreshTrigger, setRefreshTrigger] = useState(0)
  const handleSuccess = useCallback(() => setRefreshTrigger(prev => prev + 1), [])

  return (
    <div className="min-h-screen bg-[#0a0b0f] text-white" style={{backgroundImage: 'radial-gradient(ellipse 80% 50% at 50% -20%, rgba(99,102,241,0.15), transparent)'}}>
      <header className="border-b border-white/5 px-6 py-4 flex justify-between items-center backdrop-blur-sm sticky top-0 z-10 bg-[#0a0b0f]/80">
        <div className="flex items-center gap-2">
          <div className="w-7 h-7 rounded-lg bg-linear-to-br from-indigo-500 to-violet-600 flex items-center justify-center text-xs font-bold shadow-lg shadow-indigo-500/25">
            ⚡
          </div>
          <span className="text-sm font-semibold bg-linear-to-r from-white to-gray-400 bg-clip-text text-transparent tracking-tight">
            Intent Solver
          </span>
        </div>
        <ConnectButton />
      </header>
      <main className="max-w-2xl mx-auto px-4 py-10 flex flex-col gap-6">
        <CreateIntent onSuccess={handleSuccess} />
        <IntentList refreshTrigger={refreshTrigger} />
      </main>
    </div>
  )
}

export default App