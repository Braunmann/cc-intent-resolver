import { parseAbi, erc20Abi } from 'viem'
import { useState, useEffect, useRef, useMemo } from 'react'
import { useWriteContract, useWaitForTransactionReceipt, useAccount, useChainId, useSwitchChain } from 'wagmi'
import { waitForTransactionReceipt, readContract } from 'wagmi/actions'
import { config as wagmiConfig } from '../config/wagmi'
import { config } from '../config/env'

type TokenMap = Record<string, string>
type TokenRegistry = Record<string, TokenMap>

type ChainInfo = {
    id: number
    name: string
    registryKey: string
}

type AppConfig = {
    chains: ChainInfo[]
    slippage: number
    tokenDecimals: Record<string, number>
    tokenPriceKey: Record<string, string>
    tokens: TokenRegistry
}

function sortTokensWethFirst(tokenMap: TokenMap): [string, string][] {
    return Object.entries(tokenMap).sort(([a], [b]) => {
        if (a === 'WETH') return -1
        if (b === 'WETH') return 1
        return a.localeCompare(b)
    })
}

function getTokenName(tokenMap: TokenMap, addr: string): string {
    return Object.entries(tokenMap).find(([, a]) => a.toLowerCase() === addr.toLowerCase())?.[0] ?? 'WETH'
}

function getDecimals(decimalsMap: Record<string, number>, name: string): number {
    return decimalsMap[name] ?? 18
}

function formatUnits(wei: string, decimals: number): string {
    const value = BigInt(wei)
    const divisor = BigInt(10 ** decimals)
    const integer = value / divisor
    const fraction = value % divisor
    const fractionStr = fraction.toString().padStart(decimals, '0').slice(0, 4)
    return `${integer}.${fractionStr}`
}

const abi = parseAbi([
    'function createIntent(address inputToken, uint256 inputAmount, address outputToken, uint256 minOutputAmount, uint32 targetChainId, address recipient, uint64 deadline) returns (bytes32)'
])

export function CreateIntent({ onSuccess }: { onSuccess?: () => void }) {
    const { isConnected, address } = useAccount()
    const chainId = useChainId()
    const { writeContract, writeContractAsync, data: hash, isPending } = useWriteContract()
    const { isLoading: isConfirming, isSuccess } = useWaitForTransactionReceipt({ hash })
    const { switchChain } = useSwitchChain()
    const [txStatus, setTxStatus] = useState<string | null>(null)

    const [appConfig, setAppConfig] = useState<AppConfig | null>(null)
    const [tokens, setTokens] = useState<TokenRegistry>({})
    const [inputAmountDisplay, setInputAmountDisplay] = useState('0.001')
    const [isPriceFetching, setIsPriceFetching] = useState(false)
    const debounceRef = useRef<ReturnType<typeof setTimeout> | null>(null)

    const [formData, setFormData] = useState({
        inputToken: '0x0000000000000000000000000000000000000000',
        inputAmount: '1000000000000000',
        outputToken: '0x0000000000000000000000000000000000000000',
        minOutputAmount: '0',
        targetChainId: '11155420',
        recipient: '0x0000000000000000000000000000000000000000',
        deadline: Math.floor(Date.now() / 1000) + 3600
    })

    const intentHubAddress = config.contractAddresses[chainId]

    const chainToRegistryKey = useMemo<Record<number, string>>(
        () => Object.fromEntries((appConfig?.chains ?? []).map(c => [c.id, c.registryKey])),
        [appConfig]
    )

    const sourceRegistryKey = chainToRegistryKey[chainId] ?? ''
    const targetRegistryKey = chainToRegistryKey[Number(formData.targetChainId)] ?? ''
    const sourceTokens = tokens[sourceRegistryKey] ?? {}
    const targetTokens = tokens[targetRegistryKey] ?? {}

    useEffect(() => {
        fetch(`${config.apiUrl}/v1/config`)
            .then(r => r.json())
            .then((cfg: AppConfig) => {
                setAppConfig(cfg)
                setTokens(cfg.tokens)
                const srcKey = cfg.chains.find(c => c.id === chainId)?.registryKey ?? ''
                const tgtKey = cfg.chains.find(c => c.id === Number(formData.targetChainId))?.registryKey ?? ''
                const wethSrc = srcKey ? (cfg.tokens[srcKey]?.['WETH'] ?? Object.values(cfg.tokens[srcKey] ?? {})[0]) : undefined
                const wethTgt = tgtKey ? (cfg.tokens[tgtKey]?.['WETH'] ?? Object.values(cfg.tokens[tgtKey] ?? {})[0]) : undefined
                setFormData(prev => ({
                    ...prev,
                    ...(wethSrc ? { inputToken: wethSrc } : {}),
                    ...(wethTgt ? { outputToken: wethTgt } : {}),
                }))
            })
            .catch(() => {})
    }, [])

    const calcMinOutput = async (humanAmount: string, inputAddr: string, outputAddr: string, registry: TokenRegistry, srcKey?: string, tgtKey?: string) => {
        if (!appConfig) return
        const amount = parseFloat(humanAmount)
        if (!amount || isNaN(amount)) return

        const inputName = getTokenName(registry[srcKey ?? sourceRegistryKey] ?? {}, inputAddr)
        const outputName = getTokenName(registry[tgtKey ?? targetRegistryKey] ?? {}, outputAddr)

        const inputKey = appConfig.tokenPriceKey[inputName]
        const outputKey = appConfig.tokenPriceKey[outputName]
        if (!inputKey || !outputKey) return

        setIsPriceFetching(true)
        try {
            const res = await fetch(`${config.apiUrl}/v1/prices`)
            if (!res.ok) throw new Error('Failed to fetch prices')
            const prices = await res.json() as Record<string, number>

            const inputUsd = prices[inputKey]
            const outputUsd = prices[outputKey]
            if (!inputUsd || !outputUsd) return

            const expectedOutput = (amount * inputUsd) / outputUsd
            const minOutput = expectedOutput * (1 - appConfig.slippage)
            const decimals = getDecimals(appConfig.tokenDecimals, outputName)
            const minOutputUnits = Math.floor(minOutput * 10 ** decimals)

            setFormData(prev => ({
                ...prev,
                minOutputAmount: minOutputUnits.toString(),
            }))
        } catch {
        } finally {
            setIsPriceFetching(false)
        }
    }

    const handleInputAmountChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const value = e.target.value
        setInputAmountDisplay(value)

        const parsed = parseFloat(value)
        if (!isNaN(parsed) && parsed > 0) {
            const inputName = getTokenName(sourceTokens, formData.inputToken)
            const decimals = getDecimals(appConfig?.tokenDecimals ?? {}, inputName)
            const wei = BigInt(Math.floor(parsed * 10 ** decimals))
            setFormData(prev => ({ ...prev, inputAmount: wei.toString() }))
        }

        if (debounceRef.current) clearTimeout(debounceRef.current)
        debounceRef.current = setTimeout(() => {
            calcMinOutput(value, formData.inputToken, formData.outputToken, tokens)
        }, 600)
    }

    const handleTokenChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
        const { name, value } = e.target
        const updated = { ...formData, [name]: value }
        setFormData(updated)
        if (debounceRef.current) clearTimeout(debounceRef.current)
        debounceRef.current = setTimeout(() => {
            calcMinOutput(inputAmountDisplay, updated.inputToken, updated.outputToken, tokens, sourceRegistryKey, targetRegistryKey)
        }, 100)
    }

    const handleSourceChainChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
        const newChainId = Number(e.target.value)
        switchChain({ chainId: newChainId as 11155111 | 11155420 })
    }

    const handleTargetChainChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
        const newChainId = e.target.value
        const key = chainToRegistryKey[Number(newChainId)]
        const map = key ? (tokens[key] ?? {}) : {}
        const defaultToken = map['WETH'] ?? Object.values(map)[0] ?? '0x0000000000000000000000000000000000000000'
        setFormData(prev => ({ ...prev, targetChainId: newChainId, outputToken: defaultToken }))
    }

    const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
        const { name, value } = e.target
        setFormData(prev => ({ ...prev, [name]: value }))
    }

    useEffect(() => {
        if (isSuccess && onSuccess) {
            onSuccess()
        }
    }, [isSuccess, onSuccess])

    useEffect(() => {
        if (address) {
            setFormData(prev => ({ ...prev, recipient: address }))
        }
    }, [address])

    useEffect(() => {
        const key = chainToRegistryKey[chainId]
        if (key) {
            const map = tokens[key] ?? {}
            const defaultToken = map['WETH'] ?? Object.values(map)[0] ?? '0x0000000000000000000000000000000000000000'
            setFormData(prev => ({ ...prev, inputToken: defaultToken }))
        }
    }, [chainId, tokens, chainToRegistryKey])

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault()
        if (!address) return

        const neededAmount = BigInt(formData.inputAmount)
        const tokenAddress = formData.inputToken as `0x${string}`
        const sourceChainId = chainId as 11155111 | 11155420

        try {
            const [balance, allowance] = await Promise.all([
                readContract(wagmiConfig, {
                    address: tokenAddress,
                    abi: erc20Abi,
                    functionName: 'balanceOf',
                    args: [address],
                    chainId: sourceChainId,
                }),
                readContract(wagmiConfig, {
                    address: tokenAddress,
                    abi: erc20Abi,
                    functionName: 'allowance',
                    args: [address, intentHubAddress],
                    chainId: sourceChainId,
                }),
            ])


            const balanceCheck = {
                address: tokenAddress,
                abi: erc20Abi,
                functionName: 'balanceOf',
                args: [address],
                chainId: sourceChainId,
            }

            console.log('Balance check:', balanceCheck)
            console.log('Balance:', balance)


            if (balance < neededAmount) {
                setTxStatus('Insufficient token balance')
                return
            }

            if (allowance < neededAmount) {
                setTxStatus('Approving…')
                const approveTx = await writeContractAsync({
                    address: tokenAddress,
                    abi: erc20Abi,
                    functionName: 'approve',
                    args: [intentHubAddress, neededAmount],
                })
                await waitForTransactionReceipt(wagmiConfig, { hash: approveTx })
            }

            const args = [
                tokenAddress,
                neededAmount,
                formData.outputToken as `0x${string}`,
                BigInt(formData.minOutputAmount),
                Number(formData.targetChainId),
                formData.recipient as `0x${string}`,
                BigInt(formData.deadline),
            ] as const

            setTxStatus('Creating intent…')
            writeContract({
                address: intentHubAddress,
                abi,
                functionName: 'createIntent',
                args,
            }, {
                onError: () => setTxStatus(null),
            })
        } catch {
            setTxStatus(null)
        }
    }

    const inputCls = "w-full px-3 py-2.5 rounded-xl bg-white/5 border border-white/8 text-sm text-white placeholder-gray-600 font-mono focus:outline-none focus:border-indigo-500/60 focus:bg-white/8 transition-all duration-150"
    const selectCls = inputCls + " appearance-none cursor-pointer [&>option]:bg-[#0e0f14] [&>option]:text-white"
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
                        <label className={labelCls}>Source Chain</label>
                        <select
                            value={chainId}
                            onChange={handleSourceChainChange}
                            className={selectCls}
                        >
                            {(appConfig?.chains ?? []).map(chain => (
                                <option key={chain.id} value={chain.id}>
                                    {chain.name}
                                </option>
                            ))}
                        </select>
                    </div>
                    <div>
                        <label className={labelCls}>Input Token</label>
                        <select
                            name="inputToken"
                            value={formData.inputToken}
                            onChange={handleTokenChange}
                            className={selectCls}
                        >
                            {sortTokensWethFirst(sourceTokens).map(([name, addr]) => (
                                <option key={addr} value={addr}>{name}</option>
                            ))}
                        </select>
                        <span className="text-[10px] text-gray-600 mt-1 block font-mono truncate">{formData.inputToken}</span>
                    </div>
                </div>

                <div>
                    <label className={labelCls}>Input Amount</label>
                    <input
                        type="text"
                        name="inputAmountDisplay"
                        placeholder="0.0"
                        value={inputAmountDisplay}
                        onChange={handleInputAmountChange}
                        className={inputCls}
                    />
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
                        <label className={labelCls}>Target Chain</label>
                        <select
                            name="targetChainId"
                            value={formData.targetChainId}
                            onChange={handleTargetChainChange}
                            className={selectCls}
                        >
                            {(appConfig?.chains ?? []).map(chain => (
                                <option key={chain.id} value={chain.id}>
                                    {chain.name}
                                </option>
                            ))}
                        </select>
                    </div>
                    <div>
                        <label className={labelCls}>Output Token</label>
                        <select
                            name="outputToken"
                            value={formData.outputToken}
                            onChange={handleTokenChange}
                            className={selectCls}
                        >
                            {sortTokensWethFirst(targetTokens).map(([name, addr]) => (
                                <option key={addr} value={addr}>{name}</option>
                            ))}
                        </select>
                        <span className="text-[10px] text-gray-600 mt-1 block font-mono truncate">{formData.outputToken}</span>
                    </div>
                </div>

                <div>
                    <label className={labelCls}>
                        Min Output
                        <span className="ml-1 text-gray-600 normal-case">({((appConfig?.slippage ?? 0.05) * 100).toFixed(0)}% slippage)</span>
                        {isPriceFetching && (
                            <span className="ml-2 inline-block w-2.5 h-2.5 rounded-full border border-gray-500 border-t-gray-300 animate-spin align-middle"></span>
                        )}
                    </label>
                    <input
                        type="text"
                        name="minOutputAmount"
                        placeholder="wei"
                        value={formData.minOutputAmount}
                        onChange={handleChange}
                        className={inputCls}
                    />
                    {formData.minOutputAmount !== '0' && formData.minOutputAmount && (
                        <span className="text-xs text-gray-500 mt-1 block">
                            ≈ {formatUnits(formData.minOutputAmount, getDecimals(appConfig?.tokenDecimals ?? {}, getTokenName(targetTokens, formData.outputToken)))} {getTokenName(targetTokens, formData.outputToken)}
                        </span>
                    )}
                </div>

                <div>
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

                {txStatus && !isSuccess && (
                    <p className="text-xs text-indigo-300 text-center">{txStatus}</p>
                )}

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