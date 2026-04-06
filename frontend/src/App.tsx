import { ConnectButton } from './components/ConnectButton'
import { IntentList } from './components/IntentList'
// import { CreateIntent } from './components/CreateIntent'

function App() {
  return (
    <div className="min-h-screen bg-gray-950 text-white">
      <header className="border-b border-gray-800 px-6 py-4 flex justify-between items-center">
        <h1 className="text-lg font-semibold">Intent Solver</h1>
        <ConnectButton />
      </header>
      <main className="max-w-4xl mx-auto px-6 py-8 flex flex-col gap-8">
        {/* <CreateIntent /> */}
        <IntentList />
      </main>
    </div>
  )
}

export default App