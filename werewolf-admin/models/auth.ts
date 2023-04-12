export interface ISignUpRequest {
    username: string
    password: string
}

export interface ISignInRequest {
    username: string
    password: string
}

export interface ISignInResponse {
    token: string
    username: string
    userId: string
}
