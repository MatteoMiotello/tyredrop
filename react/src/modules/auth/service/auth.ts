import {store} from "../../../store/store";
import {authRefreshToken} from "../store/auth-slice";
import {AuthStatusType, UserState} from "../store/state";
import {User} from "./user";

export class Auth {
    _user: User | null;

    constructor(
        private _status: AuthStatusType,
        private _loggedIn: boolean | null,
        private _error: string | null | undefined,
        private _refreshToken: string | null,
        _user: UserState | null,
    ) {
        this._user = new User(_user);
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

    unknownStatus(): boolean {
        return this._loggedIn === null;
    }

    isLoggedIn(): boolean {
        return this._loggedIn === true;
    }

    isNotLoggedIn(): boolean {
        return this._loggedIn === false;
    }

    isUserCompleted(): boolean {
        if (this.user?.isCompleted()) {
            return true;
        }

        return false;
    }

    isUserRegistering(): boolean {
        return this.user?.isRegistering() as boolean;
    }

    tryRefreshToken() {
        const refreshToken = window.localStorage.getItem('refresh_token');

        if (refreshToken) {
            return store.dispatch(authRefreshToken(refreshToken));
        }

        return null;
    }
}