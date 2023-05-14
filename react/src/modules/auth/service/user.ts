import {jwt} from "../../../common/jwt/jwt";
import {UserState} from "../store/state";
import moment from "moment";

export enum UserStatus {
    COMPLETED,
    REGISTERING,
    NOT_CONFIRMED,
}

export const getUserStatus = ( user: UserState ): UserStatus => {
    return user.status as UserStatus;
};

export const extractFromJwt = ( accessToken: string ): User => {
    if (jwt.isExpired(accessToken)) {
        throw new Error( 'Token is expired' );
    }

    return jwt.decodeJwt<User>(accessToken) as User;
};

export class User {
    constructor(
        private _userState: UserState | null
    ) {
    }

    public isRegistering(): boolean {
        if ( !this._userState ) {
            return false;
        }

        return this._userState.status == UserStatus.REGISTERING;
    }

    public isCompleted(): boolean {
        if ( !this._userState ) {
            return false;
        }

        return this._userState.status == UserStatus.COMPLETED;
    }

    public isNotConfirmed(): boolean {
        if ( !this._userState ) {
            return true;
        }

        return this._userState.status == UserStatus.NOT_CONFIRMED;
    }

    public isTokenValid(): boolean {
        const expiration = this._userState?.exp;

        if ( expiration && Date.now() < expiration * 1000 ) {
            return true;
        }

        return false;
    }

    public getExpiration(): Date | null {
        if ( !this._userState ) {
            return null;
        }

        return moment( this._userState.exp ).toDate();
    }

    public getCompleteName(): string | null {
        if ( !this.user?.name ) {
            return this.user?.username || '';
        }

        return this.user?.name + ' ' + this._userState?.surname;
    }

    get user(): UserState | null{
        return this._userState;
    }
}