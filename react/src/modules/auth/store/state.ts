export type AuthState = {
    user: UserState | null
    refreshToken: string | null
    status: 'pending' | 'fullfilled' | 'error'| null
    error: string | null
}
export type UserState = {
    username: string
    email: string
    name?: string
    surname?: string
    role: string
    iss: string,
    exp: number
}