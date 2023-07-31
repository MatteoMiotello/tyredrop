export  interface RegisterRequest {
    email: string
    password: string
    name?: string | null
    surname?: string | null
}