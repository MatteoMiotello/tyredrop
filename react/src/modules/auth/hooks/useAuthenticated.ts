import {useSelector} from "react-redux";
import {selectAuthStatus, selectUser} from "../store/auth-selector";
import {useEffect, useState} from "react";

export const useAuthenticated = ( ) => {
    const user = useSelector( selectUser );
    const authStatus = useSelector( selectAuthStatus );
    const [ isAuthenticated, setIsAuthenticated ] = useState( false );

    useEffect( () => {
        if ( authStatus.isError() ) {
            setIsAuthenticated(false);
            return;
        }

        if ( authStatus.isFullfilled() && user === null ) {
            setIsAuthenticated( false );
            return;
        }

        const expiration = user?.exp;

        if ( expiration && Date.now() >= expiration * 1000 ) {
            setIsAuthenticated(false);
            return;
        }

        setIsAuthenticated(true);
    }, [ authStatus, user ]);

    return isAuthenticated;
};