import {jwt} from "../../../common/jwt/jwt";
import {User} from "../models/user";
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

export class UserService {
    constructor( private _user: UserState ) {
    }

    public isRegistering(): boolean {
        return this._user.status == UserStatus.REGISTERING;
    }

    public isCompleted(): boolean {
        return this._user.status == UserStatus.COMPLETED;
    }

    get user(): UserState {
        return this._user;
    }
}