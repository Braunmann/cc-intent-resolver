import { createConfig, http } from 'wagmi'
import { sepolia, optimismSepolia } from 'wagmi/chains'
import { injected } from 'wagmi/connectors'

export const config = createConfig({
  chains: [sepolia, optimismSepolia],
  connectors: [injected()],
  transports: {
    [sepolia.id]: http(),
    [optimismSepolia.id]: http(),
  },
})