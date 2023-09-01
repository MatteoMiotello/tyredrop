import React, {useEffect} from "react";
import {useDispatch} from "react-redux";
import {Outlet, useNavigate} from "react-router-dom";
import {ThunkDispatch} from "redux-thunk";
import Breadcrumbs from "./common/components-library/Breadcrumbs";
import {NotAuthorized} from "./common/errors/not-authorized";
import AdminNavbar from "./modules/admin/components/AdminNavbar";
import AdminSidebar from "./modules/admin/components/AdminSidebar";
import {useAuth} from "./modules/auth/hooks/useAuth";
import {authRefreshToken} from "./modules/auth/store/auth-slice";

const Admin: React.FC = () => {
    const auth = useAuth();
    const dispatch = useDispatch<ThunkDispatch<any, any, any>>();
    const navigate = useNavigate();

    useEffect( () => {
        if (auth.unknownStatus()) {
            const refreshToken = window.localStorage.getItem('refresh_token');

            if (refreshToken) {
                dispatch(authRefreshToken(refreshToken));
                return;
            }

            navigate('/auth/login');
            return;
        }

        if (auth.isNotLoggedIn()) {
            navigate('/auth/login');
            return;
        }

        if ( auth.isLoggedIn() ) {
            if ( !auth.user?.isAdmin() ) {
                throw NotAuthorized;
            }
        }
    }, [auth] );

    return <div className="xl:px-56">
        <div className="m-1">
            <AdminNavbar/>
        </div>
        <Breadcrumbs/>
        <div className="grid grid-cols-12 gap-1 min-h-screen">
            <AdminSidebar className="col-span-2 ml-1"/>
            <div className="col-span-10 mr-1">
                <Outlet/>
            </div>
        </div>
    </div>;
};

export default Admin;