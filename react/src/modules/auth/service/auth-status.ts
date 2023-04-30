import {AuthStatusType, UserState} from "../store/state";



export class AuthStatus {
    constructor(private _status: AuthStatusType, private _error: string | null | undefined, private _user: UserState | null) {
    }

    isError(): boolean {
        return this._status === 'error';
    }

    isEmpty(): boolean {
        return this._status === null;
    }

    isFullfilled(): boolean {
        return this._status === 'fullfilled';
    }

    isPending(): boolean {
        return this._status === 'pending';
    }

    get status(): AuthStatusType {
        return this._status;
    }

    get error(): string | null | undefined {
        return this._error;
    }

    get user(): UserState | null {
        return this._user;
    }

    isAuthenticated(): boolean {
        if ( this.isEmpty() ) {
            return false;
        }

        if ( this.isError() ) {
            return false;
        }

        if ( this.isFullfilled() && this._user === null ) {
            return false;
        }

        const expiration = this.user?.exp;

        if ( expiration && Date.now() >= expiration * 1000 ) {
            return false;
        }

        return true;
    }
}