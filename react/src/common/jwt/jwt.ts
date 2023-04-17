import {JWTPayload} from "jose";
import * as jose from "jose";

const isExpired = (jwtString: string): boolean => {
    const decoded = jose.decodeJwt(jwtString);

    if ( !decoded.exp ) {
        return true;
    }

    if (Date.now() >= decoded.exp * 1000) {
        return true;
    }

    return false;
};

const decodeJwt = <T = any>(jwtString: string): JWTPayload | T => {
    return jose.decodeJwt(jwtString);
};

export const jwt = {
    decodeJwt,
    isExpired
};