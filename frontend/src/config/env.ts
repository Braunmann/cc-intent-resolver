// src/config/env.ts
function requireEnv(key: string): string {
  const value = import.meta.env[key]
  if (!value) throw new Error(`Missing required env variable: ${key}`)
  return value
}

export const config = {
  apiUrl: import.meta.env.VITE_API_URL ?? 'http://localhost:8080',
  contractAddresses: {
    11155111: import.meta.env.VITE_CONTRACT_ADDRESS_SEPOLIA ?? '0x0000000000000000000000000000000000000000',
    11155420: import.meta.env.VITE_CONTRACT_ADDRESS_OP_SEPOLIA ?? '0x0000000000000000000000000000000000000000',
  } as Record<number, `0x${string}`>
} as const