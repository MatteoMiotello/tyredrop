import React, {useEffect} from "react";
import {Outlet, useNavigate} from "react-router-dom";
import CustomFooter from "../../common/components/CustomFooter";
import {useAuth} from "./hooks/useAuth";
import Logo from "../../common/components/Logo";

const AuthTemplate: React.FC = ( ) => {
    const auth = useAuth();
    const navigate = useNavigate();

    useEffect( () => {
        if ( auth.isLoggedIn() ) {
            navigate( '/' );
            return;
        }

        if ( auth.unknownStatus() ) {
            auth.tryRefreshToken();
            return;
        }

    }, [auth] );

    return <>
        <Logo width={75} className="m-6 absolute"/>
        <main className="min-h-screen lg:p-24 p-4 h-full w-full flex flex-col">
            <Outlet/>
        </main>
        <CustomFooter/>
    </>;
};

export default AuthTemplate;