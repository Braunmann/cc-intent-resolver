// src/config/env.ts
function requireEnv(key: string): string {
  const value = import.meta.env[key]
  if (!value) throw new Error(`Missing required env variable: ${key}`)
  return value
}

export const config = {
  apiUrl: import.meta.env.VITE_API_URL ?? 'http://localhost:8080',
  chainId: Number(import.meta.env.VITE_CHAIN_ID ?? 11155111),
  contractAddress: import.meta.env.VITE_CONTRACT_ADDRESS ?? '',
} as const