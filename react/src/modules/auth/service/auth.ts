import {store} from "../../../store/store";
import {authRefreshToken} from "../store/auth-slice";
import {AuthStatusType, UserState} from "../store/state";
import {User} from "./user";

export class Auth {
    _user: User | null;

    constructor(
        private _status: AuthStatusType,
        private _error: string | null | undefined,
        private _refreshToken: string | null,
        _user: UserState | null
    ) {
        this._user = new User( _user );
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

    get user(): User | null {
        return this._user;
    }

    get refreshToken(): string | null {
        return this._refreshToken;
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

        if ( !this.user?.isTokenValid() ) {
            return false;
        }

        if ( this.isFullfilled() && !this.isEmpty() ) {
            return true;
        }

        return false;
    }

    isUserCompleted(): boolean {
        if ( this.isFullfilled() && this.user?.isCompleted()) {
            return true;
        }

        return false;
    }

    tryRefreshToken() {
        const refreshToken = window.localStorage.getItem('refresh_token');

        if (refreshToken && this.isEmpty()) {
            store.dispatch(authRefreshToken(refreshToken));
        }
    }
}