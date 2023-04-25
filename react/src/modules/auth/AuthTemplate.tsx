import React from "react";
import {Outlet} from "react-router-dom";
import logo from "../../assets/logo-transparent.png";
import CustomFooter from "../../common/components/CustomFooter";

const AuthTemplate: React.FC = ( ) => {
    return <>
        <img src={logo} width={75} alt={"Logo"} className="m-6 absolute"/>
        <main className="lg:p-24 p-4 h-full w-full flex flex-col">
            <Outlet/>
        </main>
        <CustomFooter/>
    </>;
};

export default AuthTemplate;