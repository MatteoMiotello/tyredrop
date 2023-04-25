import {AnyAction} from "@reduxjs/toolkit";
import React from "react";
import {useDispatch, useSelector} from "react-redux";
import {useNavigate} from "react-router-dom";
import {ThunkDispatch} from "redux-thunk";
import {LoginRequest} from "../../common/backend/requests/login-request";
import Panel from "../../common/components-library/Panel";
import logo from "../../assets/logo-transparent.png";
import {Store} from "../../store/store";
import LoginForm from "./components/LoginForm";
import {authLogin} from "./store/auth-slice";
import {useAuthenticated} from "./hooks/useAuthenticated";
import Spinner from "../../common/components/Spinner";
import {selectUserStatus} from "./store/auth-selector";

export const LoginPage: React.FC = () => {
    const dispatch: ThunkDispatch<Store, any, AnyAction> = useDispatch();
    const userStatus = useSelector( selectUserStatus );
    const navigate = useNavigate();
    const isAuthenticated = useAuthenticated();

    const login = ( loginRequest: LoginRequest ) => {
        dispatch(authLogin( loginRequest ));
    };

    const onSuccess = () => {
        navigate( '/' );
    };

    if ( isAuthenticated ) {
        navigate( '/' );
    }

    return <>
            <Panel className="flex flex-col justify-center items-center my-auto">
                { (userStatus.status == 'pending') && <Spinner></Spinner> }
                <img src={logo} width={100} alt={"Logo"}/>
                <LoginForm login={login} onSuccess={onSuccess}/>
            </Panel>
    </>;
};