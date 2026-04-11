import { createConfig, http } from 'wagmi'
import { sepolia, optimismSepolia } from 'wagmi/chains'
import { injected } from 'wagmi/connectors'

export const config = createConfig({
  chains: [sepolia, optimismSepolia],
  connectors: [injected()],
  transports: {
    [sepolia.id]: http(import.meta.env.VITE_SEPOLIA_RPC_URL || undefined),
    [optimismSepolia.id]: http(import.meta.env.VITE_OP_SEPOLIA_RPC_URL || undefined),
  },
})