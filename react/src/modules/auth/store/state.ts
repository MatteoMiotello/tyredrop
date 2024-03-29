
export type AuthStatusType = 'registering' | null

export type AuthState = {
    user: UserState | null
    refreshToken: string | null
    status: AuthStatusType
    loggedIn: boolean | null
    error: string | null
    loading: boolean
}

export type UserState = {
    userID: string
    username: string
    email: string
    name?: string
    surname?: string
    avatarUrl?: string
    userCode?: string
    role: {
        name: string
        code: string
    }
    language_code: string
    status: number
    iss: string
    exp: number
}