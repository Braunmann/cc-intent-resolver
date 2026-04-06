import { parseAbi } from 'viem'
import { useState } from 'react'
import { config } from '../config/env'

import { useWriteContract, useWaitForTransactionReceipt, useConnection } from 'wagmi'

const abi = parseAbi([
    'function createIntent(address inputToken, uint256 inputAmount, address outputToken, uint256 minOutputAmount, uint32 targetChainId, address recipient, uint64 deadline) returns (bytes32)'
])

export function CreateIntent() {
    const { isConnected } = useConnection()
    const { writeContract, data: hash, isPending } = useWriteContract()
    const { isLoading: isConfirming, isSuccess } = useWaitForTransactionReceipt({ hash })

    const [formData, setFormData] = useState({
        inputToken: '0x7b79995e5f793A07Bc00c21412e50Ecae098E7f9',
        inputAmount: '1000000000000000',
        outputToken: '0x7b79995e5f793A07Bc00c21412e50Ecae098E7f9',
        minOutputAmount: '950000000000000',
        targetChainId: '11155420',
        recipient: '0xb1b090EA69b3Fe4D66a3FCabfdA4a958AAeA0e05',
        deadline: Math.floor(Date.now() / 1000) + 3600
    })

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target
        setFormData(prev => ({
            ...prev,
            [name]: name === 'inputAmount' || name === 'minOutputAmount' || name === 'targetChainId' || name === 'deadline' 
                ? value 
                : value
        }))
    }

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault()
        writeContract({
            address: config.contractAddress as `0x${string}`,
            abi,
            functionName: 'createIntent',
            args: [
                formData.inputToken as `0x${string}`,
                BigInt(formData.inputAmount),
                formData.outputToken as `0x${string}`,
                BigInt(formData.minOutputAmount),
                Number(formData.targetChainId),
                formData.recipient as `0x${string}`,
                BigInt(formData.deadline),
            ]
        })
    }
    const inputCls = "w-full px-3 py-2.5 rounded-xl bg-white/5 border border-white/8 text-sm text-white placeholder-gray-600 font-mono focus:outline-none focus:border-indigo-500/60 focus:bg-white/8 transition-all duration-150"
    const labelCls = "block text-xs font-medium text-gray-500 mb-1.5 uppercase tracking-wider"

    return (
        <div className="rounded-2xl border border-white/8 bg-white/3 p-6 shadow-xl">
            <div className="flex items-center justify-between mb-6">
                <h2 className="text-sm font-semibold text-white">New Intent</h2>
                {isSuccess && (
                    <span className="flex items-center gap-1.5 text-xs text-emerald-400 font-medium">
                        <span className="w-1.5 h-1.5 rounded-full bg-emerald-400"></span>
                        Created successfully
                    </span>
                )}
            </div>

            <form onSubmit={handleSubmit} className="flex flex-col gap-4">
                <div className="grid grid-cols-2 gap-3">
                    <div>
                        <label className={labelCls}>Input Token</label>
                        <input
                            type="text"
                            name="inputToken"
                            placeholder="0x…"
                            value={formData.inputToken}
                            onChange={handleChange}
                            className={inputCls}
                        />
                    </div>
                    <div>
                        <label className={labelCls}>Input Amount</label>
                        <input
                            type="text"
                            name="inputAmount"
                            placeholder="wei"
                            value={formData.inputAmount}
                            onChange={handleChange}
                            className={inputCls}
                        />
                    </div>
                </div>

                <div className="relative flex items-center justify-center">
                    <div className="absolute inset-0 flex items-center">
                        <div className="w-full border-t border-white/6"></div>
                    </div>
                    <div className="relative flex items-center justify-center w-7 h-7 rounded-full bg-[#0a0b0f] border border-white/10 text-gray-500 text-xs shadow-sm">
                        ↓
                    </div>
                </div>

                <div className="grid grid-cols-2 gap-3">
                    <div>
                        <label className={labelCls}>Output Token</label>
                        <input
                            type="text"
                            name="outputToken"
                            placeholder="0x…"
                            value={formData.outputToken}
                            onChange={handleChange}
                            className={inputCls}
                        />
                    </div>
                    <div>
                        <label className={labelCls}>Min Output Amount</label>
                        <input
                            type="text"
                            name="minOutputAmount"
                            placeholder="wei"
                            value={formData.minOutputAmount}
                            onChange={handleChange}
                            className={inputCls}
                        />
                    </div>
                </div>

                <div className="grid grid-cols-3 gap-3">
                    <div>
                        <label className={labelCls}>Target Chain ID</label>
                        <input
                            type="text"
                            name="targetChainId"
                            placeholder="e.g. 11155420"
                            value={formData.targetChainId}
                            onChange={handleChange}
                            className={inputCls}
                        />
                    </div>
                    <div className="col-span-2">
                        <label className={labelCls}>Recipient</label>
                        <input
                            type="text"
                            name="recipient"
                            placeholder="0x…"
                            value={formData.recipient}
                            onChange={handleChange}
                            className={inputCls}
                        />
                    </div>
                </div>

                <div>
                    <label className={labelCls}>Deadline (unix)</label>
                    <input
                        type="text"
                        name="deadline"
                        placeholder="unix timestamp"
                        value={formData.deadline}
                        onChange={handleChange}
                        className={inputCls}
                    />
                </div>

                <button
                    type="submit"
                    disabled={!isConnected || isPending || isConfirming}
                    className="mt-1 w-full py-2.5 rounded-xl bg-indigo-600 hover:bg-indigo-500 disabled:opacity-40 disabled:cursor-not-allowed text-sm font-semibold transition-all duration-150 shadow-lg shadow-indigo-500/20 active:scale-[0.99]"
                >
                    {isPending ? (
                        <span className="flex items-center justify-center gap-2">
                            <span className="w-3.5 h-3.5 rounded-full border-2 border-white/30 border-t-white animate-spin"></span>
                            Awaiting wallet…
                        </span>
                    ) : isConfirming ? (
                        <span className="flex items-center justify-center gap-2">
                            <span className="w-3.5 h-3.5 rounded-full border-2 border-white/30 border-t-white animate-spin"></span>
                            Confirming…
                        </span>
                    ) : !isConnected ? (
                        'Connect wallet to continue'
                    ) : (
                        'Create Intent'
                    )}
                </button>
            </form>
        </div>
    )
}