import {jwt} from "../../../common/jwt/jwt";
import {UserState} from "../store/state";

export enum UserStatus {
    COMPLETED,
    REGISTERING
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
    constructor( private _userState: UserState | null ) {
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

    public isTokenValid(): boolean {
        const expiration = this._userState?.exp;

        if ( expiration && Date.now() < expiration * 1000 ) {
            return true;
        }

        return false;
    }

    get user(): UserState | null{
        return this._userState;
    }
}