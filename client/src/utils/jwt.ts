import { jwtDecode } from 'jwt-decode'
import type { JwtPayload } from './../types/jwtPayload '

export const decodeJwt = (token: string): JwtPayload | null => {
  if (!token.trim()) {
    console.error('cannot decode jwt')
    return null
  }

  try {
    return jwtDecode<JwtPayload>(token)
  } catch (error) {
    console.error(error)
    return null
  }
}
