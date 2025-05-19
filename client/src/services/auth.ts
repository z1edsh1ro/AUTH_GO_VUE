import type { RegisterRequest, RegisterResponse } from '@/types/auth'
import type { LoginForm } from '@/types/auth'

export const login = async (payload: LoginForm): Promise<string> => {
  try {
    const response = await fetch('http://localhost:8080/api/auth/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(payload),
    })

    const data = await response.json()

    if (!response.ok) {
      throw new Error(data.message || 'Login failed')
    }

    if (!data.jwt) {
      throw new Error('JWT token missing in response')
    }

    return data.jwt
  } catch (error) {
    throw new Error(`Login error: ${error instanceof Error ? error.message : 'Unknown error'}`)
  }
}

export const register = async (payload: RegisterRequest): Promise<RegisterResponse> => {
  try {
    const response = await fetch('http://localhost:8080/api/auth/register', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(payload),
    })

    const data = await response.json()

    if (!response.ok) {
      throw new Error(data.message || 'Registration failed')
    }

    return data
  } catch (error) {
    throw new Error(`Login error: ${error instanceof Error ? error.message : 'Unknown error'}`)
  }
}
