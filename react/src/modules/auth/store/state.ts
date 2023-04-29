
export type AuthStatusType = 'pending' | 'fullfilled' | 'error' | 'registering' | null

export type AuthState = {
    user: UserState | null
    refreshToken: string | null
    status: AuthStatusType
    error: string | null
}

export type UserState = {
    username: string
    email: string
    name?: string
    surname?: string
    role: {
        name: string
        code: string
    }
    status: number
    iss: string
    exp: number
}