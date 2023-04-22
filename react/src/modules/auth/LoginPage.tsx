import {AnyAction} from "@reduxjs/toolkit";
import React, {useEffect, useState} from "react";
import {useDispatch, useSelector} from "react-redux";
import {useNavigate} from "react-router-dom";
import {ThunkDispatch} from "redux-thunk";
import {LoginRequest} from "../../common/backend/requests/login-request";
import Panel from "../../common/components-library/Panel";
import logo from "../../assets/logo-transparent.png";
import Spinner from "../../common/components/Spinner";
import {Store} from "../../store/store";
import LoginForm from "./components/LoginForm";
import { selectUserStatus} from "./store/auth-selector";
import {authLogin} from "./store/auth-slice";

export const LoginPage: React.FC = () => {
    const dispatch: ThunkDispatch<Store, any, AnyAction> = useDispatch();
    const userStatus = useSelector<Store>( selectUserStatus );
    const [ error, setError ] = useState(null);
    const navigate = useNavigate();
    const login = ( loginRequest: LoginRequest ) => {
        dispatch(authLogin( loginRequest ));
    };

    useEffect( () => {
        if ( userStatus.status == 'error' ) {
            setError( userStatus.error );
        }

        if ( userStatus.status == 'fullfilled' ) {
            navigate( '/' );
        }

    }, [ userStatus ] );

    return <div>
        <main className="bg-base lg:p-24 p-4">
            <Panel className="flex flex-col justify-center items-center my-auto relative">
                { (userStatus.status == 'pending') && <Spinner></Spinner> }
                <img src={logo} width={100} alt={"Logo"}/>
                <LoginForm login={login} error={error}/>
            </Panel>
        </main>
    </div>;
};