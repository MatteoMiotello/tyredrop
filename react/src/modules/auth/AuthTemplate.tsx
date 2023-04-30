import React, {useEffect} from "react";
import {Outlet, useNavigate} from "react-router-dom";
import logo from "../../assets/logo-transparent.png";
import CustomFooter from "../../common/components/CustomFooter";
import {useAuth} from "./hooks/useAuth";

const AuthTemplate: React.FC = ( ) => {
    const auth = useAuth();
    const navigate = useNavigate();

    useEffect( () => {
        if ( auth.isAuthenticated() ) {
            navigate( '/' );
        }

        auth.tryRefreshToken();
    }, [auth] );

    return <>
        <img src={logo} width={75} alt={"Logo"} className="m-6 absolute"/>
        <main className="min-h-screen lg:p-24 p-4 h-full w-full flex flex-col">
            <Outlet/>
        </main>
        <CustomFooter/>
    </>;
};

export default AuthTemplate;