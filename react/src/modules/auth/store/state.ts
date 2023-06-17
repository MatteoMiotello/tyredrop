
export type AuthStatusType = 'registering' | null

export type AuthState = {
    user: UserState | null
    refreshToken: string | null
    status: AuthStatusType
    loggedIn: boolean | null
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
    language_code: string
    status: number
    iss: string
    exp: number
}