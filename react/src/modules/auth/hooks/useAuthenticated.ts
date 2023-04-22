import {useSelector} from "react-redux";
import {selectUser, selectUserStatus} from "../store/auth-selector";
import {useEffect, useState} from "react";

export const useAuthenticated = ( ) => {
    const user = useSelector( selectUser );
    const userStatus = useSelector( selectUserStatus );
    const [ isAuthenticated, setIsAuthenticated ] = useState( false );

    useEffect( () => {
        if ( userStatus.status != 'fullfilled' ) {
            setIsAuthenticated( false );
            return; 
        }

        if ( !user ) {
            setIsAuthenticated( false );
            return;
        }

        const expiration = user?.exp;

        if ( expiration && Date.now() >= expiration * 1000 ) {
            setIsAuthenticated(false);
            return;
        }

        setIsAuthenticated(true);
    });

    return isAuthenticated;
};