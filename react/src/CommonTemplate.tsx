import React from "react";
import CustomFooter from "./common/components/CustomFooter";
import Logo from "./common/components/Logo";
import {Link, Outlet} from "react-router-dom";
import {useAuth} from "./modules/auth/hooks/useAuth";

const CommonTemplate: React.FC = () => {
    const auth = useAuth();



    return <div>
        <div className="navbar w-full">
            <Logo className="navbar-start w-24"/>
            <div className="ml-auto">
                {
                    auth.isLoggedIn() ? <Link to="/" className="btn btn-primary"> Home </Link> : <Link to="/auth/login" className="btn btn-primary"> Login </Link>
                }
            </div>
        </div>
        <div className="p-4 md:p-24">
            <Outlet/>
        </div>
        <CustomFooter/>
    </div>;
};

export default CommonTemplate;