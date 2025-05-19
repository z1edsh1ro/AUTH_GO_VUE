import type { User } from '@/types/user'

export const fetchUsers = async (token: string | null): Promise<User[]> => {
  if (!token) {
    throw new Error('No authentication token found')
  }

  const endPoint = `http://localhost:8080/api/auth/user/getAll`
  const myHeaders = new Headers()
  myHeaders.append('Authorization', `Bearer ${token}`)

  const requestOptions = {
    method: 'GET',
    headers: myHeaders,
  }

  const response = await fetch(endPoint, requestOptions)

  if (response.ok) {
    const responseJson = await response.json()
    return responseJson.data
  }

  if (response.status === 401) {
    throw new Error('Unauthorized access')
  }

  throw new Error(`Cannot fetch user data`)
}
