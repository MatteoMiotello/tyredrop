import {AuthStatusType} from "../store/state";

export class AuthStatus {
    constructor(private _status: AuthStatusType, private _error: string | null | undefined) {
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
}