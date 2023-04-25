import React from "react";
import {Outlet} from "react-router-dom";
import logo from "../../assets/logo-transparent.png";
import CustomFooter from "../../common/components/CustomFooter";

const AuthTemplate: React.FC = ( ) => {
    return <>
        <main className="lg:p-24 p-4 h-full w-full flex flex-col">
            <img src={logo} width={100} alt={"Logo"} className="mx-auto my-6"/>
            <Outlet/>
        </main>
        <CustomFooter/>
    </>;
};

export default AuthTemplate;